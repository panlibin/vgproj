package gate

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"reflect"
	"time"
	"vgproj/common/cluster"
	ec "vgproj/common/define/err_code"
	"vgproj/proto/loginrpc"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iplayer "vgproj/vggame/public/game/player"
	igate "vgproj/vggame/public/gate"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
	network "github.com/panlibin/vgnet"
)

type ConnectionStatus int32

const (
	ConnectionStatus_WaitLogin ConnectionStatus = iota
	ConnectionStatus_CheckingAccount
	ConnectionStatus_LoadingPlayer
	ConnectionStatus_WaitCreatePlayer
	ConnectionStatus_CreatingPlayer
	ConnectionStatus_Normal
	ConnectionStatus_Closed
)

type Connection struct {
	g            *Gate
	connectionId uint32
	conn         network.Connection
	pPlayer      iplayer.IPlayer
	msgRouter    igate.IMessageRouter
	msgReceiver  interface{}
	accountId    int64
	serverId     int32
	status       ConnectionStatus
}

func NewConnection(g *Gate, connectionId uint32, conn network.Connection) *Connection {
	pObj := new(Connection)
	pObj.conn = conn
	pObj.g = g
	pObj.connectionId = connectionId
	pObj.status = ConnectionStatus_WaitLogin
	pObj.msgRouter = g.msgRouter
	pObj.msgReceiver = pObj

	return pObj
}

func (c *Connection) DoRead(r io.Reader) error {
	msgBuf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	msgId := binary.BigEndian.Uint32(msgBuf[:4])
	msgType, exist := c.g.msgDesc.GetMessageType(msgId)
	if !exist {
		logger.Errorf("receive unregister message id = %v", msgId)
		return errors.New("receive unregister message")
	}
	msg := reflect.New(msgType).Interface().(proto.Message)
	err = proto.Unmarshal(msgBuf[4:], msg)
	if err != nil {
		return err
	}
	c.msgRouter.Route(msgId, c.msgReceiver, msg)

	return err
}

func (c *Connection) DoWrite(w io.Writer, m interface{}) error {
	msg, ok := m.(proto.Message)
	if !ok {
		return errors.New("message convert to proto buffer fail")
	}

	msgId := c.g.msgDesc.GetMessageId(msg)
	msgBuf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	msgLen := len(msgBuf)
	writeBuf := make([]byte, msgLen+4)
	binary.BigEndian.PutUint32(writeBuf, msgId)
	copy(writeBuf[4:], msgBuf)

	_, err = w.Write(writeBuf)
	return err
}

func (c *Connection) OnClose(err error) {
	if err != nil {
		logger.Debugf("connection closed: ", err)
	}

}

func (c *Connection) Close(err error) {
	c.conn.Close(err)
}

func (c *Connection) Write(pMsg proto.Message) {
	c.conn.Write(pMsg)
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) handleLogin(pMsgReq *msg.C2S_Login) {
	if c.status != ConnectionStatus_WaitLogin {
		return
	}
	c.status = ConnectionStatus_CheckingAccount

	var errCode int32 = ec.Unknown
	for {
		if !public.Server.IsSelfServerId(pMsgReq.ServerId) {
			errCode = ec.InvalidParam
			break
		}

		pNode := public.Server.GetCluster().GetNode(cluster.NodeLogin, 1)
		if pNode == nil {
			errCode = ec.Unknown
			break
		}

		pLoginNode := pNode.(*cluster.LoginNode)
		public.Server.AsyncTask(func([]interface{}) {
			rsp, err := pLoginNode.Login(context.Background(), &loginrpc.ReqLogin{AccountId: pMsgReq.AccountId, Token: pMsgReq.Token, ServerId: pMsgReq.ServerId})
			public.Server.SyncTask(c.loginCheckCallback, rsp, err, pMsgReq)
		})

		errCode = ec.Success
		break
	}

	if errCode != ec.Success {
		c.Write(&msg.S2C_Login{Code: errCode})
	}
}

func (c *Connection) loginCheckCallback(args []interface{}) {
	if c.status != ConnectionStatus_CheckingAccount {
		return
	}

	pMsgRsp := new(msg.S2C_Login)
	var errCode int32 = ec.Unknown
	for {
		if args[0] == nil || args[1] != nil {
			errCode = ec.Unknown
			c.status = ConnectionStatus_WaitLogin
			break
		}

		pMsgReq := args[2].(*msg.C2S_Login)

		rspLogin := args[0].(*loginrpc.RspLogin)
		if rspLogin.Code != ec.Success {
			errCode = rspLogin.Code
			c.status = ConnectionStatus_WaitLogin
			break
		}

		// 账号验证通过
		c.accountId = pMsgReq.AccountId
		c.serverId = pMsgReq.ServerId
		c.g.AddAccountSession(c)
		// 获取角色id
		pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
		playerId := pPlayerManager.GetPlayerIdByAccountId(c.accountId, c.serverId)
		t := time.Now()
		pMsgRsp.ServerTime = t.Unix()*1000 + int64(t.Nanosecond())/1000000
		_, zoneOffset := t.Local().Zone()
		pMsgRsp.Offset = int32(zoneOffset)
		pMsgRsp.OpenServerTime = public.Server.GetOpenServerTs()
		pMsgRsp.PlayerId = playerId

		if playerId > 0 {
			// 加载角色数据
			c.status = ConnectionStatus_LoadingPlayer
			pPlayerManager.GetPlayer(playerId, pMsgRsp, c.loginLoadPlayerCallback)
		} else {
			// 未创建角色
			c.status = ConnectionStatus_WaitCreatePlayer
		}

		errCode = ec.Success
		break
	}

	pMsgRsp.Code = errCode
	if pMsgRsp.PlayerId == 0 {
		c.Write(pMsgRsp)
	}
}

func (c *Connection) loginLoadPlayerCallback(ctx interface{}, pPlayer iplayer.IPlayer) {
	if c.status != ConnectionStatus_LoadingPlayer {
		return
	}
	pMsgRsp := ctx.(*msg.S2C_Login)
	if pPlayer == nil {
		// 加载角色失败
		pMsgRsp.PlayerId = 0
		pMsgRsp.Code = ec.Unknown
		c.status = ConnectionStatus_WaitLogin
		c.Write(pMsgRsp)
		return
	}

	c.status = ConnectionStatus_Normal
	pMsgRsp.Code = ec.Success
	c.pPlayer = pPlayer
	c.msgReceiver = pPlayer
	c.msgRouter = public.Server.GetGameManager().GetPlayerManager().GetMessageRouter()

	c.Write(pMsgRsp)
	pPlayer.Login(c)
}

func (c *Connection) handleCreateCharacter(pMsgReq *msg.C2S_CreateCharacter) {
	if c.status != ConnectionStatus_WaitCreatePlayer {
		return
	}
	c.status = ConnectionStatus_CreatingPlayer

	pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
	pPlayerManager.CreatePlayer(c.accountId, c.serverId, pMsgReq.Name, pMsgReq.Head, nil, c.createCharacterCallback)
}

func (c *Connection) createCharacterCallback(ctx interface{}, pPlayer iplayer.IPlayer, errCode int32) {
	if c.status != ConnectionStatus_CreatingPlayer {
		return
	}
	pMsgRsp := new(msg.S2C_CreateCharacter)
	for {
		pMsgRsp.Code = errCode

		if errCode == ec.DuplicateCreateCharacter {
			c.status = ConnectionStatus_WaitLogin
			break
		}
		if errCode != ec.Success {
			c.status = ConnectionStatus_WaitCreatePlayer
			break
		}

		c.status = ConnectionStatus_Normal
		c.pPlayer = pPlayer
		c.msgReceiver = pPlayer
		c.msgRouter = public.Server.GetGameManager().GetPlayerManager().GetMessageRouter()

		break
	}

	c.Write(pMsgRsp)
	if errCode == ec.Success && pPlayer != nil {
		pPlayer.Login(c)
	}
}
