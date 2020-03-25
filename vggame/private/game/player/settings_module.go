package player

import (
	"database/sql"
	"fmt"
	"vgproj/proto/msg"
	"vgproj/vggame/public"

	logger "github.com/panlibin/vglog"
)

type SettingsModule struct {
	*playerModule

	lan  int32
	inDb bool
}

func newSettingsModule(pPlayer *Player) *SettingsModule {
	pObj := new(SettingsModule)
	pObj.playerModule = newPlayerModule(pPlayer)

	return pObj
}

func (sm *SettingsModule) getLoadSql() string {
	return fmt.Sprintf("select language from player_settings where player_id=%d;", sm.player.GetId())
}

func (sm *SettingsModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_settings error: %v", err)
		return err
	}
	if !rows.Next() {
		return nil
	}
	err = rows.Scan(&sm.lan)
	if err != nil {
		logger.Errorf("load player_settings error: %v", err)
		return err
	} else {
		sm.inDb = true
	}

	return nil
}

func (sm *SettingsModule) onLogin() {
	pMsgSettings := new(msg.S2C_PlayerSettings)
	pMsgSettings.Lan = sm.lan
	sm.player.SendMessage(pMsgSettings)
}

func (sm *SettingsModule) GetLanguage() int32 {
	return sm.lan
}

func (sm *SettingsModule) SetLanguage(lan int32) {
	if sm.lan == lan {
		return
	}
	sm.lan = lan
	sm.save()
}

func (sm *SettingsModule) insert() {
	sm.inDb = true
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(sm.player.GetId()), sqlSettingsInsert, sm.player.GetId(), sm.lan)
}

func (sm *SettingsModule) save() {
	if sm.inDb {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(sm.player.GetId()), sqlSettingsUpdate, sm.lan, sm.player.GetId())
	} else {
		sm.insert()
	}
}
