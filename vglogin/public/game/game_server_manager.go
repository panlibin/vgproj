package igame

type GameServer struct {
	ServerId int32
	Name     string
	Status   int32
	Addr     string
}

type IGameServerManager interface {
	GrabServerList() map[int32]*GameServer
	ReleaseServerList()
	AddServer(serverId int32, name string, status int32, addr string)
	ModifyServer(serverId int32, name string, status int32, addr string)
}
