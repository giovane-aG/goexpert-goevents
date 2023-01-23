package events

import "time"

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GeyPayload() interface{}
}

type EventHandler interface {
	handle(event Event)
}

type EventManager interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(eventName string, handler EventHandler) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear() error
}
