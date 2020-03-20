package cluster

import (
	"context"
	"vgproj/proto/globalrpc"
	"vgproj/proto/loginrpc"

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

// NotifyLogout 处理角色下线
func (s *Server) PlayerLogout(context.Context, *loginrpc.NotifyLogout) (*globalrpc.Nop, error) {
	logger.Debug("NotifyLogout")

	return &globalrpc.Nop{}, nil
}
