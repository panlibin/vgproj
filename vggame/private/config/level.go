package config

import (
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgstr"
)

type LevelConfig struct {
	arrLevConf []*iconfig.LevelConf
	maxLev     int32
}

func newLevelConfig() *LevelConfig {
	pObj := new(LevelConfig)

	return pObj
}

func (lc *LevelConfig) OnLoadConfig() error {
	const sqlSelect = "select lv,exp,hero_id,beauty_id from cfg_lv"
	rows, err := public.Server.GetConfigDb().Query(0, sqlSelect)
	mapTmpLevConf := make(map[int32]*iconfig.LevelConf)
	for {
		if err != nil {
			break
		}

		var tmpHeroIds string
		var tmpBeautyIds string
		for rows.Next() {
			pLevelConf := new(iconfig.LevelConf)
			err = rows.Scan(&pLevelConf.Lev, &pLevelConf.Exp, &tmpHeroIds, &tmpBeautyIds)
			if err != nil {
				break
			}

			pLevelConf.HeroIds, _ = vgstr.SplitToInt32Array(tmpHeroIds, ",")
			pLevelConf.BeautyIds, _ = vgstr.SplitToInt32Array(tmpBeautyIds, ",")

			mapTmpLevConf[pLevelConf.Lev] = pLevelConf
			if pLevelConf.Lev > lc.maxLev {
				lc.maxLev = pLevelConf.Lev
			}
		}
		rows.Close()
		break
	}
	if err != nil {
		logger.Errorf("load cfg_science error: %v", err)
		return err
	}

	lc.arrLevConf = make([]*iconfig.LevelConf, lc.maxLev+1)
	for k, v := range mapTmpLevConf {
		lc.arrLevConf[k] = v
	}

	return nil
}

func (lc *LevelConfig) GetLevConf(lev int32) *iconfig.LevelConf {
	if lev <= 0 || lev >= int32(len(lc.arrLevConf)) {
		return nil
	}

	return lc.arrLevConf[lev]
}

func (lc *LevelConfig) GetMaxLevel() int32 {
	return lc.maxLev
}
