package cluster

import (
	"context"
	"vgproj/proto/globalrpc"
	"vgproj/proto/masterrpc"
	"vgproj/vgmaster/public"

	// logger "github.com/panlibin/vglog"
	logger "github.com/panlibin/vglog"
	"google.golang.org/grpc"
)

// Server cluster
type Server struct {
}

// Register 注册
func (s *Server) Register(rpcServer *grpc.Server) {
	masterrpc.RegisterMasterServer(rpcServer, s)
}

func (s *Server) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	public.Server.GetNodeManager().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	logger.Infof("new node type: %d, id: %v", req.Info.ServerType, req.Info.ServerId)

	return &globalrpc.Nop{}, nil
}

func (s *Server) GetServerList(context.Context, *globalrpc.ReqServerList) (rsp *globalrpc.RspServerList, err error) {
	rsp = &globalrpc.RspServerList{}
	pCluster := public.Server.GetCluster()
	mapAllNode := pCluster.GrabAllNode()
	defer pCluster.ReleaseAllNode()
	for _, mapTypeNode := range mapAllNode {
		for _, pNode := range mapTypeNode {
			pNodeInfo := new(globalrpc.ServerInfo)
			pNodeInfo.ServerType, pNodeInfo.ServerId, pNodeInfo.Ip = pNode.GetServerInfo()
			rsp.List = append(rsp.List, pNodeInfo)
		}
	}

	return
}

func (s *Server) GrabPlayerName(context.Context, *masterrpc.ReqGrabPlayerName) (*masterrpc.RspGrabPlayerName, error) {
	return &masterrpc.RspGrabPlayerName{}, nil
}

func (s *Server) GrabGuildName(context.Context, *masterrpc.ReqGrabGuildName) (*masterrpc.RspGrabGuildName, error) {
	return &masterrpc.RspGrabGuildName{}, nil
}

func (s *Server) ReleasePlayerName(context.Context, *masterrpc.NotifyReleasePlayerName) (*globalrpc.Nop, error) {
	return &globalrpc.Nop{}, nil
}

func (s *Server) ReleaseGuildName(context.Context, *masterrpc.NotifyReleaseGuildName) (*globalrpc.Nop, error) {
	return &globalrpc.Nop{}, nil
}
