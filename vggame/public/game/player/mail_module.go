package iplayer

import (
	imail "vgproj/vggame/public/game/mail"
)

type IMailModule interface {
	SendMail(pMailDef *imail.PlayerMailDef)
	SendGlobalMail(pMailDef *imail.GlobalMailDef)
}
