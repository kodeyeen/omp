package event

import (
	"reflect"
	"slices"
)

type dispatcher struct {
	listeners map[Type][]listener
}

func NewDispatcher() *dispatcher {
	return &dispatcher{
		listeners: make(map[Type][]listener),
	}
}

func Dispatch[T any](d *dispatcher, evtType Type, evt T) bool {
	listeners := d.listeners[evtType]

	for _, l := range listeners {
		handler, ok := l.handler.(func(T) bool)
		if !ok {
			continue
		}

		callNext := handler(evt)

		if l.once {
			d.Off(evtType, l.handler)
		}

		if !callNext {
			return false
		}
	}

	return true
}

func (d *dispatcher) On(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	listeners = append(listeners, listener{
		handler: handler,
		once:    false,
	})

	d.listeners[evtType] = listeners
}

func (d *dispatcher) Once(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	listeners = append(listeners, listener{
		handler: handler,
		once:    true,
	})

	d.listeners[evtType] = listeners
}

func (d *dispatcher) Off(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	idx := slices.IndexFunc(listeners, func(l listener) bool {
		return reflect.ValueOf(l.handler).Pointer() == reflect.ValueOf(handler).Pointer()
	})

	if idx == -1 {
		return
	}

	d.listeners[evtType] = slices.Delete(listeners, idx, idx+1)
}

func (d *dispatcher) Has(evtType Type) bool {
	_, ok := d.listeners[evtType]
	return ok
}
