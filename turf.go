package omp

// #include "include/turf.h"
import "C"
import (
	"errors"
	"unsafe"
)

type TurfPosition struct {
	Min Vector2
	Max Vector2
}

type Turf struct {
	handle unsafe.Pointer
}

func NewTurf(pos TurfPosition) (*Turf, error) {
	cTurf := C.turf_create(C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))

	if cTurf == nil {
		return nil, errors.New("turf limit was reached")
	}

	return &Turf{handle: cTurf}, nil
}

func FreeTurf(turf *Turf) {
	C.turf_release(turf.handle)
}

func EnableTurfCheck(turf *Turf) {
	C.turf_useCheck(turf.handle, 1)
}

func DisableTurfCheck(turf *Turf) {
	C.turf_useCheck(turf.handle, 0)
}

func (t *Turf) IsShownFor(plr *Player) bool {
	return C.turf_isShownForPlayer(t.handle, plr.handle) != 0
}

func (t *Turf) IsFlashingFor(plr *Player) bool {
	return C.turf_isFlashingForPlayer(t.handle, plr.handle) != 0
}

func (t *Turf) ShowFor(plr *Player, clr uint) {
	C.turf_showForPlayer(t.handle, plr.handle, C.uint(clr))
}

func (t *Turf) ShowForAll(clr uint) {
	C.turf_showForAll(t.handle, C.uint(clr))
}

func (t *Turf) HideFor(plr *Player) {
	C.turf_hideForPlayer(t.handle, plr.handle)
}

func (t *Turf) HideForAll() {
	C.turf_hideForAll(t.handle)
}

func (t *Turf) FlashFor(plr *Player, clr uint) {
	C.turf_flashForPlayer(t.handle, plr.handle, C.uint(clr))
}

func (t *Turf) FlashForAll(clr uint) {
	C.turf_flashForAll(t.handle, C.uint(clr))
}

func (t *Turf) StopFlashFor(plr *Player) {
	C.turf_stopFlashForPlayer(t.handle, plr.handle)
}

func (t *Turf) StopFlashForAll() {
	C.turf_stopFlashForAll(t.handle)
}

func (t *Turf) Position() TurfPosition {
	cPos := C.turf_getPosition(t.handle)

	return TurfPosition{
		Min: Vector2{
			X: float32(cPos.min.x),
			Y: float32(cPos.min.y),
		},
		Max: Vector2{
			X: float32(cPos.max.x),
			Y: float32(cPos.max.y),
		},
	}
}

func (t *Turf) SetPosition(pos TurfPosition) {
	C.turf_setPosition(t.handle, C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))
}

func (t *Turf) IsPlayerInside(plr *Player) bool {
	return C.turf_isPlayerInside(t.handle, plr.handle) != 0
}

func (t *Turf) FlashingColorFor(plr *Player) uint {
	return uint(C.turf_getFlashingColourForPlayer(t.handle, plr.handle))
}

func (t *Turf) ColorFor(plr *Player) uint {
	return uint(C.turf_getColourForPlayer(t.handle, plr.handle))
}

type PlayerTurf struct {
	handle unsafe.Pointer
	player *Player
}

func NewPlayerTurf(plr *Player, pos TurfPosition) (*PlayerTurf, error) {
	cTurf := C.playerTurf_create(plr.handle, C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))
	if cTurf == nil {
		return nil, errors.New("player turf limit was reached")
	}

	turf := &PlayerTurf{
		handle: cTurf,
		player: plr,
	}

	return turf, nil
}

func FreePlayerTurf(turf *PlayerTurf) {
	C.playerTurf_release(turf.handle, turf.player.handle)
}

func EnablePlayerTurfCheck(turf *PlayerTurf) {
	C.turf_useCheck(turf.handle, 1)
}

func DisablePlayerTurfCheck(turf *PlayerTurf) {
	C.turf_useCheck(turf.handle, 0)
}

func (t *PlayerTurf) Player() *Player {
	return t.player
}

func (t *PlayerTurf) IsShown() bool {
	return C.turf_isShownForPlayer(t.handle, t.player.handle) != 0
}

func (t *PlayerTurf) IsFlashing() bool {
	return C.turf_isFlashingForPlayer(t.handle, t.player.handle) != 0
}

func (t *PlayerTurf) Show(clr uint) {
	C.turf_showForPlayer(t.handle, t.player.handle, C.uint(clr))
}

func (t *PlayerTurf) Hide() {
	C.turf_hideForPlayer(t.handle, t.player.handle)
}

func (t *PlayerTurf) Flash(clr uint) {
	C.turf_flashForPlayer(t.handle, t.player.handle, C.uint(clr))
}

func (t *PlayerTurf) StopFlash() {
	C.turf_stopFlashForPlayer(t.handle, t.player.handle)
}

func (t *PlayerTurf) Position() TurfPosition {
	cPos := C.turf_getPosition(t.handle)

	return TurfPosition{
		Min: Vector2{
			X: float32(cPos.min.x),
			Y: float32(cPos.min.y),
		},
		Max: Vector2{
			X: float32(cPos.max.x),
			Y: float32(cPos.max.y),
		},
	}
}

func (t *PlayerTurf) SetPosition(pos TurfPosition) {
	C.turf_setPosition(t.handle, C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))
}

func (t *PlayerTurf) IsPlayerInside() bool {
	return C.turf_isPlayerInside(t.handle, t.player.handle) != 0
}

func (t *PlayerTurf) FlashingColor() uint {
	return uint(C.turf_getFlashingColourForPlayer(t.handle, t.player.handle))
}

func (t *PlayerTurf) Color() uint {
	return uint(C.turf_getColourForPlayer(t.handle, t.player.handle))
}
