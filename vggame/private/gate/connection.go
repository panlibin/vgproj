package gate

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"reflect"

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
	// pPlayer        player.IPlayer
	accountId int64
	serverId  int32
	status    ConnectionStatus
}

func NewConnection(g *Gate, connectionId uint32, conn network.Connection) *Connection {
	pObj := new(Connection)
	pObj.conn = conn
	pObj.g = g
	pObj.connectionId = connectionId
	pObj.status = ConnectionStatus_WaitLogin

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
