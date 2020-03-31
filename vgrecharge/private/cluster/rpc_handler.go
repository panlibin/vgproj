package cluster

import (
	"context"
	"vgproj/proto/globalrpc"
	"vgproj/proto/rechargerpc"
	"vgproj/vgrecharge/public"

	"google.golang.org/grpc"
)

// Server cluster
type Server struct {
}

// Register 注册
func (s *Server) Register(rpcServer *grpc.Server) {
	rechargerpc.RegisterRechargeServer(rpcServer, s)
}

func (s *Server) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}

// NotifyLogout 处理角色下线
func (s *Server) CreateOrder(ctx context.Context, req *rechargerpc.ReqCreateOrder) (rsp *rechargerpc.RspCreateOrder, err error) {
	rsp = &rechargerpc.RspCreateOrder{}

	return
}
