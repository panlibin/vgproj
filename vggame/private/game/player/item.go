package player

import (
	"time"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
)

type Item struct {
	player     *Player
	id         int32
	num        int64
	inDatabase bool
}

func newItem(player *Player, id int32) *Item {
	pObj := new(Item)
	pObj.player = player
	pObj.id = id
	return pObj
}

func (i *Item) Load(num int64) {
	i.num = num
	i.inDatabase = true
}

func (i *Item) GetId() int32 {
	return i.id
}

func (i *Item) GetNum() int64 {
	return i.num
}

func (i *Item) AddNum(num int64, source int32, sourceExt string) {
	if i.addNum(num, source, sourceExt) {
		i.notifyChange()
	}
}

func (i *Item) SubNum(num int64, source int32, sourceExt string) bool {
	ret := i.subNum(num, source, sourceExt)
	if ret {
		i.notifyChange()
	}
	return ret
}

func (i *Item) addNum(num int64, source int32, sourceExt string) bool {
	if num <= 0 {
		return false
	}
	i.changeNum(num, source, sourceExt)
	i.update()
	return true
}

func (i *Item) subNum(num int64, source int32, sourceExt string) bool {
	if num <= 0 || i.num < num {
		return false
	}
	i.changeNum(-num, source, sourceExt)
	i.update()
	return true
}

func (i *Item) notifyChange() {
	if i.num != 0 {
		pMsgNotify := new(msg.S2C_SYNC_ITEMS)
		pItem := new(msg.ROLE_ITEM)
		pItem.ItemId = i.id
		pItem.Num = i.num
		pMsgNotify.Items = append(pMsgNotify.Items, pItem)
		i.player.SendMessage(pMsgNotify)
	} else {
		pMsgNotify := new(msg.S2C_SYNC_ITEMS_DEL)
		pMsgNotify.DelIds = []int32{i.id}
		i.player.SendMessage(pMsgNotify)
	}
}

func (i *Item) changeNum(delta int64, source int32, sourceExt string) {
	i.num += delta
	public.Server.GetOaWriter().Write("log_item", i.player.GetId(), source, sourceExt, i.id, delta, i.num, time.Now())
}

func (i *Item) update() {
	if i.inDatabase {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(i.player.GetId()), sqlUpdateItem, i.num, i.player.GetId(), i.id)
	} else {
		i.inDatabase = true
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(i.player.GetId()), sqlInsertItem, i.player.GetId(), i.id, i.num)
	}
}

func (i *Item) formatToMessage() *msg.ROLE_ITEM {
	return &msg.ROLE_ITEM{ItemId: i.id, Num: i.num}
}
