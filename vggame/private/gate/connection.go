package gate

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"reflect"
	"time"
	ec "vgproj/common/define/err_code"
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
	public.Server.SyncTask(c.dispatch, msgId, msg)

	return err
}

func (c *Connection) dispatch(args []interface{}) {
	if c.status == ConnectionStatus_Closed {
		return
	}
	c.msgRouter.Route(args[0].(uint32), c.msgReceiver, args[1].(proto.Message))
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
	c.g.OnConnectionClose(c.connectionId)
	if c.status == ConnectionStatus_Closed {
		return
	}
	public.Server.SyncTask(c.doClose)
}

func (c *Connection) Close(err error) {
	if c.status == ConnectionStatus_Closed {
		return
	}
	c.conn.Close(err)
	c.doClose(nil)
}

func (c *Connection) doClose([]interface{}) {
	if c.status == ConnectionStatus_Closed {
		return
	}
	c.status = ConnectionStatus_Closed

	if c.accountId > 0 {
		c.g.RemoveAccountSession(c.accountId)
	}

	if c.pPlayer != nil {
		c.pPlayer.Logout()
	}
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

	pMsgRsp := new(msg.S2C_Login)
	var errCode int32 = ec.Unknown
	var playerId int64
	for {
		if !public.Server.IsSelfServerId(pMsgReq.ServerId) {
			errCode = ec.InvalidParam
			break
		}

		t := time.Now()
		ts := t.Unix()*1000 + int64(t.Nanosecond())/1000000
		if pMsgReq.Ts+600000 < ts {
			errCode = ec.InvalidToken
			break
		}

		token := sha256.Sum256([]byte(fmt.Sprintf("ts=%d&auth_key=%s&account_id=%d", pMsgReq.Ts, public.Server.GetAuthKey(), pMsgReq.AccountId)))
		if pMsgReq.Token != base64.URLEncoding.EncodeToString(token[:]) {
			errCode = ec.InvalidToken
			break
		}

		// 账号验证通过
		c.accountId = pMsgReq.AccountId
		c.serverId = pMsgReq.ServerId
		c.g.AddAccountSession(c)
		// 获取角色id
		pPlayerManager := public.Server.GetGameManager().GetPlayerManager()
		playerId = pPlayerManager.GetPlayerIdByAccountId(c.accountId, c.serverId)
		pMsgRsp.ServerTime = ts
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

	if errCode != ec.Success {
		c.status = ConnectionStatus_WaitLogin
		pMsgRsp.Code = errCode
		c.Write(pMsgRsp)
	} else if playerId == 0 {
		pMsgRsp.Code = errCode
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
