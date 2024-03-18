package gomp

// #include <stdlib.h>
// #include "component.h"
import "C"
import (
	"errors"
	"unsafe"
)

type PlayerNameStatus int

const (
	PlayerNameStatusUpdated PlayerNameStatus = iota
	PlayerNameStatusTaken
	PlayerNameStatusInvalid
)

type Player struct {
	handle unsafe.Pointer
}

func (p *Player) ID() int {
	return int(C.player_getID(p.handle))
}

func (p *Player) Skin() int {
	panic("not implemented")
}

func (p *Player) SetSkin(skin int) {
	panic("not implemented")
}

func (p *Player) DefaultColor() int {
	panic("not implemented")
}

// Gets the amount of ammo in a player's current weapon.
func (p *Player) Ammo() int {
	panic("not implemented")
}

// Get the player animation flags.
func (p *Player) AnimationFlags() int {
	panic("not implemented")
}

// Returns the index of any running applied animations.
func (p *Player) AnimationIndex() int {
	panic("not implemented")
}

func (p *Player) Armor() float32 {
	panic("not implemented")
}

func (p *Player) RemovedBuildingsCount() int {
	panic("not implemented")
}

func (p *Player) AspectRatio() float32 {
	panic("not implemented")
}

func (p *Player) CameraFrontVector() *Position {
	panic("not implemented")
}

func (p *Player) CameraMode() int {
	panic("not implemented")
}

func (p *Player) CameraPosition() *Position {
	panic("not implemented")
}

func (p *Player) CameraTargetActor() {
	panic("not implemented")
}

func (p *Player) CameraTargetObject() {
	panic("not implemented")
}

func (p *Player) CameraTargetPlayer() {
	panic("not implemented")
}

func (p *Player) CameraTargetPlayerObject() {
	panic("not implemented")
}

func (p *Player) CameraTargetVehicle() {
	panic("not implemented")
}

func (p *Player) CameraUpVector() {
	panic("not implemented")
}

func (p *Player) CameraZoom() {
	panic("not implemented")
}

func (p *Player) CheckPoint() {
	panic("not implemented")
}

func (p *Player) Class() {
	panic("not implemented")
}

func (p *Player) Color() {
	panic("not implemented")
}

func (p *Player) CustomSkin() {
	panic("not implemented")
}

func (p *Player) DialogID() {
	panic("not implemented")
}

func (p *Player) DialogData() {
	panic("not implemented")
}

func (p *Player) DistanceFromPoint() {
	panic("not implemented")
}

func (p *Player) DrunkLevel() {
	panic("not implemented")
}

func (p *Player) FacingAngle() {
	panic("not implemented")
}

func (p *Player) FightingStyle() {
	panic("not implemented")
}

func (p *Player) GhostMode() {
	panic("not implemented")
}

func (p *Player) Gravity() {
	panic("not implemented")
}

func (p *Player) Health() {
	panic("not implemented")
}

func (p *Player) HydraReactorAngle() {
	panic("not implemented")
}

func (p *Player) Interior() {
	panic("not implemented")
}

func (p *Player) IP() {
	panic("not implemented")
}

func (p *Player) Keys() {
	panic("not implemented")
}

func (p *Player) LandingGearState() {
	panic("not implemented")
}

func (p *Player) LastShotVectors() {
	panic("not implemented")
}

func (p *Player) LastSyncedTrailerID() {
	panic("not implemented")
}

func (p *Player) LastSyncedVehicleID() {
	panic("not implemented")
}

// Get the colour of a player's nametag and radar blip for another player.
func (p *Player) MarkerForPlayer(target *Player) {
	panic("not implemented")
}

func (p *Player) Menu(target *Player) {
	panic("not implemented")
}

func (p *Player) Money(target *Player) int {
	panic("not implemented")
}

func (p *Player) Name() string {
	cname := C.player_getName(p.handle)

	return C.GoString(cname)
}

func (p *Player) NetworkStats() {
	panic("not implemented")
}

func (p *Player) Position() Vector3 {
	pos := C.player_getPosition(p.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *Player) SetName(name string) PlayerNameStatus {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return PlayerNameStatus(C.player_setName(p.handle, cname))
}

func (p *Player) SendMessage(color int, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(p.handle, C.int(color), cmsg)
}

func (p *Player) Vehicle() (*Vehicle, error) {
	vehHandle := C.player_getVehicle(p.handle)

	if vehHandle == nil {
		return nil, errors.New("player is not in vehicle")
	}

	return &Vehicle{vehHandle}, nil
}
