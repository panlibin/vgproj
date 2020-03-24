package util

import (
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"
	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgstr"
)

type MessageDescriptor struct {
	mapMsgInfo map[uint32]reflect.Type
	mapMsgId   map[reflect.Type]uint32
	rwLock     sync.RWMutex
}

func NewMessageDescriptor() *MessageDescriptor {
	pObj := new(MessageDescriptor)
	pObj.mapMsgInfo = make(map[uint32]reflect.Type)
	pObj.mapMsgId = make(map[reflect.Type]uint32)
	return pObj
}

func (md *MessageDescriptor) Register(msg proto.Message) uint32 {
	msgType := reflect.TypeOf(msg).Elem()
	msgId := md.calcMessageId(msgType)

	md.rwLock.Lock()
	if _, exist := md.mapMsgId[msgType]; exist {
		logger.Warningf("duplicate register message %s", msgType.String())
		return msgId
	}
	if _, exist := md.mapMsgInfo[msgId]; exist {
		logger.Errorf("register message hash id collided")
		return msgId
	}

	md.mapMsgId[msgType] = msgId
	md.mapMsgInfo[msgId] = msgType
	md.rwLock.Unlock()
	return msgId
}

func (md *MessageDescriptor) GetMessageId(msg proto.Message) (msgId uint32) {
	msgType := reflect.TypeOf(msg).Elem()
	var exist bool
	md.rwLock.RLock()
	msgId, exist = md.mapMsgId[msgType]
	md.rwLock.RUnlock()
	if !exist {
		msgId = md.Register(msg)
	}
	return
}

func (md *MessageDescriptor) GetMessageType(msgId uint32) (msgType reflect.Type, exist bool) {
	md.rwLock.RLock()
	msgType, exist = md.mapMsgInfo[msgId]
	md.rwLock.RUnlock()
	return
}

func (md *MessageDescriptor) calcMessageId(msgType reflect.Type) uint32 {
	//arrMsgNameSplit := strings.Split(msgType.String(), ".")
	//return vg_str.Hash(arrMsgNameSplit[len(arrMsgNameSplit) - 1])
	return vgstr.Hash(msgType.String())
}
