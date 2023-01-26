package events

import (
	"errors"
)

var (
	EventHasAlreadyBeenRegistered = "this handler has already been registered to this event"
)

type EventManager struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventManager) Register(eventName string, handler *EventHandlerInterface) error {
	if _, hasRegisteredEvent := e.handlers[eventName]; hasRegisteredEvent {
		for _, savedHandler := range e.handlers[eventName] {
			if savedHandler == *handler {
				return errors.New(EventHasAlreadyBeenRegistered)
			}
		}
	}

	e.handlers[eventName] = append(e.handlers[eventName], *handler)
	return nil
}

func (e *EventManager) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}
