package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GeyPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventManagerInterface interface {
	Register(eventName string, handler *EventHandlerInterface) error
	Dispatch(eventName string, handler EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
