package igame

import (
	icuslan "vgproj/vggame/public/game/custom_language"
	imail "vgproj/vggame/public/game/mail"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/panlibin/virgo/util/vgevent"
)

type IGameManager interface {
	GetEventManager() *vgevent.EventManager
	GetPlayerManager() iplayer.IPlayerManager
	GetMailManager() imail.IMailManager
	GetCustomLanguageManager() icuslan.ICustomLanguageManager
}
