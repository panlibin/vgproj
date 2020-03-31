package cluster

import (
	"context"
	"time"
	"vgproj/proto/gamerpc"
	"vgproj/proto/globalrpc"
)

type GameNode struct {
	*Node
	gamerpc.GameClient
}

func NewGameNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *GameNode {
	pObj := &GameNode{
		Node: NewNode(pCluster, serverType, serverID, ip, authKey),
	}
	pObj.GameClient = gamerpc.NewGameClient(pObj.cc)

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
