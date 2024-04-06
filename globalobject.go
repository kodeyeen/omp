package gomp

import "unsafe"

type GlobalObject struct {
	handle unsafe.Pointer
}

func NewGlobalObject() *GlobalObject {
	return &GlobalObject{}
}

func (o *GlobalObject) SetPosition(pos Vector3) {
	panic("not implemented")
}

func (o *GlobalObject) Position() Vector3 {
	panic("not implemented")
}
