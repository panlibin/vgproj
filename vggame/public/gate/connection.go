package igate

import (
	"net"

	"github.com/golang/protobuf/proto"
)

type IConnection interface {
	Close(error)
	Write(proto.Message)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	GetRnd() uint64
}
