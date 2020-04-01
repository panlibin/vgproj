package recharge

import (
	"sync"
	"vgproj/vgrecharge/public"
)

type platform struct {
	mapPfOrder    map[string]*order
	mapLocalOrder map[uint64]*order
	mtx           sync.RWMutex
	dbIdx         uint32
}

func newPlatform() *platform {
	return &platform{
		mapPfOrder:    make(map[string]*order, 5000),
		mapLocalOrder: make(map[uint64]*order, 5000),
	}
}

func (pf *platform) getOrderByPfId(pfId int32, pfOrderId string) *order {
	pf.mtx.RLock()
	o, exist := pf.mapPfOrder[pfOrderId]
	pf.mtx.RUnlock()
	if exist {
		return o
	}

	pf.dbIdx++
	row := public.Server.GetDataDb().QueryRow(pf.dbIdx, "select local_order_id,receive_date,source,currency,amount,pf_product_id,local_product_id,"+
		"account_id,server_id,player_id,status,sandbox from recharge_order where pf_id=? and pf_order_id=?", pfId, pfOrderId)
	o = &order{
		pfId:      pfId,
		pfOrderId: pfOrderId,
	}
	err := row.Scan(&o.localOrderId, o.receiveDate, o.source, o.currency, o.amount, o.pfProductId, o.localProductId, o.accountId,
		o.serverId, o.playerId, o.status, o.sandbox)

	pf.mtx.Lock()
	defer pf.mtx.Unlock()
	checkOrder, exist := pf.mapPfOrder[pfOrderId]
	if exist {
		return checkOrder
	}
	checkOrder, exist = pf.mapLocalOrder[o.localOrderId]
	if exist {
		return checkOrder
	}

	if err != nil {
		return nil
	}

	pf.mapPfOrder[pfOrderId] = o
	pf.mapLocalOrder[o.localOrderId] = o

	return o
}

func (pf *platform) getOrderByLocalId(localOrderId uint64) *order {
	pf.mtx.RLock()
	o, exist := pf.mapLocalOrder[localOrderId]
	pf.mtx.RUnlock()
	if exist {
		return o
	}

	row := public.Server.GetDataDb().QueryRow(uint32(localOrderId), "select pf_id,pf_order_id,receive_date,source,currency,amount,pf_product_id,local_product_id,"+
		"account_id,server_id,player_id,status,sandbox from recharge_order where local_order_id=?", localOrderId)
	o = &order{
		localOrderId: localOrderId,
	}
	err := row.Scan(&o.pfId, &o.pfOrderId, o.receiveDate, o.source, o.currency, o.amount, o.pfProductId, o.localProductId, o.accountId,
		o.serverId, o.playerId, o.status, o.sandbox)

	pf.mtx.Lock()
	defer pf.mtx.Unlock()
	checkOrder, exist := pf.mapLocalOrder[localOrderId]
	if exist {
		return checkOrder
	}
	checkOrder, exist = pf.mapPfOrder[o.pfOrderId]
	if exist {
		return checkOrder
	}

	if err != nil {
		return nil
	}

	pf.mapLocalOrder[localOrderId] = o
	pf.mapPfOrder[o.pfOrderId] = o

	return o
}

func (pf *platform) insertOrder(o *order) bool {

	return false
}
