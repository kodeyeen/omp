package main

type DisconnectReason int
type DialogResponse int
type PlayerState int
type BodyPart int
type PlayerClickSource int
type PlayerBulletHitType int

const (
	DisconnectReasonTimeout DisconnectReason = iota
	DisconnectReasonQuit
	DisconnectReasonKicked
	DisconnectReasonCustom
	DisconnectReasonModeEnd
)

const (
	DialogResponseRight DialogResponse = iota
	DialogResponseLeft
)

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

const (
	BodyPartTorso BodyPart = iota + 3
	BodyPartGroin
	BodyPartLeftArm
	BodyPartRightArm
	BodyPartLeftLeg
	BodyPartRightLeg
	BodyPartHead
)

const (
	PlayerClickSourceScoreboard = iota
)

const (
	PlayerBulletHitTypeNone PlayerBulletHitType = iota
	PlayerBulletHitTypePlayer
	PlayerBulletHitTypeVehicle
	PlayerBulletHitTypeObject
	PlayerBulletHitTypePlayerObject
)

type GameModeInitEvent struct {
	GameMode *GameMode
}

type PlayerRequestSpawnEvent struct {
	Player *Player
}

type PlayerSpawnEvent struct {
	Player *Player
}

// Player connect events

type IncomingConnectionEvent struct {
	Player    *Player
	IPAddress string
	Port      int
}

type PlayerConnectEvent struct {
	Player *Player
}

type PlayerDisconnectEvent struct {
	Player *Player
	Reason DisconnectReason
}

type PlayerClientInitEvent struct {
	Player *Player
}

type PlayerStreamInEvent struct {
	Player    *Player
	ForPlayer *Player
}

type PlayerStreamOutEvent struct {
	Player    *Player
	ForPlayer *Player
}

type PlayerTextEvent struct {
	Player  *Player
	Message string
}

type PlayerCommandTextEvent struct {
	Player  *Player
	Message string
}

type PlayerScoreChangeEvent struct {
	Player *Player
	Score  int
}

type PlayerNameChangeEvent struct {
	Player  *Player
	OldName string
}

type PlayerInteriorChangeEvent struct {
	Player      *Player
	NewInterior uint
	OldInterior uint
}

type PlayerStateChangeEvent struct {
	Player   *Player
	NewState PlayerState
	OldState PlayerState
}

type PlayerKeyStateChangeEvent struct {
	Player  *Player
	NewKeys uint
	OldKeys uint
}

type PlayerDeathEvent struct {
	Player *Player
	Killer *Player
	Reason int
}

type PlayerTakeDamageEvent struct {
	Player *Player
	From   *Player
	Amount float32
	Weapon uint
	Part   BodyPart
}

type PlayerGiveDamageEvent struct {
	Player *Player
	To     *Player
	Amount float32
	Weapon uint
	Part   BodyPart
}

type PlayerClickMapEvent struct {
	Player   *Player
	Position *Position
}

type PlayerClickPlayerEvent struct {
	Player  *Player
	Clicked *Player
	Source  PlayerClickSource
}

type ClientCheckResponseEvent struct {
	Player     *Player
	ActionType int
	Address    int
	Results    int
}

type PlayerUpdateEvent struct {
	Player *Player
	Now    int
}

type DialogResponseEvent struct {
	Player    *Player
	DialogID  int
	Response  DialogResponse
	ListItem  int
	InputText string
}

type PlayerGiveDamageActorEvent struct {
	Player *Player
	Actor  *Player
	Amount float32
	Weapon uint
	Part   BodyPart
}

type ActorStreamInEvent struct {
	Actor     *Player
	ForPlayer *Player
}

type ActorStreamOutEvent struct {
	Actor     *Player
	ForPlayer *Player
}
