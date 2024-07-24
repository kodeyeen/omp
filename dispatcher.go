package omp

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type EventType string

type Event interface {
	Type() EventType
	Payload() any
}

type event struct {
	Event

	_type   EventType
	payload any
}

func NewEvent(_type EventType, payload any) Event {
	return &event{
		_type:   _type,
		payload: payload,
	}
}

func (e *event) Type() EventType {
	return e._type
}

func (e *event) Payload() any {
	return e.payload
}

type Dispatcher struct {
	listeners map[EventType][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[EventType][]Listener),
	}
}

func (d *Dispatcher) HandleEvent(ctx context.Context, e Event) error {
	listeners := d.listeners[e.Type()]

	g, ctx := errgroup.WithContext(ctx)

	for _, listener := range listeners {
		g.Go(func() error {
			return listener.HandleEvent(ctx, e)
		})
	}

	return g.Wait()
}

func (d *Dispatcher) Listen(_type EventType, listener Listener) {
	d.listeners[_type] = append(d.listeners[_type], listener)
}

func (d *Dispatcher) ListenFunc(_type EventType, listener func(context.Context, Event) error) {
	d.Listen(_type, ListenerFunc(listener))
}

func (d *Dispatcher) HasListener(_type EventType) bool {
	_, ok := d.listeners[_type]
	return ok
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listeners := range events {
		d.listeners[_type] = append(d.listeners[_type], listeners...)
	}
}

func (d *Dispatcher) SubscribeFunc(subscriber func() map[EventType][]Listener) {
	d.Subscribe(SubscriberFunc(subscriber))
}

type Listener interface {
	HandleEvent(context.Context, Event) error
}

type ListenerFunc func(context.Context, Event) error

func (f ListenerFunc) HandleEvent(ctx context.Context, e Event) error {
	return f(ctx, e)
}

type Subscriber interface {
	SubscribedEvents() map[EventType][]Listener
}

type SubscriberFunc func() map[EventType][]Listener

func (f SubscriberFunc) SubscribedEvents() map[EventType][]Listener {
	return f()
}
