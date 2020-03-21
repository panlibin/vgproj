package cluster

import (
	"context"
	"errors"
	"vgproj/proto/globalrpc"
	"vgproj/proto/masterrpc"
	"vgproj/vgmaster/public"

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
	logger.Debug(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)
	if req.Token != public.Server.GetAuthKey() {
		return &globalrpc.Nop{}, errors.New("err token")
	}
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}

func (s *Server) GetServerList(context.Context, *globalrpc.ReqServerList) (*globalrpc.RspServerList, error) {
	logger.Debug("Auth")

	return &globalrpc.RspServerList{}, nil
}

func (s *Server) GrabPlayerName(context.Context, *masterrpc.ReqGrabPlayerName) (*masterrpc.RspGrabPlayerName, error) {
	logger.Debug("Auth")

	return &masterrpc.RspGrabPlayerName{}, nil
}

func (s *Server) GrabGuildName(context.Context, *masterrpc.ReqGrabGuildName) (*masterrpc.RspGrabGuildName, error) {
	logger.Debug("Auth")

	return &masterrpc.RspGrabGuildName{}, nil
}

func (s *Server) ReleasePlayerName(context.Context, *masterrpc.NotifyReleasePlayerName) (*globalrpc.Nop, error) {
	logger.Debug("Auth")

	return &globalrpc.Nop{}, nil
}

func (s *Server) ReleaseGuildName(context.Context, *masterrpc.NotifyReleaseGuildName) (*globalrpc.Nop, error) {
	logger.Debug("Auth")

	return &globalrpc.Nop{}, nil
}
