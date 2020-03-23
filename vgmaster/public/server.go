package public

import (
	"vgproj/common/cluster"
	icluster "vgproj/vgmaster/public/cluster"
	igame "vgproj/vgmaster/public/game"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

type IServer interface {
	virgo.IProcedure
	GetAuthKey() string
	GetDataDb() *database.Mysql
	GetCluster() *cluster.Cluster
	GetNodeManager() icluster.INodeManager
	GetNameManager() igame.INameManager
	IsDebug() bool
}

var Server IServer
