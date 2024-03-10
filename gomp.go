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
	eventDispatcher.On(event.TypeGameModeInit, func(evt *GameModeInitEvent) {
		log.Println("GAME MODE INITIALIZED")
	})

	eventDispatcher.On(event.TypePlayerConnect, func(evt *PlayerConnectEvent) {
		plr := evt.Player

		plr.SendMessage(
			0x00FF0000,
			fmt.Sprintf("Hello %s", plr.Name()),
		)

		NewVehicle(602, Position{2161.8389, -1143.7473, 24.6501}, 266.9070)
	})
}

func On(evtType event.Type, callback any) {
	eventDispatcher.On(evtType, callback)
}

func main() {}

//export onGameModeInit
func onGameModeInit() {
	clibpath := C.CString("./components/Gomp.dll")
	defer C.free(unsafe.Pointer(clibpath))

	handle := C.loadLib(clibpath)

	C.initFuncs(handle)

	gm := &GameMode{}

	event.Dispatch(eventDispatcher, event.TypeGameModeInit, &GameModeInitEvent{
		GameMode: gm,
	})
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeIncomingConnection, &PlayerRequestSpawnEvent{
		Player: &Player{player},
	})
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeIncomingConnection, &PlayerSpawnEvent{
		Player: &Player{player},
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(eventDispatcher, event.TypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &Player{player},
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerConnect, &PlayerConnectEvent{
		Player: &Player{player},
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) {
	event.Dispatch(eventDispatcher, event.TypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: &Player{player},
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerClientInit, &PlayerClientInitEvent{
		Player: &Player{player},
	})
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player unsafe.Pointer, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &Player{player},
		ForPlayer: &Player{forPlayer},
	})
}

//export onPlayerStreamOut
func onPlayerStreamOut(player unsafe.Pointer, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &Player{player},
		ForPlayer: &Player{forPlayer},
	})
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, event.TypePlayerText, &PlayerTextEvent{
		Player:  &Player{player},
		Message: C.GoString(message),
	})
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, event.TypePlayerCommandText, &PlayerCommandTextEvent{
		Player:  &Player{player},
		Message: C.GoString(message),
	})
}

// Player shot events

// TODO:
//
////export onPlayerWeaponShot
// func onPlayerWeaponShot(player unsafe.Pointer, weapon C.uchar, hitType C.int, hitID C.ushort) {
// 	event.Dispatch(eventDispatcher, event.TypePlayerCommandText, &PlayerCommandTextEvent{
// 		Player:  &Player{player},
// 		Message: C.GoString(message),
// 	})
// }

// Player data change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerScoreChange, &PlayerScoreChangeEvent{
		Player: &Player{player},
		Score:  int(score),
	})
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName *C.char) {
	event.Dispatch(eventDispatcher, event.TypePlayerNameChange, &PlayerNameChangeEvent{
		Player:  &Player{player},
		OldName: C.GoString(oldName),
	})
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior C.uint, oldInterior C.uint) {
	event.Dispatch(eventDispatcher, event.TypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &Player{player},
		NewInterior: uint(newInterior),
		OldInterior: uint(oldInterior),
	})
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState C.int, oldState C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &Player{player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys C.uint, oldKeys C.uint) {
	event.Dispatch(eventDispatcher, event.TypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &Player{player},
		NewKeys: uint(newKeys),
		OldKeys: uint(oldKeys),
	})
}

// Player death and damage events

//export onPlayerDeath
func onPlayerDeath(player unsafe.Pointer, killer unsafe.Pointer, reason C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerDeath, &PlayerDeathEvent{
		Player: &Player{player},
		Killer: &Player{killer},
		Reason: int(reason),
	})
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player unsafe.Pointer, from unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &Player{player},
		From:   &Player{from},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player unsafe.Pointer, to unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerGiveDamage, &PlayerGiveDamageEvent{
		Player: &Player{player},
		To:     &Player{to},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(player unsafe.Pointer, x, y, z C.float) {
	event.Dispatch(eventDispatcher, event.TypePlayerClickMap, &PlayerClickMapEvent{
		Player:   &Player{player},
		Position: &Position{float32(x), float32(y), float32(z)},
	})
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player unsafe.Pointer, clicked unsafe.Pointer, source C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player:  &Player{player},
		Clicked: &Player{clicked},
		Source:  PlayerClickSource(source),
	})
}

// Player client check event

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results C.int) {
	event.Dispatch(eventDispatcher, event.TypeClientCheckResponse, &ClientCheckResponseEvent{
		Player:     &Player{player},
		ActionType: int(actionType),
		Address:    int(address),
		Results:    int(results),
	})
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerUpdate, &PlayerUpdateEvent{
		Player: &Player{player},
		Now:    int(now),
	})
}

// Player dialog events

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID C.int, response C.int, listItem C.int, inputText *C.char) {
	event.Dispatch(eventDispatcher, event.TypeDialogResponse, &DialogResponseEvent{
		Player:    &Player{player},
		DialogID:  int(dialogID),
		Response:  DialogResponse(response),
		ListItem:  int(listItem),
		InputText: C.GoString(inputText),
	})
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(player unsafe.Pointer, actor unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &Player{player},
		Actor:  &Player{actor},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onActorStreamIn
func onActorStreamIn(actor unsafe.Pointer, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

//export onActorStreamOut
func onActorStreamOut(actor unsafe.Pointer, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

// go build -o test.dll -buildmode=c-shared
