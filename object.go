package gomp

import "unsafe"

type Object struct {
	handle unsafe.Pointer
}

func NewObject() *Object {
	return &Object{}
}
