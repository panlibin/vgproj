package gate

import (
	"vgproj/common/util"
	"vgproj/proto/msg"
	"vgproj/vggame/public"

	"github.com/golang/protobuf/proto"
)

type messageRouter struct {
	mapHandler map[uint32]func([]interface{})
}

func (mr *messageRouter) init(msgDesc *util.MessageDescriptor) {
	mr.mapHandler = map[uint32]func([]interface{}){
		msgDesc.Register(&msg.C2S_Login{}):           handleLogin,
		msgDesc.Register(&msg.C2S_CreateCharacter{}): handleCreateCharacter,
	}
}

func (mr *messageRouter) Route(msgId uint32, receiver interface{}, msg proto.Message) {
	f, exist := mr.mapHandler[msgId]
	if !exist {
		return
	}
	public.Server.SyncTask(f, receiver, msg)
}

func handleLogin(args []interface{}) {
	c := args[0].(*Connection)
	m := args[1].(*msg.C2S_Login)
	c.handleLogin(m)
}

func handleCreateCharacter(args []interface{}) {
	c := args[0].(*Connection)
	m := args[1].(*msg.C2S_CreateCharacter)
	c.handleCreateCharacter(m)
}
