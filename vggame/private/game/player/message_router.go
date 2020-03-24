package player

import (
	"vgproj/common/util"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/golang/protobuf/proto"
)

type messageRouter struct {
	mapHandler map[uint32]func([]interface{})
}

func (mr *messageRouter) init(msgDesc *util.MessageDescriptor) {
	mr.mapHandler = map[uint32]func([]interface{}){
		msgDesc.Register(&msg.C2S_ROLE_ITEMS{}): handleGetRoleItems,
	}
}

func (mr *messageRouter) Route(msgId uint32, receiver interface{}, msg proto.Message) {
	f, exist := mr.mapHandler[msgId]
	if !exist {
		return
	}
	public.Server.SyncTask(f, receiver, msg)
}

func handleGetRoleItems(args []interface{}) {
	p := args[0].(*Player)
	// m := args[1].(*msg.C2S_ROLE_ITEMS)
	p.GetModule(iplayer.PlayerModule_Item)
}
