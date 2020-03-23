package igate

import (
	"net"

	"github.com/golang/protobuf/proto"
)

type IConnection interface {
	Init()
	Close(error)
	Write(proto.Message)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}
