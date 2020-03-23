package igame

import "github.com/panlibin/virgo/util/vgevent"

const (
	EventType_DailyRefresh vgevent.EventType = iota
)

type EventDailyRefresh struct {
	RefreshTs int64
}

func (e *EventDailyRefresh) GetType() vgevent.EventType {
	return EventType_DailyRefresh
}
