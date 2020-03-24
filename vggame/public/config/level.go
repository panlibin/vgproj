package iconfig

type LevelConf struct {
	Lev       int32
	Exp       int64
	HeroIds   []int32
	BeautyIds []int32
}

type ILevelConfig interface {
	IConfig
	GetLevConf(lev int32) *LevelConf
	GetMaxLevel() int32
}
