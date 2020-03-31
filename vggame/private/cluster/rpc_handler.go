package cluster

import (
	"context"
	"vgproj/proto/gamerpc"
	"vgproj/proto/globalrpc"
	"vgproj/vggame/public"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Register(rpcServer *grpc.Server) {
	gamerpc.RegisterGameServer(rpcServer, s)
}

func (s *Server) Auth(ctx context.Context, req *globalrpc.NotifyServerAuth) (*globalrpc.Nop, error) {
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}

func (s *Server) Kick(ctx context.Context, req *gamerpc.NotifyKick) (rsp *globalrpc.Nop, err error) {
	rsp = &globalrpc.Nop{}

	public.Server.GetGate().Kick(req.AccountId)

	return
}
