package recharge

import (
	"sync/atomic"
	"time"
)

const (
	platformApple      = 1
	platformGooglePlay = 2
)

type sdkParam struct {
	pfId  int32
	name  string
	appId string
	keys  []string
}

var serialId uint32

func genOrderId() uint64 {
	s := atomic.AddUint32(&serialId, 1)
	return (uint64(time.Now().Unix()) << TimeOffset) + uint64(s)
}
