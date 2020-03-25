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
//func (c *HeroSysConfig) Init(arrData []string) {
//	c.Id = SystemConfig_Hero
//	c.MaxHeroLevel, _ = strconv.ParseInt(arrData[0], 10, 64)
//	maxStar, _ := strconv.ParseInt(arrData[1], 10, 64)
//	c.MaxHeroStar = int32(maxStar)
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
//func (c *RebornSysConfig) Init(arrData []string) {
//	c.Id = SystemConfig_Reborn
//	c.LevNeedIncomeFactor1.SetNumberByString(arrData[0])
//	c.LevNeedIncomeFactor2.SetNumberByString(arrData[1])
//	c.LevNeedIncomeFactor3, _ = strconv.ParseInt(arrData[2], 10, 64)
//	c.LevNeedIncomeFactor4.SetNumberByString(arrData[3])
//	c.LevNeedIncomeFactor5.SetNumberByString(arrData[4])
//	c.DollarRewardFactor1.SetNumberByString(arrData[5])
//	c.DollarRewardFactor2.SetNumberByString(arrData[6])
//	c.DollarRewardFactor3, _ = strconv.ParseInt(arrData[7], 10, 64)
//	c.DollarRewardFactor4.SetNumberByString(arrData[8])
//	c.PropRewardFactor1, _ = strconv.ParseInt(arrData[9], 10, 64)
//	propRewardId, _ := strconv.ParseInt(arrData[10], 10, 32)
//	c.PropRewardId = int32(propRewardId)
//	c.CoolDown, _ = strconv.ParseInt(arrData[11], 10, 64)
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
//func (c *FansGiftSysConfig) Init(arrData []string) {
//	c.Id = SystemConfig_FansGift
//	c.LoginRefreshInterval, _ = strconv.ParseInt(arrData[0], 10, 64)
//	c.StayTime1, _ = strconv.ParseInt(arrData[1], 10, 64)
//	c.RefreshInterval1, _ = strconv.ParseInt(arrData[2], 10, 64)
//	c.RefreshInterval2, _ = strconv.ParseInt(arrData[4], 10, 64)
//	c.StayTime2, _ = strconv.ParseInt(arrData[3], 10, 64)
//	var tmpI64 int64
//	tmpI64, _ = strconv.ParseInt(arrData[5], 10, 32)
//	c.MaxGiftCount = int32(tmpI64)
//	tmpI64, _ = strconv.ParseInt(arrData[7], 10, 32)
//	c.StageSlotCount = int32(tmpI64)
//
//	c.RefreshMinInterval = c.LoginRefreshInterval
//	if c.RefreshInterval1 < c.RefreshMinInterval {
//		c.RefreshMinInterval = c.RefreshInterval1
//	}
//	if c.RefreshInterval2 < c.RefreshMinInterval {
//		c.RefreshMinInterval = c.RefreshInterval2
//	}
//}

type ISystemConfig interface {
	IConfig
	GetConf(id int32) IModuleSysConfig
}
