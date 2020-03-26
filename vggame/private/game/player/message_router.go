package player

import (
	"vgproj/common/util"
	"vgproj/proto/msg"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/golang/protobuf/proto"
)

type messageRouter struct {
	mapHandler map[uint32]func(*Player, proto.Message)
}

func (mr *messageRouter) init(msgDesc *util.MessageDescriptor) {
	mr.mapHandler = map[uint32]func(*Player, proto.Message){
		msgDesc.Register(&msg.C2S_ROLE_ITEMS{}): handleGetRoleItems,
	}
}

func (mr *messageRouter) Route(msgId uint32, receiver interface{}, msg proto.Message) {
	f, exist := mr.mapHandler[msgId]
	if !exist {
		return
	}
	f(receiver.(*Player), msg)
}

func handleGetRoleItems(p *Player, m proto.Message) {
	p.GetModule(iplayer.PlayerModule_Item).(*itemModule).handleGetRoleItems(m.(*msg.C2S_ROLE_ITEMS))
}
