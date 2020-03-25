package player

import (
	"database/sql"
	"fmt"
	"vgproj/proto/msg"

	logger "github.com/panlibin/vglog"
)

type propertyModule struct {
	*playerModule

	mapProp map[int32]*Property
}

func newPropertyModule(pPlayer *Player) *propertyModule {
	pObj := new(propertyModule)
	pObj.playerModule = newPlayerModule(pPlayer)

	pObj.mapProp = make(map[int32]*Property, 32)

	return pObj
}

func (pm *propertyModule) getLoadSql() string {
	return fmt.Sprintf("select prop_id,prop_value,update_ts from player_property where player_id=%d;", pm.player.GetId())
}

func (pm *propertyModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_property error: %v", err)
		return err
	}

	var tmpId int32
	var tmpVal int64
	var tmpTs int64
	for rows.Next() {
		err = rows.Scan(&tmpId, &tmpVal, &tmpTs)
		if err != nil {
			break
		}
		pProp, exist := pm.mapProp[tmpId]
		if !exist || pProp == nil {
			pProp = newProperty(pm.player, tmpId)
			pm.mapProp[tmpId] = pProp
		}
		pProp.Load(tmpVal, tmpTs)
	}

	if err != nil {
		logger.Errorf("load player_property error: %v", err)
	}

	return err
}

func (pm *propertyModule) onCreate() {

}

func (pm *propertyModule) GetProp(id int32) *Property {
	p, exist := pm.mapProp[id]
	if !exist {
		return nil
	}
	return p
}

func (pm *propertyModule) AddProp(id int32, val int64, source int32, sourceExt string) {
	p := pm.GetProp(id)
	if p == nil {
		p = newProperty(pm.player, id)
		pm.mapProp[id] = p
	}
	p.AddValue(val, source, sourceExt)
}

func (pm *propertyModule) SubProp(id int32, val int64, source int32, sourceExt string) bool {
	p := pm.GetProp(id)
	if p == nil {
		return false
	}
	return p.SubValue(val, source, sourceExt)
}

func (pm *propertyModule) AddProps(mapProp map[int32]int64, source int32, sourceExt string) {
	if len(mapProp) == 0 {
		return
	}

	pMsgProp := new(msg.S2C_SYNC_NUM)
	//arrEvent := make([]*player.EventPropertyIncrease, 0, len(mapProp))

	for id, val := range mapProp {
		if val <= 0 {
			continue
		}
		p := pm.GetProp(id)
		if p == nil {
			p = newProperty(pm.player, id)
			pm.mapProp[id] = p
		}
		p.addValue(val, source, sourceExt)
		pProp := new(msg.NUM)
		pProp.NumType = id
		pProp.Num = p.GetValue()
		pProp.Data1 = p.GetUpdateTs()
		pMsgProp.Nums = append(pMsgProp.Nums, pProp)

		//arrEvent = append(arrEvent, &player.EventPropertyIncrease{Id: id, Delta: val, CurNum: p.GetValue()})
	}

	pm.player.SendMessage(pMsgProp)

	//for _, pEvent := range arrEvent {
	//	pm.Player.GetEventManager().Dispatch(pEvent)
	//}
}

func (pm *propertyModule) IsEnoughProps(mapProp map[int32]int64) bool {
	for id, val := range mapProp {
		if val <= 0 {
			return false
		}
		p := pm.GetProp(id)
		if p == nil {
			return false
		}
		if p.GetValue() < val {
			return false
		}
	}

	return true
}

func (pm *propertyModule) SubProps(mapProp map[int32]int64, source int32, sourceExt string) bool {
	if !pm.IsEnoughProps(mapProp) {
		return false
	}

	pMsgProp := new(msg.S2C_SYNC_NUM)

	for id, val := range mapProp {
		p := pm.GetProp(id)
		p.subValue(val, source, sourceExt)
		pProp := new(msg.NUM)
		pProp.NumType = id
		pProp.Num = p.GetValue()
		pProp.Data1 = p.GetUpdateTs()
		pMsgProp.Nums = append(pMsgProp.Nums, pProp)
	}

	pm.player.SendMessage(pMsgProp)
	return true
}

func (pm *propertyModule) GetPropsNum() map[int32]int64 {
	ret := make(map[int32]int64, len(pm.mapProp))
	for k, v := range pm.mapProp {
		ret[k] = v.GetValue()
	}
	return ret
}
