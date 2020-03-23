package iplayer

type IPlayerManager interface {
	GetPlayerIdByName(name string) int64
	GetPlayerIdByAccountId(accountId int64, serverId int32) int64
	GetPlayer(playerId int64) IPlayer
	GetOnlinePlayers() map[int64]IPlayer
	CreatePlayer(accountId int64, serverId int32, name string, head int32) (int32, IPlayer)
	SetPlayerOnline(playerId int64)
	SetPlayerOffline(playerId int64)
	GetOnlinePlayerCount() int32
	GetTotalPlayerCount() int32
}
