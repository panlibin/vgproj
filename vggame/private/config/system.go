package config

import (
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"

	logger "github.com/panlibin/vglog"
)

type SystemConfig struct {
	mapSysConf map[int32]iconfig.IModuleSysConfig
}

func newSystemConfig() *SystemConfig {
	pObj := new(SystemConfig)
	pObj.mapSysConf = map[int32]iconfig.IModuleSysConfig{
		//iconfig.SystemConfig_Hero:			new(iconfig.HeroSysConfig),
		//iconfig.SystemConfig_Reborn:			new(iconfig.RebornSysConfig),
		//iconfig.SystemConfig_FansGift:		new(iconfig.FansGiftSysConfig),
	}

	return pObj
}

func (sc *SystemConfig) OnLoadConfig() error {
	const sqlSystem = "select * from cfg_config"
	rows, err := public.Server.GetConfigDb().Query(0, sqlSystem)
	for {
		if err != nil {
			break
		}

		var tmpId int32
		tmpArrStrData := make([]string, 13)
		tmpArrData := make([]interface{}, 14)
		tmpArrData[0] = &tmpId
		for i := 1; i < 14; i++ {
			tmpArrData[i] = &tmpArrStrData[i-1]
		}
		for rows.Next() {
			err = rows.Scan(tmpArrData...)
			if err != nil {
				break
			}
			pConf, exist := sc.mapSysConf[tmpId]
			if exist && pConf != nil {
				pConf.Init(tmpArrStrData)
				sc.mapSysConf[tmpId] = pConf
			}
		}
		rows.Close()
		break
	}
	if err != nil {
		logger.Errorf("load cfg_config error: %v", err)
	}
	return err
}

func (sc *SystemConfig) GetConf(id int32) iconfig.IModuleSysConfig {
	pConf, exist := sc.mapSysConf[id]
	if !exist {
		return nil
	}
	return pConf
}
