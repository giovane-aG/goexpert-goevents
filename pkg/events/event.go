package events

import "time"

type Event struct {
	name         string
	payload      interface{}
	registeredAt time.Time
}

func NewEvent(name string, payload interface{}) *Event {
	return &Event{
		name:         name,
		payload:      payload,
		registeredAt: time.Now(),
	}
}

func (e *Event) GetName() string {
	return e.name
}
func (e *Event) GetDateTime() time.Time {
	return e.registeredAt
}
func (e *Event) GetPayload() interface{} {
	return e.payload
}
