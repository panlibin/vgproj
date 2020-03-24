package iplayer

import (
	igate "vgproj/vggame/public/gate"

	"github.com/golang/protobuf/proto"
)

type IPlayer interface {
	Login(conn igate.IConnection)
	Logout()
	SendMessage(msg proto.Message)
	GetId() int64
	GetIpAddr() string
	DailyRefresh(refreshTs int64)
	GetModule(id int32) interface{}
	// AddThings(mapThings map[int32]int64, source int32, sourceExt string)
	// IsEnoughThings(mapConsume map[int32]int64) bool
	// ConsumeThings(mapConsume map[int32]int64, source int32, sourceExt string) bool
}
