package player

import (
	"database/sql"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iplayer "vgproj/vggame/public/game/player"
	igate "vgproj/vggame/public/gate"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
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
	playerId  int64
	arrModule [iplayer.PlayerModule_Count]iplayer.IPlayerModule
	status    EPlayerStatus
	conn      igate.IConnection
	arrCb     []*loadContext
	loadCb    *loadContext
}

func newPlayer(playerId int64) *Player {
	pObj := new(Player)
	pObj.playerId = playerId
	pObj.status = EPlayerStatus_Loading
	pObj.arrCb = make([]*loadContext, 0, 8)

	// pObj.arrModule[player.PlayerModule_Data] = data.NewDataModule(pObj)
	// pObj.arrModule[player.PlayerModule_Property] = property.NewPropertyModule(pObj)
	// pObj.arrModule[player.PlayerModule_Item] = item.NewItemModule(pObj)
	// pObj.arrModule[player.PlayerModule_Hero] = hero.NewHeroModule(pObj)
	// pObj.arrModule[player.PlayerModule_Mail] = mail.NewMailModule(pObj)
	// pObj.arrModule[player.PlayerModule_Vip] = vip.NewVipModule(pObj)
	// pObj.arrModule[player.PlayerModule_Settings] = settings.NewSettingsModule(pObj)

	return pObj
}

func (p *Player) insert(accountId int64, serverId int32, name string, head int32, ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	p.loadCb = &loadContext{ctx: ctx, cb: cb}
	public.Server.GetDataDb().AsyncExec(p.insertCallback, uint32(p.playerId), "insert into player_data(player_id,account_id,server_id,name,head,create_ts) values(?,?,?,?,?,?)",
		p.playerId, accountId, serverId, name, head, vgtime.Now())
}

func (p *Player) insertCallback(args []interface{}) {
	if args[1] != nil {
		p.returnWaiting(args[1].(error))
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
		sqlCollect += pModele.GetLoadSql()
	}

	public.Server.GetDataDb().AsyncQuery(p.initForward, uint32(p.playerId), sqlCollect)
}

func (p *Player) initForward(args []interface{}) {
	var rows *sql.Rows
	var err error
	if args[0] != nil {
		rows = args[0].(*sql.Rows)
	}
	defer rows.Close()
	if args[1] != nil {
		err = args[1].(error)
	}

	if err != nil {
		p.returnWaiting(err)
		return
	}

	for _, pModule := range p.arrModule {
		err = pModule.OnLoadData(rows)
		if err != nil {
			p.returnWaiting(err)
			return
		}
		rows.NextResultSet()
	}

	for _, pModule := range p.arrModule {
		pModule.OnInit1()
	}
	for _, pModule := range p.arrModule {
		pModule.OnInit2()
	}
	for _, pModule := range p.arrModule {
		pModule.OnInit3()
	}
	for _, pModule := range p.arrModule {
		pModule.OnInit4()
	}
	for _, pModule := range p.arrModule {
		pModule.OnInit5()
	}
	p.status = EPlayerStatus_Offline

	p.returnWaiting(nil)

	return
}

func (p *Player) createInit() {
	for _, pModule := range p.arrModule {
		pModule.OnCreate()
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
		pModule.OnLogin()
	}
	pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
	pPlayerManager.SetPlayerOnline(p.playerId)

	p.SendMessage(&msg.S2C_LoginFinish{})
}

func (p *Player) Logout() {
	if p.status != EPlayerStatus_Online {
		return
	}
	p.status = EPlayerStatus_Offline
	for _, pModule := range p.arrModule {
		pModule.OnLogout()
	}
	p.conn = nil
	pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
	pPlayerManager.SetPlayerOffline(p.playerId)
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
		pModule.OnRelease()
	}
}

func (p *Player) GetId() int64 {
	return p.playerId
}

func (p *Player) DailyRefresh(refreshTs int64) {
	// pData := p.GetModule(player.PlayerModule_Data).(*data.DataModule)
	// pData.DailyRefresh(refreshTs)
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
