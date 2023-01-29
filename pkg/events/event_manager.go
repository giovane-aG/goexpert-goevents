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

func (e *EventManager) Register(eventName string, handler EventHandlerInterface) error {
	if _, hasRegisteredEvent := e.handlers[eventName]; hasRegisteredEvent {
		for _, savedHandler := range e.handlers[eventName] {
			if savedHandler == handler {
				return errors.New(EventHasAlreadyBeenRegistered)
			}
		}
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

func (e *EventManager) Has(eventName string, handler *EventHandlerInterface) bool {
	if _, hasHandlers := e.handlers[eventName]; hasHandlers {
		for _, savedHandler := range e.handlers[eventName] {
			if savedHandler == *handler {
				return true
			}
		}
	}
	return false
}

func (e *EventManager) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}

func (e *EventManager) Dispatch(event EventInterface) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}

	return nil
}

func (e *EventManager) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := e.handlers[eventName]; ok {
		newHandlers := []EventHandlerInterface{}
		for _, savedHandler := range handlers {
			if savedHandler != handler {
				newHandlers = append(newHandlers, savedHandler)
			}
		}
		e.handlers[eventName] = newHandlers
	}

	return nil
}
