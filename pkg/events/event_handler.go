package events

import (
	"fmt"
	"sync"
)

type EventHandler struct {
	ID int8
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (e *EventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	fmt.Println(event.GetName())
	wg.Done()
}
