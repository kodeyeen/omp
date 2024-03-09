package event

import (
	"reflect"
	"slices"
)

type dispatcher struct {
	handlers map[Type][]handler
}

func NewDispatcher() *dispatcher {
	return &dispatcher{
		handlers: make(map[Type][]handler),
	}
}

func Dispatch[T any](d *dispatcher, evtType Type, evt T) {
	hlrs, ok := d.handlers[evtType]
	if !ok {
		return
	}

	for _, hlr := range hlrs {
		fn, ok := hlr.handle.(func(T))
		if !ok {
			continue
		}

		fn(evt)

		if hlr.once {
			d.Off(evtType, hlr.handle)
		}
	}
}

func (d *dispatcher) On(evtType Type, callback any) {
	_, ok := d.handlers[evtType]
	if !ok {
		d.handlers[evtType] = make([]handler, 0)
	}

	hlrs := d.handlers[evtType]
	hlrs = append(hlrs, handler{
		handle: callback,
		once:   false,
	})

	d.handlers[evtType] = hlrs
}

func (d *dispatcher) Once(evtType Type, callback any) {
	_, ok := d.handlers[evtType]
	if !ok {
		d.handlers[evtType] = make([]handler, 0)
	}

	hlrs := d.handlers[evtType]
	hlrs = append(hlrs, handler{
		handle: callback,
		once:   true,
	})

	d.handlers[evtType] = hlrs
}

func (d *dispatcher) Off(evtType Type, callback any) {
	hlrs, ok := d.handlers[evtType]
	if !ok {
		return
	}

	idx := slices.IndexFunc(hlrs, func(hlr handler) bool {
		return reflect.ValueOf(hlr.handle).Pointer() == reflect.ValueOf(callback).Pointer()
	})

	hlrs = append(hlrs[:idx], hlrs[idx+1:]...)

	d.handlers[evtType] = hlrs
}
