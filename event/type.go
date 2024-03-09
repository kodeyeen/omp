package event

type Type string

const (
	TypeGameModeInit Type = "gameModeInit"

	// Player connect events
	TypeIncomingConnection Type = "incomingConnection"
	TypePlayerConnect      Type = "playerConnect"
	TypePlayerDisconnect   Type = "playerDisconnect"
	TypePlayerClientInit   Type = "playerClientInit"
)
