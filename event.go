package gomp

import "github.com/kodeyeen/gomp/event"

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

const (
	EventTypeGameModeInit event.Type = "gameModeInit"

	// Player connect events
	EventTypeIncomingConnection event.Type = "incomingConnection"
	EventTypePlayerConnect      event.Type = "playerConnect"
	EventTypePlayerDisconnect   event.Type = "playerDisconnect"
	EventTypePlayerClientInit   event.Type = "playerClientInit"

	// Player stream events
	EventTypePlayerStreamIn  event.Type = "playerStreamIn"
	EventTypePlayerStreamOut event.Type = "playerStreamOut"

	// Player text events
	EventTypePlayerText        event.Type = "playerText"
	EventTypePlayerCommandText event.Type = "playerCommandText"

	// Player change events
	EventTypePlayerScoreChange    event.Type = "playerScoreChange"
	EventTypePlayerNameChange     event.Type = "playerNameChange"
	EventTypePlayerInteriorChange event.Type = "playerInteriorChange"
	EventTypePlayerStateChange    event.Type = "playerStateChange"
	EventTypePlayerKeyStateChange event.Type = "playerKeyStateChange"

	// Player death and damage events
	EventTypePlayerDeath      event.Type = "playerDeath"
	EventTypePlayerTakeDamage event.Type = "playerTakeDamage"
	EventTypePlayerGiveDamage event.Type = "playerGiveDamage"

	// Player click events
	EventTypePlayerClickMap    event.Type = "playerClickMap"
	EventTypePlayerClickPlayer event.Type = "playerClickPlayer"

	// Client check event
	EventTypeClientCheckResponse event.Type = "clientCheckResponse"

	// Player updat event
	EventTypePlayerUpdate event.Type = "playerUpdate"

	// Player dialog event
	EventTypeDialogResponse event.Type = "dialogResponse"

	// Actor events
	EventTypePlayerGiveDamageActor event.Type = "playerGiveDamageActor"
	EventTypeActorStreamIn         event.Type = "actorStreamIn"
	EventTypeActorStreamOut        event.Type = "actorStreamOut"

	// Vehicle events
	EventTypeVehicleStreamIn           event.Type = "vehicleStreamIn"
	EventTypeVehicleStreamOut          event.Type = "vehicleStreamOut"
	EventTypeVehicleDeath              event.Type = "vehicleDeath"
	EventTypePlayerEnterVehicle        event.Type = "playerEnterVehicle"
	EventTypePlayerExitVehicle         event.Type = "playerExitVehicle"
	EventTypeVehicleDamageStatusUpdate event.Type = "vehicleDamageStatusUpdate"
	EventTypeVehiclePaintJob           event.Type = "vehiclePaintJob"
	EventTypeVehicleMod                event.Type = "vehicleMod"
	EventTypeVehicleRespray            event.Type = "vehicleRespray"
	EventTypeEnterExitModShop          event.Type = "enterExitModShop"
	EventTypeVehicleSpawn              event.Type = "vehicleSpawn"
	EventTypeUnoccupiedVehicleUpdate   event.Type = "unoccupiedVehicleUpdate"
	EventTypeTrailerUpdate             event.Type = "trailerUpdate"
	EventTypeVehicleSirenStateChange   event.Type = "vehicleSirenStateChange"

	// Object events
	EventTypeObjectMoved                event.Type = "objectMoved"
	EventTypePlayerObjectMoved          event.Type = "playerObjectMoved"
	EventTypeObjectSelected             event.Type = "objectSelected"
	EventTypePlayerObjectSelected       event.Type = "playerObjectSelected"
	EventTypeObjectEdited               event.Type = "objectEdited"
	EventTypePlayerObjectEdited         event.Type = "playerObjectEdited"
	EventTypePlayerAttachedObjectEdited event.Type = "playerAttachedObjectEdited"

	// Checkpoint events
	EventTypePlayerEnterCheckpoint     event.Type = "playerEnterCheckpoint"
	EventTypePlayerLeaveCheckpoint     event.Type = "playerLeaveCheckpoint"
	EventTypePlayerEnterRaceCheckpoint event.Type = "playerEnterRaceCheckpoint"
	EventTypePlayerLeaveRaceCheckpoint event.Type = "playerLeaveRaceCheckpoint"

	// TextDraw events
	EventTypePlayerClickTextDraw                 event.Type = "playerClickTextDraw"
	EventTypePlayerClickPlayerTextDraw           event.Type = "playerClickPlayerTextDraw"
	EventTypePlayerCancelTextDrawSelection       event.Type = "playerCancelTextDrawSelection"
	EventTypePlayerCancelPlayerTextDrawSelection event.Type = "playerCancelPlayerTextDrawSelection"

	// Player model events
	EventTypePlayerFinishedDownloading event.Type = "playerFinishedDownloading"
	EventTypePlayerRequestDownload     event.Type = "playerRequestDownload"

	// Console events
	EventTypeConsoleText      event.Type = "consoleText"
	EventTypeRconLoginAttempt event.Type = "rconLoginAttempt"

	// Pickup events
	EventTypePlayerPickUpPickup event.Type = "playerPickUpPickup"

	// GangZone events
	EventTypePlayerEnterGangZone event.Type = "playerEnterGangZone"
	EventTypePlayerLeaveGangZone event.Type = "playerLeaveGangZone"
	EventTypePlayerClickGangZone event.Type = "playerClickGangZone"

	// Menu events
	EventTypePlayerSelectedMenuRow event.Type = "playerSelectedMenuRow"
	EventTypePlayerExitedMenu      event.Type = "playerExitedMenu"

	// Class events
	EventTypePlayerRequestClass event.Type = "playerRequestClass"
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

type PlayerClickTextDrawEvent struct {
	Player   *Player
	TextDraw *TextDraw
}

type PlayerClickPlayerTextDrawEvent struct {
	Player   *Player
	TextDraw *TextDraw
}

type PlayerCancelTextDrawSelectionEvent struct {
	Player *Player
}

type PlayerCancelPlayerTextDrawSelectionEvent struct {
	Player *Player
}

type PlayerFinishedDownloadingEvent struct {
	Player *Player
}

type PlayerRequestDownloadEvent struct {
	Player   *Player
	Type     int
	Checksum uint
}

type ConsoleTextEvent struct {
	Command    string
	Parameters string
}

type RconLoginAttemptEvent struct {
	Player   *Player
	Password string
	Success  bool
}

type PlayerPickUpPickupEvent struct {
	Player *Player
	Pickup *Pickup
}

type PlayerEnterGangZoneEvent struct {
	Player   *Player
	GangZone *GangZone
}

type PlayerLeaveGangZoneEvent struct {
	Player   *Player
	GangZone *GangZone
}

type PlayerClickGangZoneEvent struct {
	Player   *Player
	GangZone *GangZone
}

type PlayerSelectedMenuRowEvent struct {
	Player  *Player
	MenuRow uint8
}

type PlayerExitedMenuEvent struct {
	Player *Player
}

type PlayerRequestClassEvent struct {
	Player  *Player
	ClassID uint
}
