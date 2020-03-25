package private

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"time"
	"vgproj/common/cluster"
	"vgproj/common/oa"
	"vgproj/common/util"
	clusthdl "vgproj/vggame/private/cluster"
	"vgproj/vggame/private/config"
	"vgproj/vggame/private/game"
	"vgproj/vggame/private/gate"
	"vgproj/vggame/public"
	iconfig "vgproj/vggame/public/config"
	igame "vgproj/vggame/public/game"
	igate "vgproj/vggame/public/gate"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
	"github.com/panlibin/virgo/util/vgdir"
	"github.com/panlibin/virgo/util/vgtime"
)

type localConfig struct {
	ServerId []int32 `json:"server_id"`
	EnvDsn   string  `json:"env_dsn"`
}

type envConfig struct {
	ServerId      int32
	ArrServerId   []int32
	MapServerId   map[int32]struct{}
	ListenAddr    string
	ClusterAddr   string
	MasterAddr    string
	DataDsn       string
	DataDbConnNum int32
	OaDsn         string
	OaDbConnNum   int32
	ConfDsn       string
	AuthKey       string
	Debug         int32
}

type Server struct {
	*virgo.Procedure
	envConf envConfig

	pDataDb   *database.Mysql
	pGlobalDb *database.Mysql
	pOaDb     *database.Mysql
	pConfDb   *database.Mysql

	openServerTime time.Time
	openServerTs   int64

	pOaWriter      *oa.Writer
	pConfigManager *config.ConfigManager
	pGameManager   *game.GameManager
	pGate          *gate.Gate
	pCluster       *cluster.Cluster
}

func NewServer() *Server {
	pObj := new(Server)
	pObj.envConf.MapServerId = make(map[int32]struct{})
	return pObj
}

func (s *Server) initEnv() bool {
	fn := vgdir.ConvDirAbs("./conf/conf.json")
	localConfData, err := ioutil.ReadFile(fn)
	if err != nil {
		logger.Errorf("Server init read file error: %v", err)
		s.Stop()
		return false
	}
	localConf := localConfig{}
	err = json.Unmarshal(localConfData, &localConf)
	if err != nil {
		logger.Errorf("Server init json unmarshal error: %v", err)
		s.Stop()
		return false
	}
	s.envConf.ServerId = localConf.ServerId[0]
	s.envConf.ArrServerId = localConf.ServerId
	for _, serverId := range localConf.ServerId {
		s.envConf.MapServerId[serverId] = struct{}{}
	}

	pEnvDb := database.NewMysql(s)
	err = pEnvDb.Open(localConf.EnvDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect config database error: %v", err)
		s.Stop()
		return false
	}

	row := pEnvDb.QueryRow(0, "select listen_addr,cluster_addr,master_addr,data_dsn,data_db_conn_num"+
		",oa_dsn,oa_db_conn_num,conf_dsn,auth_key,debug from c_game_server where server_id=?", s.envConf.ServerId)
	err = row.Scan(&s.envConf.ListenAddr, &s.envConf.ClusterAddr, &s.envConf.MasterAddr, &s.envConf.DataDsn, &s.envConf.DataDbConnNum,
		&s.envConf.OaDsn, &s.envConf.OaDbConnNum, &s.envConf.ConfDsn, &s.envConf.AuthKey, &s.envConf.Debug)
	if err != nil {
		logger.Errorf("Server init scan config error: %v", err)
		s.Stop()
		return false
	}
	pEnvDb.Close()

	s.pDataDb = database.NewMysql(s)
	err = s.pDataDb.Open(s.envConf.DataDsn, s.envConf.DataDbConnNum)
	if err != nil {
		logger.Errorf("Server init connect data database error: %v", err)
		s.Stop()
		return false
	}

	s.pGlobalDb = database.NewMysql(s)
	err = s.pGlobalDb.Open(s.envConf.DataDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect data database error: %v", err)
		s.Stop()
		return false
	}

	s.pOaDb = database.NewMysql(s)
	err = s.pOaDb.Open(s.envConf.OaDsn, s.envConf.OaDbConnNum)
	if err != nil {
		logger.Errorf("Server init connect log database error: %v", err)
		s.Stop()
		return false
	}

	s.pConfDb = database.NewMysql(s)
	err = s.pConfDb.Open(s.envConf.ConfDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect config database error: %v", err)
		s.Stop()
		return false
	}

	if !s.IsDebug() {
		logger.DefaultLogger.SetSeverityLimit(logger.SeverityInfo)
	}

	return true
}

func (s *Server) OnInit(p *virgo.Procedure) {
	s.Procedure = p
	logger.Info("server init")
	if !s.initEnv() {
		return
	}

	public.Server = s
	var err error

	for {
		row := s.pDataDb.QueryRow(0, "select open_server_time from global_system")
		if err = row.Scan(&s.openServerTime); err != nil {
			if err == sql.ErrNoRows {
				s.openServerTime = time.Now()
				s.pDataDb.Exec(0, "insert into global_system values(?)", s.openServerTime)
			} else {
				break
			}
		}

		s.openServerTs = s.openServerTime.Unix() * 1000

		s.pConfigManager = config.NewConfigManager()
		if err = s.pConfigManager.LoadConfig(); err != nil {
			break
		}

		s.pOaWriter = oa.NewWriter(s, s.pOaDb, s.envConf.OaDbConnNum)

		msgDesc := util.NewMessageDescriptor()
		s.pGate = gate.NewGate(msgDesc)
		s.pGameManager = game.NewGameManager(msgDesc)
		if err = s.pGameManager.LoadData(); err != nil {
			break
		}
		if err = s.pGameManager.Init(); err != nil {
			break
		}

		s.pCluster = cluster.NewCluster(s, cluster.NodeGame, s.GetServerIdArray(), s.envConf.ClusterAddr, s.GetAuthKey())
		s.pCluster.SetHandler(&clusthdl.Server{})
		if err = s.pCluster.Start(); err != nil {
			break
		}
		s.pCluster.AddNode(cluster.NodeMaster, []int32{1}, s.envConf.MasterAddr)

		if err = s.pGate.Start(s.envConf.ListenAddr); err != nil {
			break
		}

		break
	}
	if err != nil {
		s.Stop()
		return
	}

	logger.Info("server init finish")
}

func (s *Server) OnRelease() {
	if s.pGate != nil {
		s.pGate.Stop()
	}
	if s.pCluster != nil {
		s.pCluster.Stop()
	}
	if s.pGameManager != nil {
		s.pGameManager.Release()
	}

	if s.pDataDb != nil {
		s.pDataDb.Close()
	}
	if s.pOaDb != nil {
		s.pOaDb.Close()
	}
	if s.pConfDb != nil {
		s.pConfDb.Close()
	}

	logger.DefaultLogger.Flush()
}

func (s *Server) GetServerId() int32 {
	return s.envConf.ServerId
}

func (s *Server) GetServerIdArray() []int32 {
	return s.envConf.ArrServerId
}

func (s *Server) IsSelfServerId(serverId int32) bool {
	_, exist := s.envConf.MapServerId[serverId]
	return exist
}

func (s *Server) GetAuthKey() string {
	return s.envConf.AuthKey
}

func (s *Server) GetDataDb() *database.Mysql {
	return s.pDataDb
}

func (s *Server) GetGlobalDb() *database.Mysql {
	return s.pGlobalDb
}

func (s *Server) GetOaDb() *database.Mysql {
	return s.pOaDb
}

func (s *Server) GetConfigDb() *database.Mysql {
	return s.pConfDb
}

func (s *Server) GetOpenServerTime() time.Time {
	return s.openServerTime
}

func (s *Server) GetOpenServerTs() int64 {
	return s.openServerTs
}

func (s *Server) GetOpenServerDay(ts int64) int32 {
	return int32((ts-vgtime.GetDayZeroTs(s.openServerTs))/(24*3600*1000) + 1)
}

func (s *Server) GetGameManager() igame.IGameManager {
	return s.pGameManager
}

func (s *Server) GetCluster() *cluster.Cluster {
	return s.pCluster
}

func (s *Server) GetOaWriter() *oa.Writer {
	return s.pOaWriter
}

func (s *Server) GetGate() igate.IGate {
	return s.pGate
}

func (s *Server) GetConfigManager() iconfig.IConfigManager {
	return s.pConfigManager
}

func (s *Server) IsDebug() bool {
	return s.envConf.Debug != 0
}
