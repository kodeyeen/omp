package gomp

import "unsafe"

type PlayerObject struct {
	handle unsafe.Pointer
}

func NewPlayerObject() *PlayerObject {
	return &PlayerObject{}
}

func (o *PlayerObject) SetPosition(pos Vector3) {
	panic("not implemented")
}

func (o *PlayerObject) Position() Vector3 {
	panic("not implemented")
}
