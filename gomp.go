package main

import "C"
import (
	"unsafe"
)

func main() {}

//export OnGameModeInit
func OnGameModeInit() {
	comp = getComponent()
	comp.init("./components/empty-template.dll")
}

//export OnPlayerConnect
func OnPlayerConnect(plrHandle unsafe.Pointer) {
	plr := newPlayer(plrHandle, getComponent())

	plr.SendMessage(0x00FF0000, "Hello!")
}

// go build -o test.dll -buildmode=c-shared
// set GOARCH=386
// set CGO_ENABLED=1
