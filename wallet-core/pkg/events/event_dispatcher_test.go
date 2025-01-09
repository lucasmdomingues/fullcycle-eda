package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type eventMock struct {
	mock.Mock

	Name    string
	Payload any
}

func (e *eventMock) GetDateTime() time.Time {
	return time.Now()
}

func (e *eventMock) GetName() string {
	return e.Name
}

func (e *eventMock) GetPayload() any {
	return e.Payload
}

func (e *eventMock) SetPayload(data any) {
	e.Payload = data
}

type eventHandlerMock struct {
	mock.Mock

	ID int
}

func (h *eventHandlerMock) Handle(event Event) {
	h.Called(event)
}

type EventDispatcherTestSuite struct {
	suite.Suite

	eventFoo        Event
	eventBar        Event
	handlerFoo      EventHandler
	handlerBar      EventHandler
	eventDispatcher EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()

	suite.handlerFoo = &eventHandlerMock{ID: 1}
	suite.handlerBar = &eventHandlerMock{ID: 2}

	suite.eventFoo = &eventMock{
		Name:    "foo",
		Payload: "foo",
	}

	suite.eventBar = &eventMock{
		Name:    "bar",
		Payload: "bar",
	}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Register foo
	err := suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerFoo)
	suite.NoError(err)

	handlersLength := suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.NoError(err)
	suite.Equal(1, handlersLength)

	// Register bar
	err = suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerBar)
	suite.NoError(err)

	handlersLength = suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.NoError(err)
	suite.Equal(2, handlersLength)

	// Assert
	suite.True(suite.eventDispatcher.Has(suite.eventFoo.GetName(), suite.handlerFoo))
	suite.True(suite.eventDispatcher.Has(suite.eventFoo.GetName(), suite.handlerBar))

	// Handler Already Registred Error
	err = suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerFoo)
	suite.Error(err)
	suite.ErrorIs(err, ErrHandlerAlreadyRegistred)
}

func (suite *EventDispatcherTestSuite) TestClear() {
	// Foo Event
	err := suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerFoo)
	suite.NoError(err)

	err = suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerBar)
	suite.NoError(err)

	suite.eventDispatcher.Clear()

	fooHandlersLength := suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.Equal(0, fooHandlersLength)

	// Bar Event
	err = suite.eventDispatcher.Register(suite.eventBar.GetName(), suite.handlerFoo)
	suite.NoError(err)

	err = suite.eventDispatcher.Register(suite.eventBar.GetName(), suite.handlerBar)
	suite.NoError(err)

	suite.eventDispatcher.Clear()
	barHandlersLength := suite.eventDispatcher.GetHandlersLength(suite.eventBar.GetName())
	suite.Equal(0, barHandlersLength)
}

func (suite *EventDispatcherTestSuite) TestHas() {
	err := suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerFoo)
	suite.NoError(err)

	err = suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerBar)
	suite.NoError(err)

	fooHandlersLength := suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.NoError(err)
	suite.Equal(2, fooHandlersLength)

	suite.True(suite.eventDispatcher.Has(suite.eventFoo.GetName(), suite.handlerFoo))
	suite.True(suite.eventDispatcher.Has(suite.eventFoo.GetName(), suite.handlerBar))
}

func (suite *EventDispatcherTestSuite) TestDispatch() {
	eventHandler := &eventHandlerMock{}
	eventHandler.On("Handle", suite.eventFoo)

	suite.eventDispatcher.Register(suite.eventFoo.GetName(), eventHandler)
	suite.eventDispatcher.Dispatch(suite.eventFoo)

	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) Test_Remove() {
	// Foo Event
	err := suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerFoo)
	suite.NoError(err)

	err = suite.eventDispatcher.Register(suite.eventFoo.GetName(), suite.handlerBar)
	suite.NoError(err)

	// Bar Event
	err = suite.eventDispatcher.Register(suite.eventBar.GetName(), suite.handlerFoo)
	suite.NoError(err)

	err = suite.eventDispatcher.Register(suite.eventBar.GetName(), suite.handlerBar)
	suite.NoError(err)

	// Remove
	suite.eventDispatcher.Remove(suite.eventFoo.GetName(), suite.handlerFoo)
	fooHandlersLength := suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.NoError(err)
	suite.Equal(1, fooHandlersLength)

	suite.eventDispatcher.Remove(suite.eventBar.GetName(), suite.handlerFoo)
	fooHandlersLength = suite.eventDispatcher.GetHandlersLength(suite.eventFoo.GetName())
	suite.NoError(err)
	suite.Equal(1, fooHandlersLength)
}

func TestSuite(t *testing.T) {
	suite.Run(t, &EventDispatcherTestSuite{})
}
