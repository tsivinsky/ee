package ee

import (
	"errors"
	"fmt"
)

type EventHandlerFunc = func(data ...any)

type EventHandler struct {
	Id   int
	Func EventHandlerFunc
}

type EventEmitter struct {
	lastId int
	events map[string][]EventHandler
}

// On subscribes handler to event and returns index of handler
// It also creates internal events map if struct wasn't initialized
func (emitter *EventEmitter) On(event string, handlerFunc EventHandlerFunc) int {
	if emitter.events == nil {
		emitter.lastId = 0
		emitter.events = make(map[string][]EventHandler)
	}

	handler := EventHandler{
		Id:   emitter.lastId + 1,
		Func: handlerFunc,
	}

	emitter.events[event] = append(emitter.events[event], handler)
	emitter.lastId += 1

	return handler.Id
}

// Once subscribes handler to event and removes it after first execution
// also removes event from map if handler was only one
func (emitter *EventEmitter) Once(event string, handler EventHandlerFunc) {
	id := -1

	h := func(data ...any) {
		handler(data...)

		if id != -1 {
			hasOneHandler := len(emitter.events[event]) == 1
			if hasOneHandler {
				emitter.Remove(event)
			}

			emitter.Off(event, id)
		}
	}

	id = emitter.On(event, h)
}

// Emit calls all handlers subscribed to event and returns error if event doesn't exist in map
func (emitter *EventEmitter) Emit(event string, data ...any) error {
	handlers, ok := emitter.events[event]
	if !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	for _, handler := range handlers {
		handler.Func(data...)
	}

	return nil
}

// Remove deletes event and returns error if event doesn't exist
func (emitter *EventEmitter) Remove(event string) error {
	if _, ok := emitter.events[event]; !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	delete(emitter.events, event)

	return nil
}

// Off unsubscribes handler from event and returns error if event doesn't exist
func (emitter *EventEmitter) Off(event string, handlerId int) error {
	if _, ok := emitter.events[event]; !ok {
		return errors.New(fmt.Sprintf("\"%s\" event does not exist. Register it with ee.On(\"%s\", handler)", event, event))
	}

	for i, handler := range emitter.events[event] {
		if handler.Id == handlerId {
			emitter.events[event] = append(emitter.events[event][:i], emitter.events[event][i+1:]...)
			break
		}
	}

	return nil
}

// New initializes new EventEmitter and returns pointer to it
func New() *EventEmitter {
	return &EventEmitter{
		events: make(map[string][]EventHandler),
	}
}
