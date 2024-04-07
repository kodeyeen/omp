package gomp

import "unsafe"

type TextDraw struct {
	handle unsafe.Pointer
}

func NewTextDraw() *TextDraw {
	return &TextDraw{}
}
