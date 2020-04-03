package rank

import (
	"encoding/json"
	"fmt"
	"sort"
	"vgproj/vggame/public"
	irank "vgproj/vggame/public/game/rank"
)

type RankManager struct {
	mapRankList map[int32]*RankList
}

func NewRankManager() *RankManager {
	return &RankManager{
		mapRankList: map[int32]*RankList{
			irank.RankTypeCombat: newRankList(irank.RankTypeCombat, 1000),
		},
	}
}

func (rm *RankManager) OnLoadData() error {
	rows, err := public.Server.GetGlobalDb().Query(0, sqlLoadRank)
	var tmpRankType int32
	var tmpExtra []byte
	if err != nil {
		goto endPoint
	}

	for rows.Next() {
		ro := &irank.NormalRankObject{}
		err = rows.Scan(&tmpRankType, &ro.ObjData.Id, &ro.Val, &ro.ChangeTime, &tmpExtra, ro.ObjData.Name, ro.ObjData.ServerId, ro.ObjData.Lev, ro.ObjData.TitleId)
		if err != nil {
			goto endPoint
		}
		err = json.Unmarshal(tmpExtra, &ro.Extra)
		if err != nil {
			goto endPoint
		}
		rl, exist := rm.mapRankList[tmpRankType]
		if !exist {
			err = fmt.Errorf("rank manager load data. unknown rank type %d!", tmpRankType)
			goto endPoint
		}
		rl.arrObj = append(rl.arrObj, ro)
		rl.mapObj[ro.ObjData.Id] = ro
	}

endPoint:
	return err
}

func (rm *RankManager) OnInit() error {
	for _, rl := range rm.mapRankList {
		sort.Sort(sort.Reverse(rl.arrObj))
		rl.clamp()
	}

	return nil
}

func (rm *RankManager) OnRelease() {

}

func (rm *RankManager) AddRankObject(rankType int32, ro irank.IRankObject) {
	rl, exist := rm.mapRankList[rankType]
	if !exist {
		return
	}
	rl.AddRankObject(ro)
}

func (rm *RankManager) GetObjectRank(rankType int32, id int64) int32 {
	rl, exist := rm.mapRankList[rankType]
	if !exist {
		return 0
	}
	return rl.GetObjectRank(id)
}

func (rm *RankManager) GetListByRank(rankType int32, begin int32, count int32) []irank.IRankObject {
	rl, exist := rm.mapRankList[rankType]
	if !exist {
		return nil
	}
	return rl.GetListByRank(begin, count)
}
