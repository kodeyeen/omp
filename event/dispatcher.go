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

func Dispatch[T any](d *dispatcher, evtType Type, evt T) {
	lstrs, ok := d.listeners[evtType]
	if !ok {
		return
	}

	for _, lstr := range lstrs {
		handler, ok := lstr.handler.(func(T))
		if !ok {
			continue
		}

		handler(evt)

		if lstr.once {
			d.Off(evtType, lstr.handler)
		}
	}
}

func (d *dispatcher) On(evtType Type, handler any) {
	_, ok := d.listeners[evtType]
	if !ok {
		d.listeners[evtType] = make([]listener, 0)
	}

	lstrs := d.listeners[evtType]
	lstrs = append(lstrs, listener{
		handler: handler,
		once:    false,
	})

	d.listeners[evtType] = lstrs
}

func (d *dispatcher) Once(evtType Type, handler any) {
	_, ok := d.listeners[evtType]
	if !ok {
		d.listeners[evtType] = make([]listener, 0)
	}

	lstrs := d.listeners[evtType]
	lstrs = append(lstrs, listener{
		handler: handler,
		once:    true,
	})

	d.listeners[evtType] = lstrs
}

func (d *dispatcher) Off(evtType Type, handler any) {
	lstrs, ok := d.listeners[evtType]
	if !ok {
		return
	}

	idx := slices.IndexFunc(lstrs, func(lstr listener) bool {
		return reflect.ValueOf(lstr.handler).Pointer() == reflect.ValueOf(handler).Pointer()
	})

	lstrs = append(lstrs[:idx], lstrs[idx+1:]...)

	d.listeners[evtType] = lstrs
}
