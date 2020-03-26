package account

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
	"vgproj/vglogin/public"
	iaccount "vgproj/vglogin/public/account"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgtime"
)

type Account struct {
	id           int64
	password     string
	createTime   time.Time
	isBan        int32
	banTs        int64
	banType      int32
	banDuration  int64
	mapName      map[int32]*Name
	mapCharacter map[int32]iaccount.ICharacter
	mtx          sync.Mutex
	lastErr      error
}

func NewAccount(accountId int64) *Account {
	pObj := new(Account)
	pObj.id = accountId
	pObj.mapName = make(map[int32]*Name)
	pObj.mapCharacter = make(map[int32]iaccount.ICharacter)
	return pObj
}

func (a *Account) Lock() error {
	a.mtx.Lock()
	if a.lastErr != nil {
		a.mtx.Unlock()
		return a.lastErr
	}
	return nil
}

func (a *Account) Unlock() {
	a.mtx.Unlock()
}

func (a *Account) GetId() int64 {
	return a.id
}

func (a *Account) IsBan() bool {
	if a.isBan > 0 {
		if a.banDuration > 0 {
			curTs := vgtime.Now()
			if a.banTs+a.banDuration <= curTs {
				a.isBan = 0
				a.banTs = 0
				a.banType = 0
				a.banDuration = 0
				a.updateBanInfo()
			}
		}
	}

	return a.isBan > 0
}

func (a *Account) GetBanTs() int64 {
	return a.banTs
}

func (a *Account) GetBanType() int32 {
	return a.banType
}

func (a *Account) GetBanDuration() int64 {
	return a.banDuration
}

func (a *Account) Ban(banType int32, banDuration int64) {
	a.isBan = 1
	a.banType = banType
	a.banDuration = banDuration
	a.banTs = vgtime.Now()
	a.updateBanInfo()
}

func (a *Account) SetCharacter(playerId int64, serverId int32, name string, combat int64) {
	iCharacter, exist := a.mapCharacter[serverId]
	if !exist {
		pCharacter := new(Character)
		pCharacter.id = playerId
		pCharacter.accountId = a.id
		pCharacter.serverId = serverId
		pCharacter.name = name
		pCharacter.combat = combat
		pCharacter.updateTs = vgtime.Now()
		a.mapCharacter[serverId] = pCharacter
		pCharacter.insert()
	} else {
		pCharacter := iCharacter.(*Character)
		pCharacter.setName(name)
		pCharacter.setCombat(combat)
		pCharacter.updateTs = vgtime.Now()
		pCharacter.update()
	}
}

func (a *Account) GetCharacters() map[int32]iaccount.ICharacter {
	return a.mapCharacter
}

func (a *Account) loadData() error {
	rows, err := public.Server.GetDataDb().Query(uint32(a.id), fmt.Sprintf("select `password`,create_time,is_ban,ban_ts,ban_type,ban_duration from account_info where account_id=%d;"+
		"select login_type,account_name,create_time from account_name where account_id=%d;select player_id,server_id,name,combat from character_info where account_id=%d", a.id, a.id, a.id))

	if err != nil {
		logger.Error(err)
		a.lastErr = err
		return err
	}

	defer rows.Close()

	for {
		if !rows.Next() {
			err = sql.ErrNoRows
			break
		}

		if err = rows.Scan(&a.password, &a.createTime, &a.isBan, &a.banTs, &a.banType, &a.banDuration); err != nil {
			break
		}

		break
	}

	if err != nil {
		logger.Error(err)
		a.lastErr = err
		return err
	}

	if !rows.NextResultSet() {
		logger.Error(sql.ErrNoRows)
		a.lastErr = sql.ErrNoRows
		return sql.ErrNoRows
	}

	for rows.Next() {
		pAccountName := new(Name)
		pAccountName.accountId = a.id
		if err = rows.Scan(&pAccountName.loginType, &pAccountName.name, &pAccountName.createTime); err != nil {
			break
		}
		a.addName(pAccountName)
	}

	if err != nil {
		logger.Error(err)
		a.lastErr = err
		return err
	}

	if !rows.NextResultSet() {
		logger.Error(sql.ErrNoRows)
		a.lastErr = sql.ErrNoRows
		return sql.ErrNoRows
	}

	for rows.Next() {
		pCharacter := &Character{
			accountId: a.id,
		}
		if err = rows.Scan(&pCharacter.id, &pCharacter.serverId, &pCharacter.name, &pCharacter.combat); err != nil {
			break
		}
		a.mapCharacter[pCharacter.serverId] = pCharacter
	}

	if err != nil {
		logger.Error(err)
		a.lastErr = err
		return err
	}

	return nil
}

func (a *Account) insert() error {
	_, err := public.Server.GetDataDb().Exec(uint32(a.id), "insert into account_info(account_id,`password`,create_time,online_server) values(?,?,?,?)",
		a.id, a.password, a.createTime, 0)
	if err != nil {
		logger.Error(err)
		a.lastErr = err
	}
	return err
}

func (a *Account) updatePassword() {
	public.Server.GetDataDb().Exec(uint32(a.id), "update account_info set `password`=? where account_id=?",
		a.password, a.id)
}

func (a *Account) updateBanInfo() {
	public.Server.GetDataDb().Exec(uint32(a.id), "update account_info set is_ban=?,ban_ts=?,ban_type=?,ban_duration=? where account_id=?",
		a.isBan, a.banTs, a.banType, a.banDuration, a.id)
}

func (a *Account) addName(pName *Name) {
	if _, exist := a.mapName[pName.loginType]; exist {
		logger.Errorf("duplicate account name. id: %d, login type: %d, name: %s", a.id, pName.loginType, pName.name)
		return
	}
	a.mapName[pName.loginType] = pName
}
