package player

import (
	"database/sql"
	"fmt"
	"vgproj/proto/msg"

	logger "github.com/panlibin/vglog"
)

type itemModule struct {
	*playerModule

	mapItem map[int32]*Item
}

func newItemModule(pPlayer *Player) *itemModule {
	pObj := new(itemModule)
	pObj.playerModule = newPlayerModule(pPlayer)

	pObj.mapItem = make(map[int32]*Item, 64)

	return pObj
}

func (im *itemModule) getLoadSql() string {
	return fmt.Sprintf("select item_id,item_num from player_item where player_id=%d;", im.player.GetId())
}

func (im *itemModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_item error: %v", err)
		return err
	}

	var tmpId int32
	var tmpNum int64
	for rows.Next() {
		err = rows.Scan(&tmpId, &tmpNum)
		if err != nil {
			break
		}
		pItem, exist := im.mapItem[tmpId]
		if !exist || pItem == nil {
			pItem = newItem(im.player, tmpId)
			im.mapItem[tmpId] = pItem
		}
		pItem.Load(tmpNum)
	}

	if err != nil {
		logger.Errorf("load player_item error: %v", err)
	}

	return err
}

func (im *itemModule) onCreate() {

}

func (im *itemModule) GetItem(id int32) *Item {
	pItem, exist := im.mapItem[id]
	if !exist {
		return nil
	}
	return pItem
}

func (im *itemModule) AddItem(id int32, num int64, source int32, sourceExt string) {
	pItem := im.GetItem(id)
	if pItem == nil {
		pItem = newItem(im.player, id)
		im.mapItem[id] = pItem
	}
	pItem.AddNum(num, source, sourceExt)
}

func (im *itemModule) SubItem(id int32, val int64, source int32, sourceExt string) bool {
	pItem := im.GetItem(id)
	if pItem == nil {
		return false
	}
	return pItem.SubNum(val, source, sourceExt)
}

func (im *itemModule) AddItems(mapItems map[int32]int64, source int32, sourceExt string) {
	if len(mapItems) == 0 {
		return
	}

	pMsgItems := new(msg.S2C_SYNC_ITEMS)

	for id, num := range mapItems {
		if num <= 0 {
			continue
		}
		pItem := im.GetItem(id)
		if pItem == nil {
			pItem = newItem(im.player, id)
			im.mapItem[id] = pItem
		}
		pItem.addNum(num, source, sourceExt)

		pMsgItems.Items = append(pMsgItems.Items, pItem.formatToMessage())
	}

	im.player.SendMessage(pMsgItems)
}

func (im *itemModule) IsEnoughItems(mapItems map[int32]int64) bool {
	for id, num := range mapItems {
		if num <= 0 {
			return false
		}
		iItem := im.GetItem(id)
		if iItem == nil {
			return false
		}
		if iItem.GetNum() < num {
			return false
		}
	}

	return true
}

func (im *itemModule) SubItems(mapItems map[int32]int64, source int32, sourceExt string) bool {
	if !im.IsEnoughItems(mapItems) {
		return false
	}

	pMsgItems := new(msg.S2C_SYNC_ITEMS)
	pMsgItemDel := new(msg.S2C_SYNC_ITEMS_DEL)

	for id, num := range mapItems {
		pItem := im.GetItem(id)
		pItem.subNum(num, source, sourceExt)

		if pItem.GetNum() != 0 {
			pMsgItems.Items = append(pMsgItems.Items, pItem.formatToMessage())
		} else {
			pMsgItemDel.DelIds = append(pMsgItemDel.DelIds, pItem.GetId())
		}
	}

	if len(pMsgItems.Items) > 0 {
		im.player.SendMessage(pMsgItems)
	}
	if len(pMsgItemDel.DelIds) > 0 {
		im.player.SendMessage(pMsgItemDel)
	}

	return true
}

func (im *itemModule) handleGetRoleItems(pMsgReq *msg.C2S_ROLE_ITEMS) {
	pMsgRsp := new(msg.S2C_ROLE_ITEMS)

	pMsgRsp.Items = make([]*msg.ROLE_ITEM, 0, len(im.mapItem))
	for _, pItem := range im.mapItem {
		pMsgRsp.Items = append(pMsgRsp.Items, pItem.formatToMessage())
	}

	im.player.SendMessage(pMsgRsp)
}

func (im *itemModule) handleUseItem(args []interface{}) {
	//pMsgReq := args[0].(*msg.C2S_USE_PROP)
	//pMsgRsp := new(msg.S2C_USE_PROP)
	//
	//var errCode int32 = ec.Unknown
	//for {
	//	pLevConfMgr := public.Server.GetConfigManager().GetConfig(config.Config_Level).(config.ILevelConfig)
	//	if im.lev >= pLevConfMgr.GetMaxLevel() {
	//		errCode = ec.ErrRoleMaxLv
	//		break
	//	}
	//
	//	pLevConf := pLevConfMgr.GetLevConf(im.lev)
	//	if pLevConf == nil {
	//		errCode = ec.ConfigError
	//		break
	//	}
	//
	//	if im.exp < pLevConf.Exp {
	//		errCode = ec.ErrHeroNotEnoughRoleExp
	//		break
	//	}
	//
	//	im.exp -= pLevConf.Exp
	//	oldLv := im.lev
	//	im.lev++
	//	im.updateLevAndExp()
	//
	//	public.Server.GetOaWriter().Write("log_item", im.Player.GetId(), action.RoleLevelUp, "", item.ITEM_ST_CURRENCY_EXP, -pLevConf.Exp, im.exp, time.Now())
	//	public.Server.GetOaWriter().Write("log_lev_up", im.Player.GetId(), oldLv, im.lev, time.Now())
	//
	//	errCode = ec.Success
	//
	//	break
	//}
	//
	//if errCode == ec.Success {
	//	im.Player.SendMessage(&msg.S2C_ROLE_LVUP_RET{
	//		Lv: im.GetLevel(),
	//	})
	//} else {
	//	im.Player.SendMessage(&msg.S2C_ERROR{
	//		Code: errCode,
	//	})
	//}
}
