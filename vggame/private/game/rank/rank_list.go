package rank

import (
	irank "vgproj/vggame/public/game/rank"

	logger "github.com/panlibin/vglog"
)

type RankList struct {
	maxRank int
	mapObj  map[int64]irank.IRankObject
	arrObj  sortableRankArray
}

func newRankList(maxRank int) *RankList {
	return &RankList{
		maxRank: maxRank,
		mapObj:  make(map[int64]irank.IRankObject, maxRank),
		arrObj:  make(sortableRankArray, 0, maxRank),
	}
}

func (rl *RankList) AddRankObject(ro irank.IRankObject) {
	oldRo, exist := rl.mapObj[ro.GetId()]
	if exist {
		i := rl.arrObj.Search(oldRo)
		if i >= len(rl.arrObj) {
			logger.Errorf("search rank object unexpected result!")
			return
		}

		found := false
		for ; i < len(rl.arrObj); i++ {
			if oldRo.GetId() == rl.arrObj[i].GetId() {
				found = true
				break
			}
		}
		if !found {
			logger.Errorf("search rank object unexpected result!")
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
	} else {
		j := rl.arrObj.Search(ro)
		if j < len(rl.arrObj) {
			if len(rl.arrObj) >= rl.maxRank {
				delete(rl.mapObj, rl.arrObj[len(rl.arrObj)-1].GetId())
			} else {
				rl.arrObj = append(rl.arrObj, nil)
			}
			copy(rl.arrObj[j+1:], rl.arrObj[j:])
			rl.arrObj[j] = ro
		} else {
			if j < rl.maxRank {
				rl.arrObj = append(rl.arrObj, ro)
			} else {
				return
			}
		}
	}
	rl.mapObj[ro.GetId()] = ro
}
