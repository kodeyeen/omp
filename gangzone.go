package gomp

// #include "include/gangzone.h"
import "C"
import "unsafe"

type GangzonePosition struct {
	Min Vector2
	Max Vector2
}

type Gangzone struct {
	handle unsafe.Pointer
}

func NewGangzone(pos GangzonePosition) *Gangzone {
	gz := C.gangZone_create(C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))

	return &Gangzone{handle: gz}
}

func FreeGangzone(gz *Gangzone) {
	C.gangZone_release(gz.handle)
}

func EnableGangzoneCheck(gz *Gangzone) {
	C.gangZone_useCheck(gz.handle, 1)
}

func DisableGangzoneCheck(gz *Gangzone) {
	C.gangZone_useCheck(gz.handle, 0)
}

func (g *Gangzone) IsShownFor(plr *Player) bool {
	return C.gangZone_isShownForPlayer(g.handle, plr.handle) != 0
}

func (g *Gangzone) IsFlashingFor(plr *Player) bool {
	return C.gangZone_isFlashingForPlayer(g.handle, plr.handle) != 0
}

func (g *Gangzone) ShowFor(plr *Player, clr Color) {
	C.gangZone_showForPlayer(g.handle, plr.handle, C.uint(clr))
}

func (g *Gangzone) ShowForAll(clr Color) {
	C.gangZone_showForAll(g.handle, C.uint(clr))
}

func (g *Gangzone) HideFor(plr *Player) {
	C.gangZone_hideForPlayer(g.handle, plr.handle)
}

func (g *Gangzone) HideForAll() {
	C.gangZone_hideForAll(g.handle)
}

func (g *Gangzone) FlashFor(plr *Player, clr Color) {
	C.gangZone_flashForPlayer(g.handle, plr.handle, C.uint(clr))
}

func (g *Gangzone) FlashForAll(clr Color) {
	C.gangZone_flashForAll(g.handle, C.uint(clr))
}

func (g *Gangzone) StopFlashFor(plr *Player) {
	C.gangZone_stopFlashForPlayer(g.handle, plr.handle)
}

func (g *Gangzone) StopFlashForAll() {
	C.gangZone_stopFlashForAll(g.handle)
}

func (g *Gangzone) Position() GangzonePosition {
	cPos := C.gangZone_getPosition(g.handle)

	return GangzonePosition{
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

func (g *Gangzone) SetPosition(pos GangzonePosition) {
	C.gangZone_setPosition(g.handle, C.float(pos.Min.X), C.float(pos.Min.Y), C.float(pos.Max.X), C.float(pos.Max.Y))
}

func (g *Gangzone) IsPlayerInside(plr *Player) bool {
	return C.gangZone_isPlayerInside(g.handle, plr.handle) != 0
}

func (g *Gangzone) FlashingColorFor(plr *Player) Color {
	return Color(C.gangZone_getFlashingColourForPlayer(g.handle, plr.handle))
}

func (g *Gangzone) ColorFor(plr *Player) Color {
	return Color(C.gangZone_getColourForPlayer(g.handle, plr.handle))
}
