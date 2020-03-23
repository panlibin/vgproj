package cluster

import (
	"database/sql"
	"sync"
	"vgproj/vgmaster/public"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgstr"
)

type NodeInfo struct {
	ServerType int32
	ServerId   []int32
	Ip         string
}

type NodeManager struct {
	mapNode map[int32]map[int32]*NodeInfo
	mtx     sync.Mutex
}

func NewNodeManager() *NodeManager {
	return &NodeManager{
		mapNode: make(map[int32]map[int32]*NodeInfo),
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
			nInfo := &NodeInfo{}
			err = rows.Scan(&nInfo.ServerType, &tmpServerId, &nInfo.Ip)
			if err != nil {
				break
			}
			nInfo.ServerId, err = vgstr.SplitToInt32Array(tmpServerId, ",")
			if err != nil {
				break
			}

			mapType, exist := nm.mapNode[nInfo.ServerType]
			if !exist {
				mapType = make(map[int32]*NodeInfo)
				nm.mapNode[nInfo.ServerType] = mapType
			}
			for _, sId := range nInfo.ServerId {
				oldNode, exist := mapType[sId]
				if exist {
					nm.delete(oldNode.Ip)
					for _, oldSId := range oldNode.ServerId {
						delete(mapType, oldSId)
					}
				}
				mapType[sId] = nInfo
			}
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

func (nm *NodeManager) AddNode(serverType int32, serverId []int32, ip string) {
	nm.mtx.Lock()
	defer nm.mtx.Unlock()

	mapType, exist := nm.mapNode[serverType]
	if !exist {
		mapType = make(map[int32]*NodeInfo)
		nm.mapNode[serverType] = mapType
	}
	oldNode, exist := mapType[serverId[0]]
	if exist {
		if len(oldNode.ServerId) == len(serverId) && oldNode.Ip == ip {
			dup := true
			for i, sid := range serverId {
				if sid != oldNode.ServerId[i] {
					dup = false
					break
				}
			}
			if dup {
				return
			}
		}
	}

	node := &NodeInfo{
		ServerType: serverType,
		ServerId:   serverId,
		Ip:         ip,
	}

	for _, sId := range serverId {
		oldNode, exist := mapType[sId]
		if exist {
			nm.delete(oldNode.Ip)
			for _, oldSId := range oldNode.ServerId {
				delete(mapType, oldSId)
			}
		}
		mapType[sId] = node
	}

	nm.insert(node)
}

func (nm *NodeManager) RemoveNode(serverType int32, serverId []int32) {
	nm.mtx.Lock()
	defer nm.mtx.Unlock()

	mapType, exist := nm.mapNode[serverType]
	if !exist {
		return
	}
	node, exist := mapType[serverId[0]]
	if !exist {
		return
	}
	nm.delete(node.Ip)
	for _, sId := range node.ServerId {
		delete(mapType, sId)
	}
}

func (nm *NodeManager) GetAllNode() map[int32]map[int32]*NodeInfo {
	return nm.mapNode
}

func (nm *NodeManager) insert(node *NodeInfo) {
	tmpServerId := vgstr.FormatInt32ArrayToString(node.ServerId, ",")
	public.Server.GetDataDb().AsyncExec(nil, 0, "insert into node_list values(?,?,?)", node.ServerType, tmpServerId, node.Ip)
}

func (nm *NodeManager) delete(ip string) {
	public.Server.GetDataDb().AsyncExec(nil, 0, "delete from node_list where ip=?", ip)
}
