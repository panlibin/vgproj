package account

import (
	"database/sql"
	"sync"
	"time"
	ec "vgproj/common/define/err_code"
	"vgproj/vglogin/public"
	iaccount "vgproj/vglogin/public/account"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgstr"
)

const MinAccountId = 10000
const MinCustomAccountNameLength = 6
const MaxCustomAccountNameLength = 20

type AccountManager struct {
	mapAccount       map[int64]*Account
	mapAccountByName map[int32]map[string]*Account
	maxAccountId     int64
	rwMtx            sync.RWMutex

	pFacebook   *Facebook
	pGooglePlay *GooglePlay
}

func NewAccountManager() *AccountManager {
	pObj := new(AccountManager)
	pObj.mapAccount = make(map[int64]*Account, 512)
	pObj.mapAccountByName = make(map[int32]map[string]*Account)

	pObj.pFacebook = NewFacebook()
	pObj.pGooglePlay = NewGooglePlay()
	return pObj
}

func (am *AccountManager) LoadData() error {
	rows, err := public.Server.GetDataDb().Query(0, "select max(account_id) from account_info")
	if err != nil {
		return err
	}

	if rows.Next() {
		var tmp sql.NullInt64
		if err = rows.Scan(&tmp); err != nil {
			return err
		}
		if tmp.Valid {
			am.maxAccountId = tmp.Int64
		}
	}

	if am.maxAccountId < MinAccountId {
		am.maxAccountId = MinAccountId
	}

	rows.Close()

	return nil
}

func (am *AccountManager) Register(name string, pwd string) int32 {
	if !am.isValidAccountLength(name) {
		return ec.InvalidAccountLength
	}
	if len(pwd) > 64 {
		return ec.InvalidAccountPassword
	}
	if !vgstr.IsAlphanumericOrUnderscore(name) {
		return ec.InvalidAccountName
	}

	pAccount := am.getAccountByName(iaccount.LoginTypeCustom, name)
	if pAccount != nil {
		return ec.DuplicateAccountName
	}

	_, errCode := am.createAccount(iaccount.LoginTypeCustom, name, pwd)

	return errCode
}

func (am *AccountManager) Login(loginType int32, name string, password string) (iaccount.IAccount, int32) {
	if name == "" {
		return nil, ec.InvalidParam
	}

	var pAccount *Account
	if loginType == iaccount.LoginTypeCustom {
		pAccount = am.getAccountByName(loginType, name)
		if pAccount == nil {
			return nil, ec.AccountNotFound
		}
		if err := pAccount.Lock(); err != nil {
			logger.Error(err)
			return nil, ec.Unknown
		}
		defer pAccount.Unlock()

		if pAccount.password != password {
			return pAccount, ec.WrongPassword
		}
	} else {
		var err error
		switch loginType {
		case iaccount.LoginTypeFacebook:
			name, err = am.pFacebook.login(password)
		case iaccount.LoginTypeGooglePlay:
			name, err = am.pGooglePlay.login(password)
		}

		if err != nil {
			return nil, ec.InvalidToken
		}

		pAccount = am.getAccountByName(loginType, name)
		if pAccount == nil {
			var errCode int32
			pAccount, errCode = am.createAccount(loginType, name, "")
			if errCode != ec.Success && errCode != ec.DuplicateAccountName {
				return nil, errCode
			}
		} else {
			if err := pAccount.Lock(); err != nil {
				logger.Error(err)
				return nil, ec.Unknown
			}
			defer pAccount.Unlock()
		}
	}

	pAccount.genToken()

	return pAccount, ec.Success
}

func (am *AccountManager) GetAccountByName(loginType int32, name string) iaccount.IAccount {
	pAccount := am.getAccountByName(loginType, name)
	if pAccount == nil {
		return nil
	}
	return pAccount
}

func (am *AccountManager) GetAccount(accountId int64) iaccount.IAccount {
	pAccount := am.getAccount(accountId)
	if pAccount == nil {
		return nil
	}
	return pAccount
}

func (am *AccountManager) genAccountId() int64 {
	am.maxAccountId++
	return am.maxAccountId
}

func (am *AccountManager) createAccount(loginType int32, name string, pwd string) (*Account, int32) {
	am.rwMtx.Lock()
	pAccount := am.getAccountByNameM(loginType, name)
	if pAccount != nil {
		am.rwMtx.Unlock()
		return pAccount, ec.DuplicateAccountName
	}

	accountId := am.genAccountId()
	pAccountName := NewName(loginType, name)
	pAccountName.accountId = accountId
	pAccountName.createTime = time.Now()

	pAccount = NewAccount(accountId)
	pAccount.password = pwd
	pAccount.createTime = pAccountName.createTime
	pAccount.addName(pAccountName)
	pAccount.Lock()
	defer pAccount.Unlock()

	am.addAccount(pAccount)
	am.rwMtx.Unlock()

	if pAccount.insert() != nil || pAccountName.insert() != nil {
		am.rwMtx.Lock()
		am.removeAccount(pAccount)
		am.rwMtx.Unlock()
		return nil, ec.Unknown
	}

	return pAccount, ec.Success
}

func (am *AccountManager) addAccount(pAccount *Account) {
	_, exist := am.mapAccount[pAccount.id]
	if exist {
		return
	}
	am.mapAccount[pAccount.id] = pAccount
	for _, pAccountName := range pAccount.mapName {
		am.addAccountByName(pAccountName.loginType, pAccountName.name, pAccount)
	}
}

func (am *AccountManager) addAccountByName(loginType int32, name string, pAccount *Account) {
	mapAccountName, exist := am.mapAccountByName[loginType]
	if !exist {
		mapAccountName = make(map[string]*Account, 512)
		am.mapAccountByName[loginType] = mapAccountName
	}
	if _, exist := mapAccountName[name]; exist {
		logger.Errorf("duplicate account name. id: %d, login type: %d, name: %s", pAccount.id, loginType, name)
		return
	}
	mapAccountName[name] = pAccount
}

func (am *AccountManager) removeAccount(pAccount *Account) {
	delete(am.mapAccount, pAccount.id)
	for _, pAccountName := range pAccount.mapName {
		mapAccountName, exist := am.mapAccountByName[pAccountName.loginType]
		if exist {
			delete(mapAccountName, pAccountName.name)
		}
	}
}

func (am *AccountManager) getAccountM(accountId int64) *Account {
	pAccount, exist := am.mapAccount[accountId]
	if !exist {
		return nil
	}
	return pAccount
}

func (am *AccountManager) getAccount(accountId int64) *Account {
	am.rwMtx.RLock()
	pAccount := am.getAccountM(accountId)
	am.rwMtx.RUnlock()
	if pAccount != nil {
		return pAccount
	}

	pAccount = NewAccount(accountId)
	if pAccount.loadData() != nil {
		return nil
	}

	am.rwMtx.Lock()
	pAccountCheck := am.getAccountM(accountId)
	if pAccountCheck != nil {
		am.rwMtx.Unlock()
		return pAccountCheck
	}

	am.addAccount(pAccount)
	am.rwMtx.Unlock()

	return pAccount
}

func (am *AccountManager) getAccountByNameM(loginType int32, name string) *Account {
	mapAccountName, exist := am.mapAccountByName[loginType]
	if !exist {
		return nil
	}
	pAccount, exist := mapAccountName[name]
	if !exist {
		return nil
	}
	return pAccount
}

func (am *AccountManager) getAccountByName(loginType int32, name string) *Account {
	am.rwMtx.RLock()
	pAccount := am.getAccountByNameM(loginType, name)
	am.rwMtx.RUnlock()
	if pAccount != nil {
		return pAccount
	}

	pAccountName := NewName(loginType, name)
	if pAccountName.loadData() != nil {
		return nil
	}

	return am.getAccount(pAccountName.accountId)
}

func (am *AccountManager) isValidAccountLength(name string) bool {
	lenName := len(name)
	return lenName >= MinCustomAccountNameLength && lenName <= MaxCustomAccountNameLength
}
