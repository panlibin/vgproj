package config

import (
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"

	logger "github.com/panlibin/vglog"
)

type VipConfig struct {
	arrVipConf []*iconfig.VipConf
	maxLev     int32
}

func newVipConfig() *VipConfig {
	pObj := new(VipConfig)

	return pObj
}

func (vc *VipConfig) OnLoadConfig() error {
	const sqlSelect = "select vip_lv,vip_exp,son_probability,hero_id,beauty_id,find_power,greeting_vigor,son_vigor,miracles,transfer,once_find," +
		"once_catch,once_greeting,battle_skip,id1,number1,id2,number2,id3,number3,id4,number4,id5,number5,id6,number6 from cfg_vip"
	rows, err := public.Server.GetConfigDb().Query(0, sqlSelect)
	mapTmpVipConf := make(map[int32]*iconfig.VipConf)
	for {
		if err != nil {
			break
		}

		for rows.Next() {
			tmpIds := make([]int32, 6)
			tmpNums := make([]int64, len(tmpIds))
			pVipConf := new(iconfig.VipConf)
			err = rows.Scan(&pVipConf.Lev, &pVipConf.Exp, &pVipConf.ConceptionProb, &pVipConf.HeroId, &pVipConf.BeautyId, &pVipConf.FindEnergy, &pVipConf.GreetingStamina,
				&pVipConf.ChildrenStamina, &pVipConf.MiracleTimes, &pVipConf.FateTransferTimes, &pVipConf.OnceFind, &pVipConf.OnceCatch, &pVipConf.OnceGreeting, &pVipConf.BattleSkip,
				&tmpIds[0], &tmpNums[0], &tmpIds[1], &tmpNums[1], &tmpIds[2], &tmpNums[2], &tmpIds[3], &tmpNums[3], &tmpIds[4], &tmpNums[4], &tmpIds[5], &tmpNums[5])
			if err != nil {
				break
			}

			pVipConf.Gift = make(map[int32]int64)
			for i, id := range tmpIds {
				num := tmpNums[i]
				if id > 0 && num > 0 {
					pVipConf.Gift[id] = num
				}
			}

			mapTmpVipConf[pVipConf.Lev] = pVipConf
			if pVipConf.Lev > vc.maxLev {
				vc.maxLev = pVipConf.Lev
			}
		}
		rows.Close()
		break
	}
	if err != nil {
		logger.Errorf("load cfg_vip error: %v", err)
		return err
	}

	vc.arrVipConf = make([]*iconfig.VipConf, vc.maxLev+1)
	for k, v := range mapTmpVipConf {
		vc.arrVipConf[k] = v
	}

	return nil
}

func (vc *VipConfig) GetConf(lev int32) *iconfig.VipConf {
	if lev <= 0 || lev >= int32(len(vc.arrVipConf)) {
		return nil
	}

	return vc.arrVipConf[lev]
}

func (vc *VipConfig) GetMaxLevel() int32 {
	return vc.maxLev
}
