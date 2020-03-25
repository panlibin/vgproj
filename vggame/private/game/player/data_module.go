package player

import (
	"database/sql"
	"fmt"
	"time"
	ec "vgproj/common/define/err_code"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"
	"vgproj/vggame/public/define/action"
	"vgproj/vggame/public/define/item"
	iplayer "vgproj/vggame/public/game/player"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgtime"
)

type dataModule struct {
	*playerModule

	accountId            int64
	serverId             int32
	name                 string
	head                 int32
	sex                  int32
	lev                  int32
	exp                  int64
	createTime           int64
	lastLoginIp          string
	lastLoginTime        int64
	lastLogoutTime       int64
	lastDailyRefreshTime int64
}

func newDataModule(pPlayer *Player) *dataModule {
	pObj := new(dataModule)
	pObj.playerModule = newPlayerModule(pPlayer)

	return pObj
}

func (dm *dataModule) getLoadSql() string {
	return fmt.Sprintf("select account_id,server_id,`name`,`head`,sex,lev,`exp`,create_ts,last_login_ip,last_login_ts,last_logout_ts,last_daily_refresh_ts "+
		"from player_data where player_id=%d;", dm.player.GetId())
}

func (dm *dataModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_data error: %v", err)
		return err
	}
	if !rows.Next() {
		logger.Errorf("load player_data error: %v", sql.ErrNoRows)
		return sql.ErrNoRows
	}

	err = rows.Scan(&dm.accountId, &dm.serverId, &dm.name, &dm.head, &dm.sex, &dm.lev, &dm.exp, &dm.createTime, &dm.lastLoginIp,
		&dm.lastLoginTime, &dm.lastLogoutTime, &dm.lastDailyRefreshTime)
	if err != nil {
		logger.Errorf("load player_data error: %v", err)
	}
	return err
}

func (dm *dataModule) onCreate() {
	public.Server.GetOaWriter().Write("log_create_character", dm.player.GetId(), dm.accountId, dm.serverId, dm.name, dm.head, dm.sex, time.Now())
}

func (dm *dataModule) onLogin() {
	dm.lastLoginIp = dm.player.GetIpAddr()
	dm.lastLoginTime = vgtime.Now()
	dm.lastLogoutTime = dm.lastLoginTime
	dm.updateLogin()

	public.Server.GetOaWriter().Write("log_player_login", dm.player.GetId(), dm.accountId, dm.lastLoginIp, time.Now())
}

func (dm *dataModule) onLogout() {
	dm.lastLogoutTime = vgtime.Now()
	dm.updateLogout()

	public.Server.GetOaWriter().Write("log_player_logout", dm.player.GetId(), dm.accountId, dm.lastLogoutTime-dm.lastLoginTime, time.Now())
}

func (dm *dataModule) GetName() string {
	return dm.name
}

func (dm *dataModule) GetHead() int32 {
	return dm.head
}

func (dm *dataModule) GetLevel() int32 {
	return dm.lev
}

func (dm *dataModule) GetExp() int64 {
	return dm.exp
}

func (dm *dataModule) GetCreateTs() int64 {
	return dm.createTime
}

func (dm *dataModule) GetLastLoginIp() string {
	return dm.lastLoginIp
}

func (dm *dataModule) GetLastLoginTs() int64 {
	return dm.lastLoginTime
}

func (dm *dataModule) GetLastLogoutTs() int64 {
	return dm.lastLogoutTime
}

func (dm *dataModule) SetName(name string) {
	if dm.name == name {
		return
	}
	dm.name = name
	dm.updateName()
}

func (dm *dataModule) SetHead(head int32) {
	if dm.head == head {
		return
	}
	dm.head = head
	dm.updateHead()
}

func (dm *dataModule) AddExp(exp int64, source int32, sourceExt string) {
	dm.exp += exp
	dm.updateLevAndExp()

	public.Server.GetOaWriter().Write("log_item", dm.player.GetId(), source, sourceExt, item.ITEM_ST_CURRENCY_EXP, exp, dm.exp, time.Now())

	syncNum := msg.S2C_SYNC_NUM{}
	syncNum.Nums = make([]*msg.NUM, 1)
	syncNum.Nums[0] = &msg.NUM{NumType: item.ITEM_ST_CURRENCY_EXP, Num: dm.exp}
	dm.player.SendMessage(&syncNum)
}

func (dm *dataModule) DailyRefresh(ts int64) {
	if dm.lastDailyRefreshTime >= ts {
		return
	}
	dm.lastDailyRefreshTime = ts
	dm.updateDailyRefresh()
	dm.player.GetEventManager().Dispatch(&EventDailyRefresh{RefreshTs: ts})

	public.Server.GetOaWriter().Write("log_player_login", dm.player.GetId(), dm.accountId, dm.lastLoginIp, time.Now())
}

func (dm *dataModule) updateLogin() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateLogin, dm.lastLoginIp, dm.lastLoginTime, dm.lastLogoutTime, dm.player.GetId())
}

func (dm *dataModule) updateLogout() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateLogout, dm.lastLogoutTime, dm.player.GetId())
}

func (dm *dataModule) updateDailyRefresh() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateDailyRefresh, dm.lastDailyRefreshTime, dm.player.GetId())
}

func (dm *dataModule) updateName() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateName, dm.name, dm.player.GetId())
}

func (dm *dataModule) updateHead() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateHead, dm.head, dm.player.GetId())
}

func (dm *dataModule) updateLevAndExp() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(dm.player.GetId()), sqlUpdateLevExp, dm.lev, dm.exp, dm.player.GetId())
}

func (dm *dataModule) handleGetRoleInfo(args []interface{}) {
	//pMsgReq := args[0].(*msg.C2S_ROLE_INFO)
	pMsgRsp := new(msg.S2C_ROLE_INFO)

	pMsgRsp.Uid = dm.accountId
	pMsgRsp.Rid = dm.player.GetId()
	pMsgRsp.Name = dm.name
	pMsgRsp.ServId = dm.serverId
	pMsgRsp.MemberId = dm.head
	pMsgRsp.Lv = dm.lev
	pMsgRsp.Exp = dm.exp

	pVipModule := dm.player.GetModule(iplayer.PlayerModule_Vip).(*vipModule)
	pMsgRsp.VipLv = pVipModule.GetVipLev()
	pMsgRsp.VipExp = pVipModule.GetVipExp()

	pPropModule := dm.player.GetModule(iplayer.PlayerModule_Property).(*propertyModule)
	pMsgRsp.Res = pPropModule.GetPropsNum()

	dm.player.SendMessage(pMsgRsp)
}

func (dm *dataModule) handleLevelUpReq(args []interface{}) {
	//pMsgReq := args[0].(*msg.C2S_ROLE_LVUP_REQ)

	var errCode int32 = ec.Unknown
	for {
		pLevConfMgr := public.Server.GetConfigManager().GetConfig(iconfig.Config_Level).(iconfig.ILevelConfig)
		if dm.lev >= pLevConfMgr.GetMaxLevel() {
			errCode = ec.ErrRoleMaxLv
			break
		}

		pLevConf := pLevConfMgr.GetLevConf(dm.lev)
		if pLevConf == nil {
			errCode = ec.ConfigError
			break
		}

		if dm.exp < pLevConf.Exp {
			errCode = ec.ErrHeroNotEnoughRoleExp
			break
		}

		dm.exp -= pLevConf.Exp
		oldLv := dm.lev
		dm.lev++
		dm.updateLevAndExp()

		public.Server.GetOaWriter().Write("log_item", dm.player.GetId(), action.RoleLevelUp, "", item.ITEM_ST_CURRENCY_EXP, -pLevConf.Exp, dm.exp, time.Now())
		public.Server.GetOaWriter().Write("log_lev_up", dm.player.GetId(), oldLv, dm.lev, time.Now())

		errCode = ec.Success

		break
	}

	if errCode == ec.Success {
		dm.player.SendMessage(&msg.S2C_ROLE_LVUP_RET{
			Lv: dm.GetLevel(),
		})
	} else {
		dm.player.SendMessage(&msg.S2C_ERROR{
			Code: errCode,
		})
	}
}
