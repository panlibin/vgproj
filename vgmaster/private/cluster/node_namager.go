package cluster

import (
	"database/sql"
	"vgproj/vgmaster/public"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgstr"
)

type nodeInfo struct {
	serverType int32
	serverId   []int32
	ip         string
}

type NodeManager struct {
	mapNode map[string]*nodeInfo
}

func NewNodeManager() *NodeManager {
	return &NodeManager{
		mapNode: make(map[string]*nodeInfo, 32),
	}
}

func (nm *NodeManager) Init() error {
	var err error
	for {
		var rows *sql.Rows
		rows, err = public.Server.GetDataDb().Query(0, "select server_type,server_id,ip from node_list")
		if err != nil {
			break
		}
		var tmpServerId string
		for rows.Next() {
			nInfo := &nodeInfo{}
			err = rows.Scan(&nInfo.serverType, tmpServerId, &nInfo.ip)
			if err != nil {
				break
			}
			nInfo.serverId, err = vgstr.SplitToInt32Array(tmpServerId, ",")
			if err != nil {
				break
			}

			nm.mapNode[nInfo.ip] = nInfo
		}
		rows.Close()
		if err != nil {
			break
		}
		break
	}

	if err != nil {
		logger.Errorf("node manager init error: %v", err)
	}
	return err
}

func (nm *NodeManager) addNode(serverType int32, serverId []int32, ip string) {
	oldNode, exist := nm.mapNode[ip]
	if exist && oldNode != nil {
		nm.removeNode(ip)
	}

	node := &nodeInfo{
		serverType: serverType,
		serverId:   serverId,
		ip:         ip,
	}
	nm.mapNode[ip] = node
	nm.insert(node)
}

func (nm *NodeManager) removeNode(ip string) {
	node, exist := nm.mapNode[ip]
	if exist {
		delete(nm.mapNode, ip)
		if node != nil {
			nm.delete(ip)
		}
	}
}

func (nm *NodeManager) insert(node *nodeInfo) {
	tmpServerId := vgstr.FormatInt32ArrayToString(node.serverId, ",")
	public.Server.GetDataDb().AsyncExec(nil, 0, "insert into node_list values(?,?,?)", node.serverType, tmpServerId, node.ip)
}

func (nm *NodeManager) delete(ip string) {
	public.Server.GetDataDb().AsyncExec(nil, 0, "delete from node_list where ip=?", ip)
}
