package gomp

// #include "include/class.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Class struct {
	handle unsafe.Pointer
}

func NewClass(
	team,
	skin int,
	spawnPos Vector3,
	angle float32,
	weapon1 Weapon,
	ammo1 int,
	weapon2 Weapon,
	ammo2 int,
	weapon3 Weapon,
	ammo3 int,
) (*Class, error) {
	cClass := C.class_create(&C.ClassData{
		team:    C.int(team),
		skin:    C.int(skin),
		spawnX:  C.float(spawnPos.X),
		spawnY:  C.float(spawnPos.Y),
		spawnZ:  C.float(spawnPos.Z),
		angle:   C.float(angle),
		weapon1: C.uchar(weapon1),
		ammo1:   C.uint(ammo1),
		weapon2: C.uchar(weapon2),
		ammo2:   C.uint(ammo2),
		weapon3: C.uchar(weapon3),
		ammo3:   C.uint(ammo3),
	})
	if cClass == nil {
		return nil, errors.New("class limit reached")
	}

	return &Class{handle: cClass}, nil
}

func FreeClass(cls *Class) {
	C.class_release(cls.handle)
}

func (c *Class) ID() int {
	return int(C.class_getID(c.handle))
}

func (c *Class) SetTeam(team int) {
	data := C.class_getClass(c.handle)
	data.team = C.int(team)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Team() int {
	data := C.class_getClass(c.handle)

	return int(data.team)
}

func (c *Class) SetSkin(skin int) {
	data := C.class_getClass(c.handle)
	data.skin = C.int(skin)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Skin() int {
	data := C.class_getClass(c.handle)

	return int(data.skin)
}

func (c *Class) SetSpawnPosition(pos Vector3) {
	data := C.class_getClass(c.handle)
	data.spawnX = C.float(pos.X)
	data.spawnY = C.float(pos.Y)
	data.spawnZ = C.float(pos.Z)

	C.class_setClass(c.handle, &data)
}

func (c *Class) SpawnPosition() Vector3 {
	data := C.class_getClass(c.handle)

	return Vector3{
		X: float32(data.spawnX),
		Y: float32(data.spawnY),
		Z: float32(data.spawnZ),
	}
}

func (c *Class) SetAngle(angle float32) {
	data := C.class_getClass(c.handle)
	data.angle = C.float(angle)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Angle() float32 {
	data := C.class_getClass(c.handle)

	return float32(data.angle)
}

func (c *Class) SetWeapon1(weapon1 Weapon) {
	data := C.class_getClass(c.handle)
	data.weapon1 = C.uchar(weapon1)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Weapon1() Weapon {
	data := C.class_getClass(c.handle)

	return Weapon(data.weapon1)
}

func (c *Class) SetAmmo1(ammo1 int) {
	data := C.class_getClass(c.handle)
	data.ammo1 = C.uint(ammo1)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Ammo1() int {
	data := C.class_getClass(c.handle)

	return int(data.ammo1)
}

func (c *Class) SetWeapon2(weapon2 Weapon) {
	data := C.class_getClass(c.handle)
	data.weapon2 = C.uchar(weapon2)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Weapon2() Weapon {
	data := C.class_getClass(c.handle)

	return Weapon(data.weapon2)
}

func (c *Class) SetAmmo2(ammo2 int) {
	data := C.class_getClass(c.handle)
	data.ammo2 = C.uint(ammo2)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Ammo2() int {
	data := C.class_getClass(c.handle)

	return int(data.ammo2)
}

func (c *Class) SetWeapon3(weapon3 Weapon) {
	data := C.class_getClass(c.handle)
	data.weapon3 = C.uchar(weapon3)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Weapon3() Weapon {
	data := C.class_getClass(c.handle)

	return Weapon(data.weapon3)
}

func (c *Class) SetAmmo3(ammo3 int) {
	data := C.class_getClass(c.handle)
	data.ammo3 = C.uint(ammo3)

	C.class_setClass(c.handle, &data)
}

func (c *Class) Ammo3() int {
	data := C.class_getClass(c.handle)

	return int(data.ammo3)
}
