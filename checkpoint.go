package gomp

// #include "include/checkpoint.h"
import "C"
import "unsafe"

type DefaultCheckpoint struct {
	handle unsafe.Pointer
}

func (c *DefaultCheckpoint) SetPosition(pos Vector3) {
	C.checkpoint_setPosition(c.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (c *DefaultCheckpoint) Position() Vector3 {
	cPos := C.checkpoint_getPosition(c.handle)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (c *DefaultCheckpoint) SetRadius(radius float32) {
	C.checkpoint_setRadius(c.handle, C.float(radius))
}

func (c *DefaultCheckpoint) Radius() float32 {
	return float32(C.checkpoint_getRadius(c.handle))
}

func (c *DefaultCheckpoint) IsPlayerInside() bool {
	return C.checkpoint_isPlayerInside(c.handle) != 0
}

func (c *DefaultCheckpoint) Enable() {
	C.checkpoint_enable(c.handle)
}

func (c *DefaultCheckpoint) Disable() {
	C.checkpoint_disable(c.handle)
}

func (c *DefaultCheckpoint) IsEnabled() bool {
	return C.checkpoint_isEnabled(c.handle) != 0
}

type RaceCheckpointType int

const (
	RaceCheckpointTypeNormal RaceCheckpointType = iota
	RaceCheckpointTypeFinish
	RaceCheckpointTypeNothing
	RaceCheckpointTypeAirNormal
	RaceCheckpointTypeAirFinish
	RaceCheckpointTypeAirOne
	RaceCheckpointTypeAirTwo
	RaceCheckpointTypeAirThree
	RaceCheckpointTypeAirFour
	RaceCheckpointTypeNone
)

type RaceCheckpoint struct {
	handle unsafe.Pointer
}

func (c *RaceCheckpoint) SetPosition(pos Vector3) {
	C.raceCheckpoint_setPosition(c.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (c *RaceCheckpoint) Position() Vector3 {
	cPos := C.raceCheckpoint_getPosition(c.handle)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (c *RaceCheckpoint) SetRadius(radius float32) {
	C.raceCheckpoint_setRadius(c.handle, C.float(radius))
}

func (c *RaceCheckpoint) Radius() float32 {
	return float32(C.raceCheckpoint_getRadius(c.handle))
}

func (c *RaceCheckpoint) IsPlayerInside() bool {
	return C.raceCheckpoint_isPlayerInside(c.handle) != 0
}

func (c *RaceCheckpoint) Enable() {
	C.raceCheckpoint_enable(c.handle)
}

func (c *RaceCheckpoint) Disable() {
	C.raceCheckpoint_disable(c.handle)
}

func (c *RaceCheckpoint) IsEnabled() bool {
	return C.raceCheckpoint_isEnabled(c.handle) != 0
}

func (c *RaceCheckpoint) SetType(_type RaceCheckpointType) {
	C.raceCheckpoint_setType(c.handle, C.int(_type))
}

func (c *RaceCheckpoint) Type() RaceCheckpointType {
	return RaceCheckpointType(C.raceCheckpoint_getType(c.handle))
}

func (c *RaceCheckpoint) SetNextPosition(pos Vector3) {
	C.raceCheckpoint_setNextPosition(c.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (c *RaceCheckpoint) NextPosition() Vector3 {
	cPos := C.raceCheckpoint_getNextPosition(c.handle)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}
