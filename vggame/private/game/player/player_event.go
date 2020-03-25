package player

import "github.com/panlibin/virgo/util/vgevent"

const (
	EventType_DailyRefresh vgevent.EventType = iota
	EventType_PlayerLevUp
	EventType_PlayerVipLevUp
)

type EventDailyRefresh struct {
	RefreshTs int64
}

func (e *EventDailyRefresh) GetType() vgevent.EventType {
	return EventType_DailyRefresh
}

type EventPlayerLevUp struct {
	CurLev int32
}

func (e *EventPlayerLevUp) GetType() vgevent.EventType {
	return EventType_PlayerLevUp
}

type EventPlayerVipLevUp struct {
	CurVipLev   int32
	DeltaVipLev int32
}

func (e *EventPlayerVipLevUp) GetType() vgevent.EventType {
	return EventType_PlayerVipLevUp
}
