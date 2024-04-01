package gomp

import "unsafe"

type Checkpoint interface {
	SetRadius(radius float32)
	Radius() float32
	SetPosition(pos Vector3)
	Position() Vector3
	Enable()
	Disable()
	IsEnabled() bool
	IsPlayerInside() bool
}

type DefaultCheckpoint struct {
	handle unsafe.Pointer
}

func DestroyCheckpoint(cp Checkpoint) {}

func (c *DefaultCheckpoint) SetRadius(radius float32) {
	panic("not implemented")
}

func (c *DefaultCheckpoint) Radius() float32 {
	panic("not implemented")
}

func (c *DefaultCheckpoint) SetPosition(pos Vector3) {
	panic("not implemented")
}

func (c *DefaultCheckpoint) Position() Vector3 {
	panic("not implemented")
}

func (c *DefaultCheckpoint) Enable() {
	panic("not implemented")
}

func (c *DefaultCheckpoint) Disable() {
	panic("not implemented")
}

func (c *DefaultCheckpoint) IsEnabled() bool {
	panic("not implemented")
}

func (c *DefaultCheckpoint) IsPlayerInside() bool {
	panic("not implemented")
}

type RaceCheckpoint struct {
	handle unsafe.Pointer
}
