package config

import (
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"

	logger "github.com/panlibin/vglog"
)

type ItemConfig struct {
	mapItemConf map[int32]*iconfig.ItemConf
}

func newItemConfig() *ItemConfig {
	pObj := new(ItemConfig)
	pObj.mapItemConf = make(map[int32]*iconfig.ItemConf)

	return pObj
}

func (ic *ItemConfig) OnLoadConfig() error {
	const sqlSelect = "select item_id,name,first_type,second_type,quality,lv,use_type,hero_random,hero_force,hero_brains,hero_politics,hero_charm," +
		"skill_exp,book_exp,beauty_intimacy,beauty_charm,family_money,family_contribution,attribute,reward,hero_id from cfg_item"
	rows, err := public.Server.GetConfigDb().Query(0, sqlSelect)
	for {
		if err != nil {
			break
		}

		for rows.Next() {
			pItemConf := new(iconfig.ItemConf)
			err = rows.Scan(&pItemConf.Id, &pItemConf.Name, &pItemConf.FirstType, &pItemConf.SecondType, &pItemConf.Quality, &pItemConf.Lv, &pItemConf.UseType, &pItemConf.HeroRandom,
				&pItemConf.HeroForce, &pItemConf.HeroBrains, &pItemConf.HeroPolitics, &pItemConf.HeroCharm, &pItemConf.SkillExp, &pItemConf.BookExp, &pItemConf.BeautyIntimacy,
				&pItemConf.BeautyCharm, &pItemConf.FamilyMoney, &pItemConf.FamilyContribution, &pItemConf.Attribute, &pItemConf.Reward, &pItemConf.HeroId)
			if err != nil {
				break
			}
			ic.mapItemConf[pItemConf.Id] = pItemConf
		}
		rows.Close()
		break
	}
	if err != nil {
		logger.Errorf("load cfg_item error: %v", err)
	}
	return err
}

func (ic *ItemConfig) GetConf(id int32) *iconfig.ItemConf {
	pConf, exist := ic.mapItemConf[id]
	if !exist {
		return nil
	}
	return pConf
}
