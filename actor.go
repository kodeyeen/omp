package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/actor.h"
import "C"
import (
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

func NewActor(skin int, pos Vector3, angle float32) *Actor {
	cActor := C.actor_create(C.int(skin), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(angle))

	return &Actor{handle: cActor}
}

func FreeActor(actor *Actor) {
	C.actor_release(actor.handle)
}

func (a *Actor) SetSkin(skin int) {
	C.actor_setSkin(a.handle, C.int(skin))
}

func (a *Actor) Skin() int {
	return int(C.actor_getSkin(a.handle))
}

func (a *Actor) ApplyAnimation(anim Animation) {
	cLib := C.CString(anim.Lib)
	defer C.free(unsafe.Pointer(cLib))

	cLibStr := C.String{
		buf:    cLib,
		length: C.strlen(cLib),
	}

	cName := C.CString(anim.Name)
	defer C.free(unsafe.Pointer(cName))

	cNameStr := C.String{
		buf:    cName,
		length: C.strlen(cName),
	}

	C.actor_applyAnimation(
		a.handle,
		C.float(anim.Delta),
		boolToCUchar(anim.Loop),
		boolToCUchar(anim.LockX),
		boolToCUchar(anim.LockY),
		boolToCUchar(anim.Freeze),
		C.uint(anim.Duration.Milliseconds()),
		cLibStr,
		cNameStr,
	)
}

func (a *Actor) Animation() Animation {
	cAnim := C.actor_getAnimation(a.handle)

	return Animation{
		Lib:    C.GoStringN(cAnim.lib.buf, C.int(cAnim.lib.length)),
		Name:   C.GoStringN(cAnim.name.buf, C.int(cAnim.name.length)),
		Delta:  float32(cAnim.delta),
		Loop:   cAnim.loop != 0,
		LockX:  cAnim.lockX != 0,
		LockY:  cAnim.lockY != 0,
		Freeze: cAnim.freeze != 0,
	}
}

func (a *Actor) ClearAnimations() {
	C.actor_clearAnimations(a.handle)
}

func (a *Actor) SetHealth(health float32) {
	C.actor_setHealth(a.handle, C.float(health))
}

func (a *Actor) Health() float32 {
	return float32(C.actor_getHealth(a.handle))
}

func (a *Actor) MakeInvulnerable() {
	C.actor_setInvulnerable(a.handle, 1)
}

func (a *Actor) UnmakeInvulnerable() {
	C.actor_setInvulnerable(a.handle, 0)
}

func (a *Actor) IsInvulnerable() bool {
	return C.actor_isInvulnerable(a.handle) != 0
}

func (a *Actor) IsStreamedInFor(plr *Player) bool {
	return C.actor_isStreamedInForPlayer(a.handle, plr.handle) != 0
}

func (a *Actor) SpawnData() ActorSpawnData {
	data := C.actor_getSpawnData(a.handle)

	return ActorSpawnData{
		Position: Vector3{
			X: float32(data.position.x),
			Y: float32(data.position.y),
			Z: float32(data.position.z),
		},
		FacingAngle: float32(data.facingAngle),
		Skin:        int(data.skin),
	}
}

func (a *Actor) SetPosition(pos Vector3) {
	C.actor_setPosition(a.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (a *Actor) Position() Vector3 {
	cPos := C.actor_getPosition(a.handle)

	return Vector3{
		X: float32(cPos.x),
		Y: float32(cPos.y),
		Z: float32(cPos.z),
	}
}

func (a *Actor) SetVirtualWorld(vw int) {
	C.actor_setVirtualWorld(a.handle, C.int(vw))
}

func (a *Actor) VirtualWorld() int {
	return int(C.actor_getVirtualWorld(a.handle))
}

func (a *Actor) SetFacingAngle(angle float32) {
	C.actor_setFacingAngle(a.handle, C.float(angle))
}

func (a *Actor) FacingAngle() float32 {
	return float32(C.actor_getFacingAngle(a.handle))
}
