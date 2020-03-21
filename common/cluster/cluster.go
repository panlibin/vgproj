package cluster

import (
	"net"
	"sync"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"google.golang.org/grpc"
)

// IClusterHandler handler接口
type IClusterHandler interface {
	Register(rpcServer *grpc.Server)
}

// Cluster cluster
type Cluster struct {
	ln          net.Listener
	p           virgo.IProcedure
	mapNode     map[int32]map[int32]INode
	serverType  int32
	serverID    int32
	arrServerID []int32
	ip          string
	authKey     string
	rpcServer   *grpc.Server
	rwMtx       sync.RWMutex
}

// NewCluster 创建
func NewCluster(p virgo.IProcedure, serverType int32, arrServerID []int32, ip string, authKey string) *Cluster {
	pObj := new(Cluster)
	pObj.mapNode = make(map[int32]map[int32]INode)
	pObj.p = p
	pObj.serverType = serverType
	pObj.serverID = arrServerID[0]
	pObj.arrServerID = arrServerID
	pObj.ip = ip
	pObj.authKey = authKey
	pObj.rpcServer = grpc.NewServer()
	return pObj
}

// SetHandler 设置handler
func (c *Cluster) SetHandler(h IClusterHandler) {
	h.Register(c.rpcServer)
}

// Start 启动
func (c *Cluster) Start() error {
	logger.Info("start cluster")
	var err error
	c.ln, err = net.Listen("tcp", c.ip)
	if err != nil {
		logger.Errorf("start cluster error: %v", err)
		return err
	}

	go func() {
		err = c.rpcServer.Serve(c.ln)
		if err != nil {
			logger.Errorf("start cluster error: %v", err)
		}
	}()

	logger.Infof("cluster listen on %s", c.ip)
	return nil
}

// Stop 关闭
func (c *Cluster) Stop() {
	logger.Infof("stop cluster")
	c.rpcServer.Stop()
	c.ln.Close()
	logger.Infof("stop cluster finish")
}

func (c *Cluster) GrabAllNode() map[int32]map[int32]INode {
	c.rwMtx.RLock()
	return c.mapNode
}

func (c *Cluster) ReleaseAllNode() {
	c.rwMtx.RUnlock()
}

func (c *Cluster) GetNode(serverType int32, serverId int32) INode {
	c.rwMtx.RLock()
	defer c.rwMtx.RUnlock()
	mapType, exist := c.mapNode[serverType]
	if !exist || mapType == nil {
		return nil
	}
	pNode := mapType[serverId]
	return pNode
}

func (c *Cluster) AddNode(serverType int32, arrServerId []int32, ip string) {
	c.rwMtx.Lock()
	defer c.rwMtx.Unlock()

	mapType, exist := c.mapNode[serverType]
	if !exist || mapType == nil {
		mapType = make(map[int32]INode)
		c.mapNode[serverType] = mapType
	}

	majorServerId := arrServerId[0]
	pNode, exist := mapType[majorServerId]

	if exist && pNode != nil {
		if pNode.isDuplicate(arrServerId, ip) {
			return
		}
	}

	switch serverType {
	case NodeGame:
		pNode = NewGameNode(c, serverType, arrServerId, ip, c.authKey)
	case NodeMaster:
		pNode = NewMasterNode(c, serverType, arrServerId, ip, c.authKey)
	case NodeLogin:
		pNode = NewLoginNode(c, serverType, arrServerId, ip, c.authKey)
	}

	for _, serverId := range arrServerId {
		pOldNode, exist := mapType[serverId]
		if !exist {
			mapType[serverId] = pNode
		} else {
			_, oldServerIds, _ := pOldNode.GetServerInfo()
			for _, oldId := range oldServerIds {
				delete(mapType, oldId)
			}
			pOldNode.close()
		}
	}

	return
}
