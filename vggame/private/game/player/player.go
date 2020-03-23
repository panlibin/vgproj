package player

import (
	"database/sql"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/golang/protobuf/proto"
	"github.com/panlibin/virgo/util/vgtime"
)

type EPlayerStatus int32

const (
	EPlayerStatus_Loading EPlayerStatus = iota
	EPlayerStatus_Offline
	EPlayerStatus_Online
	EPlayerStatus_Release
)

type Player struct {
	playerId  int64
	arrModule [iplayer.PlayerModule_Count]iplayer.IPlayerModule
	status    EPlayerStatus
	conn      igate.IConnection
	cb        func(*Player, error)
}

func NewPlayer(playerId int64) *Player {
	pObj := new(Player)
	pObj.playerId = playerId
	pObj.status = EPlayerStatus_Loading

	// pObj.arrModule[player.PlayerModule_Data] = data.NewDataModule(pObj)
	// pObj.arrModule[player.PlayerModule_Property] = property.NewPropertyModule(pObj)
	// pObj.arrModule[player.PlayerModule_Item] = item.NewItemModule(pObj)
	// pObj.arrModule[player.PlayerModule_Hero] = hero.NewHeroModule(pObj)
	// pObj.arrModule[player.PlayerModule_Mail] = mail.NewMailModule(pObj)
	// pObj.arrModule[player.PlayerModule_Vip] = vip.NewVipModule(pObj)
	// pObj.arrModule[player.PlayerModule_Settings] = settings.NewSettingsModule(pObj)

	return pObj
}

func (p *Player) Init(cb func(*Player, error)) {
	p.cb = cb
	var sqlCollect string
	for _, pModele := range p.arrModule {
		if len(sqlCollect) > 0 && sqlCollect[len(sqlCollect)-1] != ";" {
			sqlCollect += ";"
		}
		sqlCollect += pModele.GetLoadSql()
	}

	public.Server.GetDataDb().AsyncQuery(p.InitForward, uint32(p.playerId), sqlCollect)
}

func (p *Player) InitForward(args []interface{}) {
	var rows *sql.Rows
	var err error
	if args[0] != nil {
		rows = args[0].(*sql.Rows)
	}
	if args[1] != nil {
		err = args[1].(error)
	}

	if err != nil {

	}

	for _, pModule := range p.arrModule {
		err = pModule.OnLoadData(rows)
		if err != nil {
			p.returnWaiting(nil)
			return err
		}
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

	p.returnWaiting(p)

	return nil
}

func (p *Player) CreateInit() {
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
	for i := player.PlayerModule_Count - 1; i >= 0; i-- {
		pModule := p.arrModule[i]
		pModule.OnRelease()
	}
}

func (p *Player) GetId() int64 {
	return p.playerId
}

func (p *Player) Destroy() {
	for i, _ := range p.arrModule {
		p.arrModule[i] = nil
	}
	// p.pEventManager.Clear()
	// p.pEventManager = nil
	// p.pHandlerHolder = nil
}

func (p *Player) DailyRefresh(refreshTs int64) {
	// pData := p.GetModule(player.PlayerModule_Data).(*data.DataModule)
	// pData.DailyRefresh(refreshTs)
}

func (p *Player) RegisterMessage(pMsg proto.Message, handler func([]interface{})) {
	// p.pHandlerHolder.RegisterMessage(pMsg, handler)
}

func (p *Player) HandleMessage(msgId uint32, pMsg proto.Message) {
	f := p.pHandlerHolder.GetHandler(msgId)
	if f == nil {
		return
	}
	f([]interface{}{pMsg})
}

func (p *Player) SendMessage(msg proto.Message) {
	if p.conn != nil {
		p.conn.Write(msg)
	}
}

func (p *Player) GetModule(id int32) interface{} {
	return p.arrModule[id]
}

// func (p *Player) returnWaiting(pPlayer player.IPlayer) {
// 	worker.NewTask(func() {
// 		for _, sessionId := range p.arrWaitSession {
// 			public.Server.Resume(sessionId, pPlayer)
// 		}
// 		p.arrWaitSession = make([]uint32, 0)
// 	})
// }

// func (p *Player) waitLoadFinish() player.IPlayer {
// 	sessionId := public.Server.GetRunningSessionId()
// 	p.arrWaitSession = append(p.arrWaitSession, sessionId)
// 	ret := public.Server.Yield()
// 	if ret != nil && len(ret) > 0 {
// 		return ret[0].(player.IPlayer)
// 	} else {
// 		return nil
// 	}
// }

func (p *Player) GetIpAddr() string {
	if p.conn == nil {
		return ""
	}
	return p.conn.RemoteAddr().String()
}
