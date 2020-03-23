package igame

import (
	icuslan "vgproj/vggame/public/game/custom_language"

	"github.com/panlibin/virgo/util/vgevent"
)

type IGameManager interface {
	GetEventManager() *vgevent.EventManager
	// GetPlayerManager() player.IPlayerManager
	// GetMailManager() mail.IMailManager
	GetCustomLanguageManager() icuslan.ICustomLanguageManager
}
