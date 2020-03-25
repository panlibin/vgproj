package client

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	ec "vgproj/common/define/err_code"
	"vgproj/common/util"
	"vgproj/proto/msg"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
	network "github.com/panlibin/vgnet"
	"github.com/panlibin/vgnet/websocket"
	"github.com/panlibin/virgo/util/vgtime"
)

var msgDesc *util.MessageDescriptor

func init() {
	rand.Seed(time.Now().UnixNano())
	msgDesc = util.NewMessageDescriptor()
	msgDesc.Register(&msg.C2S_Login{})
	msgDesc.Register(&msg.S2C_Login{})
	msgDesc.Register(&msg.C2S_CreateCharacter{})
	msgDesc.Register(&msg.S2C_CreateCharacter{})
	msgDesc.Register(&msg.S2C_LoginFinish{})
}

type Client struct {
	loginType   int32
	accountName string
	password    string
	accountId   int64
	token       string
	status      int32
	gsAddr      string
	pWg         *sync.WaitGroup

	conn network.Connection

	arrBehavior []func()
	waitChan    chan proto.Message
	waitMsgId   uint32
}

func NewClient(loginType int32, accountName string, password string, pWg *sync.WaitGroup) *Client {
	pObj := new(Client)
	pObj.loginType = loginType
	pObj.accountName = accountName
	pObj.password = password
	pObj.status = ClientStatus_LoginAccount
	pObj.pWg = pWg
	pObj.waitChan = make(chan proto.Message)
	pWg.Add(1)

	pObj.arrBehavior = make([]func(), ClientStatus_Count)
	pObj.arrBehavior[ClientStatus_LoginAccount] = pObj.LoginAccount
	pObj.arrBehavior[ClientStatus_RegisterAccount] = pObj.RegisterAccount
	pObj.arrBehavior[ClientStatus_GetServerInfo] = pObj.GetServerInfo
	pObj.arrBehavior[ClientStatus_LoginGame] = pObj.LoginGame
	pObj.arrBehavior[ClientStatus_CreateCharacter] = pObj.CreateCharacter

	return pObj
}

func (c *Client) Run() {
	for {
		if c.status == ClientStatus_Close {
			if c.conn != nil {
				c.conn.Close(nil)
				c.conn = nil
			}
			c.pWg.Done()
			return
		} else {
			c.arrBehavior[c.status]()
		}
	}
}

func (c *Client) SetStatus(status int32) {
	c.status = status
}

func (c *Client) GetStatus() int32 {
	return c.status
}

func (c *Client) calcSign(params []string) string {
	strSrc := fmt.Sprintf("%s%s", strings.Join(params, ""), GlobalConfig.ClientKey)
	sum := md5.Sum([]byte(strSrc))
	return hex.EncodeToString(append(sum[:]))
}

type LoginAccountRsp struct {
	Code      int32  `json:"code"`
	AccountId int64  `json:"account_id"`
	Token     string `json:"token"`
}

func (c *Client) LoginAccount() {
	curTs := vgtime.Now()
	sign := c.calcSign([]string{strconv.Itoa(int(c.loginType)), c.accountName, c.password, strconv.FormatInt(curTs, 10)})
	resp, err := http.Get(fmt.Sprintf("%s/login?login_type=%d&account_name=%s&password=%s&time=%d&sign=%s",
		GlobalConfig.LoginServerAddr, c.loginType, c.accountName, c.password, curTs, sign))
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}
	defer resp.Body.Close()
	rspBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	rsp := LoginAccountRsp{}
	err = json.Unmarshal(rspBuf, &rsp)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	if rsp.Code == ec.AccountNotFound {
		logger.Debugf("%s: account not found, try register", c.accountName)
		c.SetStatus(ClientStatus_RegisterAccount)
	} else if rsp.Code == ec.Success {
		c.accountId = rsp.AccountId
		c.token = rsp.Token
		logger.Debugf("%s: login account suc", c.accountName)
		c.SetStatus(ClientStatus_GetServerInfo)
	} else {
		logger.Debugf("%s: login account fail", c.accountName)
		c.SetStatus(ClientStatus_Close)
	}
}

type CodeRsp struct {
	Code int32 `json:"code"`
}

func (c *Client) RegisterAccount() {
	curTs := vgtime.Now()
	sign := c.calcSign([]string{c.accountName, c.password, strconv.FormatInt(curTs, 10)})
	resp, err := http.Get(fmt.Sprintf("%s/register?account_name=%s&password=%s&time=%d&sign=%s",
		GlobalConfig.LoginServerAddr, c.accountName, c.password, curTs, sign))
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}
	defer resp.Body.Close()
	rspBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	rsp := CodeRsp{}
	err = json.Unmarshal(rspBuf, &rsp)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	if rsp.Code == ec.Success {
		logger.Debugf("%s: register account suc", c.accountName)
		c.SetStatus(ClientStatus_LoginAccount)
	} else {
		logger.Debugf("%s: register account fail", c.accountName)
		c.SetStatus(ClientStatus_Close)
	}
}

type GameServer struct {
	ServerId int32
	Name     string
	Status   int32
	Addr     string
}

type ServerInfoRsp struct {
	Code   int32                 `json:"code"`
	Server map[int32]*GameServer `json:"server"`
}

func (c *Client) GetServerInfo() {
	curTs := vgtime.Now()
	sign := c.calcSign([]string{strconv.FormatInt(curTs, 10)})
	resp, err := http.Get(fmt.Sprintf("%s/server_list?&time=%d&sign=%s", GlobalConfig.LoginServerAddr, curTs, sign))
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}
	defer resp.Body.Close()
	rspBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	rsp := ServerInfoRsp{}
	err = json.Unmarshal(rspBuf, &rsp)
	if err != nil {
		logger.Error(err)
		c.SetStatus(ClientStatus_Close)
		return
	}

	if rsp.Code == ec.Success {
		logger.Debugf("%s: get server info suc", c.accountName)
		c.gsAddr = "ws://" + rsp.Server[1].Addr
		c.SetStatus(ClientStatus_LoginGame)
	} else {
		logger.Debugf("%s: get server info fail", c.accountName)
		c.SetStatus(ClientStatus_Close)
	}
}

func (c *Client) LoginGame() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	bConnectSuc := false
	websocket.Dialer.Dial(c.gsAddr, func(conn network.Connection, e error) {
		if e != nil {
			return
		}
		conn.Accept(c)
		c.conn = conn
		bConnectSuc = true

		wg.Done()
	})
	wg.Wait()

	if !bConnectSuc {
		c.SetStatus(ClientStatus_Close)
	}

	pMsgReq := new(msg.C2S_Login)
	pMsgReq.AccountId = c.accountId
	pMsgReq.Token = c.token
	pMsgReq.ServerId = 1
	c.Write(pMsgReq)
	pMsgRsp := c.WaitResponse(&msg.S2C_Login{}).(*msg.S2C_Login)
	if pMsgRsp.Code != ec.Success {
		logger.Debugf("%s: login game server fail", c.accountName)
		c.SetStatus(ClientStatus_Close)
		return
	}
	logger.Debugf("%s: login game server suc", c.accountName)
	if pMsgRsp.PlayerId == 0 {
		c.SetStatus(ClientStatus_CreateCharacter)
		return
	} else {
		c.SetStatus(ClientStatus_Close)
		return
	}
}

func (c *Client) CreateCharacter() {
	pMsgReq := new(msg.C2S_CreateCharacter)
	pMsgReq.Name = strconv.FormatUint(rand.Uint64(), 16)
	c.Write(pMsgReq)
	pMsgRsp := c.WaitResponse(&msg.S2C_CreateCharacter{}).(*msg.S2C_CreateCharacter)
	if pMsgRsp.Code != ec.Success {
		c.SetStatus(ClientStatus_Close)
		return
	} else {
		logger.Debugf("%s: create character success", c.accountName)
		c.SetStatus(ClientStatus_Close)
		return
	}
}

func (c *Client) Write(msg proto.Message) {
	c.conn.Write(msg)
}

func (c *Client) WaitResponse(msg proto.Message) proto.Message {
	c.waitMsgId = msgDesc.GetMessageId(msg)
	rsp := <-c.waitChan
	c.waitMsgId = 0
	return rsp
}

func (c *Client) OnClose(err error) {
	if err != nil {
		logger.Debug(err)
	}
}

func (c *Client) DoRead(r io.Reader) error {
	msgBuf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	msgId := binary.BigEndian.Uint32(msgBuf[:4])
	msgType, exist := msgDesc.GetMessageType(msgId)
	if !exist {
		return nil
	}
	msg := reflect.New(msgType).Interface().(proto.Message)
	err = proto.Unmarshal(msgBuf[4:], msg)
	if err != nil {
		return err
	}
	if msgId == c.waitMsgId {
		c.waitChan <- msg
	}

	return nil
}

func (c *Client) DoWrite(w io.Writer, m interface{}) error {
	msg, ok := m.(proto.Message)
	if !ok {
		return errors.New("message convert to proto buffer fail")
	}

	msgId := msgDesc.GetMessageId(msg)
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
