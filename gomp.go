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

// Client check event

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

// Player dialog event

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
func onPlayerGiveDamageActor(player, actor unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &Player{player},
		Actor:  &Player{actor},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onActorStreamIn
func onActorStreamIn(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

//export onActorStreamOut
func onActorStreamOut(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &Vehicle{vehicle},
		ForPlayer: &Player{player},
	})
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{vehicle},
		ForPlayer: &Player{player},
	})
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &Vehicle{vehicle},
		Killer:  &Player{killer},
	})
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger C.int) {
	event.Dispatch(eventDispatcher, event.TypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &Player{player},
		Vehicle:     &Vehicle{vehicle},
		IsPassenger: int(isPassenger) != 0,
	})
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
	})
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{vehicle},
		Player:  &Player{player},
	})
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob C.int) {
	event.Dispatch(eventDispatcher, event.TypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &Player{player},
		Vehicle:  &Vehicle{vehicle},
		PaintJob: int(paintJob),
	})
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component C.int) {
	event.Dispatch(eventDispatcher, event.TypeVehicleMod, &VehicleModEvent{
		Player:    &Player{player},
		Vehicle:   &Vehicle{vehicle},
		Component: int(component),
	})
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 C.int) {
	event.Dispatch(eventDispatcher, event.TypeVehicleRespray, &VehicleResprayEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
		Color:   VehicleColor{int(color1), int(color2)},
	})
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit, interiorID C.int) {
	event.Dispatch(eventDispatcher, event.TypeEnterExitModShop, &EnterExitModShopEvent{
		Player:     &Player{player},
		EnterExit:  int(enterexit) != 0,
		InteriorID: int(interiorID),
	})
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &Vehicle{vehicle},
	})
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, seat C.uchar, posX, posY, posZ, velX, velY, velZ C.float) {
	event.Dispatch(eventDispatcher, event.TypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle:  &Vehicle{vehicle},
		Seat:     int(seat),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
		Velocity: Vector3{float32(velX), float32(velY), float32(velZ)},
	})
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
	})
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState C.uchar) {
	event.Dispatch(eventDispatcher, event.TypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &Player{player},
		Vehicle:    &Vehicle{vehicle},
		SirenState: int(sirenState),
	})
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypeObjectMoved, &ObjectMovedEvent{
		Object: &Object{object},
	})
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerObjectMoved, &PlayerObjectMovedEvent{
		Player: &Player{player},
		Object: &Object{object},
	})
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, event.TypeObjectSelected, &ObjectSelectedEvent{
		Player:   &Player{player},
		Object:   &Object{object},
		Model:    int(model),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
	})
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, event.TypePlayerObjectSelected, &PlayerObjectSelectedEvent{
		Player:   &Player{player},
		Object:   &Object{object},
		Model:    int(model),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
	})
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, event.TypeObjectEdited, &ObjectEditedEvent{
		Player:   &Player{player},
		Object:   &Object{object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{float32(offsetX), float32(offsetY), float32(offsetZ)},
		Rotation: Vector3{float32(rotX), float32(rotY), float32(rotZ)},
	})
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, event.TypePlayerObjectEdited, &PlayerObjectEditedEvent{
		Player:   &Player{player},
		Object:   &Object{object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{float32(offsetX), float32(offsetY), float32(offsetZ)},
		Rotation: Vector3{float32(rotX), float32(rotY), float32(rotZ)},
	})
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index, saved, model, bone C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, scaleX, scaleY, scaleZ C.float) {
	event.Dispatch(eventDispatcher, event.TypePlayerAttachedObjectEdited, &PlayerAttachedObjectEditedEvent{
		Player:   &Player{player},
		Index:    int(index),
		Saved:    int(saved),
		Model:    int(model),
		Bone:     int(bone),
		Offset:   Vector3{float32(offsetX), float32(offsetY), float32(offsetZ)},
		Rotation: Vector3{float32(rotX), float32(rotY), float32(rotZ)},
		Scale:    Vector3{float32(scaleX), float32(scaleY), float32(scaleZ)},
	})
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, event.TypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{player},
	})
}

// go build -o test.dll -buildmode=c-shared
