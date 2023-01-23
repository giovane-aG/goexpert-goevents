package events

import "time"

type Event struct {
	Name         string
	Payload      interface{}
	RegisteredAt time.Time
}

func NewEvent(name string, payload interface{}) *Event {
	return &Event{
		Name:         name,
		Payload:      payload,
		RegisteredAt: time.Now(),
	}
}

func (e *Event) GetName() string {
	return e.Name
}
func (e *Event) GetDateTime() time.Time {
	return e.RegisteredAt
}
func (e *Event) GetPayload() interface{} {
	return e.Payload
}
