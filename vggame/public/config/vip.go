package iconfig

type VipConf struct {
	Lev               int32
	Exp               int64
	ConceptionProb    int32
	FindEnergy        int32
	GreetingStamina   int32
	ChildrenStamina   int32
	MiracleTimes      int32
	FateTransferTimes int32
	OnceFind          int32
	OnceCatch         int32
	OnceGreeting      int32
	BattleSkip        int32
	HeroId            int32
	BeautyId          int32
	Gift              map[int32]int64
}

type IVipConfig interface {
	IConfig
	GetConf(lv int32) *VipConf
	GetMaxLevel() int32
}
