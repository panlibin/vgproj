package gate

import (
	"vgproj/common/util"
	"vgproj/proto/msg"

	"github.com/golang/protobuf/proto"
)

type messageRouter struct {
	mapHandler map[uint32]func(*Connection, proto.Message)
}

func (mr *messageRouter) init(msgDesc *util.MessageDescriptor) {
	mr.mapHandler = map[uint32]func(*Connection, proto.Message){
		msgDesc.Register(&msg.C2S_Login{}):           handleLogin,
		msgDesc.Register(&msg.C2S_CreateCharacter{}): handleCreateCharacter,
	}
}

func (mr *messageRouter) Route(msgId uint32, receiver interface{}, msg proto.Message) {
	f, exist := mr.mapHandler[msgId]
	if !exist {
		return
	}
	f(receiver.(*Connection), msg)
}

func handleLogin(c *Connection, m proto.Message) {
	c.handleLogin(m.(*msg.C2S_Login))
}

func handleCreateCharacter(c *Connection, m proto.Message) {
	c.handleCreateCharacter(m.(*msg.C2S_CreateCharacter))
}
