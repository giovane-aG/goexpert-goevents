package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	suite.eventHandler2 = NewEventHandler()
	suite.eventManager = NewEventManager()
}

// function to run the test suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventManagerTestSuite))
}

func (suite *EventManagerTestSuite) Test_Register() {

	err := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler1)
	assert.Nil(suite.T(), err)

	err2 := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler2)
	assert.Nil(suite.T(), err2)

	assert.Equal(suite.T(), len(suite.eventManager.handlers[suite.event1.GetName()]), 2)

	assert.Equal(suite.T(), len(suite.eventManager.handlers[suite.event1.GetName()]), 2)

	assert.Equal(suite.T(), suite.eventManager.handlers[suite.event1.GetName()][0], suite.eventHandler1)
	assert.Equal(suite.T(), suite.eventManager.handlers[suite.event1.GetName()][1], suite.eventHandler2)
}
