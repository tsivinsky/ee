package ee

import (
	"errors"
	"fmt"
)

type EventHandler = func(data ...any)

type EventEmitter struct {
	events map[string][]EventHandler
}

func (emitter *EventEmitter) On(event string, handler EventHandler) int {
	if emitter.events == nil {
		emitter.events = make(map[string][]EventHandler)
	}

	emitter.events[event] = append(emitter.events[event], handler)
	index := len(emitter.events[event]) - 1

	return index
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

func (emitter *EventEmitter) Remove(event string) error {
	if _, ok := emitter.events[event]; !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	delete(emitter.events, event)

	return nil
}

func (emitter *EventEmitter) Off(event string, index int) error {
	if _, ok := emitter.events[event]; !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	emitter.events[event] = append(emitter.events[event][:index], emitter.events[event][index+1:]...)

	return nil
}

func New() *EventEmitter {
	return &EventEmitter{
		events: make(map[string][]EventHandler),
	}
}
