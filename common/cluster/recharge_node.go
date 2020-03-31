package cluster

import (
	"context"
	"time"
	"vgproj/proto/globalrpc"
	"vgproj/proto/rechargerpc"
)

// RechargeNode 充值服节点
type RechargeNode struct {
	*Node
	rechargerpc.RechargeClient
}

// NewRechargeNode 创建充值节点
func NewRechargeNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *RechargeNode {
	pObj := &RechargeNode{
		Node: NewNode(pCluster, serverType, serverID, ip, authKey),
	}
	pObj.RechargeClient = rechargerpc.NewRechargeClient(pObj.cc)

	reqAuth := &globalrpc.NotifyServerAuth{
		Info: &globalrpc.ServerInfo{
			ServerType: pCluster.serverType,
			ServerId:   pCluster.arrServerID,
			Ip:         pCluster.ip,
		},
		Token: authKey,
	}

	go func() {
		for !pObj.quit {
			_, err := pObj.Auth(context.Background(), reqAuth)
			if err == nil {
				break
			}
			time.Sleep(time.Second * 3)
		}
	}()

	return pObj
}
