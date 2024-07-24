package omp

import (
	"time"
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

type PlayerBullet struct {
	Origin  Vector3
	HitPos  Vector3
	Offset  Vector3
	Weapon  Weapon
	HitType int
	HitID   int
}

type UnoccupiedVehicleUpdate struct {
	Seat     int
	Position Vector3
	Velocity Vector3
}

const (
	EventTypeGameModeInit EventType = "gameModeInit"
	EventTypeGameModeExit EventType = "gameModeExit"

	// Actor events
	EventTypePlayerGiveDamageActor EventType = "playerGiveDamageActor"
	EventTypeActorStreamOut        EventType = "actorStreamOut"
	EventTypeActorStreamIn         EventType = "actorStreamIn"

	// Checkpoint events
	EventTypePlayerEnterCheckpoint     EventType = "playerEnterCheckpoint"
	EventTypePlayerLeaveCheckpoint     EventType = "playerLeaveCheckpoint"
	EventTypePlayerEnterRaceCheckpoint EventType = "playerEnterRaceCheckpoint"
	EventTypePlayerLeaveRaceCheckpoint EventType = "playerLeaveRaceCheckpoint"

	// Class events
	EventTypePlayerRequestClass EventType = "playerRequestClass"

	// Console events
	EventTypeConsoleText      EventType = "consoleText"
	EventTypeRconLoginAttempt EventType = "rconLoginAttempt"

	// Custom model events
	EventTypePlayerFinishedDownloading EventType = "playerFinishedDownloading"
	EventTypePlayerRequestDownload     EventType = "playerRequestDownload"

	// Player dialog event
	EventTypeDialogResponse EventType = "dialogResponse"
	EventTypeDialogShow     EventType = "dialogShow"
	EventTypeDialogHide     EventType = "dialogHide"

	// Turf events
	EventTypePlayerEnterTurf       EventType = "playerEnterTurf"
	EventTypePlayerEnterPlayerTurf EventType = "playerEnterPlayerTurf"
	EventTypePlayerLeaveTurf       EventType = "playerLeaveTurf"
	EventTypePlayerLeavePlayerTurf EventType = "playerLeavePlayerTurf"
	EventTypePlayerClickTurf       EventType = "playerClickTurf"
	EventTypePlayerClickPlayerTurf EventType = "playerClickPlayerTurf"

	// Menu events
	EventTypePlayerSelectedMenuRow EventType = "playerSelectedMenuRow"
	EventTypePlayerExitedMenu      EventType = "playerExitedMenu"

	// Object events
	EventTypeObjectMoved            EventType = "objectMoved"
	EventTypePlayerObjectMoved      EventType = "playerObjectMoved"
	EventTypeObjectSelected         EventType = "objectSelected"
	EventTypePlayerObjectSelected   EventType = "playerObjectSelected"
	EventTypeObjectEdited           EventType = "objectEdited"
	EventTypePlayerObjectEdited     EventType = "playerObjectEdited"
	EventTypePlayerAttachmentEdited EventType = "playerAttachmentEdited"

	// Pickup events
	EventTypePlayerPickUpPickup       EventType = "playerPickUpPickup"
	EventTypePlayerPickUpPlayerPickup EventType = "playerPickUpPlayerPickup"

	// Player spawn events
	EventTypePlayerRequestSpawn EventType = "playerRequestSpawn"
	EventTypePlayerSpawn        EventType = "playerSpawn"

	// Player connect events
	EventTypeIncomingConnection EventType = "incomingConnection"
	EventTypePlayerConnect      EventType = "playerConnect"
	EventTypePlayerDisconnect   EventType = "playerDisconnect"
	EventTypePlayerClientInit   EventType = "playerClientInit"

	// Player stream events
	EventTypePlayerStreamIn  EventType = "playerStreamIn"
	EventTypePlayerStreamOut EventType = "playerStreamOut"

	// Player text events
	EventTypePlayerText EventType = "playerText"

	// Player shot events
	EventTypePlayerShotMissed       EventType = "playerShotMissed"
	EventTypePlayerShotPlayer       EventType = "playerShotPlayer"
	EventTypePlayerShotVehicle      EventType = "playerShotVehicle"
	EventTypePlayerShotObject       EventType = "playerShotObject"
	EventTypePlayerShotPlayerObject EventType = "playerShotPlayerObject"

	// Player change events
	EventTypePlayerScoreChange    EventType = "playerScoreChange"
	EventTypePlayerNameChange     EventType = "playerNameChange"
	EventTypePlayerInteriorChange EventType = "playerInteriorChange"
	EventTypePlayerStateChange    EventType = "playerStateChange"
	EventTypePlayerKeyStateChange EventType = "playerKeyStateChange"

	// Player damage events
	EventTypePlayerDeath      EventType = "playerDeath"
	EventTypePlayerTakeDamage EventType = "playerTakeDamage"
	EventTypePlayerGiveDamage EventType = "playerGiveDamage"

	// Player click events
	EventTypePlayerClickMap    EventType = "playerClickMap"
	EventTypePlayerClickPlayer EventType = "playerClickPlayer"

	// Player check events
	EventTypeClientCheckResponse EventType = "clientCheckResponse"

	// Player update event
	EventTypePlayerUpdate EventType = "playerUpdate"

	// TextDraw events
	EventTypePlayerClickTextDraw                 EventType = "playerClickTextDraw"
	EventTypePlayerClickPlayerTextDraw           EventType = "playerClickPlayerTextDraw"
	EventTypePlayerCancelTextDrawSelection       EventType = "playerCancelTextDrawSelection"
	EventTypePlayerCancelPlayerTextDrawSelection EventType = "playerCancelPlayerTextDrawSelection"

	// Vehicle events
	EventTypeVehicleStreamIn           EventType = "vehicleStreamIn"
	EventTypeVehicleStreamOut          EventType = "vehicleStreamOut"
	EventTypeVehicleDeath              EventType = "vehicleDeath"
	EventTypePlayerEnterVehicle        EventType = "playerEnterVehicle"
	EventTypePlayerExitVehicle         EventType = "playerExitVehicle"
	EventTypeVehicleDamageStatusUpdate EventType = "vehicleDamageStatusUpdate"
	EventTypeVehiclePaintJob           EventType = "vehiclePaintJob"
	EventTypeVehicleMod                EventType = "vehicleMod"
	EventTypeVehicleRespray            EventType = "vehicleRespray"
	EventTypeEnterExitModShop          EventType = "enterExitModShop"
	EventTypeVehicleSpawn              EventType = "vehicleSpawn"
	EventTypeUnoccupiedVehicleUpdate   EventType = "unoccupiedVehicleUpdate"
	EventTypeTrailerUpdate             EventType = "trailerUpdate"
	EventTypeVehicleSirenStateChange   EventType = "vehicleSirenStateChange"
)

type GameModeInitEvent struct {
}

type GameModeExitEvent struct {
}

// Actor events

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

// Custom model events

type PlayerFinishedDownloadingEvent struct {
	Player *Player
}

type PlayerRequestDownloadEvent struct {
	Player         *Player
	Type, Checksum int
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
	MenuRow uint8
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
	Player *Player
	To     *Player
	Amount float32
	Weapon Weapon
	Part   BodyPart
}

// Player click events

type PlayerClickMapEvent struct {
	Player   *Player
	Position Vector3
}

type PlayerClickPlayerEvent struct {
	Player  *Player
	Clicked *Player
	Source  PlayerClickSource
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
	Vehicle *Vehicle
}

type VehicleSirenStateChangeEvent struct {
	Player     *Player
	Vehicle    *Vehicle
	SirenState int
}
