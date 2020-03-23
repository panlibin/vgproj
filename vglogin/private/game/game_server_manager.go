package game

import (
	"database/sql"
	"sync"
	"vgproj/vglogin/public"
	igame "vgproj/vglogin/public/game"

	logger "github.com/panlibin/vglog"
)

type GameServerManager struct {
	mapServer map[int32]*igame.GameServer
	rwMtx     sync.RWMutex
}

func NewGameServerManager() *GameServerManager {
	pObj := new(GameServerManager)
	pObj.mapServer = make(map[int32]*igame.GameServer)

	return pObj
}

func (g *GameServerManager) Init() error {
	var err error
	for {
		var rows *sql.Rows
		rows, err = public.Server.GetDataDb().Query(g.getDbIdx(), "select server_id,name,status,addr from server_list")
		if err != nil {
			break
		}

		for rows.Next() {
			pGameServer := new(igame.GameServer)
			err = rows.Scan(&pGameServer.ServerId, &pGameServer.Name, &pGameServer.Status, &pGameServer.Addr)
			if err != nil {
				break
			}
			g.mapServer[pGameServer.ServerId] = pGameServer
		}
		rows.Close()
		if err != nil {
			break
		}

		break
	}

	if err != nil {
		logger.Errorf("game server manager init error: %v", err)
	}

	return err
}

func (g *GameServerManager) GetServer(id int32) *igame.GameServer {
	g.rwMtx.RLock()
	defer g.rwMtx.RUnlock()
	pServerInfo, exist := g.mapServer[id]
	if !exist || pServerInfo == nil {
		return nil
	}
	return pServerInfo
}

func (g *GameServerManager) GrabServerList() map[int32]*igame.GameServer {
	g.rwMtx.RLock()
	return g.mapServer
}

func (g *GameServerManager) ReleaseServerList() {
	g.rwMtx.RUnlock()
}

func (g *GameServerManager) AddServer(serverId int32, name string, status int32, addr string) {
	pGameServer := new(igame.GameServer)
	pGameServer.ServerId = serverId
	pGameServer.Name = name
	pGameServer.Status = status
	pGameServer.Addr = addr
	g.rwMtx.Lock()
	g.mapServer[pGameServer.ServerId] = pGameServer
	g.rwMtx.Unlock()
	g.insertServer(pGameServer)
}

func (g *GameServerManager) ModifyServer(serverId int32, name string, status int32, addr string) {
	g.rwMtx.Lock()
	pGameServer, exist := g.mapServer[serverId]
	if !exist {
		g.rwMtx.Unlock()
		return
	}
	pGameServer.Name = name
	pGameServer.Status = status
	pGameServer.Addr = addr
	g.rwMtx.Unlock()
	g.updateServer(pGameServer)
}

func (g *GameServerManager) insertServer(pGameServer *igame.GameServer) {
	const strInsertSql = "insert into server_list values(?,?,?,?)"
	public.Server.GetDataDb().AsyncExec(nil, g.getDbIdx(), strInsertSql, pGameServer.ServerId, pGameServer.Name, pGameServer.Status, pGameServer.Addr)
}

func (g *GameServerManager) updateServer(pGameServer *igame.GameServer) {
	const strUpdateSql = "update server_list set name=?,status=?,addr=? where server_id=?"
	public.Server.GetDataDb().AsyncExec(nil, g.getDbIdx(), strUpdateSql, pGameServer.Name, pGameServer.Status, pGameServer.Addr, pGameServer.ServerId)
}

func (g *GameServerManager) getDbIdx() uint32 {
	return 0
}
