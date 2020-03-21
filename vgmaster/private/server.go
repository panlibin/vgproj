package private

import (
	"encoding/json"
	"io/ioutil"
	"vgproj/common/cluster"
	clusthdl "vgproj/vgmaster/private/cluster"
	"vgproj/vgmaster/private/game"
	"vgproj/vgmaster/public"
	igame "vgproj/vgmaster/public/game"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
	"github.com/panlibin/virgo/util/vgdir"
)

type localConfig struct {
	ServerId int32  `json:"server_id"`
	EnvDsn   string `json:"env_dsn"`
}

type envConfig struct {
	ServerId      int32
	ListenAddr    string
	ClusterAddr   string
	DataDsn       string
	DataDbConnNum int32
	AuthKey       string
	Debug         int32
}

type Server struct {
	*virgo.Procedure
	envConf envConfig

	pDataDb *database.Mysql

	pCluster     *cluster.Cluster
	pNodeManager *clusthdl.NodeManager
	pNameManager *game.NameManager
}

func NewServer() *Server {
	pObj := new(Server)
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
	s.envConf.ServerId = localConf.ServerId

	pEnvDb := database.NewMysql(s)
	err = pEnvDb.Open(localConf.EnvDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect config database error: %v", err)
		s.Stop()
		return false
	}

	row := pEnvDb.QueryRow(0, "select listen_addr,cluster_addr,data_dsn,data_db_conn_num,auth_key,debug"+
		" from c_master_server where server_id=?", s.envConf.ServerId)
	err = row.Scan(&s.envConf.ListenAddr, &s.envConf.ClusterAddr, &s.envConf.DataDsn, &s.envConf.DataDbConnNum,
		&s.envConf.AuthKey, &s.envConf.Debug)
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

	if !s.IsDebug() {
		logger.DefaultLogger.SetSeverityLimit(logger.SeverityInfo)
	}

	return true
}

func (s *Server) OnInit(p *virgo.Procedure) {
	logger.Infof("server init")
	if !s.initEnv() {
		return
	}

	public.Server = s
	var err error

	for {
		s.pNameManager = game.NewNameManager()
		if err = s.pNameManager.Init(); err != nil {
			break
		}

		s.pNodeManager = clusthdl.NewNodeManager()
		if err = s.pNodeManager.Init(); err != nil {
			break
		}

		s.pCluster = cluster.NewCluster(s, cluster.NodeMaster, []int32{s.envConf.ServerId}, s.envConf.ClusterAddr, s.GetAuthKey())
		s.pCluster.SetHandler(&clusthdl.Server{})
		if err = s.pCluster.Start(); err != nil {
			break
		}

		break
	}
	if err != nil {
		s.Stop()
		return
	}

	logger.Infof("server init finish")
}

func (s *Server) OnRelease() {
	if s.pCluster != nil {
		s.pCluster.Stop()
	}

	if s.pDataDb != nil {
		s.pDataDb.Close()
	}
}

func (s *Server) GetServerId() int32 {
	return s.envConf.ServerId
}

func (s *Server) GetAuthKey() string {
	return s.envConf.AuthKey
}

func (s *Server) GetDataDb() *database.Mysql {
	return s.pDataDb
}

func (s *Server) GetCluster() *cluster.Cluster {
	return s.pCluster
}

func (s *Server) GetNameManager() igame.INameManager {
	return s.pNameManager
}

func (s *Server) IsDebug() bool {
	return s.envConf.Debug != 0
}
