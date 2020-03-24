package igate

import "github.com/golang/protobuf/proto"

type IGate interface {
	Kick(accountId int64) bool
}

type IMessageRouter interface {
	Route(msgId uint32, receiver interface{}, msg proto.Message)
}
