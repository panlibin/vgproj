package public

import (
	"vgproj/common/cluster"
	iaccount "vgproj/vglogin/public/account"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

// IServer 服务器
type IServer interface {
	virgo.IProcedure
	GetAccountManager() iaccount.IAccountManager
	GetCluster() *cluster.Cluster
	GetDataDb() *database.Mysql
	GetClientKey() string
	CheckTime() bool
	IsDebug() bool
	GetAuthKey() string
}

// Server 服务器全局变量
var Server IServer
