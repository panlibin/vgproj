package game

import (
	"time"
	cuslan "vgproj/vggame/private/game/custom_language"
	"vgproj/vggame/public"
	igame "vgproj/vggame/public/game"
	icuslan "vgproj/vggame/public/game/custom_language"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgevent"
	"github.com/panlibin/virgo/util/vgtime"
)

type GameManager struct {
	arrModule     [Module_Count]IModule
	pEventManager *vgevent.EventManager
	dailyTimer    *time.Timer
}

func NewGameManager() *GameManager {
	pObj := new(GameManager)
	// pObj.arrModule[Module_Player] = player.NewPlayerManager(pMsgDesc)
	pObj.arrModule[Module_CustomLanguage] = cuslan.NewCustomLanguageManager()
	// pObj.arrModule[Module_Mail] = mail.NewMailManager()
	pObj.pEventManager = vgevent.NewEventManager()

	return pObj
}

func (gm *GameManager) LoadData() error {
	var err error
	for idx, pModule := range gm.arrModule {
		err = pModule.OnLoadData()
		if err != nil {
			logger.Errorf("%v load data error: %v", idx, err)
			return err
		}
	}
	return nil
}

func (gm *GameManager) Init() error {
	var err error
	for idx, pModule := range gm.arrModule {
		err = pModule.OnInit()
		if err != nil {
			logger.Errorf("%v init error: %v", idx, err)
			return err
		}
	}

	gm.createDailyTimer(nil)

	return nil
}

func (gm *GameManager) Release() {
	if gm.dailyTimer != nil {
		gm.dailyTimer.Stop()
		gm.dailyTimer = nil
	}

	for i := Module_Count - 1; i >= 0; i-- {
		pModule := gm.arrModule[i]
		pModule.OnRelease()
	}
}

func (gm *GameManager) GetEventManager() *vgevent.EventManager {
	return gm.pEventManager
}

// func (gm *GameManager) GetPlayerManager() pub_player.IPlayerManager {
// 	return gm.arrModule[Module_Player].(pub_player.IPlayerManager)
// }

// func (gm *GameManager) GetMailManager() pub_mail.IMailManager {
// 	return gm.arrModule[Module_Mail].(pub_mail.IMailManager)
// }

func (gm *GameManager) GetCustomLanguageManager() icuslan.ICustomLanguageManager {
	return gm.arrModule[Module_CustomLanguage].(icuslan.ICustomLanguageManager)
}

func (gm *GameManager) createDailyTimer(args []interface{}) {
	nextRefreshTs := vgtime.NextDailyRefreshTs()
	curTs := vgtime.Now()
	gm.dailyTimer = public.Server.AfterFunc(time.Duration(nextRefreshTs-curTs)*time.Millisecond, gm.createDailyTimer, nextRefreshTs)
	gm.pEventManager.Dispatch(&igame.EventDailyRefresh{RefreshTs: args[0].(int64)})
}
