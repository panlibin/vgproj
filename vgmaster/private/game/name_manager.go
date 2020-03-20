package game

import (
	"database/sql"
	"vgproj/vgmaster/public"
)

type NameManager struct {
	mapPlayerName map[string]struct{}
	mapGuildName  map[string]struct{}
}

func NewNameManager() *NameManager {
	pObj := new(NameManager)
	pObj.mapPlayerName = make(map[string]struct{}, 512)
	pObj.mapGuildName = make(map[string]struct{}, 512)
	return pObj
}

func (this *NameManager) Init() error {
	var err error
	for {
		var rows *sql.Rows
		rows, err = public.Server.GetDataDb().Query(this.getDbIdx(), "select `name` from global_player_name")
		if err != nil {
			break
		}
		var tmpName string
		for rows.Next() {
			err = rows.Scan(&tmpName)
			if err != nil {
				rows.Close()
				break
			}
			this.mapPlayerName[tmpName] = struct{}{}
		}
		rows.Close()
		if err != nil {
			break
		}

		rows, err = public.Server.GetDataDb().Query(this.getDbIdx(), "select `name` from global_guild_name")
		if err != nil {
			break
		}
		for rows.Next() {
			err = rows.Scan(&tmpName)
			if err != nil {
				rows.Close()
				break
			}
			this.mapGuildName[tmpName] = struct{}{}
		}
		rows.Close()
		if err != nil {
			break
		}

		break
	}

	if err != nil {
		logger.Errorf("name manager init error: %v", err)
	}
	return err
}

func (this *NameManager) Release() {

}

func (this *NameManager) GrabPlayerName(name string) bool {
	_, exist := this.mapPlayerName[name]
	if exist {
		return false
	}
	this.mapPlayerName[name] = struct{}{}
	this.insertPlayerName(name)
	return true
}

func (this *NameManager) ReleasePlayerName(name string) {
	delete(this.mapPlayerName, name)
	this.deletePlayerName(name)
}

func (this *NameManager) GrabGuildName(name string) bool {
	_, exist := this.mapGuildName[name]
	if exist {
		return false
	}
	this.mapGuildName[name] = struct{}{}
	this.insertGuildName(name)
	return true
}

func (this *NameManager) ReleaseGuildName(name string) {
	delete(this.mapGuildName, name)
	this.deleteGuildName(name)
}

func (this *NameManager) insertPlayerName(name string) {
	public.Server.GetDataDb().AsyncExec(this.getDbIdx(), "insert into global_player_name values(?)", name)
}

func (this *NameManager) deletePlayerName(name string) {
	public.Server.GetDataDb().AsyncExec(this.getDbIdx(), "delete from global_player_name where `name`=?", name)
}

func (this *NameManager) insertGuildName(name string) {
	public.Server.GetDataDb().AsyncExec(this.getDbIdx(), "insert into global_guild_name values(?)", name)
}

func (this *NameManager) deleteGuildName(name string) {
	public.Server.GetDataDb().AsyncExec(this.getDbIdx(), "delete from global_guild_name where `name`=?", name)
}

func (this *NameManager) getDbIdx() uint32 {
	return 1
}
