package gomp

// #include "include/gamemode.h"
import "C"

type GameMode struct {
}

func (gm *GameMode) UseManualEngineAndLights() {
	C.useManualEngineAndLights()
}

func (gm *GameMode) DisableInteriorEnterExits() {
	panic("not implemented")
}
