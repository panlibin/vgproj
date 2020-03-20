package cluster

import (
	"time"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"google.golang.org/grpc"
)

// INode 节点接口
type INode interface {
	GetServerInfo() (serverType int32, serverID []int32, ip string)
}

// Node 节点基类
type Node struct {
	serverType    int32
	majorServerID int32
	arrServerID   []int32
	ip            string
	p             virgo.IProcedure
	pCluster      *Cluster
	authKey       string
	cc            *grpc.ClientConn
}

// NewNode 创建节点
func NewNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *Node {
	pObj := new(Node)
	pObj.serverType = serverType
	pObj.majorServerID = serverID[0]
	pObj.arrServerID = serverID
	pObj.ip = ip
	pObj.pCluster = pCluster
	pObj.p = pCluster.p
	pObj.authKey = authKey
	pObj.connect()
	return pObj
}

// Connect 连接节点
func (n *Node) connect() {
	var err error
	for {
		n.cc, err = grpc.Dial(n.ip, grpc.WithInsecure(), grpc.WithAuthority(n.authKey), grpc.WithBackoffMaxDelay(time.Second*3))
		if err != nil {
			logger.Errorf("node connect error: %v", err)
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}

// GetServerInfo 获取服务器类型,ID,IP
func (n *Node) GetServerInfo() (int32, []int32, string) {
	return n.serverType, n.arrServerID, n.ip
}
