package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	ec "vgproj/common/define/err_code"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
	network "github.com/panlibin/vgnet"
	"github.com/panlibin/virgo/util/vgtime"
)

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
		c.gsAddr = rsp.Server[1].Addr
		c.SetStatus(ClientStatus_Close)
	} else {
		logger.Debugf("%s: get server info fail", c.accountName)
		c.SetStatus(ClientStatus_Close)
	}
}
