package player

import (
	"database/sql"
	"fmt"

	logger "github.com/panlibin/vglog"
)

type heroModule struct {
	*playerModule

	mapHero map[int32]*Hero
}

func newHeroModule(pPlayer *Player) *heroModule {
	pObj := new(heroModule)
	pObj.playerModule = newPlayerModule(pPlayer)
	pObj.mapHero = make(map[int32]*Hero)

	//pPlayer.RegisterMessage(&c_gs.C2S_UpgradeHeroLevReq{}, pObj.handleUpgradeHeroLev)
	//pPlayer.RegisterMessage(&c_gs.C2S_UpgradeHeroStarReq{}, pObj.handleUpgradeHeroStar)

	//pPlayer.GetEventManager().Register(player.EventType_PropertyIncrease, pObj.onPropertyIncrease)

	return pObj
}

func (hm *heroModule) getLoadSql() string {
	return fmt.Sprintf("select hero_id,star,lev from player_hero where player_id=%d;", hm.player.GetId())
}

func (hm *heroModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_hero error: %v", err)
		return err
	}
	for rows.Next() {
		pHero := newHero(hm.player)
		// err := rows.Scan(&pHero.id, &pHero.star, &pHero.lev)
		// if err != nil {
		// 	return err
		// }
		pHero.inDb = true
		hm.mapHero[pHero.id] = pHero
	}

	return nil
}

func (hm *heroModule) onCreate() {

}

//func (hm *heroModule) GetHero(id int32) hero.IHero {
//	pHero := hm.getHero(id)
//	if pHero == nil {
//		return nil
//	}
//	return pHero
//}
//
//func (hm *heroModule) AddHero(id int32) {
//	pHeroConfMgr := public.Server.GetConfigManager().GetConfig(config.Config_Hero).(config.IHeroConfig)
//	pHeroConf := pHeroConfMgr.GetConf(id)
//	if pHeroConf == nil {
//		return
//	}
//
//	pHero := hm.getHero(id)
//	if pHero != nil {
//		return
//	}
//
//	pHero = newHero(hm.Player)
//	pHero.id = id
//	pHero.save()
//	hm.mapHero[id] = pHero
//}

func (hm *heroModule) getHero(id int32) *Hero {
	pHero, exist := hm.mapHero[id]
	if !exist {
		return nil
	}
	return pHero
}
