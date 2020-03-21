package game

import (
	"database/sql"
	"vgproj/vgmaster/public"

	logger "github.com/panlibin/vglog"
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

func (nm *NameManager) Init() error {
	var err error
	for {
		var rows *sql.Rows
		rows, err = public.Server.GetDataDb().Query(nm.getDbIdx(), "select `name` from global_player_name")
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
			nm.mapPlayerName[tmpName] = struct{}{}
		}
		rows.Close()
		if err != nil {
			break
		}

		rows, err = public.Server.GetDataDb().Query(nm.getDbIdx(), "select `name` from global_guild_name")
		if err != nil {
			break
		}
		for rows.Next() {
			err = rows.Scan(&tmpName)
			if err != nil {
				rows.Close()
				break
			}
			nm.mapGuildName[tmpName] = struct{}{}
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

func (nm *NameManager) Release() {

}

func (nm *NameManager) GrabPlayerName(name string) bool {
	_, exist := nm.mapPlayerName[name]
	if exist {
		return false
	}
	nm.mapPlayerName[name] = struct{}{}
	nm.insertPlayerName(name)
	return true
}

func (nm *NameManager) ReleasePlayerName(name string) {
	delete(nm.mapPlayerName, name)
	nm.deletePlayerName(name)
}

func (nm *NameManager) GrabGuildName(name string) bool {
	_, exist := nm.mapGuildName[name]
	if exist {
		return false
	}
	nm.mapGuildName[name] = struct{}{}
	nm.insertGuildName(name)
	return true
}

func (nm *NameManager) ReleaseGuildName(name string) {
	delete(nm.mapGuildName, name)
	nm.deleteGuildName(name)
}

func (nm *NameManager) insertPlayerName(name string) {
	public.Server.GetDataDb().AsyncExec(nil, nm.getDbIdx(), "insert into global_player_name values(?)", name)
}

func (nm *NameManager) deletePlayerName(name string) {
	public.Server.GetDataDb().AsyncExec(nil, nm.getDbIdx(), "delete from global_player_name where `name`=?", name)
}

func (nm *NameManager) insertGuildName(name string) {
	public.Server.GetDataDb().AsyncExec(nil, nm.getDbIdx(), "insert into global_guild_name values(?)", name)
}

func (nm *NameManager) deleteGuildName(name string) {
	public.Server.GetDataDb().AsyncExec(nil, nm.getDbIdx(), "delete from global_guild_name where `name`=?", name)
}

func (nm *NameManager) getDbIdx() uint32 {
	return 1
}
