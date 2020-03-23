package account

import (
	"math/rand"
	"time"
	"vgproj/vglogin/public"

	logger "github.com/panlibin/vglog"
)

type Name struct {
	loginType  int32
	name       string
	accountId  int64
	createTime time.Time
}

func NewName(loginType int32, name string) *Name {
	pObj := new(Name)
	pObj.loginType = loginType
	pObj.name = name
	return pObj
}

func (n *Name) loadData() error {
	row := public.Server.GetDataDb().QueryRow(rand.Uint32(), "select account_id,create_time from account_name where login_type=? and account_name=?",
		n.loginType, n.name)

	if err := row.Scan(&n.accountId, &n.createTime); err != nil {
		// logger.Error(err)
		return err
	}

	return nil
}

func (n *Name) insert() error {
	_, err := public.Server.GetDataDb().Exec(uint32(n.accountId), "insert into account_name values(?,?,?,?)",
		n.loginType, n.name, n.accountId, n.createTime)
	if err != nil {
		logger.Error(err)
	}
	return err
}
