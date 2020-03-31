package cluster

import (
	"context"
	"time"
	"vgproj/proto/globalrpc"
	"vgproj/proto/loginrpc"
)

// LoginNode 登录服节点
type LoginNode struct {
	*Node
	loginrpc.LoginClient
}

// NewLoginNode 创建登陆节点
func NewLoginNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *LoginNode {
	pObj := &LoginNode{
		Node: NewNode(pCluster, serverType, serverID, ip, authKey),
	}
	pObj.LoginClient = loginrpc.NewLoginClient(pObj.cc)

	reqAuth := &globalrpc.NotifyServerAuth{
		Info: &globalrpc.ServerInfo{
			ServerType: pCluster.serverType,
			ServerId:   pCluster.arrServerID,
			Ip:         pCluster.ip,
		},
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
