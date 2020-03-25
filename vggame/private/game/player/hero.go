package player

import (
	"vgproj/vggame/public"
)

type Hero struct {
	player *Player

	id  int32
	lev int64

	inDb bool
}

func newHero(pPlayer *Player) *Hero {
	pObj := new(Hero)
	pObj.player = pPlayer
	pObj.lev = 1
	pObj.inDb = false
	return pObj
}

func (h *Hero) GetId() int32 {
	return h.id
}

func (h *Hero) GetLevel() int64 {
	return h.lev
}

func (h *Hero) addLevel(addLev int64) {
	h.lev += addLev
	h.save()
}

func (h *Hero) insert() {
	h.inDb = true
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(h.player.GetId()), sqlInsertHero, h.player.GetId(), h.id, h.lev)
}

func (h *Hero) save() {
	if h.inDb {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(h.player.GetId()), sqlUpdateHero, h.lev, h.player.GetId(), h.id)
	} else {
		h.insert()
	}
}
