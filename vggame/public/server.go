package public

import (
	"time"
	"vgproj/common/cluster"
	"vgproj/common/oa"
	iconfig "vgproj/vggame/public/config"
	igame "vgproj/vggame/public/game"
	igate "vgproj/vggame/public/gate"

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
	GetGameManager() igame.IGameManager
	GetCluster() *cluster.Cluster
	GetOaWriter() *oa.Writer
	GetGate() igate.IGate
	GetConfigManager() iconfig.IConfigManager
	IsDebug() bool
}

var Server IServer
