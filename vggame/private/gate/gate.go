package gate

import (
	"sync"
	ec "vgproj/common/define/err_code"
	"vgproj/proto/msg"

	logger "github.com/panlibin/vglog"
	network "github.com/panlibin/vgnet"
	"github.com/panlibin/vgnet/websocket"
)

type Gate struct {
	listener          *websocket.Listener
	mtx               sync.Mutex
	maxConnectionId   uint32
	mapConnection     map[uint32]*Connection
	mapAccountSession map[int64]*Connection
	msgDesc           *MessageDescriptor
}

func NewGate() *Gate {
	pObj := new(Gate)
	pObj.mapConnection = make(map[uint32]*Connection, 1024)
	pObj.mapAccountSession = make(map[int64]*Connection, 1024)
	pObj.msgDesc = NewMessageDescriptor()

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

	g.mtx.Lock()
	for _, pConnection := range g.mapConnection {
		pConnection.Close(nil)
	}
	g.mtx.Unlock()

	logger.Info("stop gate finish")
}

func (g *Gate) OnNewConnection(conn network.Connection) {
	g.mtx.Lock()
	connectionId := g.GenConnectionId()
	pConnection := NewConnection(g, connectionId, conn)
	g.mapConnection[connectionId] = pConnection
	g.mtx.Unlock()

	conn.Accept(pConnection)
	logger.Debugf("new connection from %v, connectionId = %v", conn.RemoteAddr(), connectionId)
}

func (g *Gate) OnConnectionClose(connectionId uint32) {
	g.mtx.Lock()
	delete(g.mapConnection, connectionId)
	g.mtx.Unlock()

	logger.Debugf("connection closed. remove connectionId = %v", connectionId)
}

func (g *Gate) AddAccountSession(pConn *Connection) {
	pOldConn, exist := g.mapAccountSession[pConn.accountId]
	if exist {
		pOldConn.Write(&msg.S2C_Disconnect{Code: ec.LoginOnOtherTerminal})
		pOldConn.Close(nil)
	}
	g.mapAccountSession[pConn.accountId] = pConn
}

func (g *Gate) RemoveAccountSession(accountId int64) {
	delete(g.mapAccountSession, accountId)
}

func (g *Gate) Kick(accountId int64) bool {
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
