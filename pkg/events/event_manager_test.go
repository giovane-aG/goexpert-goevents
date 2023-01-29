package events

import (
	"testing"

	"github.com/stretchr/testify/mock"
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

	err := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler1)
	err2 := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler2)

	suite.Nil(err)
	suite.Nil(err2)
	suite.Equal(len(suite.eventManager.handlers[suite.event1.GetName()]), 2)
	suite.Equal(suite.eventManager.handlers[suite.event1.GetName()][0], suite.eventHandler1)
	suite.Equal(suite.eventManager.handlers[suite.event1.GetName()][1], suite.eventHandler2)
}

func (suite *EventManagerTestSuite) Test_RegisterSameHandler() {
	err := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler1)
	err2 := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler1)

	suite.Nil(err)
	suite.NotNil(err2)

}

func (suite *EventManagerTestSuite) Test_Clear() {
	err := suite.eventManager.Register(suite.event3.GetName(), suite.eventHandler3)
	suite.eventManager.Clear()

	suite.Nil(err)
	suite.Equal(len(suite.eventManager.handlers), 0)
}

func (suite *EventManagerTestSuite) Test_Has() {
	err := suite.eventManager.Register(suite.event3.GetName(), suite.eventHandler3)
	hasEvent := suite.eventManager.Has(suite.event3.GetName(), &suite.eventHandler3)

	suite.Nil(err)
	suite.True(hasEvent)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventManagerTestSuite) Test_Dispatch() {
	mockedEventHandler := &MockHandler{}
	mockedEventHandler.On("Handle", suite.event1)

	suite.eventManager.Register(suite.event1.GetName(), mockedEventHandler)
	suite.eventManager.Dispatch(suite.event1)

	mockedEventHandler.AssertExpectations(suite.T())
	mockedEventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventManagerTestSuite) Test_Remove() {
	// register two handlers
	err1 := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler1)
	err2 := suite.eventManager.Register(suite.event1.GetName(), suite.eventHandler2)

	suite.Nil(err1)
	suite.Nil(err2)
	suite.Equal(len(suite.eventManager.handlers[suite.event1.GetName()]), 2)

	// removes the handler
	err3 := suite.eventManager.Remove(suite.event1.GetName(), suite.eventHandler1)

	suite.Nil(err3)
	suite.Equal(len(suite.eventManager.handlers[suite.event1.GetName()]), 1)
}
