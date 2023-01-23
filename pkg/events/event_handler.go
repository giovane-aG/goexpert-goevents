package events

import "fmt"

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (e *EventHandler) Handle(event EventInterface) {
	fmt.Println(event.GetName())
}
