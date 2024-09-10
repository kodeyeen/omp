package omp

import (
	"time"

	"github.com/kodeyeen/event"
)

type DisconnectReason int

const (
	DisconnectReasonTimeout DisconnectReason = iota
	DisconnectReasonQuit
	DisconnectReasonKicked
	DisconnectReasonCustom
	DisconnectReasonModeEnd
)

type DialogResponse int

const (
	DialogResponseRight DialogResponse = iota
	DialogResponseLeft
)

type BodyPart int

const (
	BodyPartTorso BodyPart = iota + 3
	BodyPartGroin
	BodyPartLeftArm
	BodyPartRightArm
	BodyPartLeftLeg
	BodyPartRightLeg
	BodyPartHead
)

type PlayerClickSource int

const (
	PlayerClickSourceScoreboard = iota
)

type PlayerBulletHitType int

const (
	PlayerBulletHitTypeNone PlayerBulletHitType = iota
	PlayerBulletHitTypePlayer
	PlayerBulletHitTypeVehicle
	PlayerBulletHitTypeObject
	PlayerBulletHitTypePlayerObject
)

type ObjectEditResponse int

const (
	ObjectEditResponseCancel = iota
	ObjectEditResponseFinal
	ObjectEditResponseUpdate
)

type DownloadRequestType int

const (
	DownloadRequestTypeUnknown DownloadRequestType = iota - 1
	DownloadRequestTypeEmpty
	DownloadRequestTypeModelFile
	DownloadRequestTypeTextureFile
)

type PlayerBullet struct {
	Origin      Vector3
	HitPosition Vector3
	Offset      Vector3
	Weapon      Weapon
	HitType     int
	HitID       int
}

type UnoccupiedVehicleUpdate struct {
	Seat     int
	Position Vector3
	Velocity Vector3
}

const (
	EventTypeGameModeInit event.Type = "gameModeInit"
	EventTypeGameModeExit event.Type = "gameModeExit"

	// Actor events
	EventTypePlayerGiveDamageActor event.Type = "playerGiveDamageActor"
	EventTypeActorStreamOut        event.Type = "actorStreamOut"
	EventTypeActorStreamIn         event.Type = "actorStreamIn"

	// Checkpoint events
	EventTypePlayerEnterCheckpoint     event.Type = "playerEnterCheckpoint"
	EventTypePlayerLeaveCheckpoint     event.Type = "playerLeaveCheckpoint"
	EventTypePlayerEnterRaceCheckpoint event.Type = "playerEnterRaceCheckpoint"
	EventTypePlayerLeaveRaceCheckpoint event.Type = "playerLeaveRaceCheckpoint"

	// Class events
	EventTypePlayerRequestClass event.Type = "playerRequestClass"

	// Console events
	EventTypeConsoleText      event.Type = "consoleText"
	EventTypeRconLoginAttempt event.Type = "rconLoginAttempt"

	// Core events
	EventTypeTick event.Type = "tick"

	// Custom model events
	EventTypePlayerFinishedDownloading event.Type = "playerFinishedDownloading"
	EventTypePlayerRequestDownload     event.Type = "playerRequestDownload"

	// Player dialog event
	EventTypeDialogResponse event.Type = "dialogResponse"
	EventTypeDialogShow     event.Type = "dialogShow"
	EventTypeDialogHide     event.Type = "dialogHide"

	// Turf events
	EventTypePlayerEnterTurf       event.Type = "playerEnterTurf"
	EventTypePlayerEnterPlayerTurf event.Type = "playerEnterPlayerTurf"
	EventTypePlayerLeaveTurf       event.Type = "playerLeaveTurf"
	EventTypePlayerLeavePlayerTurf event.Type = "playerLeavePlayerTurf"
	EventTypePlayerClickTurf       event.Type = "playerClickTurf"
	EventTypePlayerClickPlayerTurf event.Type = "playerClickPlayerTurf"

	// Menu events
	EventTypePlayerSelectedMenuRow event.Type = "playerSelectedMenuRow"
	EventTypePlayerExitedMenu      event.Type = "playerExitedMenu"

	// Object events
	EventTypeObjectMoved            event.Type = "objectMoved"
	EventTypePlayerObjectMoved      event.Type = "playerObjectMoved"
	EventTypeObjectSelected         event.Type = "objectSelected"
	EventTypePlayerObjectSelected   event.Type = "playerObjectSelected"
	EventTypeObjectEdited           event.Type = "objectEdited"
	EventTypePlayerObjectEdited     event.Type = "playerObjectEdited"
	EventTypePlayerAttachmentEdited event.Type = "playerAttachmentEdited"

	// Pickup events
	EventTypePlayerPickUpPickup       event.Type = "playerPickUpPickup"
	EventTypePlayerPickUpPlayerPickup event.Type = "playerPickUpPlayerPickup"

	// Player spawn events
	EventTypePlayerRequestSpawn event.Type = "playerRequestSpawn"
	EventTypePlayerSpawn        event.Type = "playerSpawn"

	// Player connect events
	EventTypeIncomingConnection event.Type = "incomingConnection"
	EventTypePlayerConnect      event.Type = "playerConnect"
	EventTypePlayerDisconnect   event.Type = "playerDisconnect"
	EventTypePlayerClientInit   event.Type = "playerClientInit"

	// Player stream events
	EventTypePlayerStreamIn  event.Type = "playerStreamIn"
	EventTypePlayerStreamOut event.Type = "playerStreamOut"

	// Player text events
	EventTypePlayerText event.Type = "playerText"

	// Player shot events
	EventTypePlayerShotMissed       event.Type = "playerShotMissed"
	EventTypePlayerShotPlayer       event.Type = "playerShotPlayer"
	EventTypePlayerShotVehicle      event.Type = "playerShotVehicle"
	EventTypePlayerShotObject       event.Type = "playerShotObject"
	EventTypePlayerShotPlayerObject event.Type = "playerShotPlayerObject"

	// Player change events
	EventTypePlayerScoreChange    event.Type = "playerScoreChange"
	EventTypePlayerNameChange     event.Type = "playerNameChange"
	EventTypePlayerInteriorChange event.Type = "playerInteriorChange"
	EventTypePlayerStateChange    event.Type = "playerStateChange"
	EventTypePlayerKeyStateChange event.Type = "playerKeyStateChange"

	// Player damage events
	EventTypePlayerDeath      event.Type = "playerDeath"
	EventTypePlayerTakeDamage event.Type = "playerTakeDamage"
	EventTypePlayerGiveDamage event.Type = "playerGiveDamage"

	// Player click events
	EventTypePlayerClickMap    event.Type = "playerClickMap"
	EventTypePlayerClickPlayer event.Type = "playerClickPlayer"

	// Player check events
	EventTypeClientCheckResponse event.Type = "clientCheckResponse"

	// Player update event
	EventTypePlayerUpdate event.Type = "playerUpdate"

	// TextDraw events
	EventTypePlayerClickTextDraw                 event.Type = "playerClickTextDraw"
	EventTypePlayerClickPlayerTextDraw           event.Type = "playerClickPlayerTextDraw"
	EventTypePlayerCancelTextDrawSelection       event.Type = "playerCancelTextDrawSelection"
	EventTypePlayerCancelPlayerTextDrawSelection event.Type = "playerCancelPlayerTextDrawSelection"

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
)

type GameModeInitEvent struct {
}

type GameModeExitEvent struct {
}

// Actor events

type PlayerGiveDamageActorEvent struct {
	Player   *Player
	Actor    *Player
	Amount   float32
	Weapon   Weapon
	BodyPart BodyPart
}

type ActorStreamInEvent struct {
	Actor     *Player
	ForPlayer *Player
}

type ActorStreamOutEvent struct {
	Actor     *Player
	ForPlayer *Player
}

// Checkpoint events

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

// Class events

type PlayerRequestClassEvent struct {
	Player *Player
	Class  *Class
}

// Console events

type ConsoleTextEvent struct {
	Command    string
	Parameters string
}

type RconLoginAttemptEvent struct {
	Player   *Player
	Password string
	Success  bool
}

type TickEvent struct {
	Elapsed time.Duration
	Now     time.Time
}

// Custom model events

type PlayerFinishedDownloadingEvent struct {
	Player *Player
}

type PlayerRequestDownloadEvent struct {
	Player   *Player
	Type     DownloadRequestType
	Checksum int
}

// Dialog events

type MessageDialogResponseEvent struct {
	Player   *Player
	Response DialogResponse
}

type InputDialogResponseEvent struct {
	Player    *Player
	Response  DialogResponse
	InputText string
}

type ListDialogResponseEvent struct {
	Player     *Player
	Response   DialogResponse
	ItemNumber int
	Item       string
}

type TabListDialogResponseEvent struct {
	Player     *Player
	Response   DialogResponse
	ItemNumber int
	Item       TabListItem
}

type DialogShowEvent struct {
	Player *Player
}

type DialogHideEvent struct {
	Player *Player
}

// Turf events

type PlayerEnterTurfEvent struct {
	Player *Player
	Turf   *Turf
}

type PlayerEnterPlayerTurfEvent struct {
	Player *Player
	Turf   *PlayerTurf
}

type PlayerLeaveTurfEvent struct {
	Player *Player
	Turf   *Turf
}

type PlayerLeavePlayerTurfEvent struct {
	Player *Player
	Turf   *PlayerTurf
}

type PlayerClickTurfEvent struct {
	Player *Player
	Turf   *Turf
}

type PlayerClickPlayerTurfEvent struct {
	Player *Player
	Turf   *PlayerTurf
}

// Menu events

type PlayerSelectedMenuRowEvent struct {
	Player  *Player
	MenuRow int
}

type PlayerExitedMenuEvent struct {
	Player *Player
}

// Object events

type ObjectMovedEvent struct {
	Object *Object
}

type PlayerObjectMovedEvent struct {
	Player *Player
	Object *PlayerObject
}

type ObjectSelectedEvent struct {
	Player   *Player
	Object   *Object
	Model    int
	Position Vector3
}

type PlayerObjectSelectedEvent struct {
	Player   *Player
	Object   *PlayerObject
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
	Object   *PlayerObject
	Response ObjectEditResponse
	Offset   Vector3
	Rotation Vector3
}

type PlayerAttachmentEdited struct {
	Player     *Player
	Index      int
	Saved      bool
	Attachment PlayerAttachment
}

// Pickup events

type PlayerPickUpPickupEvent struct {
	Player *Player
	Pickup *Pickup
}

type PlayerPickUpPlayerPickupEvent struct {
	Player *Player
	Pickup *PlayerPickup
}

// Player spawn events

type PlayerRequestSpawnEvent struct {
	Player *Player
}

type PlayerSpawnEvent struct {
	Player *Player
}

// Player connect events

type IncomingConnectionEvent struct {
	Player    *Player
	ipAddress string
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

// Player stream events

type PlayerStreamInEvent struct {
	Player    *Player
	ForPlayer *Player
}

type PlayerStreamOutEvent struct {
	Player    *Player
	ForPlayer *Player
}

// Player text events

type PlayerTextEvent struct {
	Player  *Player
	Message string
}

// Player shot events

type PlayerShotMissedEvent struct {
	Player *Player
	Bullet PlayerBullet
}

type PlayerShotPlayerEvent struct {
	Player *Player
	Target *Player
	Bullet PlayerBullet
}

type PlayerShotVehicleEvent struct {
	Player *Player
	Target *Vehicle
	Bullet PlayerBullet
}

type PlayerShotObjectEvent struct {
	Player *Player
	Target *Object
	Bullet PlayerBullet
}

type PlayerShotPlayerObjectEvent struct {
	Player *Player
	Target *PlayerObject
	Bullet PlayerBullet
}

// Player change events

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
	NewInterior int
	OldInterior int
}

type PlayerStateChangeEvent struct {
	Player   *Player
	NewState PlayerState
	OldState PlayerState
}

type PlayerKeyStateChangeEvent struct {
	Player  *Player
	NewKeys int
	OldKeys int
}

// Player damage events

type PlayerDeathEvent struct {
	Player *Player
	Killer *Player
	Reason int
}

type PlayerTakeDamageEvent struct {
	Player *Player
	From   *Player
	Amount float32
	Weapon Weapon
	Part   BodyPart
}

type PlayerGiveDamageEvent struct {
	Player   *Player
	To       *Player
	Amount   float32
	Weapon   Weapon
	BodyPart BodyPart
}

// Player click events

type PlayerClickMapEvent struct {
	Player   *Player
	Position Vector3
}

type PlayerClickPlayerEvent struct {
	Player *Player
	Target *Player
	Source PlayerClickSource
}

// Player check events

type ClientCheckResponseEvent struct {
	Player     *Player
	ActionType int
	Address    int
	Results    int
}

// Player update events

type PlayerUpdateEvent struct {
	Player *Player
	Now    time.Time
}

// Textdraw events

type PlayerClickTextDrawEvent struct {
	Player   *Player
	Textdraw *Textdraw
}

type PlayerClickPlayerTextDrawEvent struct {
	Player   *Player
	Textdraw *PlayerTextdraw
}

type PlayerCancelTextDrawSelectionEvent struct {
	Player *Player
}

type PlayerCancelPlayerTextDrawSelectionEvent struct {
	Player *Player
}

// Vehicle events

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
	Vehicle *Vehicle
	Player  *Player
	Update  UnoccupiedVehicleUpdate
}

type TrailerUpdateEvent struct {
	Player  *Player
	Trailer *Vehicle
}

type VehicleSirenStateChangeEvent struct {
	Player     *Player
	Vehicle    *Vehicle
	SirenState int
}
