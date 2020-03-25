package game

import (
	"time"
	"vgproj/common/util"
	cuslan "vgproj/vggame/private/game/custom_language"
	"vgproj/vggame/private/game/mail"
	"vgproj/vggame/private/game/player"
	"vgproj/vggame/public"
	igame "vgproj/vggame/public/game"
	icuslan "vgproj/vggame/public/game/custom_language"
	imail "vgproj/vggame/public/game/mail"
	iplayer "vgproj/vggame/public/game/player"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgevent"
	"github.com/panlibin/virgo/util/vgtime"
)

type GameManager struct {
	arrModule     [Module_Count]IModule
	pEventManager *vgevent.EventManager
	dailyTimer    *time.Timer
}

func NewGameManager(msgDesc *util.MessageDescriptor) *GameManager {
	pObj := new(GameManager)
	pObj.arrModule[Module_Player] = player.NewPlayerManager(msgDesc)
	pObj.arrModule[Module_CustomLanguage] = cuslan.NewCustomLanguageManager()
	pObj.arrModule[Module_Mail] = mail.NewMailManager()
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

func (gm *GameManager) GetPlayerManager() iplayer.IPlayerManager {
	return gm.arrModule[Module_Player].(iplayer.IPlayerManager)
}

func (gm *GameManager) GetMailManager() imail.IMailManager {
	return gm.arrModule[Module_Mail].(imail.IMailManager)
}

func (gm *GameManager) GetCustomLanguageManager() icuslan.ICustomLanguageManager {
	return gm.arrModule[Module_CustomLanguage].(icuslan.ICustomLanguageManager)
}

func (gm *GameManager) createDailyTimer(args []interface{}) {
	nextRefreshTs := vgtime.NextDailyRefreshTs()
	curTs := vgtime.Now()
	gm.dailyTimer = public.Server.AfterFunc(time.Duration(nextRefreshTs-curTs)*time.Millisecond, gm.createDailyTimer, nextRefreshTs)
	if args != nil {
		gm.pEventManager.Dispatch(&igame.EventDailyRefresh{RefreshTs: args[0].(int64)})
	}
}
