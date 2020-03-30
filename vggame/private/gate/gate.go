package gate

import (
	"fmt"
	"sync"
	ec "vgproj/common/define/err_code"
	"vgproj/common/util"
	"vgproj/proto/msg"
	"vgproj/vggame/private/gate/ws"
	wsutil "vgproj/vggame/private/gate/ws/util"

	"github.com/panlibin/gnet"
	logger "github.com/panlibin/vglog"
)

type Gate struct {
	*gnet.EventServer
	listener          *gnet.GServer
	connMtx           sync.Mutex
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
	pObj.EventServer = &gnet.EventServer{}

	return pObj
}

func (g *Gate) Start(addr string) error {
	logger.Info("start gate")

	g.listener = &gnet.GServer{}

	pCodec := &WsCodec{g: g}
	pCusLog := &CustomLogger{Logger: &logger.DefaultLogger}

	err := g.listener.Serve(g, fmt.Sprintf("tcp://%s", addr), gnet.WithMulticore(true), gnet.WithCodec(pCodec), gnet.WithLogger(pCusLog))
	if err != nil {
		logger.Errorf("start gate error: %v", err)
	} else {
		logger.Infof("gate listen on %s", addr)
	}

	return err
}

func (g *Gate) Stop() {
	logger.Info("stop gate")

	if g.listener != nil {
		g.listener.SignalShutdown()
		g.listener.WaitShutdown()
	}

	g.connMtx.Lock()
	for _, pConnection := range g.mapConnection {
		pConnection.Close(nil)
	}
	g.connMtx.Unlock()

	logger.Info("stop gate finish")
}

func (g *Gate) OnNewConnection(conn gnet.Conn) *Connection {
	g.connMtx.Lock()
	connectionId := g.GenConnectionId()
	pConnection := NewConnection(g, connectionId, conn)
	g.mapConnection[connectionId] = pConnection
	g.connMtx.Unlock()

	return pConnection
	// logger.Debugf("new connection from %v, connectionId = %v", conn.RemoteAddr(), connectionId)
}

func (g *Gate) OnConnectionClose(connectionId uint32) {
	g.connMtx.Lock()
	delete(g.mapConnection, connectionId)
	g.connMtx.Unlock()

	// logger.Debugf("connection closed. remove connectionId = %v", connectionId)
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

func (g *Gate) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	ctx := c.Context()
	if ctx == nil {
		return
	}
	ctx.(*Connection).OnClose(err)
	return
}

func (g *Gate) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	ctx := c.Context()
	if ctx == nil {
		return
	}
	pConnection := ctx.(*Connection)

	if pConnection.wantType == wantClose {
		action = gnet.Close
		return
	}

	if !pConnection.upgraded {
		pConnection.upgraded = true
		out = frame
		return
	}

	switch pConnection.header.OpCode {
	case ws.OpBinary:
		err := pConnection.handleMessage(frame)
		if err != nil {
			action = gnet.Close
		}
	case ws.OpClose:
		out, _ = wsutil.HandleClose(pConnection.header, frame)
		action = gnet.Close
	case ws.OpPing:
		out, _ = wsutil.HandlePing(frame)
	case ws.OpPong:
		out, _ = wsutil.HandlePong(frame)
	}

	return
}
