package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) error {
	// Simulate some processing
	return nil
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	EventDispatcher EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.EventDispatcher = NewEventDispatcher()
	suite.event = TestEvent{Name: "test.event", Payload: "test payload"}
	suite.event2 = TestEvent{Name: "test.event2", Payload: "test payload 2"}
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.EventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NotNil(err)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal("handler already registered", err.Error())
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Clear()
	suite.Equal(0, len(suite.EventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	//Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) error {
	args := m.Called(event)
	wg.Done()
	return args.Error(0)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][0])

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	suite.EventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Clear()
	suite.Equal(0, len(suite.EventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", mock.AnythingOfType("*events.TestEvent")).Return(nil)

	eh2 := &MockHandler{}
	eh2.On("Handle", mock.AnythingOfType("*events.TestEvent")).Return(nil)

	suite.EventDispatcher.Register(suite.event.GetName(), eh)
	suite.EventDispatcher.Register(suite.event.GetName(), eh2)

	suite.EventDispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)

	eh2.AssertExpectations(suite.T())
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
