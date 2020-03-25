package player

import (
	"database/sql"
	"fmt"
	ec "vgproj/common/define/err_code"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"

	logger "github.com/panlibin/vglog"
)

type vipModule struct {
	*playerModule

	vipLev int32
	vipExp int32
	inDb   bool

	mapVipGiftStatus map[int32]int32
}

func newVipModule(pPlayer *Player) *vipModule {
	pObj := new(vipModule)
	pObj.playerModule = newPlayerModule(pPlayer)
	pObj.mapVipGiftStatus = make(map[int32]int32, 16)

	return pObj
}

func (vm *vipModule) getLoadSql() string {
	return fmt.Sprintf("select vip_lev,vip_exp from player_vip where player_id=%d;select vip_lev from player_vip_gift where player_id=%d;",
		vm.player.GetId(), vm.player.GetId())
}

func (vm *vipModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_vip error: %v", err)
		return err
	}
	if !rows.Next() {
		return nil
	}
	err = rows.Scan(&vm.vipLev, &vm.vipExp)
	if err != nil {
		logger.Errorf("load player_vip error: %v", err)
		return err
	} else {
		vm.inDb = true
	}

	rows.NextResultSet()
	err = rows.Err()
	if err != nil {
		logger.Errorf("load player_vip_gift error: %v", err)
		return err
	}
	var tmpVipLev int32
	for rows.Next() {
		err = rows.Scan(&tmpVipLev)
		if err != nil {
			break
		}
		vm.mapVipGiftStatus[tmpVipLev] = 1
	}

	if err != nil {
		logger.Errorf("load player_vip_gift error: %v", err)
	}

	return err
}

func (vm *vipModule) GetVipLev() int32 {
	return vm.vipLev
}

func (vm *vipModule) GetVipExp() int32 {
	return vm.vipExp
}

func (vm *vipModule) AddVipLev(lev int32) {
	if lev <= 0 {
		return
	}
	beforeLev := vm.vipLev
	vm.vipLev += lev
	pVipConfMgr := public.Server.GetConfigManager().GetConfig(iconfig.Config_Vip).(iconfig.IVipConfig)
	if vm.vipLev > pVipConfMgr.GetMaxLevel() {
		vm.vipLev = pVipConfMgr.GetMaxLevel()
	}
	deltaLev := vm.vipLev - beforeLev
	if deltaLev > 0 {
		vm.updateVipLevExp()
		vm.syncVipLevExp()
		vm.player.GetEventManager().Dispatch(&EventPlayerVipLevUp{CurVipLev: vm.vipLev, DeltaVipLev: deltaLev})
	}
}

func (vm *vipModule) AddVipExp(exp int32) {
	if exp <= 0 {
		return
	}
	vm.vipExp += exp

	var deltaLev int32 = 0
	pVipConfMgr := public.Server.GetConfigManager().GetConfig(iconfig.Config_Vip).(iconfig.IVipConfig)
	for {
		pNextVipConf := pVipConfMgr.GetConf(vm.vipLev + 1)
		if pNextVipConf == nil || vm.vipExp < int32(pNextVipConf.Exp) {
			break
		}
		vm.vipLev++
		deltaLev++
	}

	vm.updateVipLevExp()
	vm.syncVipLevExp()
	if deltaLev > 0 {
		vm.player.GetEventManager().Dispatch(&EventPlayerVipLevUp{CurVipLev: vm.vipLev, DeltaVipLev: deltaLev})
	}
}

func (vm *vipModule) syncVipLevExp() {
	pMsgSync := new(msg.S2C_VIP_NOTIFY)
	pMsgSync.VipLev = vm.vipLev
	pMsgSync.VipExp = vm.vipExp
	vm.player.SendMessage(pMsgSync)
}

func (vm *vipModule) insert() {
	vm.inDb = true
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(vm.player.GetId()), sqlInsert, vm.player.GetId(), vm.vipLev, vm.vipExp)
}

func (vm *vipModule) updateVipLevExp() {
	if vm.inDb {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(vm.player.GetId()), sqlUpdate, vm.vipLev, vm.vipExp, vm.player.GetId())
	} else {
		vm.insert()
	}
}

func (vm *vipModule) insertGiftStatus(vipLev int32) {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(vm.player.GetId()), sqlInsertVipGift, vm.player.GetId(), vipLev)
}

func (vm *vipModule) handleGetVipGiftStatus(args []interface{}) {
	//pMsgReq := args[0].(*msg.C2S_VIP_GET_GIFT_STATUS)
	pMsgRsp := new(msg.S2C_VIP_GET_GIFT_STATUS)
	pMsgRsp.Status = make([]*msg.VIP_GIFT_STATUS, 0, len(vm.mapVipGiftStatus))
	for k, _ := range vm.mapVipGiftStatus {
		pMsgRsp.Status = append(pMsgRsp.Status, &msg.VIP_GIFT_STATUS{VipLev: k, GetReward: 1})
	}
	pMsgRsp.Code = ec.Success
	vm.player.SendMessage(pMsgRsp)
}

func (vm *vipModule) handleCollectVipGift(args []interface{}) {
	pMsgReq := args[0].(*msg.C2S_VIP_COLLECT_GIFT)
	pMsgRsp := new(msg.S2C_VIP_COLLECT_GIFT)

	var errCode int32 = ec.Unknown
	pMsgRsp.GiftLev = pMsgReq.GiftLev
	for {
		if vm.vipLev < pMsgReq.GiftLev {
			errCode = ec.ErrVipLevNotEnough
			break
		}

		if _, exist := vm.mapVipGiftStatus[pMsgReq.GiftLev]; exist {
			errCode = ec.ErrVipGiftAlreadyGot
			break
		}

		pVipConfMgr := public.Server.GetConfigManager().GetConfig(iconfig.Config_Vip).(iconfig.IVipConfig)
		pVipConf := pVipConfMgr.GetConf(vm.vipLev)
		if pVipConf == nil {
			errCode = ec.ConfigError
			break
		}

		vm.mapVipGiftStatus[pMsgReq.GiftLev] = 1
		vm.insertGiftStatus(pMsgReq.GiftLev)

		if pVipConf.HeroId > 0 {

		}
		if pVipConf.BeautyId > 0 {

		}
		if len(pVipConf.Gift) > 0 {
			// vm.player.AddThings(pVipConf.Gift, action.VipGift, "")
		}

		errCode = ec.Success
		break
	}

	pMsgRsp.Code = errCode
	vm.player.SendMessage(pMsgRsp)
}
