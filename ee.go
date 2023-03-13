package ee

import (
	"errors"
	"fmt"
)

type EventHandler = func(data ...any)

type EventEmitter struct {
	events map[string][]EventHandler
}

func (emitter *EventEmitter) On(event string, handler EventHandler) {
	if emitter.events == nil {
		emitter.events = make(map[string][]EventHandler)
	}

	if emitter.events[event] == nil {
		emitter.events[event] = []EventHandler{}
	}

	emitter.events[event] = append(emitter.events[event], handler)
}

func (emitter *EventEmitter) Emit(event string, data ...any) error {
	handlers, ok := emitter.events[event]
	if !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	for _, handler := range handlers {
		handler(data...)
	}

	return nil
}

func New() *EventEmitter {
	return &EventEmitter{
		events: make(map[string][]EventHandler),
	}
}
