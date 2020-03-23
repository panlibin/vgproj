package client

const (
	ClientStatus_LoginAccount int32 = iota
	ClientStatus_RegisterAccount
	ClientStatus_GetServerInfo
	ClientStatus_LoginGame
	ClientStatus_CreateCharacter
	ClientStatus_Idle
	ClientStatus_Disconnect
	ClientStatus_Close
	ClientStatus_Count
)
