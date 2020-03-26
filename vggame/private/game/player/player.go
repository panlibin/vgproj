package player

import (
	"context"
	"database/sql"
	"vgproj/common/cluster"
	"vgproj/proto/loginrpc"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"
	"vgproj/vggame/public/define/item"
	iplayer "vgproj/vggame/public/game/player"
	igate "vgproj/vggame/public/gate"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgevent"
	"github.com/panlibin/virgo/util/vgtime"
)

type EPlayerStatus int32

const (
	EPlayerStatus_Loading EPlayerStatus = iota
	EPlayerStatus_Offline
	EPlayerStatus_Online
	EPlayerStatus_Release
)

type loadContext struct {
	cb  func(interface{}, iplayer.IPlayer)
	ctx interface{}
}

type Player struct {
	playerId      int64
	rnd           uint64
	arrModule     [iplayer.PlayerModule_Count]iPlayerModule
	status        EPlayerStatus
	conn          igate.IConnection
	arrCb         []*loadContext
	loadCb        *loadContext
	pEventManager *vgevent.EventManager
}

func newPlayer(playerId int64) *Player {
	pObj := new(Player)
	pObj.playerId = playerId
	pObj.status = EPlayerStatus_Loading
	pObj.arrCb = make([]*loadContext, 0, 8)
	pObj.pEventManager = vgevent.NewEventManager()

	pObj.arrModule[iplayer.PlayerModule_Data] = newDataModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Property] = newPropertyModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Item] = newItemModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Hero] = newHeroModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Mail] = newMailModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Vip] = newVipModule(pObj)
	pObj.arrModule[iplayer.PlayerModule_Settings] = newSettingsModule(pObj)

	return pObj
}

func (p *Player) insert(accountId int64, serverId int32, name string, head int32, ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	p.loadCb = &loadContext{ctx: ctx, cb: cb}
	public.Server.GetDataDb().AsyncExec(nil, p.insertCallback, uint32(p.playerId), "insert into player_data(player_id,account_id,server_id,name,head,sex,lev,`exp`,create_ts)"+
		" values(?,?,?,?,?,0,1,0,?)", p.playerId, accountId, serverId, name, head, vgtime.Now())
}

func (p *Player) insertCallback(args []interface{}) {
	if args[2] != nil {
		p.returnWaiting(args[2].(error))
	} else {
		cbCtx := p.loadCb
		p.loadCb = nil
		cbCtx.cb(cbCtx.ctx, p)
	}
}

func (p *Player) init(ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	p.loadCb = &loadContext{ctx: ctx, cb: cb}
	var sqlCollect string
	for _, pModele := range p.arrModule {
		if len(sqlCollect) > 0 && sqlCollect[len(sqlCollect)-1] != ';' {
			sqlCollect += ";"
		}
		sqlCollect += pModele.getLoadSql()
	}

	public.Server.GetDataDb().AsyncQuery(nil, p.initForward, uint32(p.playerId), sqlCollect)
}

func (p *Player) initForward(args []interface{}) {
	var rows *sql.Rows
	var err error
	if args[1] != nil {
		rows = args[1].(*sql.Rows)
	}
	if args[2] != nil {
		err = args[2].(error)
	}

	if err != nil {
		p.returnWaiting(err)
		return
	}
	defer rows.Close()

	for _, pModule := range p.arrModule {
		err = pModule.onLoadData(rows)
		if err != nil {
			p.returnWaiting(err)
			return
		}
		rows.NextResultSet()
	}

	for _, pModule := range p.arrModule {
		pModule.onInit1()
	}
	for _, pModule := range p.arrModule {
		pModule.onInit2()
	}
	for _, pModule := range p.arrModule {
		pModule.onInit3()
	}
	for _, pModule := range p.arrModule {
		pModule.onInit4()
	}
	for _, pModule := range p.arrModule {
		pModule.onInit5()
	}
	p.status = EPlayerStatus_Offline

	p.returnWaiting(nil)

	return
}

func (p *Player) createInit() {
	for _, pModule := range p.arrModule {
		pModule.onCreate()
	}
}

func (p *Player) Login(conn igate.IConnection) {
	if p.status != EPlayerStatus_Offline {
		return
	}
	p.DailyRefresh(vgtime.GetDayZeroTs(vgtime.Now()))
	p.status = EPlayerStatus_Online
	p.conn = conn
	for _, pModule := range p.arrModule {
		pModule.onLogin()
	}
	pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
	pPlayerManager.SetPlayerOnline(p.playerId)

	logger.Debugf("player login id %d", p.playerId)
	p.SendMessage(&msg.S2C_LoginFinish{})
}

func (p *Player) Logout() {
	if p.status != EPlayerStatus_Online {
		return
	}
	p.status = EPlayerStatus_Offline
	for _, pModule := range p.arrModule {
		pModule.onLogout()
	}

	pDataModule := p.GetModule(iplayer.PlayerModule_Data).(*dataModule)
	accountId := pDataModule.accountId
	serverId := pDataModule.serverId
	playerId := p.playerId
	name := pDataModule.name

	p.conn = nil
	pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
	pPlayerManager.SetPlayerOffline(p.playerId)
	logger.Debugf("player logout id %d", p.playerId)

	public.Server.AsyncTask(func([]interface{}) {
		pNode := public.Server.GetCluster().GetNode(cluster.NodeLogin, 1)
		if pNode != nil {
			pLoginNode := pNode.(*cluster.LoginNode)
			pLoginNode.PlayerLogout(context.Background(), &loginrpc.NotifyLogout{
				AccountId: accountId,
				ServerId:  serverId,
				PlayerId:  playerId,
				Name:      name,
				Combat:    0,
			})
		}
	})
}

func (p *Player) Release() {
	if p.status == EPlayerStatus_Online {
		p.Logout()
	}
	if p.status != EPlayerStatus_Offline {
		return
	}
	p.status = EPlayerStatus_Release
	for i := iplayer.PlayerModule_Count - 1; i >= 0; i-- {
		pModule := p.arrModule[i]
		pModule.onRelease()
	}
}

func (p *Player) GetId() int64 {
	return p.playerId
}

func (p *Player) DailyRefresh(refreshTs int64) {
	pData := p.GetModule(iplayer.PlayerModule_Data).(*dataModule)
	pData.DailyRefresh(refreshTs)
}

func (p *Player) SendMessage(msg proto.Message) {
	if p.conn != nil {
		p.conn.Write(msg)
	}
}

func (p *Player) GetModule(id int32) interface{} {
	return p.arrModule[id]
}

func (p *Player) returnWaiting(err error) {
	var retPlayer iplayer.IPlayer
	if err == nil {
		retPlayer = p
	} else {
		logger.Errorf("load player error %v, %d", err, p.playerId)
	}
	cbCtx := p.loadCb
	p.loadCb = nil
	if cbCtx != nil {
		cbCtx.cb(cbCtx.ctx, retPlayer)
	}
	arr := p.arrCb
	p.arrCb = make([]*loadContext, 0)
	for _, cbCtx := range arr {
		cbCtx.cb(cbCtx.ctx, retPlayer)
	}
}

func (p *Player) waitLoadFinish(ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	p.arrCb = append(p.arrCb, &loadContext{ctx: ctx, cb: cb})
}

func (p *Player) GetIpAddr() string {
	if p.conn == nil {
		return ""
	}
	return p.conn.RemoteAddr().String()
}

func (p *Player) GetEventManager() *vgevent.EventManager {
	return p.pEventManager
}

func (p *Player) splitThings(mapThings map[int32]int64) (mapAttr map[int32]int64, mapProp map[int32]int64, mapItem map[int32]int64, mapTitle map[int32]int64) {
	mapAttr = make(map[int32]int64)
	mapProp = make(map[int32]int64)
	mapItem = make(map[int32]int64)

	pItemConfMgr := public.Server.GetConfigManager().GetConfig(iconfig.Config_Item).(iconfig.IItemConfig)
	for k, v := range mapThings {
		if item.IsTitle(k) {
			mapTitle[k] = 1
		} else {
			pItemConf := pItemConfMgr.GetConf(k)
			if pItemConf == nil {
				logger.Warningf("unknown item id %d", k)
				continue
			}
			if pItemConf.FirstType == item.ITEM_FT_CURRENCY {
				mapProp[k] = v
			} else if pItemConf.FirstType == item.ITEM_FT_EXP {
				mapAttr[k] = v
			} else {
				mapItem[k] = v
			}
		}
	}

	return
}

func (p *Player) AddThings(mapThings map[int32]int64, source int32, sourceExt string) {
	mapAttr, mapProp, mapItem, mapTitle := p.splitThings(mapThings)

	if len(mapProp) > 0 {
		pPropModule := p.GetModule(iplayer.PlayerModule_Property).(*propertyModule)
		pPropModule.AddProps(mapProp, source, sourceExt)
	}

	if len(mapItem) > 0 {
		pItemModule := p.GetModule(iplayer.PlayerModule_Item).(*itemModule)
		pItemModule.AddItems(mapItem, source, sourceExt)
	}

	if len(mapAttr) > 0 {
		for k, v := range mapAttr {
			switch k {
			case item.ITEM_ST_CURRENCY_EXP:
				pDataModule := p.GetModule(iplayer.PlayerModule_Data).(*dataModule)
				pDataModule.AddExp(v, source, sourceExt)
			case item.ITEM_ST_CURRENCY_VIP_LV:
				pVipModule := p.GetModule(iplayer.PlayerModule_Vip).(*vipModule)
				pVipModule.AddVipLev(int32(v))
			case item.ITEM_ST_CURRENCY_VIP_EXP:
				pVipModule := p.GetModule(iplayer.PlayerModule_Vip).(*vipModule)
				pVipModule.AddVipExp(int32(v))
			}
		}
	}

	if len(mapTitle) > 0 {
		for k, v := range mapTitle {
			logger.Debug(k, v)
		}
	}

	return
}

func (p *Player) IsEnoughThings(mapConsume map[int32]int64) bool {
	_, mapProps, mapItems, _ := p.splitThings(mapConsume)
	pPropModule := p.GetModule(iplayer.PlayerModule_Property).(*propertyModule)
	pItemModule := p.GetModule(iplayer.PlayerModule_Item).(*itemModule)

	if len(mapItems) > 0 {
		if !pItemModule.IsEnoughItems(mapItems) {
			return false
		}
	}
	if len(mapProps) > 0 {
		if !pPropModule.IsEnoughProps(mapProps) {
			return false
		}
	}

	return true
}

func (p *Player) ConsumeThings(mapConsume map[int32]int64, source int32, sourceExt string) bool {
	_, mapProps, mapItems, _ := p.splitThings(mapConsume)

	pPropModule := p.GetModule(iplayer.PlayerModule_Property).(*propertyModule)
	pItemModule := p.GetModule(iplayer.PlayerModule_Item).(*itemModule)

	if len(mapItems) > 0 {
		if !pItemModule.IsEnoughItems(mapItems) {
			return false
		}
	}
	if len(mapProps) > 0 {
		if !pPropModule.IsEnoughProps(mapProps) {
			return false
		}
	}

	if len(mapItems) > 0 {
		pItemModule.SubItems(mapItems, source, sourceExt)
	}
	if len(mapProps) > 0 {
		pPropModule.SubProps(mapProps, source, sourceExt)
	}

	return true
}
