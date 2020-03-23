package public

import (
	"time"
	"vgproj/common/cluster"
	"vgproj/common/oa"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

type IServer interface {
	virgo.IProcedure
	GetServerId() int32
	GetServerIdArray() []int32
	IsSelfServerId(serverId int32) bool
	GetAuthKey() string
	GetDataDb() *database.Mysql
	GetGlobalDb() *database.Mysql
	GetOaDb() *database.Mysql
	GetConfigDb() *database.Mysql
	GetOpenServerTime() time.Time
	GetOpenServerTs() int64
	GetOpenServerDay(ts int64) int32
	// GetGameManager() game.IGameManager
	GetCluster() *cluster.Cluster
	GetOaWriter() *oa.Writer
	// GetGate() gate.IGate
	// GetConfigManager() config.IConfigManager
	IsDebug() bool
}

var Server IServer
