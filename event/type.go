package event

type Type string

const (
	TypeGameModeInit Type = "gameModeInit"

	// Player connect events
	TypeIncomingConnection Type = "incomingConnection"
	TypePlayerConnect      Type = "playerConnect"
	TypePlayerDisconnect   Type = "playerDisconnect"
	TypePlayerClientInit   Type = "playerClientInit"

	// Player stream events
	TypePlayerStreamIn  Type = "playerStreamIn"
	TypePlayerStreamOut Type = "playerStreamOut"

	// Player text events
	TypePlayerText        Type = "playerText"
	TypePlayerCommandText Type = "playerCommandText"

	// Player change events
	TypePlayerScoreChange    Type = "playerScoreChange"
	TypePlayerNameChange     Type = "playerNameChange"
	TypePlayerInteriorChange Type = "playerInteriorChange"
	TypePlayerStateChange    Type = "playerStateChange"
	TypePlayerKeyStateChange Type = "playerKeyStateChange"

	// Player death and damage events
	TypePlayerDeath      Type = "playerDeath"
	TypePlayerTakeDamage Type = "playerTakeDamage"
	TypePlayerGiveDamage Type = "playerGiveDamage"

	// Player click events
	TypePlayerClickMap    Type = "playerClickMap"
	TypePlayerClickPlayer Type = "playerClickPlayer"

	// Client check event
	TypeClientCheckResponse Type = "clientCheckResponse"

	// Player updat event
	TypePlayerUpdate Type = "playerUpdate"

	// Player dialog event
	TypeDialogResponse Type = "dialogResponse"

	// Actor events
	TypePlayerGiveDamageActor Type = "playerGiveDamageActor"
	TypeActorStreamIn         Type = "actorStreamIn"
	TypeActorStreamOut        Type = "actorStreamOut"

	// Vehicle events
	TypeVehicleStreamIn           Type = "vehicleStreamIn"
	TypeVehicleStreamOut          Type = "vehicleStreamOut"
	TypeVehicleDeath              Type = "vehicleDeath"
	TypePlayerEnterVehicle        Type = "playerEnterVehicle"
	TypePlayerExitVehicle         Type = "playerExitVehicle"
	TypeVehicleDamageStatusUpdate Type = "vehicleDamageStatusUpdate"
	TypeVehiclePaintJob           Type = "vehiclePaintJob"
	TypeVehicleMod                Type = "vehicleMod"
	TypeVehicleRespray            Type = "vehicleRespray"
	TypeEnterExitModShop          Type = "enterExitModShop"
	TypeVehicleSpawn              Type = "vehicleSpawn"
	TypeUnoccupiedVehicleUpdate   Type = "unoccupiedVehicleUpdate"
	TypeTrailerUpdate             Type = "trailerUpdate"
	TypeVehicleSirenStateChange   Type = "vehicleSirenStateChange"

	// Object events
	TypeObjectMoved                Type = "objectMoved"
	TypePlayerObjectMoved          Type = "playerObjectMoved"
	TypeObjectSelected             Type = "objectSelected"
	TypePlayerObjectSelected       Type = "playerObjectSelected"
	TypeObjectEdited               Type = "objectEdited"
	TypePlayerObjectEdited         Type = "playerObjectEdited"
	TypePlayerAttachedObjectEdited Type = "playerAttachedObjectEdited"

	// Checkpoint events
	TypePlayerEnterCheckpoint     Type = "playerEnterCheckpoint"
	TypePlayerLeaveCheckpoint     Type = "playerLeaveCheckpoint"
	TypePlayerEnterRaceCheckpoint Type = "playerEnterRaceCheckpoint"
	TypePlayerLeaveRaceCheckpoint Type = "playerLeaveRaceCheckpoint"
)
