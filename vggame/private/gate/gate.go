package gate

import (
	"sync"
	ec "vgproj/common/define/err_code"
	"vgproj/common/util"
	"vgproj/proto/msg"

	logger "github.com/panlibin/vglog"
	network "github.com/panlibin/vgnet"
	"github.com/panlibin/vgnet/websocket"
)

type Gate struct {
	listener          *websocket.Listener
	connMtx           sync.Mutex
	acntMtx           sync.Mutex
	maxConnectionId   uint32
	mapConnection     map[uint32]*Connection
	mapAccountSession map[int64]*Connection
	msgDesc           *util.MessageDescriptor
	msgRouter         *messageRouter
}

func NewGate(msgDesc *util.MessageDescriptor) *Gate {
	pObj := new(Gate)
	pObj.mapConnection = make(map[uint32]*Connection, 1024)
	pObj.mapAccountSession = make(map[int64]*Connection, 1024)
	pObj.msgDesc = msgDesc
	pObj.msgRouter = &messageRouter{}
	pObj.msgRouter.init(msgDesc)

	return pObj
}

func (g *Gate) Start(addr string) error {
	logger.Info("start gate")

	g.listener = &websocket.Listener{
		Addr:            addr,
		NewConnCallback: g.OnNewConnection,
	}

	err := g.listener.Start()
	if err != nil {
		logger.Errorf("start gate error: %v", err)
	} else {
		logger.Infof("gate listen on %s", addr)
	}

	return err
}

func (g *Gate) Stop() {
	logger.Info("stop gate")

	g.listener.Stop()

	g.connMtx.Lock()
	for _, pConnection := range g.mapConnection {
		pConnection.Close(nil)
	}
	g.connMtx.Unlock()

	logger.Info("stop gate finish")
}

func (g *Gate) OnNewConnection(conn network.Connection) {
	g.connMtx.Lock()
	connectionId := g.GenConnectionId()
	pConnection := NewConnection(g, connectionId, conn)
	g.mapConnection[connectionId] = pConnection
	g.connMtx.Unlock()

	conn.Accept(pConnection)
	// logger.Debugf("new connection from %v, connectionId = %v", conn.RemoteAddr(), connectionId)
}

func (g *Gate) OnConnectionClose(connectionId uint32) {
	g.connMtx.Lock()
	delete(g.mapConnection, connectionId)
	g.connMtx.Unlock()

	// logger.Debugf("connection closed. remove connectionId = %v", connectionId)
}

func (g *Gate) AddAccountSession(pConn *Connection) {
	g.acntMtx.Lock()
	pOldConn, exist := g.mapAccountSession[pConn.accountId]
	if exist {
		pOldConn.Write(&msg.S2C_Disconnect{Code: ec.LoginOnOtherTerminal})
		pOldConn.Close(nil)
	}
	g.mapAccountSession[pConn.accountId] = pConn
	g.acntMtx.Unlock()
}

func (g *Gate) RemoveAccountSession(accountId int64) {
	g.acntMtx.Lock()
	delete(g.mapAccountSession, accountId)
	g.acntMtx.Unlock()
}

func (g *Gate) Kick(accountId int64) bool {
	g.acntMtx.Lock()
	defer g.acntMtx.Unlock()

	pConn, exist := g.mapAccountSession[accountId]
	if exist {
		pConn.Write(&msg.S2C_Disconnect{Code: ec.LoginOnOtherTerminal})
		pConn.Close(nil)
		delete(g.mapAccountSession, accountId)
		return true
	} else {
		return false
	}
}

func (g *Gate) GenConnectionId() uint32 {
	g.maxConnectionId++
	return g.maxConnectionId
}
