package recharge

import (
	"errors"
	"sync/atomic"
	"time"
	"vgproj/vgrecharge/public"

	logger "github.com/panlibin/vglog"
)

const (
	SerialIdMask = 0xFFFFFFFF
	TimeOffset   = 32
	TimeMask     = 0xFFFFFFFF
)

var (
	ErrCreatePlatformImplementFail = errors.New("create platform implement fail!")
)

type platformImplement interface {
	verify(accountId int64, serverId int32, playerId int64, pfProductId string, jsonParams []byte) error
}

type RechargeManager struct {
	mapSdkParam          map[int32]*sdkParam
	mapPlatformImplement map[int32]platformImplement
	serialId             uint32
}

func NewRechargeManager() *RechargeManager {
	return &RechargeManager{
		mapSdkParam:          make(map[int32]*sdkParam, 8),
		mapPlatformImplement: make(map[int32]platformImplement, 8),
	}
}

func (rm *RechargeManager) Init() error {
	rows, err := public.Server.GetDataDb().Query(0, "select * from sdk_param")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		p := &sdkParam{
			keys: make([]string, 4),
		}
		err = rows.Scan(&p.pfId, &p.name, &p.appId, &p.keys[0], &p.keys[1], &p.keys[2], &p.keys[3])
		if err != nil {
			return err
		}
		rm.mapSdkParam[p.pfId] = p
	}

	for _, p := range rm.mapSdkParam {
		var pi platformImplement
		switch p.pfId {
		case platformApple:
			pi = newAppleIAP(p)
		case platformGooglePlay:
			pi = newGooglePlay(p)
		default:
			logger.Warningf("RechargeManager init unknown platform id %d, name %s.", p.pfId, p.name)
		}
		if pi == nil {
			return ErrCreatePlatformImplementFail
		}
	}

	return nil
}

func (rm *RechargeManager) genOrderId() uint64 {
	s := atomic.AddUint32(&rm.serialId, 1)
	return (uint64(time.Now().Unix()) << TimeOffset) + uint64(s)
}
