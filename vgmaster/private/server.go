package private

import (
	"encoding/json"
	"io/ioutil"
	"vgproj/common/cluster"
	"vgproj/vgmaster/private/game"

	"vgproj/vgmaster/public"

	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
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
	pNameManager *game.NameManager
}

func NewServer() *Server {
	pObj := new(Server)
	return pObj
}

func (this *Server) initEnv() bool {
	fn := vg_dir.ConvDirAbs("./conf/conf.json")
	localConfData, err := ioutil.ReadFile(fn)
	if err != nil {
		logger.Errorf("Server init read file error: %v", err)
		this.Stop()
		return false
	}
	localConf := localConfig{}
	err = json.Unmarshal(localConfData, &localConf)
	if err != nil {
		logger.Errorf("Server init json unmarshal error: %v", err)
		this.Stop()
		return false
	}
	this.envConf.ServerId = localConf.ServerId

	pEnvDb := database.NewMysql()
	err = pEnvDb.Open(localConf.EnvDsn, 1)
	if err != nil {
		logger.Errorf("Server init connect config database error: %v", err)
		this.Stop()
		return false
	}

	row := pEnvDb.QueryRow(0, "select listen_addr,cluster_addr,data_dsn,data_db_conn_num,auth_key,debug"+
		" from c_master_server where server_id=?", this.envConf.ServerId)
	err = row.Scan(&this.envConf.ListenAddr, &this.envConf.ClusterAddr, &this.envConf.DataDsn, &this.envConf.DataDbConnNum,
		&this.envConf.AuthKey, &this.envConf.Debug)
	if err != nil {
		logger.Errorf("Server init scan config error: %v", err)
		this.Stop()
		return false
	}
	pEnvDb.Close()

	this.pDataDb = database.NewMysql()
	err = this.pDataDb.Open(this.envConf.DataDsn, this.envConf.DataDbConnNum)
	if err != nil {
		logger.Errorf("Server init connect data database error: %v", err)
		this.Stop()
		return false
	}

	if !this.IsDebug() {
		logger.SetSeverityLimit(logger.SeverityInfo)
	}

	return true
}

func (this *Server) OnInit(p *virgo.Procedure) {
	logger.Infof("server init")
	if !this.initEnv() {
		return
	}

	public.Server = this
	var err error

	for {
		this.pNameManager = game.NewNameManager()
		if err = this.pNameManager.Init(); err != nil {
			break
		}

		this.pCluster = cluster.NewCluster(this, cluster.NodeMaster, []int32{this.envConf.ServerId}, this.envConf.ClusterAddr, this.GetAuthKey())
		this.pCluster.SetHandler(cluster2.NewMasterHandler(this.pCluster.MsgDesc))
		if err = this.pCluster.Start(); err != nil {
			break
		}

		break
	}
	if err != nil {
		this.Stop()
		return
	}

	logger.Infof("server init finish")
}

func (this *Server) OnRelease() {
	if this.pCluster != nil {
		this.pCluster.Stop()
	}

	if this.pDataDb != nil {
		this.pDataDb.Close()
	}
}

func (this *Server) GetServerId() int32 {
	return this.envConf.ServerId
}

func (this *Server) GetAuthKey() string {
	return this.envConf.AuthKey
}

func (this *Server) GetDataDb() *database.Mysql {
	return this.pDataDb
}

func (this *Server) GetCluster() *cluster.Cluster {
	return this.pCluster
}

func (this *Server) GetNameManager() pub_game.INameManager {
	return this.pNameManager
}

func (this *Server) IsDebug() bool {
	return this.envConf.Debug != 0
}
