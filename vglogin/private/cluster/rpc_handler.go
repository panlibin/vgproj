package cluster

import (
	"context"
	"vgproj/proto/globalrpc"
	"vgproj/proto/loginrpc"
	"vgproj/vglogin/public"

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
	public.Server.GetCluster().AddNode(req.Info.ServerType, req.Info.ServerId, req.Info.Ip)

	return &globalrpc.Nop{}, nil
}

// NotifyLogout 处理角色下线
func (s *Server) PlayerLogout(ctx context.Context, req *loginrpc.NotifyLogout) (rsp *globalrpc.Nop, err error) {
	rsp = &globalrpc.Nop{}
	pAccount := public.Server.GetAccountManager().GetAccount(req.AccountId)
	if pAccount == nil {
		return
	}
	if pAccount.Lock() != nil {
		return
	}
	defer pAccount.Unlock()
	if req.PlayerId > 0 {
		pAccount.SetCharacter(req.PlayerId, req.ServerId, req.Name, req.Combat)
	}

	return
}

// func (s *Server) Login(ctx context.Context, req *loginrpc.ReqLogin) (rsp *loginrpc.RspLogin, err error) {
// 	rsp = &loginrpc.RspLogin{}

// 	pAccount := public.Server.GetAccountManager().GetAccount(req.AccountId)
// 	if pAccount == nil {
// 		rsp.Code = ec.AccountNotFound
// 		return
// 	}
// 	if pAccount.Lock() != nil {
// 		rsp.Code = ec.Unknown
// 		return
// 	}
// 	defer pAccount.Unlock()

// 	if req.Token == "" || pAccount.GetToken() != req.Token {
// 		rsp.Code = ec.InvalidToken
// 		return
// 	}
// 	curTs := vgtime.Now()
// 	if pAccount.GetTokenExpireTs() < curTs {
// 		rsp.Code = ec.InvalidToken
// 		return
// 	}

// 	if pAccount.IsBan() {
// 		rsp.Code = ec.AccountBanned
// 		return
// 	}

// 	pAccount.GenRnd()
// 	onlineServer := pAccount.GetOnlineServer()
// 	pAccount.Login(req.ServerId)

// 	if onlineServer > 0 {
// 		pNode := public.Server.GetCluster().GetNode(cluster.NodeGame, onlineServer)
// 		if pNode != nil {
// 			pGameNode := pNode.(*cluster.GameNode)
// 			pGameNode.Kick(context.Background(), &gamerpc.NotifyKick{AccountId: req.AccountId})
// 		}
// 	}

// 	rsp.Rnd = pAccount.GetRnd()
// 	rsp.Code = ec.Success

// 	return
// }
