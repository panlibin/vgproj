package cluster

import (
	"context"
	"errors"
	"vgproj/proto/globalrpc"
	"vgproj/proto/loginrpc"
	"vgproj/vglogin/public"

	logger "github.com/panlibin/vglog"
	"google.golang.org/grpc"
)

// Server cluster
type Server struct {
}

// Register 注册
func (s *Server) Register(rpcServer *grpc.Server) {
	loginrpc.RegisterLoginServer(rpcServer, s)
}

func (s *Server) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	if req.Token != public.Server.GetAuthKey() {
		return &globalrpc.Nop{}, errors.New("err token")
	}
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}

// NotifyLogout 处理角色下线
func (s *Server) PlayerLogout(context.Context, *loginrpc.NotifyLogout) (*globalrpc.Nop, error) {
	logger.Debug("NotifyLogout")

	return &globalrpc.Nop{}, nil
}
