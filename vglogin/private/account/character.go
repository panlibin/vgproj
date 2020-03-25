package account

import (
	"vgproj/vglogin/public"

	"github.com/panlibin/virgo/util/vgtime"
)

type Character struct {
	id         int64
	accountId  int64
	serverId   int32
	name       string
	combat     int64
	updateTs   int64
	needUpdate bool
}

func (c *Character) GetPlayerId() int64 {
	return c.id
}

func (c *Character) GetName() string {
	return c.name
}

func (c *Character) GetCombat() int64 {
	return c.combat
}

func (c *Character) GetUpdateTs() int64 {
	return c.updateTs
}

func (c *Character) setName(name string) {
	if name == c.name {
		return
	}
	c.name = name
	c.needUpdate = true
}

func (c *Character) setCombat(combat int64) {
	if combat == c.combat {
		return
	}
	c.combat = combat
	c.needUpdate = true
}

func (c *Character) insert() {
	const strInsertSql = "insert into character_info values(?,?,?,?,?,?)"
	public.Server.GetDataDb().Exec(uint32(c.accountId), strInsertSql, c.id, c.accountId,
		c.serverId, c.name, c.combat, vgtime.Now())
}

func (c *Character) update() {
	if !c.needUpdate {
		return
	}
	const strUpdateSql = "update character_info set name=?,combat=? where player_id=?"
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(c.accountId), strUpdateSql, c.name, c.combat, c.id)
	c.needUpdate = false
}
