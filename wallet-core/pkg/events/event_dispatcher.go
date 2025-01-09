package events

import (
	"errors"
	"sync"
)

var (
	ErrHandlerAlreadyRegistred = errors.New("handler already registred")
	ErrHandlerNotFound         = errors.New("no handler found")
)

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Dispatch(event Event) error
	Clear() error
	GetHandlersLength(eventName string) int
}

type eventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() EventDispatcher {
	return &eventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (e *eventDispatcher) Register(eventName string, handler EventHandler) error {
	handlers, ok := e.handlers[eventName]
	if ok {
		for _, h := range handlers {
			if h == handler {
				return ErrHandlerAlreadyRegistred
			}
		}
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

func (e *eventDispatcher) GetHandlersLength(eventName string) int {
	handlers, ok := e.handlers[eventName]
	if !ok {
		return 0
	}
	return len(handlers)
}

func (e *eventDispatcher) Clear() error {
	clear(e.handlers)
	return nil
}

func (e *eventDispatcher) Has(eventName string, handler EventHandler) bool {
	handlers, ok := e.handlers[eventName]
	if ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (e *eventDispatcher) Dispatch(event Event) error {
	handlers, ok := e.handlers[event.GetName()]
	if ok {
		wg := &sync.WaitGroup{}
		wg.Add(len(handlers))

		for _, h := range handlers {
			go func() {
				defer wg.Done()
				h.Handle(event)
			}()
		}

		wg.Wait()
	}

	return nil
}

func (e *eventDispatcher) Remove(eventName string, handler EventHandler) error {
	handlers, ok := e.handlers[eventName]
	if ok {
		for i, h := range handlers {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
			}
		}
	}

	return nil
}
