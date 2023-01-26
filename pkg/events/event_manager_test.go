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

	eventHandler1 EventHandlerInterface
	eventHandler2 EventHandlerInterface
	eventHandler3 EventHandlerInterface

	eventManager *EventManager
}

func (suite *EventManagerTestSuite) SetupTest() {
	suite.event1 = NewEvent("test event", "test payload")
	suite.event3 = NewEvent("event3", "payload3")
	suite.eventHandler1 = NewEventHandler()
	suite.eventHandler2 = NewEventHandler()
	suite.eventManager = NewEventManager()
}

// function to run the test suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventManagerTestSuite))
}

func (suite *EventManagerTestSuite) Test_Register() {

	err := suite.eventManager.Register(suite.event1.GetName(), &suite.eventHandler1)
	err2 := suite.eventManager.Register(suite.event1.GetName(), &suite.eventHandler2)

	suite.Nil(err)
	suite.Nil(err2)
	suite.Equal(len(suite.eventManager.handlers[suite.event1.GetName()]), 2)
	suite.Equal(suite.eventManager.handlers[suite.event1.GetName()][0], suite.eventHandler1)
	suite.Equal(suite.eventManager.handlers[suite.event1.GetName()][1], suite.eventHandler2)
}

func (suite *EventManagerTestSuite) Test_RegisterSameHandler() {
	pointer := &suite.eventHandler1
	err := suite.eventManager.Register(suite.event1.GetName(), pointer)
	err2 := suite.eventManager.Register(suite.event1.GetName(), pointer)

	suite.Nil(err)
	suite.NotNil(err2)

}

func (suite *EventManagerTestSuite) Test_Clear() {
	err := suite.eventManager.Register(suite.event3.GetName(), &suite.eventHandler3)
	suite.eventManager.Clear()

	suite.Nil(err)
	suite.Equal(len(suite.eventManager.handlers), 0)
}
