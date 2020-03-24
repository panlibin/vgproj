package cluster

import (
	"context"
	"time"
	"vgproj/proto/globalrpc"
	"vgproj/proto/masterrpc"
)

type MasterNode struct {
	*Node
	masterrpc.MasterClient
}

func NewMasterNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *MasterNode {
	pObj := &MasterNode{
		Node: NewNode(pCluster, serverType, serverID, ip, authKey),
	}
	pObj.MasterClient = masterrpc.NewMasterClient(pObj.cc)

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

		for !pObj.quit {
			rspServerList, err := pObj.GetServerList(context.Background(), &globalrpc.ReqServerList{})
			if err == nil {
				for _, serverInfo := range rspServerList.List {
					if pObj.pCluster.ip == serverInfo.Ip || serverInfo.ServerType == NodeMaster {
						continue
					}
					pObj.pCluster.AddNode(serverInfo.ServerType, serverInfo.ServerId, serverInfo.Ip)
				}

				break
			}
			time.Sleep(time.Second * 3)
		}
	}()

	return pObj
}
