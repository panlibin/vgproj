package iaccount

// IAccountManager 账号管理接口
type IAccountManager interface {
	GetAccount(accountID int64) IAccount
	GetAccountByName(loginType int32, name string) IAccount
	Register(name string, pwd string) int32
	Login(loginType int32, name string, password string) (IAccount, int32)
}
