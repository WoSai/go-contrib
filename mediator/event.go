package mediator

import (
	"context"
	"sync/atomic"
)

type (
	// EventType 事件类型描述
	EventType string

	// EventHandler 事件处理函数
	EventHandler func(context.Context, Event)

	// Event 事件接口
	Event interface {
		Type() EventType
	}

	EventCollection struct {
		events []Event
		raised int32
	}
)

func NewEventCollection() *EventCollection {
	return &EventCollection{events: make([]Event, 0)}
}

func (es *EventCollection) Add(ev Event) {
	if atomic.LoadInt32(&es.raised) == 0 {
		es.events = append(es.events, ev)
	}
}

func (es *EventCollection) Raise(ctx context.Context, m Mediator) {
	if atomic.CompareAndSwapInt32(&es.raised, 0, 1) {
		for _, event := range es.events {
			m.Dispatch(ctx, event)
		}
	}
}
