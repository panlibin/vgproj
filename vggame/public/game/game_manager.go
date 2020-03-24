package igame

import (
	icuslan "vgproj/vggame/public/game/custom_language"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/panlibin/virgo/util/vgevent"
)

type IGameManager interface {
	GetEventManager() *vgevent.EventManager
	GetPlayerManager() iplayer.IPlayerManager
	// GetMailManager() mail.IMailManager
	GetCustomLanguageManager() icuslan.ICustomLanguageManager
}
