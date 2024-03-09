package main

type DisconnectReason int

const (
	DisconnectReasonTimeout DisconnectReason = iota
	DisconnectReasonQuit
	DisconnectReasonKicked
	DisconnectReasonCustom
	DisconnectReasonModeEnd
)

type gameModeInitEvent struct {
	GameMode *GameMode
}

// Player connect events

type incomingConnectionEvent struct {
	Player    *Player
	IPAddress string
	Port      int
}

type playerConnectEvent struct {
	Player *Player
}

type playerDisconnectEvent struct {
	Player *Player
	Reason DisconnectReason
}

type playerClientInitEvent struct {
	Player *Player
}
