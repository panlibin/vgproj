package iplayer

import (
	igate "vgproj/vggame/public/gate"

	"github.com/golang/protobuf/proto"
	"github.com/panlibin/virgo/util/vgevent"
)

type IPlayer interface {
	Login(conn igate.IConnection)
	Logout()
	RegisterMessage(pMsg proto.Message, handler func([]interface{}))
	HandleMessage(msgId uint32, pMsg proto.Message)
	SendMessage(msg proto.Message)
	GetId() int64
	GetIpAddr() string
	DailyRefresh(refreshTs int64)
	GetModule(id int32) interface{}
	GetEventManager() *vgevent.EventManager
	AddThings(mapThings map[int32]int64, source int32, sourceExt string)
	IsEnoughThings(mapConsume map[int32]int64) bool
	ConsumeThings(mapConsume map[int32]int64, source int32, sourceExt string) bool
}
