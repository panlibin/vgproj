package player

import "database/sql"

type iPlayerModule interface {
	getLoadSql() string
	onLoadData(*sql.Rows) error
	onInit1()
	onInit2()
	onInit3()
	onInit4()
	onInit5()
	onCreate()
	onLogin()
	onLogout()
	onRelease()
}

type playerModule struct {
	player *Player
}

func newPlayerModule(pPlayer *Player) *playerModule {
	pObj := new(playerModule)
	pObj.player = pPlayer
	return pObj
}

func (pm *playerModule) getLoadSql() string         { return "" }
func (pm *playerModule) onLoadData(*sql.Rows) error { return nil }
func (pm *playerModule) onInit1()                   {}
func (pm *playerModule) onInit2()                   {}
func (pm *playerModule) onInit3()                   {}
func (pm *playerModule) onInit4()                   {}
func (pm *playerModule) onInit5()                   {}
func (pm *playerModule) onCreate()                  {}
func (pm *playerModule) onLogin()                   {}
func (pm *playerModule) onLogout()                  {}
func (pm *playerModule) onRelease()                 {}
