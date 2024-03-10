package event

type Type string

const (
	TypeGameModeInit Type = "gameModeInit"

	// Player connect events
	TypeIncomingConnection    Type = "incomingConnection"
	TypePlayerConnect         Type = "playerConnect"
	TypePlayerDisconnect      Type = "playerDisconnect"
	TypePlayerClientInit      Type = "playerClientInit"
	TypePlayerStreamIn        Type = "playerStreamIn"
	TypePlayerStreamOut       Type = "playerStreamOut"
	TypePlayerText            Type = "playerText"
	TypePlayerCommandText     Type = "playerCommandText"
	TypePlayerScoreChange     Type = "playerScoreChange"
	TypePlayerNameChange      Type = "playerNameChange"
	TypePlayerInteriorChange  Type = "playerInteriorChange"
	TypePlayerStateChange     Type = "playerStateChange"
	TypePlayerKeyStateChange  Type = "playerKeyStateChange"
	TypePlayerDeath           Type = "playerDeath"
	TypePlayerTakeDamage      Type = "playerTakeDamage"
	TypePlayerGiveDamage      Type = "playerGiveDamage"
	TypePlayerClickMap        Type = "playerClickMap"
	TypePlayerClickPlayer     Type = "playerClickPlayer"
	TypeClientCheckResponse   Type = "clientCheckResponse"
	TypePlayerUpdate          Type = "playerUpdate"
	TypeDialogResponse        Type = "dialogResponse"
	TypePlayerGiveDamageActor Type = "playerGiveDamageActor"
	TypeActorStreamIn         Type = "actorStreamIn"
	TypeActorStreamOut        Type = "actorStreamOut"
)
