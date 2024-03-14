package gomp

// #include <stdlib.h>
// #include "component.h"
import "C"
import (
	"unsafe"
)

type Player struct {
	handle unsafe.Pointer
}

func (p *Player) ID() int {
	// return int(C.player_getID(p.handle))
	return 0
}

func (p *Player) Skin() int {
	return 0
}

func (p *Player) SetSkin(skin int) {

}

func (p *Player) DefaultColor() int {
	return 0
}

// Gets the amount of ammo in a player's current weapon.
func (p *Player) Ammo() int {
	return 0
}

// Get the player animation flags.
func (p *Player) AnimationFlags() int {
	return 0
}

// Returns the index of any running applied animations.
func (p *Player) AnimationIndex() int {
	return 0
}

func (p *Player) Armor() float32 {
	return 0.0
}

func (p *Player) RemovedBuildingsCount() int {
	return 0
}

func (p *Player) AspectRatio() float32 {
	return 0.0
}

func (p *Player) CameraFrontVector() *Position {
	return nil
}

func (p *Player) CameraMode() int {
	return 0
}

func (p *Player) CameraPosition() *Position {
	return nil
}

func (p *Player) CameraTargetActor() {

}

func (p *Player) CameraTargetObject() {

}

func (p *Player) CameraTargetPlayer() {

}

func (p *Player) CameraTargetPlayerObject() {

}

func (p *Player) CameraTargetVehicle() {

}

func (p *Player) CameraUpVector() {

}

func (p *Player) CameraZoom() {

}

func (p *Player) CheckPoint() {

}

func (p *Player) Class() {

}

func (p *Player) Color() {

}

func (p *Player) CustomSkin() {

}

func (p *Player) DialogID() {

}

func (p *Player) DialogData() {

}

func (p *Player) DistanceFromPoint() {

}

func (p *Player) DrunkLevel() {

}

func (p *Player) FacingAngle() {

}

func (p *Player) FightingStyle() {

}

func (p *Player) GhostMode() {

}

func (p *Player) Gravity() {

}

func (p *Player) Health() {

}

func (p *Player) HydraReactorAngle() {

}

func (p *Player) Interior() {

}

func (p *Player) IP() {

}

func (p *Player) Keys() {

}

func (p *Player) LandingGearState() {

}

func (p *Player) LastShotVectors() {

}

func (p *Player) LastSyncedTrailerID() {

}

func (p *Player) LastSyncedVehicleID() {

}

// Get the colour of a player's nametag and radar blip for another player.
func (p *Player) MarkerForPlayer(target *Player) {

}

func (p *Player) Menu(target *Player) {

}

func (p *Player) Money(target *Player) int {
	return 0
}

func (p *Player) Name() string {
	cname := C.player_getName(p.handle)

	return C.GoString(cname)
}

func (p *Player) SetName(name string) {

}

func (p *Player) NetworkStats() {

}

func (p *Player) SendMessage(color int, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(p.handle, C.int(color), cmsg)
}
