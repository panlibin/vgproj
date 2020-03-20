package public

import (
	"vgproj/common/cluster"
	igame "vgproj/vgmaster/public/game"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

type IServer interface {
	virgo.IProcedure
	GetAuthKey() string
	GetDataDb() *database.Mysql
	GetCluster() *cluster.Cluster
	GetNameManager() igame.INameManager
	GetGameServerManager() igame.IGameServerManager
	IsDebug() bool
}

var Server IServer
