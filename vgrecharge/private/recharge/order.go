package recharge

import (
	"sync"
	"time"
	"vgproj/vgrecharge/public"
)

const (
	orderStatusWaitPay     = 0
	orderStatusWaitDeliver = 1
	orderStatusDone        = 2
	orderStatusException   = 3
)

type order struct {
	localOrderId   uint64
	pfId           int32
	pfOrderId      string
	receiveDate    time.Time
	source         string
	currency       string
	amount         int64
	pfProductId    string
	localProductId int32
	accountId      int64
	serverId       int32
	playerId       int64
	status         int32
	sandbox        int32

	mtx sync.Mutex
}

func (o *order) insert() (err error) {
	_, err = public.Server.GetDataDb().Exec(uint32(o.localOrderId), "insert into recharge_order values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		o.localOrderId, o.pfId, o.pfOrderId, o.receiveDate, o.source, o.currency, o.amount, o.pfProductId, o.localProductId, o.accountId,
		o.serverId, o.playerId, o.status, o.sandbox)
	return
}

func (o *order) update() (err error) {
	_, err = public.Server.GetDataDb().Exec(uint32(o.localOrderId), "update recharge_order set pf_order_id=?,receive_date=?,source=?"+
		",currency=?,amount=?,status=?,sandbox=? where local_order_id=?", o.pfOrderId, o.receiveDate, o.source, o.currency, o.amount,
		o.status, o.sandbox, o.localOrderId)
	return
}
