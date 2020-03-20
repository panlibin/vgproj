package public

import (
	iaccount "vgproj/vglogin/public/account"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

// IServer 服务器
type IServer interface {
	virgo.IProcedure
	GetAccountManager() iaccount.IAccountManager
	GetDataDb() *database.Mysql
	GetClientKey() string
	CheckTime() bool
	IsDebug() bool
}

// Server 服务器全局变量
var Server IServer
