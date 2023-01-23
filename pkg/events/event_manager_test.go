package events

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// defines a test suite
type EventManagerTestSuite struct {
	suite.Suite

	event1 *Event
	event2 *Event
	event3 *Event

	eventHandler1 *EventHandler
	eventHandler2 *EventHandler
	eventHandler3 *EventHandler

	eventManager *EventManager
}

func (suite *EventManagerTestSuite) SetupTest() {
	suite.event1 = NewEvent("test event", "test payload")
	suite.eventHandler1 = NewEventHandler()
	suite.eventManager = NewEventManager()
}

// function to run the test suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventManagerTestSuite))
}
