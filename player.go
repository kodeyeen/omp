package gomp

import (
	"time"
	"unsafe"
)

type Player interface {
	Handle() unsafe.Pointer
	SetName(name string) PlayerNameStatus
	Name() string
	GiveWeapon(weapon Weapon, ammo int)
	SendMessage(msg string, color int)
	Position() Vector3
	Vehicle() (*DefaultVehicle, error)
}

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
