package main

// #include "component.h"
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

		NewVehicle(602, Position{2161.8389, -1143.7473, 24.6501}, 266.9070)
	})
}

func main() {}

//export onGameModeInit
func onGameModeInit() {
	clibpath := C.CString("./components/Gomp.dll")
	defer C.free(unsafe.Pointer(clibpath))

	handle := C.loadLib(clibpath)

	C.initFuncs(handle)

	gm := &GameMode{}

	event.Dispatch(eventDispatcher, event.TypeGameModeInit, &gameModeInitEvent{
		GameMode: gm,
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(plrHandle unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(eventDispatcher, event.TypeIncomingConnection, &incomingConnectionEvent{
		Player:    &Player{plrHandle},
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(plrHandle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerConnect, &playerConnectEvent{
		Player: &Player{plrHandle},
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(plrHandle unsafe.Pointer, reason int) {
	event.Dispatch(eventDispatcher, event.TypePlayerDisconnect, &playerDisconnectEvent{
		Player: &Player{plrHandle},
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(plrHandle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerClientInit, &playerClientInitEvent{
		Player: &Player{plrHandle},
	})
}

// go build -o test.dll -buildmode=c-shared
