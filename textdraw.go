package gomp

import "unsafe"

type TextDraw struct {
	handle       unsafe.Pointer
	playerHandle unsafe.Pointer
}

func NewTextDraw() *TextDraw {
	return &TextDraw{}
}

func NewPlayerTextDraw(player *Player) *TextDraw {
	return &TextDraw{}
}
