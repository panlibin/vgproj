package iplayer

import "database/sql"

const (
	PlayerModule_Data int32 = iota
	PlayerModule_Property
	PlayerModule_Item
	PlayerModule_Hero
	PlayerModule_Mail
	PlayerModule_Vip
	PlayerModule_Settings

	PlayerModule_Count
)

type IPlayerModule interface {
	GetLoadSql() string
	OnLoadData(*sql.Rows) error
	OnInit1()
	OnInit2()
	OnInit3()
	OnInit4()
	OnInit5()
	OnCreate()
	OnLogin()
	OnLogout()
	OnRelease()
}

type PlayerModule struct {
	Player IPlayer
}

func NewPlayerModule(pPlayer IPlayer) *PlayerModule {
	pObj := new(PlayerModule)
	pObj.Player = pPlayer
	return pObj
}

func (pm *PlayerModule) OnLoadData() error { return nil }
func (pm *PlayerModule) OnInit1()          {}
func (pm *PlayerModule) OnInit2()          {}
func (pm *PlayerModule) OnInit3()          {}
func (pm *PlayerModule) OnInit4()          {}
func (pm *PlayerModule) OnInit5()          {}
func (pm *PlayerModule) OnCreate()         {}
func (pm *PlayerModule) OnLogin()          {}
func (pm *PlayerModule) OnLogout()         {}
func (pm *PlayerModule) OnRelease()        {}
