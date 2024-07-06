package omp

// #include <stdlib.h>
// #include "include/wrappers.h"
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

type ActorSpawnData struct {
	Position    Vector3
	FacingAngle float32
	Skin        int
}

type Actor struct {
	handle unsafe.Pointer
}

func NewActor(skin int, pos Vector3, angle float32) (*Actor, error) {
	var cID C.int
	cActor := C.Actor_Create(C.int(skin), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle), &cID)
	if cActor == nil {
		return nil, errors.New("actor limit was reached")
	}

	return &Actor{handle: cActor}, nil
}

func FreeActor(actor *Actor) {
	C.Actor_Destroy(actor.handle)
}

func (a *Actor) SetSkin(skin int) {
	C.Actor_SetSkin(a.handle, C.int(skin))
}

func (a *Actor) Skin() int {
	return int(C.Actor_GetSkin(a.handle))
}

func (a *Actor) ApplyAnimation(anim Animation) {
	cLib := C.CString(anim.Lib)
	defer C.free(unsafe.Pointer(cLib))

	cName := C.CString(anim.Name)
	defer C.free(unsafe.Pointer(cName))

	C.Actor_ApplyAnimation(
		a.handle,
		cName,
		cLib,
		C.float(anim.Delta),
		C.bool(anim.Loop),
		C.bool(anim.LockX),
		C.bool(anim.LockY),
		C.bool(anim.Freeze),
		C.int(anim.Duration.Milliseconds()),
	)
}

func (a *Actor) Animation() Animation {
	var (
		cLib    C.struct_CAPIStringView
		cName   C.struct_CAPIStringView
		cDelta  C.float
		cLoop   C.bool
		cLockX  C.bool
		cLockY  C.bool
		cFreeze C.bool
		cTime   C.int
	)

	C.Actor_GetAnimation(a.handle, &cLib, &cName, &cDelta, &cLoop, &cLockX, &cLockY, &cFreeze, &cTime)

	return Animation{
		Lib:      C.GoStringN(cLib.data, C.int(cLib.len)),
		Name:     C.GoStringN(cName.data, C.int(cName.len)),
		Delta:    float32(cDelta),
		Loop:     bool(cLoop),
		LockX:    bool(cLockX),
		LockY:    bool(cLockY),
		Freeze:   bool(cFreeze),
		Duration: time.Duration(cTime) * time.Millisecond,
	}
}

func (a *Actor) ClearAnimations() {
	C.Actor_ClearAnimations(a.handle)
}

func (a *Actor) SetHealth(health float32) {
	C.Actor_SetHealth(a.handle, C.float(health))
}

func (a *Actor) Health() float32 {
	return float32(C.Actor_GetHealth(a.handle))
}

func (a *Actor) MakeInvulnerable() {
	C.Actor_SetInvulnerable(a.handle, true)
}

func (a *Actor) UnmakeInvulnerable() {
	C.Actor_SetInvulnerable(a.handle, false)
}

func (a *Actor) IsInvulnerable() bool {
	return bool(C.Actor_IsInvulnerable(a.handle))
}

func (a *Actor) IsStreamedInFor(player *Player) bool {
	return bool(C.Actor_IsStreamedInFor(a.handle, player.handle))
}

func (a *Actor) SpawnData() ActorSpawnData {
	var cPosX, cPosY, cPosZ, cAngle C.float
	var cSkin C.int

	C.Actor_GetSpawnInfo(a.handle, &cPosX, &cPosY, &cPosZ, &cAngle, &cSkin)

	return ActorSpawnData{
		Position: Vector3{
			X: float32(cPosX),
			Y: float32(cPosY),
			Z: float32(cPosZ),
		},
		FacingAngle: float32(cAngle),
		Skin:        int(cSkin),
	}
}

func (a *Actor) SetPosition(pos Vector3) {
	C.Actor_SetPos(a.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (a *Actor) Position() Vector3 {
	var cPosX, cPosY, cPosZ C.float

	C.Actor_GetPos(a.handle, &cPosX, &cPosY, &cPosZ)

	return Vector3{
		X: float32(cPosX),
		Y: float32(cPosY),
		Z: float32(cPosZ),
	}
}

func (a *Actor) SetVirtualWorld(vw int) {
	C.Actor_SetVirtualWorld(a.handle, C.int(vw))
}

func (a *Actor) VirtualWorld() int {
	return int(C.Actor_GetVirtualWorld(a.handle))
}

func (a *Actor) SetFacingAngle(angle float32) {
	C.Actor_SetFacingAngle(a.handle, C.float(angle))
}

func (a *Actor) FacingAngle() float32 {
	return float32(C.Actor_GetFacingAngle(a.handle))
}
