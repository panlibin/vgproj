package private

import (
	"encoding/json"
	"io/ioutil"
	"vgproj/common/cluster"
	clusthdl "vgproj/vgrecharge/private/cluster"
	"vgproj/vgrecharge/private/http"
	"vgproj/vgrecharge/private/recharge"
	"vgproj/vgrecharge/public"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
	"github.com/panlibin/virgo/util/nethelper"
	"github.com/panlibin/virgo/util/vgdir"
)

type localConfig struct {
	ServerID int32  `json:"server_id"`
	EnvDsn   string `json:"env_dsn"`
}

type envConfig struct {
	ServerID      int32
	ListenAddr    string
	ClusterAddr   string
	MasterAddr    string
	DataDsn       string
	DataDbConnNum int32
	ClientKey     string
	AuthKey       string
	CheckTime     int32
	Debug         int32
}

// Server 服务器
type Server struct {
	*virgo.Procedure
	envConf envConfig

	pDataDb *database.Mysql

	pHTTPServer      *nethelper.HTTPServer
	pCluster         *cluster.Cluster
	pRechargeManager *recharge.RechargeManager
}

// NewServer 创建服务器
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

	pEnvDb := database.NewMysql(s)
	err = pEnvDb.Open(localConf.EnvDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect config database error: %v", err)
		s.Stop()
		return false
	}

	row := pEnvDb.QueryRow(0, "select server_id,listen_addr,cluster_addr,master_addr,data_dsn,data_db_conn_num,client_key,auth_key,check_time,debug from c_recharge_server where server_id=?",
		localConf.ServerID)
	err = row.Scan(&s.envConf.ServerID, &s.envConf.ListenAddr, &s.envConf.ClusterAddr, &s.envConf.MasterAddr, &s.envConf.DataDsn, &s.envConf.DataDbConnNum, &s.envConf.ClientKey,
		&s.envConf.AuthKey, &s.envConf.CheckTime, &s.envConf.Debug)
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

// OnInit 初始化
func (s *Server) OnInit(p *virgo.Procedure) {
	s.Procedure = p
	logger.Infof("server init")
	if !s.initEnv() {
		return
	}

	public.Server = s
	var err error

	for {
		s.pRechargeManager = recharge.NewRechargeManager()
		if err = s.pRechargeManager.Init(); err != nil {
			break
		}

		s.pCluster = cluster.NewCluster(s, cluster.NodeRecharge, []int32{s.envConf.ServerID}, s.envConf.ClusterAddr, s.envConf.AuthKey)
		s.pCluster.SetHandler(&clusthdl.Server{})
		if err = s.pCluster.Start(); err != nil {
			break
		}
		s.pCluster.AddNode(cluster.NodeMaster, []int32{1}, s.envConf.MasterAddr)

		s.pHTTPServer = http.NewHTTPServer()
		if err = s.pHTTPServer.Start(s.envConf.ListenAddr, "conf/server.crt", "conf/server.key"); err != nil {
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

// OnRelease 销毁
func (s *Server) OnRelease() {
	if s.pHTTPServer != nil {
		s.pHTTPServer.Stop()
	}

	if s.pCluster != nil {
		s.pCluster.Stop()
	}

	if s.pDataDb != nil {
		s.pDataDb.Close()
	}

	logger.DefaultLogger.Flush()
}

func (s *Server) GetCluster() *cluster.Cluster {
	return s.pCluster
}

// GetDataDb 获取数据服务器
func (s *Server) GetDataDb() *database.Mysql {
	return s.pDataDb
}

func (s *Server) GetClientKey() string {
	return s.envConf.ClientKey
}

func (s *Server) GetAuthKey() string {
	return s.envConf.AuthKey
}

// IsDebug 是否调试状态
func (s *Server) IsDebug() bool {
	return s.envConf.Debug != 0
}

// CheckTime 验证签名时间
func (s *Server) CheckTime() bool {
	return s.envConf.CheckTime != 0
}
