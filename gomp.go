package main

import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/kodeyeen/gomp/event"
)

var eventDispatcher = event.NewDispatcher()

func init() {
	eventDispatcher.On(event.TypeGameModeInit, func(evt *gameModeInitEvent) {
		log.Println("GAME MODE INITIALIZED")
	})

	eventDispatcher.On(event.TypePlayerConnect, func(evt *playerConnectEvent) {
		plr := evt.Player

		plr.SendMessage(
			0x00FF0000,
			fmt.Sprintf("Hello %s", plr.Name()),
		)
	})
}

func main() {}

//export onGameModeInit
func onGameModeInit() {
	comp = getComponent()
	comp.init("./components/Gomp.dll")

	gm := &GameMode{}

	event.Dispatch(eventDispatcher, event.TypeGameModeInit, &gameModeInitEvent{
		GameMode: gm,
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(plrHandle unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(eventDispatcher, event.TypeIncomingConnection, &incomingConnectionEvent{
		Player:    newPlayer(plrHandle, getComponent()),
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(plrHandle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerConnect, &playerConnectEvent{
		Player: newPlayer(plrHandle, getComponent()),
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(plrHandle unsafe.Pointer, reason int) {
	event.Dispatch(eventDispatcher, event.TypePlayerDisconnect, &playerDisconnectEvent{
		Player: newPlayer(plrHandle, getComponent()),
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(plrHandle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerClientInit, &playerClientInitEvent{
		Player: newPlayer(plrHandle, getComponent()),
	})
}

// go build -o test.dll -buildmode=c-shared
