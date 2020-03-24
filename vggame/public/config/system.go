package iconfig

const (
	SystemConfig_Hero = 1
)

type IModuleSysConfig interface {
	Init(arrData []string)
}

//
//type HeroSysConfig struct {
//	Id int32
//	MaxHeroLevel int64
//	MaxHeroStar int32
//}
//
//func (this *HeroSysConfig) Init(arrData []string) {
//	this.Id = SystemConfig_Hero
//	this.MaxHeroLevel, _ = strconv.ParseInt(arrData[0], 10, 64)
//	maxStar, _ := strconv.ParseInt(arrData[1], 10, 64)
//	this.MaxHeroStar = int32(maxStar)
//}
//
//type RebornSysConfig struct {
//	Id int32
//	LevNeedIncomeFactor1 vg_number.BigIntSimilar
//	LevNeedIncomeFactor2 vg_number.BigIntSimilar
//	LevNeedIncomeFactor3 int64
//	LevNeedIncomeFactor4 vg_number.BigIntSimilar
//	LevNeedIncomeFactor5 vg_number.BigIntSimilar
//	DollarRewardFactor1 vg_number.BigIntSimilar
//	DollarRewardFactor2 vg_number.BigIntSimilar
//	DollarRewardFactor3 int64
//	DollarRewardFactor4 vg_number.BigIntSimilar
//	PropRewardFactor1 int64
//	PropRewardId int32
//	CoolDown int64
//}
//
//func (this *RebornSysConfig) Init(arrData []string) {
//	this.Id = SystemConfig_Reborn
//	this.LevNeedIncomeFactor1.SetNumberByString(arrData[0])
//	this.LevNeedIncomeFactor2.SetNumberByString(arrData[1])
//	this.LevNeedIncomeFactor3, _ = strconv.ParseInt(arrData[2], 10, 64)
//	this.LevNeedIncomeFactor4.SetNumberByString(arrData[3])
//	this.LevNeedIncomeFactor5.SetNumberByString(arrData[4])
//	this.DollarRewardFactor1.SetNumberByString(arrData[5])
//	this.DollarRewardFactor2.SetNumberByString(arrData[6])
//	this.DollarRewardFactor3, _ = strconv.ParseInt(arrData[7], 10, 64)
//	this.DollarRewardFactor4.SetNumberByString(arrData[8])
//	this.PropRewardFactor1, _ = strconv.ParseInt(arrData[9], 10, 64)
//	propRewardId, _ := strconv.ParseInt(arrData[10], 10, 32)
//	this.PropRewardId = int32(propRewardId)
//	this.CoolDown, _ = strconv.ParseInt(arrData[11], 10, 64)
//}
//
//type FansGiftSysConfig struct {
//	Id int32
//	LoginRefreshInterval int64
//	StayTime1 int64
//	StayTime2 int64
//	RefreshInterval1 int64
//	RefreshInterval2 int64
//	RefreshMinInterval int64
//	MaxGiftCount int32
//	StageSlotCount int32
//}
//
//func (this *FansGiftSysConfig) Init(arrData []string) {
//	this.Id = SystemConfig_FansGift
//	this.LoginRefreshInterval, _ = strconv.ParseInt(arrData[0], 10, 64)
//	this.StayTime1, _ = strconv.ParseInt(arrData[1], 10, 64)
//	this.RefreshInterval1, _ = strconv.ParseInt(arrData[2], 10, 64)
//	this.RefreshInterval2, _ = strconv.ParseInt(arrData[4], 10, 64)
//	this.StayTime2, _ = strconv.ParseInt(arrData[3], 10, 64)
//	var tmpI64 int64
//	tmpI64, _ = strconv.ParseInt(arrData[5], 10, 32)
//	this.MaxGiftCount = int32(tmpI64)
//	tmpI64, _ = strconv.ParseInt(arrData[7], 10, 32)
//	this.StageSlotCount = int32(tmpI64)
//
//	this.RefreshMinInterval = this.LoginRefreshInterval
//	if this.RefreshInterval1 < this.RefreshMinInterval {
//		this.RefreshMinInterval = this.RefreshInterval1
//	}
//	if this.RefreshInterval2 < this.RefreshMinInterval {
//		this.RefreshMinInterval = this.RefreshInterval2
//	}
//}

type ISystemConfig interface {
	IConfig
	GetConf(id int32) IModuleSysConfig
}
