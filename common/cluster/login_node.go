package cluster

import "vgproj/proto/loginrpc"

// LoginNode 登录服节点
type LoginNode struct {
	*Node
	rpcClient loginrpc.LoginClient
	loginrpc.LoginClient
}

// NewLoginNode 创建登陆节点
func NewLoginNode(pCluster *Cluster, serverType int32, serverID []int32, ip string, authKey string) *LoginNode {
	pObj := &LoginNode{
		Node: NewNode(pCluster, serverType, serverID, ip, authKey),
	}
	pObj.LoginClient = loginrpc.NewLoginClient(pObj.cc)
	return pObj
}

// GetRPCClient 获取RPC
func (ln *LoginNode) GetRPCClient() loginrpc.LoginClient {
	return ln.rpcClient
}
