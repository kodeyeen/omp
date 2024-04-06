package gomp

import "unsafe"

type GlobalTextDraw struct {
	handle unsafe.Pointer
}

func NewGlobalTextDraw() *GlobalTextDraw {
	return &GlobalTextDraw{}
}
