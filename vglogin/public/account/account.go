package iaccount

type ICharacter interface {
	GetPlayerId() int64
	GetName() string
	GetCombat() int64
	GetUpdateTs() int64
}

// IAccount 账号接口
type IAccount interface {
	Lock() error
	Unlock()
	GetId() int64
	IsBan() bool
	GetBanTs() int64
	GetBanType() int32
	GetBanDuration() int64
	Ban(banType int32, banDuration int64)
	GetToken() string
	GetTokenExpireTs() int64
	GetRnd() uint64
	GenRnd()
	GetOnlineServer() int32
	Login(serverId int32)
	Logout(rnd uint64)
	SetCharacter(playerId int64, serverId int32, name string, combat int64)
	GetCharacters() map[int32]ICharacter
}
