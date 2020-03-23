package icluster

type INodeManager interface {
	AddNode(serverType int32, serverId []int32, ip string)
	RemoveNode(serverType int32, serverId []int32)
}
