package irank

type IRankObject interface {
	GetId() int64
	GetObjectData() *ObjectData
	GetValue() int64
	GetChangeTime() int64
	GetExtra() []int64
	Less(IRankObject) bool
	LessOrEqual(IRankObject) bool
}

type ObjectData struct {
	Id   int64
	Name string
}

type NormalRankObject struct {
	ObjData    ObjectData
	Val        int64
	ChangeTime int64
}

func (ro *NormalRankObject) GetId() int64 {
	return ro.ObjData.Id
}

func (ro *NormalRankObject) GetObjectData() *ObjectData {
	return &ro.ObjData
}

func (ro *NormalRankObject) GetValue() int64 {
	return ro.Val
}

func (ro *NormalRankObject) GetChangeTime() int64 {
	return ro.ChangeTime
}

func (ro *NormalRankObject) GetExtra() []int64 {
	return nil
}

func (ro *NormalRankObject) Less(t IRankObject) bool {
	if ro.Val != t.GetValue() {
		return ro.Val < t.GetValue()
	} else {
		return ro.ChangeTime > t.GetChangeTime()
	}
}

func (ro *NormalRankObject) LessOrEqual(t IRankObject) bool {
	if ro.Val != t.GetValue() {
		return ro.Val <= t.GetValue()
	} else {
		return ro.ChangeTime >= t.GetChangeTime()
	}
}
