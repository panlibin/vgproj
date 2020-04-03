package rank

import (
	"encoding/json"
	"vgproj/vggame/public"
	irank "vgproj/vggame/public/game/rank"

	logger "github.com/panlibin/vglog"
)

type RankList struct {
	rankType int32
	maxRank  int
	mapObj   map[int64]irank.IRankObject
	arrObj   sortableRankArray
}

func newRankList(rankType int32, maxRank int) *RankList {
	return &RankList{
		rankType: rankType,
		maxRank:  maxRank,
		mapObj:   make(map[int64]irank.IRankObject, maxRank),
		arrObj:   make(sortableRankArray, 0, maxRank),
	}
}

func (rl *RankList) AddRankObject(ro irank.IRankObject) {
	oldRo, exist := rl.mapObj[ro.GetId()]
	if exist {
		i := rl.searchObject(oldRo)
		if i >= len(rl.arrObj) {
			return
		}

		j := rl.arrObj.Search(ro)
		if i == j {
			rl.arrObj[i] = ro
		} else if i > j {
			copy(rl.arrObj[j+1:i+1], rl.arrObj[j:i])
			rl.arrObj[j] = ro
		} else {
			copy(rl.arrObj[i:j-1], rl.arrObj[i+1:j])
			rl.arrObj[j-1] = ro
		}
		rl.updateRankObject(ro)
	} else {
		j := rl.arrObj.Search(ro)
		if j < len(rl.arrObj) {
			if len(rl.arrObj) >= rl.maxRank {
				// 可插入,排行榜满,删除最后一名
				remObj := rl.arrObj[len(rl.arrObj)-1]
				delete(rl.mapObj, remObj.GetId())
				rl.deleteRankObject(remObj.GetId())
			} else {
				// 可插入,排行榜未满,直接插入
				rl.arrObj = append(rl.arrObj, nil)
			}
			copy(rl.arrObj[j+1:], rl.arrObj[j:])
			rl.arrObj[j] = ro
			rl.insertRankObject(ro)
		} else {
			if j < rl.maxRank {
				// 排在最后一名,排行榜未满,可插入
				rl.arrObj = append(rl.arrObj, ro)
				rl.insertRankObject(ro)
			} else {
				// 排在最后一名,排行榜已满,退出
				return
			}
		}
	}
	rl.mapObj[ro.GetId()] = ro
}

func (rl *RankList) searchObject(ro irank.IRankObject) (i int) {
	i = rl.arrObj.Search(ro)

	found := false
	for ; i < len(rl.arrObj); i++ {
		if ro.GetId() == rl.arrObj[i].GetId() {
			found = true
			break
		}
	}
	if !found {
		logger.Errorf("search rank object unexpected result!")
	}
	return
}

func (rl *RankList) GetObjectRank(id int64) int32 {
	ro, exist := rl.mapObj[id]
	if !exist {
		return 0
	}
	i := rl.searchObject(ro)
	if i >= len(rl.arrObj) {
		return 0
	}
	return int32(i) + 1
}

func (rl *RankList) GetListByRank(begin, count int32) []irank.IRankObject {
	arrLen := len(rl.arrObj)
	if begin < 1 || int(begin) > arrLen || count <= 0 {
		return nil
	}
	bi := int(begin) - 1
	ei := bi + int(count)
	if ei > arrLen {
		ei = arrLen
	}
	return rl.arrObj[bi:ei]
}

func (rl *RankList) clamp() {
	if len(rl.arrObj) <= rl.maxRank {
		return
	}
	tmpDel := rl.arrObj[rl.maxRank:]
	rl.arrObj = rl.arrObj[:rl.maxRank]
	for _, ro := range tmpDel {
		delete(rl.mapObj, ro.GetId())
		rl.deleteRankObject(ro.GetId())
	}
}

func (rl *RankList) insertRankObject(ro irank.IRankObject) {
	rd := ro.GetObjectData()
	strExtra, err := json.Marshal(ro.GetExtra)
	if err != nil {
		logger.Error(err)
		return
	}
	public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlInsertRankObj, rl.rankType, rd.Id, ro.GetValue(), ro.GetChangeTime(),
		strExtra, rd.Name, rd.ServerId, rd.Lev, rd.TitleId)
}

func (rl *RankList) updateRankObject(ro irank.IRankObject) {
	rd := ro.GetObjectData()
	strExtra, err := json.Marshal(ro.GetExtra)
	if err != nil {
		logger.Error(err)
		return
	}
	public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlUpdateRankObj, ro.GetValue(), ro.GetChangeTime(), strExtra, rd.Name,
		rd.ServerId, rd.Lev, rd.TitleId, rl.rankType, rd.Id)
}

func (rl *RankList) deleteRankObject(objId int64) {
	public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlDeleteRankObj, rl.rankType, objId)
}
