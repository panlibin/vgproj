package rank

import (
	"sort"
	irank "vgproj/vggame/public/game/rank"
)

type sortableRankArray []irank.IRankObject

func (ra sortableRankArray) Len() int {
	return len(ra)
}

func (ra sortableRankArray) Swap(i, j int) {
	ra[i], ra[j] = ra[j], ra[i]
}

func (ra sortableRankArray) Less(i, j int) bool {
	return ra[i].Less(ra[j])
}

func (ra sortableRankArray) Search(obj irank.IRankObject) int {
	return sort.Search(len(ra), func(i int) bool {
		return ra[i].LessOrEqual(obj)
	})
}
