package omp

// #include "include/player.h"
// #include "include/vehicle.h"
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

const (
	PlayerKeyAction          = 1
	PlayerKeyCrouch          = 2
	PlayerKeyFire            = 4
	PlayerKeySprint          = 8
	PlayerKeySecondaryAttack = 16
	PlayerKeyJump            = 32
	PlayerKeyLookRight       = 64
	PlayerKeyHandbrake       = 128
	PlayerKeyAim             = 128
	PlayerKeyLookLeft        = 256
	PlayerKeyLookBehind      = 512
	PlayerKeySubmission      = 512
	PlayerKeyWalk            = 1024
	PlayerKeyAnalogUp        = 2048
	PlayerKeyAnalogDown      = 4096
	PlayerKeyAnalogLeft      = 8192
	PlayerKeyAnalogRight     = 16384
	PlayerKeyYes             = 65536
	PlayerKeyNo              = 131072
	PlayerKeyCtrlBack        = 262144
	PlayerKeyUp              = -128
	PlayerKeyDown            = 128
	PlayerKeyLeft            = -128
	PlayerKeyRight           = 128
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

type PlayerAttachment struct {
	ModelID            int
	Bone               PlayerBone
	Offset, Rot, Scale Vector3
	Color1, Color2     uint
}

type PlayerBone int

const (
	PlayerBoneSpine PlayerBone = iota + 1
	PlayerBoneHead
	PlayerBoneLeftUpperArm
	PlayerBoneRightUpperArm
	PlayerBoneLeftHand
	PlayerBoneRightHand
	PlayerBoneLeftThigh
	PlayerBoneRightThigh
	PlayerBoneLeftFoot
	PlayerBoneRightFoot
	PlayerBoneRightCalf
	PlayerBoneLeftCalf
	PlayerBoneLeftForearm
	PlayerBoneRightForearm
	PlayerBoneLeftClavicle
	PlayerBoneRightClavicle
	PlayerBoneNeck
	PlayerBoneJaw
)

func Players() []*Player {
	playerArr := C.player_getAll()
	defer C.freeArray(playerArr)

	players := make([]*Player, 0, playerArr.length)
	handles := unsafe.Slice(playerArr.buf, int(playerArr.length))

	for _, handle := range handles {
		players = append(players, &Player{handle: handle})
	}

	return players
}

func SendDeathMessage(killer *Player, killee *Player, weapon int) {
	var cKiller unsafe.Pointer

	if killer != nil {
		cKiller = killer.handle
	}

	if killee == nil {
		C.player_sendEmptyDeathMessageToAll()
	} else {
		C.player_sendDeathMessageToAll(cKiller, killee.handle, C.int(weapon))
	}
}

func ShowGameTextForAll(msg string, delay time.Duration, style int) {
	cMsg := newCString(msg)
	defer freeCString(cMsg)

	C.player_sendGameTextToAll(cMsg, C.int(delay.Milliseconds()), C.int(style))
}

type Player struct {
	handle unsafe.Pointer
}

func (p *Player) ID() int {
	return int(C.player_getID(p.handle))
}

// Kick kicks the player from the server.
// They will have to quit the game and re-connect if they wish to continue playing.
func (p *Player) Kick() {
	C.player_kick(p.handle)
}

// Ban bans the player with a reason.
func (p *Player) Ban(reason string) {
	cReason := newCString(reason)
	defer freeCString(cReason)

	C.player_ban(p.handle, cReason)
}

// IsBot reports whether the player is an actual player or a bot (NPC).
func (p *Player) IsBot() bool {
	return C.player_isBot(p.handle) != 0
}

// Ping returns the ping of the player.
// The ping measures the amount of time it takes for the server to 'ping' the client
// and for the client to send the message back.
func (p *Player) Ping() int {
	return int(C.player_getPing(p.handle))
}

// Spawn (re)spawns the player.
func (p *Player) Spawn() {
	C.player_spawn(p.handle)
}

// IsSpawned reports whether the player is spawned.
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

// SetPositionFindZ sets the player's position then adjusts the player's z-coordinate
// to the nearest solid ground under the position.
func (p *Player) SetPositionFindZ(pos Vector3) {
	C.player_setPositionFindZ(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

// SetCameraPosition sets the camera to a specific position for the player.
func (p *Player) SetCameraPosition(pos Vector3) {
	C.player_setCameraPosition(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

// CameraPosition returns the position of the player's camera.
func (p *Player) CameraPosition() Vector3 {
	aimData := C.player_getAimData(p.handle)

	return Vector3{
		X: float32(aimData.camPos.x),
		Y: float32(aimData.camPos.y),
		Z: float32(aimData.camPos.z),
	}
}

// SetCameraLookAt sets the direction the player's camera looks at.
// Generally meant to be used in combination with SetCameraPosition.
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

// SetCameraBehind restores the camera to a place behind the player,
// after using a function like SetCameraPosition.
func (p *Player) SetCameraBehind() {
	C.player_setCameraBehind(p.handle)
}

// InterpolateCameraPosition moves the player's camera from one position to another, within the set time.
// Useful for scripted cut scenes.
func (p *Player) InterpolateCameraPosition(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraPosition(p.handle, C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

// InterpolateCameraLookAt interpolates the player's camera's 'look at' point between two coordinates with a set speed.
// Can be be used with InterpolateCameraPosition.
func (p *Player) InterpolateCameraLookAt(from Vector3, to Vector3, time int, cutType PlayerCameraCutType) {
	C.player_interpolateCameraLookAt(p.handle, C.float(from.X), C.float(from.Y), C.float(from.Z), C.float(to.X), C.float(to.Y), C.float(to.Z), C.int(time), C.int(cutType))
}

// AttachCameraToObject attaches the player camera to objects.
func (p *Player) AttachCameraToObject(obj *Object) {
	C.player_attachCameraToObject(p.handle, obj.handle)
}

// SetName sets the name of the player.
func (p *Player) SetName(name string) PlayerNameStatus {
	cName := newCString(name)
	defer freeCString(cName)

	return PlayerNameStatus(C.player_setName(p.handle, cName))
}

// Name returns the player's name.
func (p *Player) Name() string {
	name := C.player_getName(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

// Serial fetches the CI of the player, this is linked to their SAMP/GTA on their computer.
func (p *Player) Serial() string {
	name := C.player_getSerial(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

// GiveWeapon gives the player a weapon with a specified amount of ammo.
func (p *Player) GiveWeapon(weapon Weapon, ammo int) {
	C.player_giveWeapon(p.handle, C.uchar(weapon), C.uint(ammo))
}

// RemoveWeapon removes a specified weapon from the player.
func (p *Player) RemoveWeapon(weapon Weapon) {
	C.player_removeWeapon(p.handle, C.uchar(weapon))
}

// SetWeaponAmmo sets the ammo of the player's weapon.
func (p *Player) SetWeaponAmmo(weapon Weapon, ammo int) {
	C.player_setWeaponAmmo(p.handle, C.uchar(weapon), C.uint(ammo))
}

func (p *Player) WeaponSlots() []*WeaponSlot {
	panic("not implemented")
}

// WeaponSlot returns the weapon slot at a specified slot index
// An error is returned if invalid slot index is specified
func (p *Player) WeaponSlot(slotIdx WeaponSlotIndex) (WeaponSlot, error) {
	if slotIdx < WeaponSlotIndexUnknown || slotIdx > WeaponSlotIndexDetonator {
		return WeaponSlot{}, errors.New("invalid slot index")
	}

	slot := C.player_getWeaponSlot(p.handle, C.int(slotIdx))

	return WeaponSlot{
		Weapon: Weapon(slot.id),
		Ammo:   int(slot.ammo),
	}, nil
}

// ResetWeapons removes all weapons from a player.
func (p *Player) ResetWeapons() {
	C.player_resetWeapons(p.handle)
}

// SetArmedWeapon sets which weapon (that the player already has) the player is holding.
func (p *Player) SetArmedWeapon(weapon Weapon) {
	C.player_setArmedWeapon(p.handle, C.uint(weapon))
}

// ArmedWeapon returns the weapon the player is currently holding.
func (p *Player) ArmedWeapon() Weapon {
	return Weapon(C.player_getArmedWeapon(p.handle))
}

// ArmedWeaponAmmo returns the amount of ammo in the player's current weapon.
func (p *Player) ArmedWeaponAmmo() int {
	return int(C.player_getArmedWeaponAmmo(p.handle))
}

// SetShopName loads or unloads an interior script for the player (for example the ammunation menu).
func (p *Player) SetShopName(name string) {
	cName := newCString(name)
	defer freeCString(cName)

	C.player_setShopName(p.handle, cName)
}

// ShopName returns an interior script loaded or unloaded for the player
func (p *Player) ShopName() string {
	name := C.player_getShopName(p.handle)

	return C.GoStringN(name.buf, C.int(name.length))
}

// SetDrunkLevel sets the drunk level of the player which makes the player's camera sway and vehicles hard to control.
func (p *Player) SetDrunkLevel(level int) {
	C.player_setDrunkLevel(p.handle, C.int(level))
}

// DrunkLevel returns the player's level of drunkenness.
// If the level is less than 2000, the player is sober.
// The player's level of drunkness goes down slowly automatically (26 levels per second)
// but will always reach 2000 at the end.
// The higher drunkenness levels affect the player's camera, and the car driving handling.
// The level of drunkenness increases when the player drinks from a bottle
// (You can use SetSpecialAction to give them bottles).
func (p *Player) DrunkLevel() int {
	return int(C.player_getDrunkLevel(p.handle))
}

// SetColor sets the color of the player's nametag and marker (radar blip).
func (p *Player) SetColor(color uint) {
	C.player_setColour(p.handle, C.uint(color))
}

// Color returns the color of the player's name and radar marker. Only works after SetColor.
func (p *Player) Color() uint {
	return uint(C.player_getColour(p.handle))
}

// SetColorFor sets another player's color for this player
func (p *Player) SetColorFor(other *Player, color uint) {
	C.player_setOtherColour(p.handle, other.handle, C.uint(color))
}

// ColorFor returns another player's color for this player
func (p *Player) ColorFor(other *Player) (uint, error) {
	var cColor C.uint
	hasSpecificColor := C.player_getOtherColour(p.handle, other.handle, &cColor) != 0

	if !hasSpecificColor {
		return 0, errors.New("player has no specific color")
	}

	return uint(cColor), nil
}

// Freeze freezes the player so that it cannot control their character.
// The player will also be unable to move their camera.
func (p *Player) Freeze() {
	C.player_setControllable(p.handle, 0)
}

// Unfreeze unfreezes the player so that it can control their character.
func (p *Player) Unfreeze() {
	C.player_setControllable(p.handle, 1)
}

// IsFrozen reports whether the player can control their character.
func (p *Player) IsFrozen() bool {
	return C.player_getControllable(p.handle) != 0
}

// EnableSpectating toggle the player to be in spectator mode.
// While in spectator mode the player can spectate (watch) other players and vehicles.
// After using this function, either SpectatePlayer or SpectateVehicle needs to be used.
func (p *Player) EnableSpectating() {
	C.player_setSpectating(p.handle, 1)
}

// DisableSpectating disables the player's spectator mode.
func (p *Player) DisableSpectating() {
	C.player_setSpectating(p.handle, 0)
}

// SetWantedLevel sets the player's wanted level (6 brown stars under HUD).
func (p *Player) SetWantedLevel(level int) {
	C.player_setWantedLevel(p.handle, C.uint(level))
}

// WantedLevel returns the wanted level of the player.
func (p *Player) WantedLevel() int {
	return int(C.player_getWantedLevel(p.handle))
}

// PlaySound plays the specified sound for the player.
func (p *Player) PlaySound(sound int, pos Vector3) {
	C.player_playSound(p.handle, C.uint(sound), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

// LastPlayedSound returns the sound that was last played for the player.
func (p *Player) LastPlayedSound() int {
	return int(C.player_lastPlayedSound(p.handle))
}

// PlayAudio plays an 'audio stream' for the player. Normal audio files also work (e.g. MP3).
func (p *Player) PlayAudio(url string, usePos bool, pos Vector3, distance float32) {
	cUrl := newCString(url)
	defer freeCString(cUrl)

	C.player_playAudio(p.handle, cUrl, newCUchar(usePos), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(distance))
}

// PlayCrimeReport plays a crime report for the player - just like in single-player when CJ commits a crime.
func (p *Player) PlayCrimeReport(suspect *Player, crime int) {
	C.player_playerCrimeReport(p.handle, suspect.handle, C.int(crime))
}

// StopAudio stops the current audio stream for the player.
func (p *Player) StopAudio() {
	C.player_stopAudio(p.handle)
}

// LastPlayedAudio returns the player's last played audio URL.
func (p *Player) LastPlayedAudio() string {
	audio := C.player_lastPlayedAudio(p.handle)

	return C.GoStringN(audio.buf, C.int(audio.length))
}

// CreateExplosion creates an explosion that is only visible to the player.
// This can be used to isolate explosions from other players or to make them only appear in specific virtual worlds.
func (p *Player) CreateExplosion(_type int, radius float32, pos Vector3) {
	C.player_createExplosion(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.float(radius))
}

// SendDeathMessage adds a death to the 'killfeed' on the right-hand side of the screen for the player.
func (p *Player) SendDeathMessage(killer *Player, killee *Player, weapon int) {
	C.player_sendDeathMessage(p.handle, killer.handle, killee.handle, C.int(weapon))
}

func (p *Player) SendEmptyDeathMessage() {
	C.player_sendEmptyDeathMessage(p.handle)
}

// RemoveDefaultObjects removes a standard San Andreas model for the player within a specified range.
func (p *Player) RemoveDefaultObjects(model int, radius float32, pos Vector3) {
	C.player_removeDefaultObjects(p.handle, C.uint(model), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.float(radius))
}

// ForceClassSelection the player to go back to class selection.
func (p *Player) ForceClassSelection() {
	C.player_forceClassSelection(p.handle)
}

// SetMoney sets the player's money
func (p *Player) SetMoney(money int) {
	C.player_setMoney(p.handle, C.int(money))
}

// GiveMoney gives money to or takes money from the player.
func (p *Player) GiveMoney(money int) {
	C.player_giveMoney(p.handle, C.int(money))
}

// ResetMoney resets the player's money to $0.
func (p *Player) ResetMoney() {
	C.player_resetMoney(p.handle)
}

// Money checks how much money the player has.
func (p *Player) Money() int {
	return int(C.player_getMoney(p.handle))
}

// SetMapIcon places an icon/marker on the player's map.
// Can be used to mark locations such as banks and hospitals to players.
func (p *Player) SetMapIcon(ID int, _type int, color uint, style int, pos Vector3) {
	C.player_setMapIcon(p.handle, C.int(ID), C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.int(_type), C.uint(color), C.int(style))
}

// UnsetMapIcon removes a map icon for the player.
func (p *Player) UnsetMapIcon(ID int) {
	C.player_unsetMapIcon(p.handle, C.int(ID))
}

// EnableStuntBonuses enables stunt bonuses for the player.
func (p *Player) EnableStuntBonuses() {
	C.player_useStuntBonuses(p.handle, 1)
}

// DisableStuntBonuses disables stunt bonuses for the player.
func (p *Player) DisableStuntBonuses() {
	C.player_useStuntBonuses(p.handle, 0)
}

// ShowNameTagFor enables the drawing of player nametags, healthbars and armor bars which display above their head.
// For use of a similar function like this on a global level, ShowNameTags function.
func (p *Player) ShowNameTagFor(other *Player) {
	C.player_toggleOtherNameTag(other.handle, p.handle, 1)
}

// HideNameTagFor disables the drawing of player nametags, healthbars and armor bars which display above their head.
func (p *Player) HideNameTagFor(other *Player) {
	C.player_toggleOtherNameTag(other.handle, p.handle, 0)
}

// SetTime sets the game time for the player.
// If the player's clock is enabled (EnableClock) the time displayed by it will update automatically.
func (p *Player) SetTime(time PlayerTime) {
	C.player_setTime(p.handle, C.int(time.Hours), C.int(time.Minutes))
}

// Time returns the player's current game time.
// Set by SetWorldTime, or by the game automatically if EnableClock is used.
func (p *Player) Time() PlayerTime {
	ctime := C.player_getTime(p.handle)

	return PlayerTime{
		Hours:   int(ctime.hours),
		Minutes: int(ctime.minutes),
	}
}

// ShowClock shows the in-game clock (top-right corner) for the player.
// When this is enabled, time will progress at 1 minute per second.
// Weather will also interpolate (slowly change over time) when set using SetWeather/(*Player).SetWeather.
func (p *Player) ShowClock() {
	C.player_useClock(p.handle, 1)
}

// HideClock hides the in-game clock (top-right corner) for the player.
func (p *Player) HideClock() {
	C.player_useClock(p.handle, 0)
}

// IsClockShown reports whether the in-game clock is shown for the player.
func (p *Player) IsClockShown() bool {
	return C.player_hasClock(p.handle) != 0
}

// EnableWidescreen enables player's widescreen.
func (p *Player) EnableWidescreen() {
	C.player_useWidescreen(p.handle, 1)
}

// DisableWidescreen disables player's widescreen.
func (p *Player) DisableWidescreen() {
	C.player_useWidescreen(p.handle, 0)
}

// IsWidescreenEnabled reports whether player's widescreen is enabled.
func (p *Player) IsWidescreenEnabled() bool {
	return C.player_hasWidescreen(p.handle) != 0
}

// SetHealth sets the health of the player.
func (p *Player) SetHealth(health float32) {
	C.player_setHealth(p.handle, C.float(health))
}

// Health allows you to retrieve the health of the player.
// Useful for cheat detection, among other things.
func (p *Player) Health() float32 {
	return float32(C.player_getHealth(p.handle))
}

// SetScore sets the player's score.
// Players' scores are shown in the scoreboard (shown by holding the TAB key).
func (p *Player) SetScore(score int) {
	C.player_setScore(p.handle, C.int(score))
}

// Score returns the player's score as it was set using SetScore.
func (p *Player) Score() int {
	return int(C.player_getScore(p.handle))
}

// SetArmor sets the player's armor level.
func (p *Player) SetArmor(armor float32) {
	C.player_setArmour(p.handle, C.float(armor))
}

// Armor returns the armor level of the player.
func (p *Player) Armor() float32 {
	return float32(C.player_getArmour(p.handle))
}

// SetGravity sets the player's gravity.
func (p *Player) SetGravity(gravity float32) {
	C.player_setGravity(p.handle, C.float(gravity))
}

// Gravity returns the player's gravity.
func (p *Player) Gravity() float32 {
	return float32(C.player_getGravity(p.handle))
}

// SetWorldTime sets the game time for the player.
// If the player's clock is enabled (ShowClock) the time displayed by it will update automatically.
func (p *Player) SetWorldTime(time int) {
	C.player_setWorldTime(p.handle, C.int(time))
}

// ApplyAnimation applies an animation to the player.
func (p *Player) ApplyAnimation(anim Animation, syncType PlayerAnimationSyncType) {
	cLib := newCString(anim.Lib)
	defer freeCString(cLib)

	cName := newCString(anim.Name)
	defer freeCString(cName)

	C.player_applyAnimation(
		p.handle,
		C.float(anim.Delta),
		newCUchar(anim.Loop),
		newCUchar(anim.LockX),
		newCUchar(anim.LockY),
		newCUchar(anim.Freeze),
		C.uint(anim.Duration.Milliseconds()),
		cLib,
		cName,
		C.int(syncType),
	)
}

// ClearAnimations clears all animations for the player.
// It also cancels all current tasks such as jetpacking, parachuting,
// entering vehicles, driving (removes player out of vehicle), swimming, etc.
func (p *Player) ClearAnimations(syncType PlayerAnimationSyncType) {
	// TODO player_clearAnimations
	C.player_clearTasks(p.handle, C.PlayerAnimationSyncType(syncType))
}

// AnimationIndex returns the index of any running applied animations.
func (p *Player) AnimationIndex() int {
	animData := C.player_getAnimationData(p.handle)

	return int(animData.ID)
}

// AnimationFlags returns the player animation flags.
func (p *Player) AnimationFlags() int {
	animData := C.player_getAnimationData(p.handle)

	return int(animData.flags)
}

// IsStreamedInFor reports whether the player is streamed in another player's client.
func (p *Player) IsStreamedInFor(other *Player) bool {
	return C.player_isStreamedInForPlayer(p.handle, other.handle) != 0
}

// State returns the player's current state.
func (p *Player) State() PlayerState {
	return PlayerState(C.player_getState(p.handle))
}

// SetTeam sets the team of the player.
func (p *Player) SetTeam(team int) {
	C.player_setTeam(p.handle, C.int(team))
}

// Team returns the ID of the team the player is on.
func (p *Player) Team() int {
	return int(C.player_getTeam(p.handle))
}

// SetSkin sets the skin of the player. A player's skin is their character model.
func (p *Player) SetSkin(skin int) {
	C.player_setSkin(p.handle, C.int(skin), 1)
}

// Skin returns the class of the players skin.
func (p *Player) Skin() int {
	return int(C.player_getSkin(p.handle))
}

// SetChatBubble creates a chat bubble above the player's name tag.
func (p *Player) SetChatBubble(text string, color uint, drawDist float32, expire time.Duration) {
	cText := newCString(text)
	defer freeCString(cText)

	C.player_setChatBubble(p.handle, cText, C.uint(color), C.float(drawDist), C.int(expire.Milliseconds()))
}

// SendClientMessage sends a message to the player with a chosen color in the chat.
// The whole line in the chatbox will be in the set color unless color embedding is used.
func (p *Player) SendClientMessage(msg string, color uint) {
	cMsg := newCString(msg)
	defer freeCString(cMsg)

	C.player_sendClientMessage(p.handle, C.uint(color), cMsg)
}

// SendMessageFrom sends a message in the name of a player to this player on the server.
// The message will appear in the chat box, but can only be seen by the player.
// The line will start with the sender's name in their color, followed by the message in white.
func (p *Player) SendMessageFrom(sender *Player, msg string) {
	cMsg := newCString(msg)
	defer freeCString(cMsg)

	C.player_sendChatMessage(p.handle, sender.handle, cMsg)
}

// ShowGameText shows 'game text' (on-screen text) for a certain length of time for the player.
func (p *Player) ShowGameText(msg string, delay time.Duration, style int) {
	cMsg := newCString(msg)
	defer freeCString(cMsg)

	C.player_sendGameText(p.handle, cMsg, C.int(delay.Milliseconds()), C.int(style))
}

// HideGameText stops showing a gametext style to the player.
func (p *Player) HideGameText(style int) {
	C.player_hideGameText(p.handle, C.int(style))
}

// IsGameTextShown reports whether the player currently have text in the given gametext style displayed.
func (p *Player) IsGameTextShown(style int) {
	C.player_hasGameText(p.handle, C.int(style))
}

// GameText returns all the information on the given game text style.
func (p *Player) GameText(style int) *PlayerGameText {
	var cText C.String
	var delay C.int
	var remaining C.int

	C.player_getGameText(p.handle, C.int(style), &cText, &delay, &remaining)

	return &PlayerGameText{
		Text:      C.GoStringN(cText.buf, C.int(cText.length)),
		Delay:     time.Duration(delay) * time.Millisecond,
		Remaining: time.Duration(remaining) * time.Millisecond,
	}
}

// SetWeather sets the player's weather.
func (p *Player) SetWeather(weather int) {
	C.player_setWeather(p.handle, C.int(weather))
}

// Weather returns the player's weather.
func (p *Player) Weather() int {
	return int(C.player_getWeather(p.handle))
}

// SetWorldBounds sets the world boundaries for the player.
// Players can not go out of the boundaries (they will be pushed back in).
func (p *Player) SetWorldBounds(bounds Vector4) {
	C.player_setWorldBounds(p.handle, C.float(bounds.X), C.float(bounds.Y), C.float(bounds.Z), C.float(bounds.W))
}

// UnsetWorldBounds resets the player's world boundaries to default world boundaries.
func (p *Player) UnsetWorldBounds() {
	C.player_setWorldBounds(p.handle, C.float(20000.0), C.float(-20000.0), C.float(20000.0), C.float(-20000.0))
}

// WorldBounds returns the player's world boundaries.
func (p *Player) WorldBounds() Vector4 {
	bounds := C.player_getWorldBounds(p.handle)

	return Vector4{
		X: float32(bounds.x),
		Y: float32(bounds.y),
		Z: float32(bounds.z),
		W: float32(bounds.w),
	}
}

// SetFightingStyle sets the player's special fighting style.
// To use in-game, aim and press the 'secondary attack' key (ENTER by default).
func (p *Player) SetFightingStyle(style PlayerFightingStyle) {
	C.player_setFightingStyle(p.handle, C.int(style))
}

// FightingStyle returns the fighting style the player currently using.
func (p *Player) FightingStyle() PlayerFightingStyle {
	return PlayerFightingStyle(C.player_getFightingStyle(p.handle))
}

// SetSkillLevel sets the skill level of a certain weapon type for the player.
func (p *Player) SetSkillLevel(skill PlayerWeaponSkill, level int) {
	C.player_setSkillLevel(p.handle, C.int(skill), C.int(level))
}

// SetAction sets the players special action.
func (p *Player) SetAction(action PlayerSpecialAction) {
	C.player_setAction(p.handle, C.int(action))
}

// Action returns the player's current special action.
func (p *Player) Action() PlayerSpecialAction {
	return PlayerSpecialAction(C.player_getAction(p.handle))
}

// SetVelocity sets the player's velocity on the X, Y and Z axes.
func (p *Player) SetVelocity(velocity Vector3) {
	C.player_setVelocity(p.handle, C.float(velocity.X), C.float(velocity.Y), C.float(velocity.Z))
}

// Returns the velocity (speed) of the player on the X, Y and Z axes.
func (p *Player) Velocity() Vector3 {
	vel := C.player_getVelocity(p.handle)

	return Vector3{
		X: float32(vel.x),
		Y: float32(vel.y),
		Z: float32(vel.z),
	}
}

// SetInterior sets the player's interior.
// A list of currently known interiors and their positions can be found here:
// https://www.open.mp/docs/scripting/resources/interiorids
func (p *Player) SetInterior(interior int) {
	C.player_setInterior(p.handle, C.uint(interior))
}

// Interior returns the player's current interior.
// A list of currently known interiors with their positions can be found on this page:
// https://www.open.mp/docs/scripting/resources/interiorids
func (p *Player) Interior() int {
	return int(C.player_getInterior(p.handle))
}

// KeyData checks which keys the player is pressing.
func (p *Player) KeyData() PlayerKeyData {
	data := C.player_getKeyData(p.handle)

	return PlayerKeyData{
		Keys:      int(data.keys),
		UpDown:    int(data.upDown),
		LeftRight: int(data.leftRight),
	}
}

// WeaponState checks the state of the player's weapon.
func (p *Player) WeaponState() PlayerWeaponState {
	aimData := C.player_getAimData(p.handle)

	return PlayerWeaponState(aimData.weaponState)
}

// CameraAspectRatio returns the aspect ratio of the player's camera.
func (p *Player) CameraAspectRatio() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.aspectRatio)
}

// CameraFrontVector returns the current direction of player's aiming in 3-D space,
// the coords are relative to the camera position, see CameraPosition method.
func (p *Player) CameraFrontVector() Vector3 {
	aimData := C.player_getAimData(p.handle)

	return Vector3{
		X: float32(aimData.camFrontVector.x),
		Y: float32(aimData.camFrontVector.y),
		Z: float32(aimData.camFrontVector.z),
	}
}

// TODO constants

// CameraMode returns the current GTA camera mode for the requested player.
// The camera modes are useful in determining whether a player is aiming, doing a passenger driveby etc.
func (p *Player) CameraMode() int {
	aimData := C.player_getAimData(p.handle)

	return int(aimData.camMode)
}

// CameraZoom returns the game camera zoom level for the player.
func (p *Player) CameraZoom() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.camZoom)
}

// AimZ returns the player's Z Aim (related to the camera and aiming).
func (p *Player) AimZ() float32 {
	aimData := C.player_getAimData(p.handle)

	return float32(aimData.aimZ)
}

// TODO getPlayerBulletData

// EnableCameraTargeting enables camera targetting functions for the player.
// Disabled by default to save bandwidth.
func (p *Player) EnableCameraTargeting() {
	C.player_useCameraTargeting(p.handle, 1)
}

// DisableCameraTargeting disables camera targetting functions for the player.
// Disabled by default to save bandwidth.
func (p *Player) DisableCameraTargeting() {
	C.player_useCameraTargeting(p.handle, 0)
}

// IsCameraTargetingEnabled reports whether camera targetting functions are enabled for the player.
// Disabled by default to save bandwidth.
func (p *Player) IsCameraTargetingEnabled() bool {
	return C.player_hasCameraTargeting(p.handle) != 0
}

// RemoveFromVehicle removes/ejects the player from their vehicle.
func (p *Player) RemoveFromVehicle(force bool) {
	C.player_removeFromVehicle(p.handle, newCUchar(force))
}

// CameraTargetActor returns the player the player is looking at (if any).
func (p *Player) CameraTargetPlayer() *Player {
	player := C.player_getCameraTargetPlayer(p.handle)

	return &Player{handle: player}
}

// CameraTargetActor returns the vehicle the player is looking at (if any).
func (p *Player) CameraTargetVehicle() *Vehicle {
	vehicle := C.player_getCameraTargetVehicle(p.handle)

	return &Vehicle{handle: vehicle}
}

// CameraTargetActor returns the object the player is looking at (if any).
func (p *Player) CameraTargetObject() *Object {
	object := C.player_getCameraTargetObject(p.handle)

	return &Object{handle: object}
}

// CameraTargetActor returns the actor the player is looking at (if any).
func (p *Player) CameraTargetActor() *Actor {
	actor := C.player_getCameraTargetActor(p.handle)

	return &Actor{handle: actor}
}

// TargetPlayer checks who the player is aiming at.
func (p *Player) TargetPlayer() *Player {
	player := C.player_getTargetPlayer(p.handle)

	return &Player{handle: player}
}

// TargetActor returns the actor which is aimed by the player.
func (p *Player) TargetActor() *Actor {
	actor := C.player_getTargetActor(p.handle)

	return &Actor{handle: actor}
}

// EnableRemoteVehicleCollisions enables collisions between occupied vehicles for the player.
func (p *Player) EnableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.handle, 1)
}

// DisableRemoteVehicleCollisions disables collisions between occupied vehicles for the player.
func (p *Player) DisableRemoteVehicleCollisions() {
	C.player_setRemoteVehicleCollisions(p.handle, 0)
}

// SpectatePlayer makes the player spectate (watch) another player.
func (p *Player) SpectatePlayer(player *Player, mode PlayerSpectateMode) {
	C.player_spectatePlayer(p.handle, player.handle, C.int(mode))
}

// SpectateVehicle sets the player to spectate another vehicle.
// Their camera will be attached to the vehicle as if they are driving it.
func (p *Player) SpectateVehicle(vehicle *Vehicle, mode PlayerSpectateMode) {
	C.player_spectateVehicle(p.handle, vehicle.handle, C.int(mode))
}

// SpectatingPlayer returns the player the player is spectating (watching).
func (p *Player) SpectatingPlayer() (*Player, error) {
	specData := C.player_getSpectateData(p.handle)

	if specData._type != 2 {
		return nil, errors.New("player is not spectating a player")
	}

	cPlayer := C.player_getByID(specData.spectateID)

	return &Player{handle: cPlayer}, nil
}

// SpectatingVehicle returns the vehicle the player is spectating (watching).
func (p *Player) SpectatingVehicle() (*Vehicle, error) {
	specData := C.player_getSpectateData(p.handle)

	if specData._type != 1 {
		return nil, errors.New("player is not spectating a vehicle")
	}

	cVehicle := C.vehicle_getByID(specData.spectateID)

	return &Vehicle{handle: cVehicle}, nil
}

// TODO callback

// SendClientCheck performs a memory check on the client.
func (p *Player) SendClientCheck(actionType, address, offset, count int) {
	C.player_sendClientCheck(p.handle, C.int(actionType), C.int(address), C.int(offset), C.int(count))
}

// EnableGhostMode enables the player's ghost mode.
// Ghost mode disables the collision between player models.
func (p *Player) EnableGhostMode() {
	C.player_toggleGhostMode(p.handle, 1)
}

// DisableGhostMode disables the player's ghost mode.
// Ghost mode disables the collision between player models.
func (p *Player) DisableGhostMode() {
	C.player_toggleGhostMode(p.handle, 0)
}

// IsGhostModeEnabled reports whether ghost mode is enabled.
func (p *Player) IsGhostModeEnabled() bool {
	return C.player_isGhostModeEnabled(p.handle) != 0
}

// RemovedBuildingCount returns the number of removed buildings for the player.
func (p *Player) RemovedBuildingCount() int {
	return int(C.player_getDefaultObjectsRemoved(p.handle))
}

// AllowWeapons allows weapons for the player.
func (p *Player) AllowWeapons() {
	C.player_allowWeapons(p.handle, 1)
}

// DisallowWeapons disallows weapons for the player.
func (p *Player) DisallowWeapons() {
	C.player_allowWeapons(p.handle, 0)
}

// AreWeaponsAllowed reports whether weapons are allowed for this player.
func (p *Player) AreWeaponsAllowed() bool {
	return C.player_areWeaponsAllowed(p.handle) != 0
}

// AllowTeleport enables the teleporting ability for the player by right-clicking on the map.
func (p *Player) AllowTeleport() {
	C.player_allowTeleport(p.handle, 1)
}

// DisallowTeleport disables the teleporting ability for the player by right-clicking on the map.
func (p *Player) DisallowTeleport() {
	C.player_allowTeleport(p.handle, 0)
}

// Reports whether the teleporting ability for the player by right-clicking on the map is enabled.
func (p *Player) IsTeleportAllowed() bool {
	return C.player_isTeleportAllowed(p.handle) != 0
}

// IsUsingOfficialClient reports whether the player is using the official SA-MP client.
func (p *Player) IsUsingOfficialClient() bool {
	return C.player_isUsingOfficialClient(p.handle) != 0
}

// entity

// SetPosition sets the player's position.
func (p *Player) SetPosition(pos Vector3) {
	C.player_setPosition(p.handle, C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

// Position returns the player's position.
func (p *Player) Position() Vector3 {
	pos := C.player_getPosition(p.handle)

	return Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}

// Rotation returns the players rotation on all axes as a quaternion.
func (p *Player) Rotation() Vector4 {
	rquat := C.player_getRotation(p.handle)

	return Vector4{
		X: float32(rquat.x),
		Y: float32(rquat.y),
		Z: float32(rquat.z),
		W: float32(rquat.w),
	}
}

// SetVirtualWorld sets the virtual world of the player.
// They can only see other players or vehicles that are in that same world.
func (p *Player) SetVirtualWorld(vw int) {
	C.player_setVirtualWorld(p.handle, C.int(vw))
}

// VirtualWorld returns the virtual world of the player.
func (p *Player) VirtualWorld() int {
	return int(C.player_getVirtualWorld(p.handle))
}

// console data

// MakeAdmin makes the player as an RCON admin.
func (p *Player) MakeAdmin() {
	C.player_setConsoleAccessibility(p.handle, 1)
}

// MakeAdmin unmakes the player as an RCON admin.
func (p *Player) UnmakeAdmin() {
	C.player_setConsoleAccessibility(p.handle, 0)
}

// IsAdmin reports whether the player is an RCON admin.
func (p *Player) IsAdmin() bool {
	return C.player_hasConsoleAccess(p.handle) != 0
}

// checkpoint data

// DefaultCheckpoint returns the player's default checkpoint.
func (p *Player) DefaultCheckpoint() *DefaultCheckpoint {
	cp := C.player_getCheckpoint(p.handle)

	return &DefaultCheckpoint{handle: cp}
}

// RaceCheckpoint returns the player's race checkpoint.
func (p *Player) RaceCheckpoint() *RaceCheckpoint {
	cp := C.player_getRaceCheckpoint(p.handle)

	return &RaceCheckpoint{handle: cp}
}

// custom models data

// CustomSkin returns the class of the players custom skin downloaded from the server.
func (p *Player) CustomSkin() int {
	return int(C.player_getCustomSkin(p.handle))
}

func (p *Player) RedirectDownload(url string) error {
	cURL := newCString(url)
	defer freeCString(cURL)

	cOk := C.player_redirectDownload(p.handle, cURL)
	if cOk == 0 {
		return errors.New("failed to redirect download")
	}

	return nil
}

func (p *Player) showDialog(style dialogStyle, title, body, button1, button2 string) {
	cTitle := newCString(title)
	defer freeCString(cTitle)

	cBody := newCString(body)
	defer freeCString(cBody)

	cButton1 := newCString(button1)
	defer freeCString(cButton1)

	cButton2 := newCString(button2)
	defer freeCString(cButton2)

	C.player_showDialog(p.handle, C.int(999), C.int(style), cTitle, cBody, cButton1, cButton2)
}

func (p *Player) hideDialog() {
	C.player_hideDialog(p.handle)
}

// Dialog returns the dialog that is shown to the player.
func (p *Player) Dialog() (dialog, error) {
	dlg, ok := activeDialogs[p.ID()]
	if !ok {
		return nil, errors.New("player has no active dialog")
	}

	return dlg, nil
}

// network data

// IP returns the player's IP address.
func (p *Player) IP() string {
	IP := C.player_getIp(p.handle)

	return C.GoStringN(IP.buf, C.int(IP.length))
}

// RawIP returns the player's Raw IP address (v4).
func (p *Player) RawIP() int {
	return int(C.player_getRawIp(p.handle))
}

// vehicle data

// Vehicle returns the vehicle the player is currently in.
func (p *Player) Vehicle() (*Vehicle, error) {
	vehicle := C.player_getVehicle(p.handle)

	if vehicle == nil {
		return nil, errors.New("player is not in a vehicle")
	}

	return &Vehicle{handle: vehicle}, nil
}

// VehicleSeat returns the seat the player is in.
func (p *Player) VehicleSeat() int {
	return int(C.player_getSeat(p.handle))
}

// object data

// BeginObjectEditing allows the player to edit an object (position and rotation)
// using their mouse on a GUI (Graphical User Interface).
func (p *Player) BeginObjectEditing(obj *Object) {
	C.player_beginObjectEditing(p.handle, obj.handle)
}

// EndObjectEditing cancels object edition mode for the player.
func (p *Player) EndObjectEditing() {
	C.player_endObjectEditing(p.handle)
}

// IsEditingObject reports whether the player is in object edition mode.
func (p *Player) IsEditingObject() bool {
	return C.player_isEditingObject(p.handle) != 0
}

// BeginObjectSelecting displays the cursor and allows the player to select an object.
// EventTypeObjectSelected is called when the player selects an object.
func (p *Player) BeginObjectSelecting() {
	C.player_beginObjectSelecting(p.handle)
}

// IsSelectingObject reports whether the player is in object selection mode.
func (p *Player) IsSelectingObject() bool {
	return C.player_isSelectingObject(p.handle) != 0
}

// SetAttachment attaches an object to a specific bone on the player.
func (p *Player) SetAttachment(slotIdx int, attachment PlayerAttachment) {
	C.player_setAttachedObject(
		p.handle,
		C.int(slotIdx),
		C.int(attachment.ModelID),
		C.int(attachment.Bone),
		C.float(attachment.Offset.X),
		C.float(attachment.Offset.Y),
		C.float(attachment.Offset.Z),
		C.float(attachment.Rot.X),
		C.float(attachment.Rot.Y),
		C.float(attachment.Rot.Z),
		C.float(attachment.Scale.X),
		C.float(attachment.Scale.Y),
		C.float(attachment.Scale.Z),
		C.uint(attachment.Color1),
		C.uint(attachment.Color2),
	)
}

// Attachment returns the player attachment object data by slot index.
func (p *Player) Attachment(slotIdx int) PlayerAttachment {
	obj := C.player_getAttachedObject(p.handle, C.int(slotIdx))

	return PlayerAttachment{
		ModelID: int(obj.model),
		Bone:    PlayerBone(obj.bone),
		Offset: Vector3{
			X: float32(obj.offset.x),
			Y: float32(obj.offset.y),
			Z: float32(obj.offset.z),
		},
		Rot: Vector3{
			X: float32(obj.rotation.x),
			Y: float32(obj.rotation.y),
			Z: float32(obj.rotation.z),
		},
		Scale: Vector3{
			X: float32(obj.scale.x),
			Y: float32(obj.scale.y),
			Z: float32(obj.scale.z),
		},
		Color1: uint(obj.colour1),
		Color2: uint(obj.colour2),
	}
}

// RemoveAttachment removes an attached object from the player.
func (p *Player) RemoveAttachment(slotIdx int) {
	C.player_removeAttachedObject(p.handle, C.int(slotIdx))
}

// EditAttachment enters edition mode for an attached object.
func (p *Player) EditAttachment(slotIdx int) {
	C.player_editAttachedObject(p.handle, C.int(slotIdx))
}

// HasAttachment reports whether the player has an attachment in the specified slot index.
func (p *Player) HasAttachment(slotIdx int) bool {
	return C.player_hasAttachedObject(p.handle, C.int(slotIdx)) != 0
}

// misc

// DistanceFrom calculates the distance between the player and a map coordinate.
func (p *Player) DistanceFrom(point Vector3) float32 {
	return float32(C.player_getDistanceFromPoint(p.handle, C.float(point.X), C.float(point.Y), C.float(point.Z)))
}

// IsInRangeOf reports whether the player is in range of a point.
func (p *Player) IsInRangeOf(point Vector3, _range float32) bool {
	return C.player_isInRangeOfPoint(p.handle, C.float(_range), C.float(point.X), C.float(point.Y), C.float(point.Z)) != 0
}

// SetFacingAngle sets the player's facing angle (Z rotation).
func (p *Player) SetFacingAngle(angle float32) {
	C.player_setFacingAngle(p.handle, C.float(angle))
}

// FacingAngle returns the angle the player is facing.
func (p *Player) FacingAngle() float32 {
	return float32(C.player_getFacingAngle(p.handle))
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
