package omp

import (
	"context"
	"reflect"
)

type EventEmitter interface {
	Emit(ctx context.Context, e any) error

	addListener(t reflect.Type, l listener)
	hasListeners(t reflect.Type) bool
}

type listener func(context.Context, any) error

type eventEmitter struct {
	listeners map[reflect.Type][]listener
}

func newEventEmitter() EventEmitter {
	return &eventEmitter{
		listeners: make(map[reflect.Type][]listener),
	}
}

func (em *eventEmitter) Emit(ctx context.Context, e any) error {
	t := reflect.TypeOf(e)

	for _, ln := range em.listeners[t] {
		err := ln(ctx, e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (em *eventEmitter) addListener(t reflect.Type, l listener) {
	em.listeners[t] = append(em.listeners[t], l)
}

func (em *eventEmitter) hasListeners(t reflect.Type) bool {
	return len(em.listeners[t]) > 0
}
