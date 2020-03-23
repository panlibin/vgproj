package cluster

import (
	"context"
	"errors"
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
	if req.Token != public.Server.GetAuthKey() {
		return &globalrpc.Nop{}, errors.New("err token")
	}
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}
