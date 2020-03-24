package iplayer

import igate "vgproj/vggame/public/gate"

type IPlayerManager interface {
	GetPlayerIdByName(name string) int64
	GetPlayerIdByAccountId(accountId int64, serverId int32) int64
	GetPlayer(playerId int64, ctx interface{}, cb func(interface{}, IPlayer))
	GetPlayers(arrPlayerId []int64, ctx interface{}, cb func(interface{}, map[int64]IPlayer))
	GetOnlinePlayers() map[int64]IPlayer
	CreatePlayer(accountId int64, serverId int32, name string, head int32, ctx interface{}, cb func(ctx interface{}, pPlayer IPlayer, errCode int32))
	SetPlayerOnline(playerId int64)
	SetPlayerOffline(playerId int64)
	GetOnlinePlayerCount() int32
	GetTotalPlayerCount() int32
	GetMessageRouter() igate.IMessageRouter
}
