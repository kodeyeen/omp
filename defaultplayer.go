package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/player.h"
import "C"
import (
	"errors"
	"time"
	"unsafe"
)

type DefaultPlayer struct {
	handle unsafe.Pointer
}

func (p *DefaultPlayer) Handle() unsafe.Pointer {
	return p.handle
}

func (p *DefaultPlayer) ID() int {
	return int(C.player_getID(p.Handle()))
}

func (p *DefaultPlayer) Kick() {
	C.player_kick(p.Handle())
}

func (p *DefaultPlayer) Ban(reason string) {
	creason := C.CString(reason)
	defer C.free(unsafe.Pointer(creason))

	C.player_ban(p.Handle(), C.String{
		buf:    creason,
		length: C.strlen(creason),
	})
}

func (p *DefaultPlayer) IsBot() bool {
	return C.player_isBot(p.Handle()) != 0
}

func (p *DefaultPlayer) Ping() int {
	return int(C.player_getPing(p.Handle()))
}

func (p *DefaultPlayer) Spawn() {
	C.player_spawn(p.Handle())
}

func (p *DefaultPlayer) IsSpawned() bool {
	return C.player_isSpawned(p.Handle()) != 0
}

func (p *DefaultPlayer) ClientVersion() int {
	return int(C.player_getClientVersion(p.Handle()))
}

func (p *DefaultPlayer) ClientVersionName() string {
	verName := C.player_getClientVersionName(p.Handle())

	return C.GoStringN(verName.buf, C.int(verName.length))
}

func (p *DefaultPlayer) SetPositionFindZ(pos Vector3) {
	C.player_setPositionFindZ(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *DefaultPlayer) SetCameraPosition(pos Vector3) {
	C.player_setPositionFindZ(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *DefaultPlayer) CameraPosition() Vector3 {
	aimData := C.player_getAimData(p.Handle())

	return Vector3{
		X: float32(aimData.camPos.x),
		Y: float32(aimData.camPos.y),
		Z: float32(aimData.camPos.z),
	}
}

func (p *DefaultPlayer) SetCameraLookAt(pos Vector3, cutType PlayerCameraCutType) {
	C.player_setCameraLookAt(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(cutType))
}

func (p *DefaultPlayer) CameraLookAt() Vector3 {
	pos := C.player_getCameraLookAt(p.Handle())

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *DefaultPlayer) SetCameraBehind() {
	C.player_setCameraBehind(p.Handle())
}

func (p *DefaultPlayer) InterpolateCameraPosition(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraPosition(p.Handle(), C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

func (p *DefaultPlayer) InterpolateCameraLookAt(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraLookAt(p.Handle(), C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

func (p *DefaultPlayer) AttachCameraToObject() {
	panic("not implemented")
}

func (p *DefaultPlayer) SetName(name string) PlayerNameStatus {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	status := C.player_setName(p.Handle(), C.String{
		buf:    cname,
		length: C.strlen(cname),
	})

	return PlayerNameStatus(status)
}

func (p *DefaultPlayer) Name() string {
	name := C.player_getName(p.Handle())

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *DefaultPlayer) Serial() string {
	name := C.player_getSerial(p.Handle())

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *DefaultPlayer) GiveWeapon(weapon Weapon, ammo int) {
	C.player_giveWeapon(p.Handle(), C.WeaponSlotData{
		id:   C.uchar(weapon),
		ammo: C.uint(ammo),
	})
}

func (p *DefaultPlayer) RemoveWeapon(weapon Weapon) {
	C.player_removeWeapon(p.Handle(), C.uchar(weapon))
}

func (p *DefaultPlayer) SetWeaponAmmo(weapon Weapon, ammo int) {
	C.player_setWeaponAmmo(p.Handle(), C.WeaponSlotData{
		id:   C.uchar(weapon),
		ammo: C.uint(ammo),
	})
}

func (p *DefaultPlayer) WeaponSlots() []*WeaponSlot {
	panic("not implemented")
}

func (p *DefaultPlayer) WeaponSlot(_type WeaponSlotIndex) (WeaponSlot, error) {
	if _type < WeaponSlotIndexUnknown || _type > WeaponSlotIndexDetonator {
		return WeaponSlot{}, errors.New("invalid slot type")
	}

	slot := C.player_getWeaponSlot(p.Handle(), C.int(_type))

	return WeaponSlot{
		Weapon: Weapon(slot.id),
		Ammo:   int(slot.ammo),
	}, nil
}

func (p *DefaultPlayer) ResetWeapons() {
	C.player_resetWeapons(p.Handle())
}

func (p *DefaultPlayer) SetArmedWeapon(weapon Weapon) {
	C.player_setArmedWeapon(p.Handle(), C.uint(weapon))
}

func (p *DefaultPlayer) ArmedWeapon() Weapon {
	return Weapon(C.player_getArmedWeapon(p.Handle()))
}

// Gets the amount of ammo in a player's current weapon.
func (p *DefaultPlayer) ArmedWeaponAmmo() int {
	return int(C.player_getArmedWeaponAmmo(p.Handle()))
}

func (p *DefaultPlayer) SetShopName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.player_setShopName(p.Handle(), C.String{
		buf:    cname,
		length: C.strlen(cname),
	})
}

func (p *DefaultPlayer) ShopName() string {
	name := C.player_getShopName(p.Handle())

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *DefaultPlayer) SetDrunkLevel(level int) {
	C.player_setDrunkLevel(p.Handle(), C.int(level))
}

func (p *DefaultPlayer) DrunkLevel() int {
	return int(C.player_getDrunkLevel(p.Handle()))
}

func (p *DefaultPlayer) SetColor(color int) {
	C.player_setColour(p.Handle(), C.uint(color))
}

func (p *DefaultPlayer) Color() int {
	return int(C.player_getColour(p.Handle()))
}

func (p *DefaultPlayer) SetOtherColor(other Player, color int) {
	C.player_setOtherColour(p.Handle(), other.Handle(), C.uint(color))
}

// Get the colour of a player's nametag and radar blip for another player.
func (p *DefaultPlayer) OtherColor(other Player) (int, error) {
	var ccolor C.uint
	hasSpecificColor := C.player_getOtherColour(p.Handle(), other.Handle(), &ccolor) != 0

	if !hasSpecificColor {
		return 0, errors.New("player has no specific color")
	}

	return int(ccolor), nil
}

func (p *DefaultPlayer) Freeze() {
	C.player_setControllable(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) Unfreeze() {
	C.player_setControllable(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) IsFrozen() bool {
	return C.player_getControllable(p.Handle()) != 0
}

func (p *DefaultPlayer) ToggleSpectating(enable bool) {
	if enable {
		C.player_setSpectating(p.Handle(), C.int(1))
	} else {
		C.player_setSpectating(p.Handle(), C.int(0))
	}
}

func (p *DefaultPlayer) SetWantedLevel(level int) {
	C.player_setWantedLevel(p.Handle(), C.uint(level))
}

func (p *DefaultPlayer) WantedLevel() int {
	return int(C.player_getWantedLevel(p.Handle()))
}

func (p *DefaultPlayer) PlaySound(sound int, pos Vector3) {
	C.player_playSound(p.Handle(), C.uint(sound), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *DefaultPlayer) LastPlayedSound() int {
	return int(C.player_lastPlayedSound(p.Handle()))
}

func (p *DefaultPlayer) PlayAudio(url string, usePos bool, pos Vector3, distance float32) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))

	cstr := C.String{
		buf:    curl,
		length: C.strlen(curl),
	}

	cusePos := C.int(0)
	if usePos {
		cusePos = C.int(1)
	}

	C.player_playAudio(p.Handle(), cstr, cusePos, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(distance))
}

func (p *DefaultPlayer) PlayCrimeReport(suspect Player, crime int) {
	C.player_playerCrimeReport(p.Handle(), suspect.Handle(), C.int(crime))
}

func (p *DefaultPlayer) StopAudio() {
	C.player_stopAudio(p.Handle())
}

func (p *DefaultPlayer) LastPlayedAudio() string {
	audio := C.player_lastPlayedAudio(p.Handle())

	return C.GoStringN(audio.buf, C.int(audio.length))
}

// TODO type constants
func (p *DefaultPlayer) CreateExplosion(_type int, radius float32, pos Vector3) {
	C.player_createExplosion(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.float(radius))
}

func (p *DefaultPlayer) SendDeathMessage(player Player, killer Player, weapon int) {
	C.player_sendDeathMessage(p.Handle(), player.Handle(), killer.Handle(), C.int(weapon))
}

func (p *DefaultPlayer) SendEmptyDeathMessage() {
	C.player_sendEmptyDeathMessage(p.Handle())
}

func (p *DefaultPlayer) RemoveDefaultObjects(model int, radius float32, pos Vector3) {
	C.player_removeDefaultObjects(p.Handle(), C.uint(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius))
}

func (p *DefaultPlayer) ForceClassSelection() {
	C.player_forceClassSelection(p.Handle())
}

func (p *DefaultPlayer) SetMoney(money int) {
	C.player_setMoney(p.Handle(), C.int(money))
}

func (p *DefaultPlayer) GiveMoney(money int) {
	C.player_giveMoney(p.Handle(), C.int(money))
}

func (p *DefaultPlayer) ResetMoney() {
	C.player_resetMoney(p.Handle())
}

func (p *DefaultPlayer) Money() int {
	return int(C.player_getMoney(p.Handle()))
}

func (p *DefaultPlayer) SetMapIcon(ID int, pos Vector3, _type, color, style int) {
	C.player_setMapIcon(p.Handle(), C.int(ID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.uint(color), C.int(style))
}

func (p *DefaultPlayer) UnsetMapIcon(ID int) {
	C.player_unsetMapIcon(p.Handle(), C.int(ID))
}

func (p *DefaultPlayer) EnableStuntBonuses() {
	C.player_useStuntBonuses(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisableStuntBonuses() {
	C.player_useStuntBonuses(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) ShowNameTagFor(other Player) {
	C.player_toggleOtherNameTag(other.Handle(), p.Handle(), C.int(1))
}

func (p *DefaultPlayer) HideNameTagFor(other Player) {
	C.player_toggleOtherNameTag(other.Handle(), p.Handle(), C.int(0))
}

func (p *DefaultPlayer) SetTime(time PlayerTime) {
	C.player_setTime(p.Handle(), C.int(time.Hours), C.int(time.Minutes))
}

func (p *DefaultPlayer) Time() PlayerTime {
	ctime := C.player_getTime(p.Handle())

	return PlayerTime{
		Hours:   int(ctime.hours),
		Minutes: int(ctime.minutes),
	}
}

func (p *DefaultPlayer) ShowClock() {
	C.player_useClock(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) HideClock() {
	C.player_useClock(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsClockShown() bool {
	return C.player_hasClock(p.Handle()) != 0
}

func (p *DefaultPlayer) EnableWidescreen() {
	C.player_useWidescreen(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisableWidescreen() {
	C.player_useWidescreen(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsWidescreenEnabled() bool {
	return C.player_hasWidescreen(p.Handle()) != 0
}

func (p *DefaultPlayer) SetHealth(health float32) {
	C.player_setHealth(p.Handle(), C.float(health))
}

func (p *DefaultPlayer) Health() float32 {
	return float32(C.player_getHealth(p.Handle()))
}

func (p *DefaultPlayer) SetScore(score int) {
	C.player_setScore(p.Handle(), C.int(score))
}

func (p *DefaultPlayer) Score() int {
	return int(C.player_getScore(p.Handle()))
}

func (p *DefaultPlayer) SetArmor(armor float32) {
	C.player_setArmour(p.Handle(), C.float(armor))
}

func (p *DefaultPlayer) Armor() float32 {
	return float32(C.player_getArmour(p.Handle()))
}

func (p *DefaultPlayer) SetGravity(gravity float32) {
	C.player_setGravity(p.Handle(), C.float(gravity))
}

func (p *DefaultPlayer) Gravity() float32 {
	return float32(C.player_getGravity(p.Handle()))
}

func (p *DefaultPlayer) SetWorldTime(time int) {
	C.player_setWorldTime(p.Handle(), C.int(time))
}

func (p *DefaultPlayer) ApplyAnimation(delta float32, loop, lockX, lockY, freeze bool, duration time.Duration, lib, name string, syncType PlayerAnimationSyncType) {
	// TODO
	// C.player_applyAnimation(p.Handle(), ...)
}

func (p *DefaultPlayer) ClearAnimations(syncType PlayerAnimationSyncType) {
	// TODO player_clearAnimations
	C.player_clearTasks(p.Handle(), C.PlayerAnimationSyncType(syncType))
}

// Returns the index of any running applied animations.
func (p *DefaultPlayer) AnimationIndex() int {
	animData := C.player_getAnimationData(p.Handle())

	return int(animData.ID)
}

func (p *DefaultPlayer) AnimationFlags() int {
	animData := C.player_getAnimationData(p.Handle())

	return int(animData.flags)
}

func (p *DefaultPlayer) IsStreamedInFor(other Player) bool {
	return C.player_isStreamedInForPlayer(p.Handle(), other.Handle()) != 0
}

func (p *DefaultPlayer) State(other *Player) PlayerState {
	return PlayerState(C.player_getState(p.Handle()))
}

func (p *DefaultPlayer) SetTeam(team int) {
	C.player_setTeam(p.Handle(), C.int(team))
}

func (p *DefaultPlayer) Team() int {
	return int(C.player_getTeam(p.Handle()))
}

func (p *DefaultPlayer) Skin() int {
	return int(C.player_getSkin(p.Handle()))
}

func (p *DefaultPlayer) SetSkin(skin int) {
	C.player_setSkin(p.Handle(), C.int(skin), C.int(1))
}

func (p *DefaultPlayer) SetChatBubble(text string, color int, drawDist float32, expire time.Duration) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cstr := C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	}

	C.player_setChatBubble(p.Handle(), cstr, C.uint(color), C.float(drawDist), C.int(expire.Milliseconds()))
}

func (p *DefaultPlayer) SendMessage(msg string, color int) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(p.Handle(), C.uint(color), C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	})
}

func (p *DefaultPlayer) SendMessageFrom(other Player, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendChatMessage(p.Handle(), other.Handle(), C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	})
}

func (p *DefaultPlayer) ShowGameText(msg string, delay time.Duration, style int) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	cstr := C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	}

	C.player_sendGameText(p.Handle(), cstr, C.int(delay.Milliseconds()), C.int(style))
}

func (p *DefaultPlayer) HideGameText(style int) {
	C.player_hideGameText(p.Handle(), C.int(style))
}

func (p *DefaultPlayer) IsGameTextShown(style int) {
	C.player_hasGameText(p.Handle(), C.int(style))
}

func (p *DefaultPlayer) GameText(style int) *PlayerGameText {
	var cstr C.String
	var delay C.int
	var remaining C.int

	C.player_getGameText(p.Handle(), C.int(style), &cstr, &delay, &remaining)

	return &PlayerGameText{
		Text:      C.GoStringN(cstr.buf, C.int(cstr.length)),
		Delay:     time.Duration(delay) * time.Millisecond,
		Remaining: time.Duration(remaining) * time.Millisecond,
	}
}

func (p *DefaultPlayer) SetWeather(weather int) {
	C.player_setWeather(p.Handle(), C.int(weather))
}

func (p *DefaultPlayer) Weather() int {
	return int(C.player_getWeather(p.Handle()))
}

func (p *DefaultPlayer) SetWorldBounds(bounds Vector4) {
	C.player_setWorldBounds(p.Handle(), C.float(bounds.X), C.float(bounds.Y), C.float(bounds.Z), C.float(bounds.W))
}

func (p *DefaultPlayer) UnsetWorldBounds() {
	C.player_setWorldBounds(p.Handle(), C.float(20000.0), C.float(-20000.0), C.float(20000.0), C.float(-20000.0))
}

func (p *DefaultPlayer) WorldBounds() Vector4 {
	bounds := C.player_getWorldBounds(p.Handle())

	return Vector4{
		X: float32(bounds.x),
		Y: float32(bounds.y),
		Z: float32(bounds.z),
		W: float32(bounds.w),
	}
}

func (p *DefaultPlayer) SetFightingStyle(style PlayerFightingStyle) {
	C.player_setFightingStyle(p.Handle(), C.int(style))
}

func (p *DefaultPlayer) FightingStyle() PlayerFightingStyle {
	return PlayerFightingStyle(C.player_getFightingStyle(p.Handle()))
}

func (p *DefaultPlayer) SetSkillLevel(skill PlayerWeaponSkill, level int) {
	C.player_setSkillLevel(p.Handle(), C.int(skill), C.int(level))
}

func (p *DefaultPlayer) SetAction(action PlayerSpecialAction) {
	C.player_setAction(p.Handle(), C.int(action))
}

func (p *DefaultPlayer) Action() PlayerSpecialAction {
	return PlayerSpecialAction(C.player_getAction(p.Handle()))
}

func (p *DefaultPlayer) SetVelocity(velocity Vector3) {
	C.player_setVelocity(p.Handle(), C.float(velocity.X), C.float(velocity.Y), C.float(velocity.Z))
}

func (p *DefaultPlayer) Velocity() Vector3 {
	vel := C.player_getVelocity(p.Handle())

	return Vector3{
		X: float32(vel.x),
		Y: float32(vel.y),
		Z: float32(vel.z),
	}
}

func (p *DefaultPlayer) SetInterior(interior int) {
	C.player_setInterior(p.Handle(), C.uint(interior))
}

func (p *DefaultPlayer) Interior() int {
	return int(C.player_getInterior(p.Handle()))
}

func (p *DefaultPlayer) KeyData() PlayerKeyData {
	data := C.player_getKeyData(p.Handle())

	return PlayerKeyData{
		Keys:      int(data.keys),
		UpDown:    int(data.upDown),
		LeftRight: int(data.leftRight),
	}
}

func (p *DefaultPlayer) WeaponState() PlayerWeaponState {
	aimData := C.player_getAimData(p.Handle())

	return PlayerWeaponState(aimData.weaponState)
}

func (p *DefaultPlayer) CameraAspectRatio() float32 {
	aimData := C.player_getAimData(p.Handle())

	return float32(aimData.aspectRatio)
}

func (p *DefaultPlayer) CameraFrontVector() Vector3 {
	aimData := C.player_getAimData(p.Handle())

	return Vector3{
		X: float32(aimData.camFrontVector.x),
		Y: float32(aimData.camFrontVector.y),
		Z: float32(aimData.camFrontVector.z),
	}
}

// TODO constants
func (p *DefaultPlayer) CameraMode() int {
	aimData := C.player_getAimData(p.Handle())

	return int(aimData.camMode)
}

func (p *DefaultPlayer) CameraZoom() float32 {
	aimData := C.player_getAimData(p.Handle())

	return float32(aimData.camZoom)
}

func (p *DefaultPlayer) AimZ() float32 {
	aimData := C.player_getAimData(p.Handle())

	return float32(aimData.aimZ)
}

// TODO getPlayerBulletData

func (p *DefaultPlayer) EnableCameraTargetting() {
	C.player_useCameraTargetting(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisableCameraTargetting() {
	C.player_useCameraTargetting(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsCameraTargettingEnabled() bool {
	return C.player_hasCameraTargetting(p.Handle()) != 0
}

func (p *DefaultPlayer) RemoveFromVehicle(force bool) {
	var cforce C.int
	if force {
		cforce = C.int(1)
	}

	C.player_removeFromVehicle(p.Handle(), cforce)
}

func (p *DefaultPlayer) CameraTargetPlayer() *DefaultPlayer {
	player := C.player_getCameraTargetPlayer(p.Handle())

	return &DefaultPlayer{player}
}

func (p *DefaultPlayer) CameraTargetVehicle() *DefaultVehicle {
	vehicle := C.player_getCameraTargetVehicle(p.Handle())

	return &DefaultVehicle{handle: vehicle}
}

func (p *DefaultPlayer) CameraTargetObject() Object {
	object := C.player_getCameraTargetObject(p.Handle())

	return &GlobalObject{handle: object}
}

func (p *DefaultPlayer) CameraTargetActor() *Actor {
	actor := C.player_getCameraTargetActor(p.Handle())

	return &Actor{actor}
}

func (p *DefaultPlayer) TargetPlayer() *DefaultPlayer {
	player := C.player_getTargetPlayer(p.Handle())

	return &DefaultPlayer{player}
}

func (p *DefaultPlayer) TargetActor() *Actor {
	actor := C.player_getTargetActor(p.Handle())

	return &Actor{actor}
}

func (p *DefaultPlayer) EnableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) SpectatePlayer(player Player, mode PlayerSpectateMode) {
	C.player_spectatePlayer(p.Handle(), player.Handle(), C.int(mode))
}

func (p *DefaultPlayer) SpectateVehicle(vehicle Vehicle, mode PlayerSpectateMode) {
	C.player_spectateVehicle(p.Handle(), vehicle.Handle(), C.int(mode))
}

// TODO
// func (p *DefaultPlayer) SpectatingPlayer() *DefaultPlayer {
// 	specData := C.player_getSpectateData(p.Handle())
// }

// func (p *DefaultPlayer) SpectatingVehicle() *Vehicle {
// 	specData := C.player_getSpectateData(p.Handle())
// }

// TODO callback
func (p *DefaultPlayer) SendClientCheck(actionType, address, offset, count int) {
	C.player_sendClientCheck(p.Handle(), C.int(actionType), C.int(address), C.int(offset), C.int(count))
}

func (p *DefaultPlayer) EnableGhostMode() {
	C.player_toggleGhostMode(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisableGhostMode() {
	C.player_toggleGhostMode(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsGhostModeEnabled() bool {
	return C.player_isGhostModeEnabled(p.Handle()) != 0
}

func (p *DefaultPlayer) RemovedBuildingCount() int {
	return int(C.player_getDefaultObjectsRemoved(p.Handle()))
}

func (p *DefaultPlayer) AllowWeapons() {
	C.player_allowWeapons(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisallowWeapons() {
	C.player_allowWeapons(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) AreWeaponsAllowed() bool {
	return C.player_areWeaponsAllowed(p.Handle()) != 0
}

func (p *DefaultPlayer) AllowTeleport() {
	C.player_allowTeleport(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) DisallowTeleport() {
	C.player_allowTeleport(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsTeleportAllowed() bool {
	return C.player_isTeleportAllowed(p.Handle()) != 0
}

func (p *DefaultPlayer) IsUsingOfficialClient() bool {
	return C.player_isUsingOfficialClient(p.Handle()) != 0
}

// entity

func (p *DefaultPlayer) SetPosition(pos Vector3) {
	C.player_setPosition(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *DefaultPlayer) Position() Vector3 {
	pos := C.player_getPosition(p.Handle())

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *DefaultPlayer) Rotation() Vector4 {
	rquat := C.player_getRotation(p.Handle())

	return Vector4{
		X: float32(rquat.x),
		Y: float32(rquat.y),
		Z: float32(rquat.z),
		W: float32(rquat.w),
	}
}

func (p *DefaultPlayer) SetVirtualWorld(vw int) {
	C.player_setVirtualWorld(p.Handle(), C.int(vw))
}

func (p *DefaultPlayer) VirtualWorld() int {
	return int(C.player_getVirtualWorld(p.Handle()))
}

// checkpoint data

func (p *DefaultPlayer) NewDefaultCheckpoint(radius float32, pos Vector3) *DefaultCheckpoint {
	cp := C.player_setCheckpoint(p.Handle(), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius))

	return &DefaultCheckpoint{cp}
}

// console data

func (p *DefaultPlayer) MakeAdmin() {
	C.player_setConsoleAccessibility(p.Handle(), C.int(1))
}

func (p *DefaultPlayer) UnmakeAdmin() {
	C.player_setConsoleAccessibility(p.Handle(), C.int(0))
}

func (p *DefaultPlayer) IsAdmin() bool {
	return C.player_hasConsoleAccess(p.Handle()) != 0
}

// custom models data

func (p *DefaultPlayer) CustomSkin() int {
	return int(C.player_getCustomSkin(p.Handle()))
}

// network data

func (p *DefaultPlayer) IP() string {
	IP := C.player_getIp(p.Handle())

	return C.GoStringN(IP.buf, C.int(IP.length))
}

func (p *DefaultPlayer) RawIP() int {
	return int(C.player_getRawIp(p.Handle()))
}

// vehicle data

func (p *DefaultPlayer) Vehicle() (*DefaultVehicle, error) {
	vehicle := C.player_getVehicle(p.Handle())

	if vehicle == nil {
		return nil, errors.New("player is not in a vehicle")
	}

	return &DefaultVehicle{handle: vehicle}, nil
}

func (p *DefaultPlayer) VehicleSeat() int {
	return int(C.player_getSeat(p.Handle()))
}

// misc

func (p *DefaultPlayer) DistanceFrom(point Vector3) float32 {
	return float32(C.player_getDistanceFromPoint(p.Handle(), C.float(point.X), C.float(point.Y), C.float(point.Z)))
}

func (p *DefaultPlayer) IsInRangeOf(point Vector3, _range float32) bool {
	return C.player_isInRangeOfPoint(p.Handle(), C.float(_range), C.float(point.X), C.float(point.Y), C.float(point.Z)) != 0
}

func (p *DefaultPlayer) SetFacingAngle(angle float32) {
	C.player_setFacingAngle(p.Handle(), C.float(angle))
}

func (p *DefaultPlayer) FacingAngle() float32 {
	return float32(C.player_getFacingAngle(p.Handle()))
}

func (p *DefaultPlayer) CheckPoint() {
	panic("not implemented")
}

func (p *DefaultPlayer) DialogID() {
	panic("not implemented")
}

func (p *DefaultPlayer) DialogData() {
	panic("not implemented")
}

func (p *DefaultPlayer) Menu(target *Player) {
	panic("not implemented")
}

func (p *DefaultPlayer) NetworkStats() {
	panic("not implemented")
}

// TODO
// ...PlayerCheckpoint...
// ...PlayerRaceCheckpoint...

// GetPlayerAttachedObject

// RedirectDownload
// GetPlayerDialogID
// GetPlayerDialogData

// GetPlayerLastShotVectors

// GetPlayerMenu
// GetPlayerObject[X]
// GetPlayerPickup[X]

// GetPlayerSkillLevel
// GetPlayerSurfing[X]
// GetPlayerWeaponData
// HidePlayerDialog

// SetPickupForPlayer

// SetPlayerAttachedObject
// Show[X]ForPlayer

// StartRecordingPlayerData
// StopRecordingPlayerData
// UsePlayerGangZoneCheck
