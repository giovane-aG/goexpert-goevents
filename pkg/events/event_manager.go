package events

import (
	"errors"
	"sync"
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
	wg := &sync.WaitGroup{}
	wg.Add(len(e.handlers[event.GetName()]))

	if handlers, ok := e.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			go handler.Handle(event, wg)
		}
	}

	wg.Wait()
	return nil
}

func (e *EventManager) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := e.handlers[eventName]; ok {
		for i, savedHandler := range handlers {
			if savedHandler == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
			}
		}
	}

	return nil
}
