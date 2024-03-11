package main

type DisconnectReason int
type DialogResponse int
type PlayerState int
type BodyPart int
type PlayerClickSource int
type PlayerBulletHitType int
type ObjectEditResponse int

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

const (
	ObjectEditResponseCancel = iota
	ObjectEditResponseFinal
	ObjectEditResponseUpdate
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

type VehicleStreamInEvent struct {
	Vehicle   *Vehicle
	ForPlayer *Player
}

type VehicleStreamOutEvent struct {
	Vehicle   *Vehicle
	ForPlayer *Player
}

type VehicleDeathEvent struct {
	Vehicle *Vehicle
	Killer  *Player
}

type PlayerEnterVehicleEvent struct {
	Player      *Player
	Vehicle     *Vehicle
	IsPassenger bool
}

type PlayerExitVehicleEvent struct {
	Player  *Player
	Vehicle *Vehicle
}

type VehicleDamageStatusUpdateEvent struct {
	Player  *Player
	Vehicle *Vehicle
}

type VehiclePaintJobEvent struct {
	Player   *Player
	Vehicle  *Vehicle
	PaintJob int
}

type VehicleModEvent struct {
	Player    *Player
	Vehicle   *Vehicle
	Component int
}

type VehicleResprayEvent struct {
	Player  *Player
	Vehicle *Vehicle
	Color   VehicleColor
}

type EnterExitModShopEvent struct {
	Player     *Player
	EnterExit  bool
	InteriorID int
}

type VehicleSpawnEvent struct {
	Vehicle *Vehicle
}

type UnoccupiedVehicleUpdateEvent struct {
	Vehicle  *Vehicle
	Player   *Player
	Seat     int
	Position Vector3
	Velocity Vector3
}

type TrailerUpdateEvent struct {
	Player  *Player
	Vehicle *Vehicle
}

type VehicleSirenStateChangeEvent struct {
	Player     *Player
	Vehicle    *Vehicle
	SirenState int
}

type ObjectMovedEvent struct {
	Object *Object
}

type PlayerObjectMovedEvent struct {
	Player *Player
	Object *Object
}

type ObjectSelectedEvent struct {
	Player   *Player
	Object   *Object
	Model    int
	Position Vector3
}

type PlayerObjectSelectedEvent struct {
	Player   *Player
	Object   *Object
	Model    int
	Position Vector3
}

type ObjectEditedEvent struct {
	Player   *Player
	Object   *Object
	Response ObjectEditResponse
	Offset   Vector3
	Rotation Vector3
}

type PlayerObjectEditedEvent struct {
	Player   *Player
	Object   *Object
	Response ObjectEditResponse
	Offset   Vector3
	Rotation Vector3
}

type PlayerAttachedObjectEditedEvent struct {
	Player   *Player
	Index    int
	Saved    int
	Model    int
	Bone     int
	Offset   Vector3
	Rotation Vector3
	Scale    Vector3
}

type PlayerEnterCheckpointEvent struct {
	Player *Player
}

type PlayerLeaveCheckpointEvent struct {
	Player *Player
}

type PlayerEnterRaceCheckpointEvent struct {
	Player *Player
}

type PlayerLeaveRaceCheckpointEvent struct {
	Player *Player
}
