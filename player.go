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

type PlayerNameStatus int

const (
	PlayerNameStatusUpdated PlayerNameStatus = iota
	PlayerNameStatusTaken
	PlayerNameStatusInvalid
)

type PlayerCameraCutType int

const (
	PlayerCameraCutTypeCut PlayerCameraCutType = iota
	PlayerCameraCutTypeMove
)

type PlayerTime struct {
	Hours   int
	Minutes int
}

type PlayerState int

const (
	PlayerStateNone PlayerState = iota
	PlayerStateOnFoot
	PlayerStateDriver
	PlayerStatePassenger
	PlayerStateExitVehicle
	PlayerStateEnterVehicleDriver
	PlayerStateEnterVehiclePassenger
	PlayerStateWasted
	PlayerStateSpawned
	PlayerStateSpectating
)

type PlayerGameText struct {
	Text      string
	Delay     time.Duration
	Remaining time.Duration
}

type PlayerFightingStyle int

const (
	PlayerFightingStyleNormal   PlayerFightingStyle = 4
	PlayerFightingStyleBoxing   PlayerFightingStyle = 5
	PlayerFightingStyleKungfu   PlayerFightingStyle = 6
	PlayerFightingStyleKneeHead PlayerFightingStyle = 7
	PlayerFightingStyleGrabKick PlayerFightingStyle = 15
	PlayerFightingStyleElbow    PlayerFightingStyle = 16
)

type PlayerWeaponSkill int

const (
	PlayerWeaponSkillPistol = iota
	PlayerWeaponSkillSilencedPistol
	PlayerWeaponSkillDesertEagle
	PlayerWeaponSkillShotgun
	PlayerWeaponSkillSawnOff
	PlayerWeaponSkillSPAS12
	PlayerWeaponSkillUzi
	PlayerWeaponSkillMP5
	PlayerWeaponSkillAK47
	PlayerWeaponSkillM4
	PlayerWeaponSkillSniper
)

type PlayerSpecialAction int

const (
	PlayerSpecialActionNone PlayerSpecialAction = iota
	PlayerSpecialActionDuck
	PlayerSpecialActionJetpack
	PlayerSpecialActionEnterVehicle
	PlayerSpecialActionExitVehicle
	PlayerSpecialActionDance1
	PlayerSpecialActionDance2
	PlayerSpecialActionDance3
	PlayerSpecialActionDance4
)

const (
	PlayerSpecialActionHandsUp PlayerSpecialAction = iota + 10
	PlayerSpecialActionCellphone
	PlayerSpecialActionSitting
	PlayerSpecialActionStopCellphone
)

const (
	PlayerSpecialActionBeer PlayerSpecialAction = iota + 20
	PlayerSpecialActionSmoke
	PlayerSpecialActionWine
	PlayerSpecialActionSprunk
	PlayerSpecialActionCuffed
	PlayerSpecialActionCarry
)

const PlayerSpecialActionPissing PlayerSpecialAction = 68

type PlayerKeyData struct {
	Keys      int
	UpDown    int
	LeftRight int
}

type PlayerKey int

const (
	PlayerKeyAction          PlayerKey = 1
	PlayerKeyCrouch          PlayerKey = 2
	PlayerKeyFire            PlayerKey = 4
	PlayerKeySprint          PlayerKey = 8
	PlayerKeySecondaryAttack PlayerKey = 16
	PlayerKeyJump            PlayerKey = 32
	PlayerKeyLookRight       PlayerKey = 64
	PlayerKeyHandbrake       PlayerKey = 128
	PlayerKeyAim             PlayerKey = 128
	PlayerKeyLookLeft        PlayerKey = 256
	PlayerKeyLookBehind      PlayerKey = 512
	PlayerKeySubmission      PlayerKey = 512
	PlayerKeyWalk            PlayerKey = 1024
	PlayerKeyAnalogUp        PlayerKey = 2048
	PlayerKeyAnalogDown      PlayerKey = 4096
	PlayerKeyAnalogLeft      PlayerKey = 8192
	PlayerKeyAnalogRight     PlayerKey = 16384
	PlayerKeyYes             PlayerKey = 65536
	PlayerKeyNo              PlayerKey = 131072
	PlayerKeyCtrlBack        PlayerKey = 262144
	PlayerKeyUp              PlayerKey = -128
	PlayerKeyDown            PlayerKey = 128
	PlayerKeyLeft            PlayerKey = -128
	PlayerKeyRight           PlayerKey = 128
)

type PlayerWeaponState int

const (
	PlayerWeaponStateUnknown     PlayerWeaponState = -1
	PlayerWeaponStateNoBullets   PlayerWeaponState = 0
	PlayerWeaponStateLastBullet  PlayerWeaponState = 1
	PlayerWeaponStateMoreBullets PlayerWeaponState = 2
	PlayerWeaponStateReloading   PlayerWeaponState = 3
)

type PlayerSpectateMode int

const (
	PlayerSpectateModeNormal PlayerSpectateMode = iota + 1
	PlayerSpectateModeFixed
	PlayerSpectateModeSide
)

type PlayerAnimationSyncType int

const (
	PlayerAnimationSyncTypeNoSync PlayerAnimationSyncType = iota
	PlayerAnimationSyncTypeSync
	PlayerAnimationSyncTypeSyncOthers
)

type PlayerMarkerMode int

const (
	PlayerMarkerModeOff PlayerMarkerMode = iota
	PlayerMarkerModeGlobal
	PlayerMarkerModeStreamed
)

type Player struct {
	handle unsafe.Pointer
}

func (p *Player) ID() int {
	return int(C.player_getID(p.handle))
}

func (p *Player) Kick() {
	C.player_kick(p.handle)
}

func (p *Player) Ban(reason string) {
	creason := C.CString(reason)
	defer C.free(unsafe.Pointer(creason))

	C.player_ban(p.handle, C.String{
		buf:    creason,
		length: C.strlen(creason),
	})
}

func (p *Player) IsBot() bool {
	return C.player_isBot(p.handle) != 0
}

func (p *Player) Ping() int {
	return int(C.player_getPing(p.handle))
}

func (p *Player) Spawn() {
	C.player_spawn(p.handle)
}

func (p *Player) IsSpawned() bool {
	return C.player_isSpawned(p.handle) != 0
}

func (p *Player) ClientVersion() int {
	return int(C.player_getClientVersion(p.handle))
}

func (p *Player) ClientVersionName() string {
	verName := C.player_getClientVersionName(p.handle)

	return C.GoStringN(verName.buf, C.int(verName.length))
}

func (p *Player) SetPositionFindZ(pos Vector3) {
	C.player_setPositionFindZ(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *Player) SetCameraPosition(pos Vector3) {
	C.player_setPositionFindZ(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *Player) CameraPosition() Vector3 {
	aimData := C.player_getAimData(p.handle)

	return Vector3{
		X: float32(aimData.camPos.x),
		Y: float32(aimData.camPos.y),
		Z: float32(aimData.camPos.z),
	}
}

func (p *Player) SetCameraLookAt(pos Vector3, cutType PlayerCameraCutType) {
	C.player_setCameraLookAt(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(cutType))
}

func (p *Player) CameraLookAt() Vector3 {
	pos := C.player_getCameraLookAt(p.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *Player) SetCameraBehind() {
	C.player_setCameraBehind(p.handle)
}

func (p *Player) InterpolateCameraPosition(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraPosition(p.handle, C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

func (p *Player) InterpolateCameraLookAt(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraLookAt(p.handle, C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

func (p *Player) AttachCameraToObject() {
	panic("not implemented")
}

func (p *Player) SetName(name string) PlayerNameStatus {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	status := C.player_setName(p.handle, C.String{
		buf:    cname,
		length: C.strlen(cname),
	})

	return PlayerNameStatus(status)
}

func (p *Player) Name() string {
	name := C.player_getName(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *Player) Serial() string {
	name := C.player_getSerial(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *Player) GiveWeapon(weapon Weapon, ammo int) {
	C.player_giveWeapon(p.handle, C.WeaponSlotData{
		id:   C.uchar(weapon),
		ammo: C.uint(ammo),
	})
}

func (p *Player) RemoveWeapon(weapon Weapon) {
	C.player_removeWeapon(p.handle, C.uchar(weapon))
}

func (p *Player) SetWeaponAmmo(weapon Weapon, ammo int) {
	C.player_setWeaponAmmo(p.handle, C.WeaponSlotData{
		id:   C.uchar(weapon),
		ammo: C.uint(ammo),
	})
}

func (p *Player) WeaponSlots() []*WeaponSlot {
	panic("not implemented")
}

func (p *Player) WeaponSlot(_type WeaponSlotIndex) (WeaponSlot, error) {
	if _type < WeaponSlotIndexUnknown || _type > WeaponSlotIndexDetonator {
		return WeaponSlot{}, errors.New("invalid slot type")
	}

	slot := C.player_getWeaponSlot(p.handle, C.int(_type))

	return WeaponSlot{
		Weapon: Weapon(slot.id),
		Ammo:   int(slot.ammo),
	}, nil
}

func (p *Player) ResetWeapons() {
	C.player_resetWeapons(p.handle)
}

func (p *Player) SetArmedWeapon(weapon Weapon) {
	C.player_setArmedWeapon(p.handle, C.uint(weapon))
}

func (p *Player) ArmedWeapon() Weapon {
	return Weapon(C.player_getArmedWeapon(p.handle))
}

// Gets the amount of ammo in a player's current weapon.
func (p *Player) ArmedWeaponAmmo() int {
	return int(C.player_getArmedWeaponAmmo(p.handle))
}

func (p *Player) SetShopName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.player_setShopName(p.handle, C.String{
		buf:    cname,
		length: C.strlen(cname),
	})
}

func (p *Player) ShopName() string {
	name := C.player_getShopName(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

func (p *Player) SetDrunkLevel(level int) {
	C.player_setDrunkLevel(p.handle, C.int(level))
}

func (p *Player) DrunkLevel() int {
	return int(C.player_getDrunkLevel(p.handle))
}

func (p *Player) SetColor(color uint) {
	C.player_setColour(p.handle, C.uint(color))
}

func (p *Player) Color() uint {
	return uint(C.player_getColour(p.handle))
}

func (p *Player) SetOtherColor(other *Player, color uint) {
	C.player_setOtherColour(p.handle, other.handle, C.uint(color))
}

// Get the colour of a player's nametag and radar blip for another player.
func (p *Player) OtherColor(other *Player) (int, error) {
	var ccolor C.uint
	hasSpecificColor := C.player_getOtherColour(p.handle, other.handle, &ccolor) != 0

	if !hasSpecificColor {
		return 0, errors.New("player has no specific color")
	}

	return int(ccolor), nil
}

func (p *Player) Freeze() {
	C.player_setControllable(p.handle, 0)
}

func (p *Player) Unfreeze() {
	C.player_setControllable(p.handle, 1)
}

func (p *Player) IsFrozen() bool {
	return C.player_getControllable(p.handle) != 0
}

func (p *Player) EnableSpectating() {
	C.player_setSpectating(p.handle, 1)
}

func (p *Player) DisableSpectating() {
	C.player_setSpectating(p.handle, 0)
}

func (p *Player) SetWantedLevel(level int) {
	C.player_setWantedLevel(p.handle, C.uint(level))
}

func (p *Player) WantedLevel() int {
	return int(C.player_getWantedLevel(p.handle))
}

func (p *Player) PlaySound(sound int, pos Vector3) {
	C.player_playSound(p.handle, C.uint(sound), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *Player) LastPlayedSound() int {
	return int(C.player_lastPlayedSound(p.handle))
}

func (p *Player) PlayAudio(url string, usePos bool, pos Vector3, distance float32) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))

	cstr := C.String{
		buf:    curl,
		length: C.strlen(curl),
	}

	var cusePos C.int
	if usePos {
		cusePos = 1
	}

	C.player_playAudio(p.handle, cstr, cusePos, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(distance))
}

func (p *Player) PlayCrimeReport(suspect *Player, crime int) {
	C.player_playerCrimeReport(p.handle, suspect.handle, C.int(crime))
}

func (p *Player) StopAudio() {
	C.player_stopAudio(p.handle)
}

func (p *Player) LastPlayedAudio() string {
	audio := C.player_lastPlayedAudio(p.handle)

	return C.GoStringN(audio.buf, C.int(audio.length))
}

// TODO type constants
func (p *Player) CreateExplosion(_type int, radius float32, pos Vector3) {
	C.player_createExplosion(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.float(radius))
}

func (p *Player) SendDeathMessage(player *Player, killer *Player, weapon int) {
	C.player_sendDeathMessage(p.handle, player.handle, killer.handle, C.int(weapon))
}

func (p *Player) SendEmptyDeathMessage() {
	C.player_sendEmptyDeathMessage(p.handle)
}

func (p *Player) RemoveDefaultObjects(model int, radius float32, pos Vector3) {
	C.player_removeDefaultObjects(p.handle, C.uint(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius))
}

func (p *Player) ForceClassSelection() {
	C.player_forceClassSelection(p.handle)
}

func (p *Player) SetMoney(money int) {
	C.player_setMoney(p.handle, C.int(money))
}

func (p *Player) GiveMoney(money int) {
	C.player_giveMoney(p.handle, C.int(money))
}

func (p *Player) ResetMoney() {
	C.player_resetMoney(p.handle)
}

func (p *Player) Money() int {
	return int(C.player_getMoney(p.handle))
}

func (p *Player) SetMapIcon(ID int, pos Vector3, _type, color, style int) {
	C.player_setMapIcon(p.handle, C.int(ID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.uint(color), C.int(style))
}

func (p *Player) UnsetMapIcon(ID int) {
	C.player_unsetMapIcon(p.handle, C.int(ID))
}

func (p *Player) EnableStuntBonuses() {
	C.player_useStuntBonuses(p.handle, 1)
}

func (p *Player) DisableStuntBonuses() {
	C.player_useStuntBonuses(p.handle, 0)
}

func (p *Player) ShowNameTagFor(other *Player) {
	C.player_toggleOtherNameTag(other.handle, p.handle, 1)
}

func (p *Player) HideNameTagFor(other *Player) {
	C.player_toggleOtherNameTag(other.handle, p.handle, 0)
}

func (p *Player) SetTime(time PlayerTime) {
	C.player_setTime(p.handle, C.int(time.Hours), C.int(time.Minutes))
}

func (p *Player) Time() PlayerTime {
	ctime := C.player_getTime(p.handle)

	return PlayerTime{
		Hours:   int(ctime.hours),
		Minutes: int(ctime.minutes),
	}
}

func (p *Player) ShowClock() {
	C.player_useClock(p.handle, 1)
}

func (p *Player) HideClock() {
	C.player_useClock(p.handle, 0)
}

func (p *Player) IsClockShown() bool {
	return C.player_hasClock(p.handle) != 0
}

func (p *Player) EnableWidescreen() {
	C.player_useWidescreen(p.handle, 1)
}

func (p *Player) DisableWidescreen() {
	C.player_useWidescreen(p.handle, 0)
}

func (p *Player) IsWidescreenEnabled() bool {
	return C.player_hasWidescreen(p.handle) != 0
}

func (p *Player) SetHealth(health float32) {
	C.player_setHealth(p.handle, C.float(health))
}

func (p *Player) Health() float32 {
	return float32(C.player_getHealth(p.handle))
}

func (p *Player) SetScore(score int) {
	C.player_setScore(p.handle, C.int(score))
}

func (p *Player) Score() int {
	return int(C.player_getScore(p.handle))
}

func (p *Player) SetArmor(armor float32) {
	C.player_setArmour(p.handle, C.float(armor))
}

func (p *Player) Armor() float32 {
	return float32(C.player_getArmour(p.handle))
}

func (p *Player) SetGravity(gravity float32) {
	C.player_setGravity(p.handle, C.float(gravity))
}

func (p *Player) Gravity() float32 {
	return float32(C.player_getGravity(p.handle))
}

func (p *Player) SetWorldTime(time int) {
	C.player_setWorldTime(p.handle, C.int(time))
}

func (p *Player) ApplyAnimation(delta float32, loop, lockX, lockY, freeze bool, duration time.Duration, lib, name string, syncType PlayerAnimationSyncType) {
	// TODO
	// C.player_applyAnimation(p.handle, ...)
}

func (p *Player) ClearAnimations(syncType PlayerAnimationSyncType) {
	// TODO player_clearAnimations
	C.player_clearTasks(p.handle, C.PlayerAnimationSyncType(syncType))
}

// Returns the index of any running applied animations.
func (p *Player) AnimationIndex() int {
	animData := C.player_getAnimationData(p.handle)

	return int(animData.ID)
}

func (p *Player) AnimationFlags() int {
	animData := C.player_getAnimationData(p.handle)

	return int(animData.flags)
}

func (p *Player) IsStreamedInFor(other *Player) bool {
	return C.player_isStreamedInForPlayer(p.handle, other.handle) != 0
}

func (p *Player) State() PlayerState {
	return PlayerState(C.player_getState(p.handle))
}

func (p *Player) SetTeam(team int) {
	C.player_setTeam(p.handle, C.int(team))
}

func (p *Player) Team() int {
	return int(C.player_getTeam(p.handle))
}

func (p *Player) Skin() int {
	return int(C.player_getSkin(p.handle))
}

func (p *Player) SetSkin(skin int) {
	C.player_setSkin(p.handle, C.int(skin), 1)
}

func (p *Player) SetChatBubble(text string, color uint, drawDist float32, expire time.Duration) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cstr := C.String{
		buf:    ctext,
		length: C.strlen(ctext),
	}

	C.player_setChatBubble(p.handle, cstr, C.uint(color), C.float(drawDist), C.int(expire.Milliseconds()))
}

func (p *Player) SendMessage(msg string, color Color) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(p.handle, C.uint(color), C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	})
}

func (p *Player) SendMessageFrom(other *Player, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendChatMessage(p.handle, other.handle, C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	})
}

func (p *Player) ShowGameText(msg string, delay time.Duration, style int) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	cstr := C.String{
		buf:    cmsg,
		length: C.strlen(cmsg),
	}

	C.player_sendGameText(p.handle, cstr, C.int(delay.Milliseconds()), C.int(style))
}

func (p *Player) HideGameText(style int) {
	C.player_hideGameText(p.handle, C.int(style))
}

func (p *Player) IsGameTextShown(style int) {
	C.player_hasGameText(p.handle, C.int(style))
}

func (p *Player) GameText(style int) *PlayerGameText {
	var cstr C.String
	var delay C.int
	var remaining C.int

	C.player_getGameText(p.handle, C.int(style), &cstr, &delay, &remaining)

	return &PlayerGameText{
		Text:      C.GoStringN(cstr.buf, C.int(cstr.length)),
		Delay:     time.Duration(delay) * time.Millisecond,
		Remaining: time.Duration(remaining) * time.Millisecond,
	}
}

func (p *Player) SetWeather(weather int) {
	C.player_setWeather(p.handle, C.int(weather))
}

func (p *Player) Weather() int {
	return int(C.player_getWeather(p.handle))
}

func (p *Player) SetWorldBounds(bounds Vector4) {
	C.player_setWorldBounds(p.handle, C.float(bounds.X), C.float(bounds.Y), C.float(bounds.Z), C.float(bounds.W))
}

func (p *Player) UnsetWorldBounds() {
	C.player_setWorldBounds(p.handle, C.float(20000.0), C.float(-20000.0), C.float(20000.0), C.float(-20000.0))
}

func (p *Player) WorldBounds() Vector4 {
	bounds := C.player_getWorldBounds(p.handle)

	return Vector4{
		X: float32(bounds.x),
		Y: float32(bounds.y),
		Z: float32(bounds.z),
		W: float32(bounds.w),
	}
}

func (p *Player) SetFightingStyle(style PlayerFightingStyle) {
	C.player_setFightingStyle(p.handle, C.int(style))
}

func (p *Player) FightingStyle() PlayerFightingStyle {
	return PlayerFightingStyle(C.player_getFightingStyle(p.handle))
}

func (p *Player) SetSkillLevel(skill PlayerWeaponSkill, level int) {
	C.player_setSkillLevel(p.handle, C.int(skill), C.int(level))
}

func (p *Player) SetAction(action PlayerSpecialAction) {
	C.player_setAction(p.handle, C.int(action))
}

func (p *Player) Action() PlayerSpecialAction {
	return PlayerSpecialAction(C.player_getAction(p.handle))
}

func (p *Player) SetVelocity(velocity Vector3) {
	C.player_setVelocity(p.handle, C.float(velocity.X), C.float(velocity.Y), C.float(velocity.Z))
}

func (p *Player) Velocity() Vector3 {
	vel := C.player_getVelocity(p.handle)

	return Vector3{
		X: float32(vel.x),
		Y: float32(vel.y),
		Z: float32(vel.z),
	}
}

func (p *Player) SetInterior(interior int) {
	C.player_setInterior(p.handle, C.uint(interior))
}

func (p *Player) Interior() int {
	return int(C.player_getInterior(p.handle))
}

func (p *Player) KeyData() PlayerKeyData {
	data := C.player_getKeyData(p.handle)

	return PlayerKeyData{
		Keys:      int(data.keys),
		UpDown:    int(data.upDown),
		LeftRight: int(data.leftRight),
	}
}

func (p *Player) WeaponState() PlayerWeaponState {
	aimData := C.player_getAimData(p.handle)

	return PlayerWeaponState(aimData.weaponState)
}

func (p *Player) CameraAspectRatio() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.aspectRatio)
}

func (p *Player) CameraFrontVector() Vector3 {
	aimData := C.player_getAimData(p.handle)

	return Vector3{
		X: float32(aimData.camFrontVector.x),
		Y: float32(aimData.camFrontVector.y),
		Z: float32(aimData.camFrontVector.z),
	}
}

// TODO constants
func (p *Player) CameraMode() int {
	aimData := C.player_getAimData(p.handle)

	return int(aimData.camMode)
}

func (p *Player) CameraZoom() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.camZoom)
}

func (p *Player) AimZ() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.aimZ)
}

// TODO getPlayerBulletData

func (p *Player) EnableCameraTargetting() {
	C.player_useCameraTargetting(p.handle, 1)
}

func (p *Player) DisableCameraTargetting() {
	C.player_useCameraTargetting(p.handle, 0)
}

func (p *Player) IsCameraTargettingEnabled() bool {
	return C.player_hasCameraTargetting(p.handle) != 0
}

func (p *Player) RemoveFromVehicle(force bool) {
	var cforce C.int
	if force {
		cforce = 1
	}

	C.player_removeFromVehicle(p.handle, cforce)
}

func (p *Player) CameraTargetPlayer() *Player {
	player := C.player_getCameraTargetPlayer(p.handle)

	return &Player{player}
}

func (p *Player) CameraTargetVehicle() *Vehicle {
	vehicle := C.player_getCameraTargetVehicle(p.handle)

	return &Vehicle{vehicle}
}

func (p *Player) CameraTargetObject() *Object {
	object := C.player_getCameraTargetObject(p.handle)

	return &Object{object}
}

func (p *Player) CameraTargetActor() *Actor {
	actor := C.player_getCameraTargetActor(p.handle)

	return &Actor{actor}
}

func (p *Player) TargetPlayer() *Player {
	player := C.player_getTargetPlayer(p.handle)

	return &Player{player}
}

func (p *Player) TargetActor() *Actor {
	actor := C.player_getTargetActor(p.handle)

	return &Actor{actor}
}

func (p *Player) EnableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.handle, 1)
}

func (p *Player) DisableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.handle, 0)
}

func (p *Player) SpectatePlayer(player *Player, mode PlayerSpectateMode) {
	C.player_spectatePlayer(p.handle, player.handle, C.int(mode))
}

func (p *Player) SpectateVehicle(vehicle *Vehicle, mode PlayerSpectateMode) {
	C.player_spectateVehicle(p.handle, vehicle.handle, C.int(mode))
}

// TODO
// func (p *Player) SpectatingPlayer() *Player {
// 	specData := C.player_getSpectateData(p.handle)
// }

// func (p *Player) SpectatingVehicle() *Vehicle {
// 	specData := C.player_getSpectateData(p.handle)
// }

// TODO callback
func (p *Player) SendClientCheck(actionType, address, offset, count int) {
	C.player_sendClientCheck(p.handle, C.int(actionType), C.int(address), C.int(offset), C.int(count))
}

func (p *Player) EnableGhostMode() {
	C.player_toggleGhostMode(p.handle, 1)
}

func (p *Player) DisableGhostMode() {
	C.player_toggleGhostMode(p.handle, 0)
}

func (p *Player) IsGhostModeEnabled() bool {
	return C.player_isGhostModeEnabled(p.handle) != 0
}

func (p *Player) RemovedBuildingCount() int {
	return int(C.player_getDefaultObjectsRemoved(p.handle))
}

func (p *Player) AllowWeapons() {
	C.player_allowWeapons(p.handle, 1)
}

func (p *Player) DisallowWeapons() {
	C.player_allowWeapons(p.handle, 0)
}

func (p *Player) AreWeaponsAllowed() bool {
	return C.player_areWeaponsAllowed(p.handle) != 0
}

func (p *Player) AllowTeleport() {
	C.player_allowTeleport(p.handle, 1)
}

func (p *Player) DisallowTeleport() {
	C.player_allowTeleport(p.handle, 0)
}

func (p *Player) IsTeleportAllowed() bool {
	return C.player_isTeleportAllowed(p.handle) != 0
}

func (p *Player) IsUsingOfficialClient() bool {
	return C.player_isUsingOfficialClient(p.handle) != 0
}

// entity

func (p *Player) SetPosition(pos Vector3) {
	C.player_setPosition(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (p *Player) Position() Vector3 {
	pos := C.player_getPosition(p.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

func (p *Player) Rotation() Vector4 {
	rquat := C.player_getRotation(p.handle)

	return Vector4{
		X: float32(rquat.x),
		Y: float32(rquat.y),
		Z: float32(rquat.z),
		W: float32(rquat.w),
	}
}

func (p *Player) SetVirtualWorld(vw int) {
	C.player_setVirtualWorld(p.handle, C.int(vw))
}

func (p *Player) VirtualWorld() int {
	return int(C.player_getVirtualWorld(p.handle))
}

// checkpoint data

func (p *Player) NewDefaultCheckpoint(radius float32, pos Vector3) *DefaultCheckpoint {
	cp := C.player_setCheckpoint(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius))

	return &DefaultCheckpoint{cp}
}

// console data

func (p *Player) MakeAdmin() {
	C.player_setConsoleAccessibility(p.handle, 1)
}

func (p *Player) UnmakeAdmin() {
	C.player_setConsoleAccessibility(p.handle, 0)
}

func (p *Player) IsAdmin() bool {
	return C.player_hasConsoleAccess(p.handle) != 0
}

// custom models data

func (p *Player) CustomSkin() int {
	return int(C.player_getCustomSkin(p.handle))
}

// network data

func (p *Player) IP() string {
	IP := C.player_getIp(p.handle)

	return C.GoStringN(IP.buf, C.int(IP.length))
}

func (p *Player) RawIP() int {
	return int(C.player_getRawIp(p.handle))
}

// vehicle data

func (p *Player) Vehicle() (*Vehicle, error) {
	vehicle := C.player_getVehicle(p.handle)

	if vehicle == nil {
		return nil, errors.New("player is not in a vehicle")
	}

	return &Vehicle{vehicle}, nil
}

func (p *Player) VehicleSeat() int {
	return int(C.player_getSeat(p.handle))
}

// misc

func (p *Player) DistanceFrom(point Vector3) float32 {
	return float32(C.player_getDistanceFromPoint(p.handle, C.float(point.X), C.float(point.Y), C.float(point.Z)))
}

func (p *Player) IsInRangeOf(point Vector3, _range float32) bool {
	return C.player_isInRangeOfPoint(p.handle, C.float(_range), C.float(point.X), C.float(point.Y), C.float(point.Z)) != 0
}

func (p *Player) SetFacingAngle(angle float32) {
	C.player_setFacingAngle(p.handle, C.float(angle))
}

func (p *Player) FacingAngle() float32 {
	return float32(C.player_getFacingAngle(p.handle))
}

func (p *Player) CheckPoint() {
	panic("not implemented")
}

func (p *Player) DialogID() {
	panic("not implemented")
}

func (p *Player) DialogData() {
	panic("not implemented")
}

func (p *Player) Menu(target *Player) {
	panic("not implemented")
}

func (p *Player) NetworkStats() {
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
