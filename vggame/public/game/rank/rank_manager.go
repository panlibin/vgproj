package irank

type IRankManager interface {
	AddRankObject(rankType int32, ro IRankObject)
	GetObjectRank(rankType int32, id int64) int32
	GetListByRank(rankType int32, begin int32, count int32) []IRankObject
}
