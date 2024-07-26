package omp

import (
	"context"

	"github.com/kodeyeen/event"
)

// These are wrappers around event package.

type baseEvent = event.Event
type baseDispatcher = *event.Dispatcher

type EventType event.Type

type Event interface {
	Type() EventType
	Payload() any
}

type payloadEvt struct {
	baseEvent
}

func NewEvent(_type EventType, payload any) Event {
	return &payloadEvt{
		baseEvent: event.WithPayload(event.Type(_type), payload),
	}
}

func (e *payloadEvt) Type() EventType {
	return EventType(e.baseEvent.Type())
}

func (e *payloadEvt) Payload() any {
	return e.baseEvent.Payload()
}

type Dispatcher struct {
	baseDispatcher
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		baseDispatcher: event.NewDispatcher(),
	}
}

func (d *Dispatcher) HandleEvent(ctx context.Context, e Event) error {
	pkgEvent := event.WithPayload(event.Type(e.Type()), e.Payload())
	return d.baseDispatcher.HandleEvent(ctx, pkgEvent)
}

func (d *Dispatcher) Listen(_type EventType, listener Listener) {
	pkgListener := event.ListenerFunc(func(ctx context.Context, e event.Event) error {
		ompEvent := NewEvent(EventType(e.Type()), e.Payload())
		return listener.HandleEvent(ctx, ompEvent)
	})

	d.baseDispatcher.Listen(event.Type(_type), pkgListener)
}

func (d *Dispatcher) ListenFunc(_type EventType, listener func(context.Context, Event) error) {
	d.Listen(_type, ListenerFunc(listener))
}

func (d *Dispatcher) HasListener(_type EventType) bool {
	return d.baseDispatcher.HasListener(event.Type(_type))
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	d.baseDispatcher.Subscribe(subscriber)
}

func (d *Dispatcher) SubscribeFunc(subscriber func() any) {
	d.Subscribe(SubscriberFunc(subscriber))
}

type Listener interface {
	HandleEvent(ctx context.Context, e Event) error
}

type ListenerFunc func(ctx context.Context, e Event) error

func (f ListenerFunc) HandleEvent(ctx context.Context, e Event) error {
	return f(ctx, e)
}

type Subscriber interface {
	event.Subscriber
}

type SubscriberFunc event.SubscriberFunc

func (f SubscriberFunc) SubscribedEvents() any {
	return f()
}
