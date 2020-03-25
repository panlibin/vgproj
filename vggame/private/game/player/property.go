package player

import (
	"time"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
)

type Property struct {
	player     *Player
	id         int32
	val        int64
	updateTs   int64
	inDatabase bool
}

func newProperty(player *Player, id int32) *Property {
	pObj := new(Property)
	pObj.player = player
	pObj.id = id
	return pObj
}

func (p *Property) Load(val int64, updateTs int64) {
	p.val = val
	p.updateTs = updateTs
	p.inDatabase = true
}

func (p *Property) GetId() int32 {
	return p.id
}

func (p *Property) GetValue() int64 {
	return p.val
}

func (p *Property) AddValue(val int64, source int32, sourceExt string) {
	if p.addValue(val, source, sourceExt) {
		p.notifyChange()
		//p.player.GetEventManager().Dispatch(&player.EventPropertyIncrease{Id: p.id, Delta: val, CurNum: p.val})
	}
}

func (p *Property) SubValue(val int64, source int32, sourceExt string) bool {
	ret := p.subValue(val, source, sourceExt)
	if ret {
		p.notifyChange()
	}
	return ret
}

func (p *Property) SetValue(val int64, source int32, sourceExt string) {
	delta := val - p.val
	p.changeValue(delta, source, sourceExt)
	p.update()
	p.notifyChange()
}

func (p *Property) GetUpdateTs() int64 {
	return p.updateTs
}

func (p *Property) addValue(val int64, source int32, sourceExt string) bool {
	if val <= 0 {
		return false
	}
	p.changeValue(val, source, sourceExt)
	p.update()
	return true
}

func (p *Property) subValue(val int64, source int32, sourceExt string) bool {
	if val <= 0 || p.val < val {
		return false
	}
	p.changeValue(-val, source, sourceExt)
	p.update()
	return true
}

func (p *Property) notifyChange() {
	pMsgNotify := new(msg.S2C_SYNC_NUM)
	pProp := new(msg.NUM)
	pProp.NumType = p.id
	pProp.Num = p.val
	pProp.Data1 = p.updateTs
	pMsgNotify.Nums = append(pMsgNotify.Nums, pProp)
	p.player.SendMessage(pMsgNotify)
}

func (p *Property) changeValue(delta int64, source int32, sourceExt string) {
	p.val += delta
	public.Server.GetOaWriter().Write("log_item", p.player.GetId(), source, sourceExt, p.id, delta, p.val, time.Now())
}

func (p *Property) update() {
	if p.inDatabase {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(p.player.GetId()), sqlUpdateProp, p.val, p.updateTs, p.player.GetId(), p.id)
	} else {
		p.inDatabase = true
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(p.player.GetId()), sqlInsertProp, p.player.GetId(), p.id, p.val, p.updateTs)
	}
}
