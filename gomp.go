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

//export onPlayerConnect
func onPlayerConnect(plrHandle unsafe.Pointer) {
	plr := newPlayer(plrHandle, getComponent())

	event.Dispatch(eventDispatcher, event.TypePlayerConnect, &playerConnectEvent{
		Player: plr,
	})
}

// go build -o test.dll -buildmode=c-shared
