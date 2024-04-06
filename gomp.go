package gomp

// #include <stdlib.h>
// #include "include/gomp.h"
import "C"
import (
	"unsafe"

	"github.com/kodeyeen/gomp/event"
)

// go build -o test.dll -buildmode=c-shared

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Vector2 struct {
	X float32
	Y float32
}

var eventDispatcher = event.NewDispatcher()

func On(evtType event.Type, handler any) {
	eventDispatcher.On(evtType, handler)
}

func Once(evtType event.Type, handler any) {
	eventDispatcher.Once(evtType, handler)
}

func Off(evtType event.Type, handler any) {
	eventDispatcher.Off(evtType, handler)
}

//export onGameModeInit
func onGameModeInit() {
	clibpath := C.CString("./components/Gomp.dll")
	defer C.free(unsafe.Pointer(clibpath))

	handle := C.openLib(clibpath)

	C.initFuncs(handle)

	gm := &GameMode{}

	event.Dispatch(eventDispatcher, EventTypeGameModeInit, &GameModeInitEvent{
		GameMode: gm,
	})
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerRequesSpawn, &PlayerRequestSpawnEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeIncomingConnection, &PlayerSpawnEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(eventDispatcher, EventTypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &DefaultPlayer{handle: player},
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerConnect, &PlayerConnectEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) {
	event.Dispatch(eventDispatcher, EventTypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: &DefaultPlayer{handle: player},
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClientInit, &PlayerClientInitEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &DefaultPlayer{handle: player},
		ForPlayer: &DefaultPlayer{handle: forPlayer},
	})
}

//export onPlayerStreamOut
func onPlayerStreamOut(player, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &DefaultPlayer{handle: player},
		ForPlayer: &DefaultPlayer{handle: forPlayer},
	})
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerText, &PlayerTextEvent{
		Player:  &DefaultPlayer{handle: player},
		Message: C.GoString(message),
	})
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerCommandText, &PlayerCommandTextEvent{
		Player:  &DefaultPlayer{handle: player},
		Command: C.GoString(message),
	})
}

// Player shot events

// TODO:
//
////export onPlayerWeaponShot
// func onPlayerWeaponShot(player unsafe.Pointer, weapon C.uchar, hitType C.int, hitID C.ushort) {
// 	event.Dispatch(eventDispatcher, EventTypePlayerCommandText, &PlayerCommandTextEvent{
// 		Player:  &DefaultPlayer{handle: player},
// 		Message: C.GoString(message),
// 	})
// }

// Player data change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerScoreChange, &PlayerScoreChangeEvent{
		Player: &DefaultPlayer{handle: player},
		Score:  int(score),
	})
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerNameChange, &PlayerNameChangeEvent{
		Player:  &DefaultPlayer{handle: player},
		OldName: C.GoString(oldName),
	})
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior, oldInterior C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &DefaultPlayer{handle: player},
		NewInterior: uint(newInterior),
		OldInterior: uint(oldInterior),
	})
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState, oldState C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &DefaultPlayer{handle: player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys, oldKeys C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &DefaultPlayer{handle: player},
		NewKeys: uint(newKeys),
		OldKeys: uint(oldKeys),
	})
}

// Player death and damage events

//export onPlayerDeath
func onPlayerDeath(player, killer unsafe.Pointer, reason C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerDeath, &PlayerDeathEvent{
		Player: &DefaultPlayer{handle: player},
		Killer: &DefaultPlayer{handle: killer},
		Reason: int(reason),
	})
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player, from unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &DefaultPlayer{handle: player},
		From:   &DefaultPlayer{handle: from},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player, to unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerGiveDamage, &PlayerGiveDamageEvent{
		Player: &DefaultPlayer{handle: player},
		To:     &DefaultPlayer{handle: to},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(player unsafe.Pointer, x, y, z C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickMap, &PlayerClickMapEvent{
		Player:   &DefaultPlayer{handle: player},
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
	})
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player, clicked unsafe.Pointer, source C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player:  &DefaultPlayer{handle: player},
		Clicked: &DefaultPlayer{handle: clicked},
		Source:  PlayerClickSource(source),
	})
}

// Client check event

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results C.int) {
	event.Dispatch(eventDispatcher, EventTypeClientCheckResponse, &ClientCheckResponseEvent{
		Player:     &DefaultPlayer{handle: player},
		ActionType: int(actionType),
		Address:    int(address),
		Results:    int(results),
	})
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerUpdate, &PlayerUpdateEvent{
		Player: &DefaultPlayer{handle: player},
		Now:    int(now),
	})
}

// Player dialog event

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID, response, listItem C.int, inputText *C.char) {
	event.Dispatch(eventDispatcher, EventTypeDialogResponse, &DialogResponseEvent{
		Player:    &DefaultPlayer{handle: player},
		DialogID:  int(dialogID),
		Response:  DialogResponse(response),
		ListItem:  int(listItem),
		InputText: C.GoString(inputText),
	})
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(player, actor unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &DefaultPlayer{handle: player},
		Actor:  &DefaultPlayer{handle: actor},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onActorStreamIn
func onActorStreamIn(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &DefaultPlayer{handle: actor},
		ForPlayer: &DefaultPlayer{handle: player},
	})
}

//export onActorStreamOut
func onActorStreamOut(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &DefaultPlayer{handle: actor},
		ForPlayer: &DefaultPlayer{handle: player},
	})
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &DefaultVehicle{handle: vehicle},
		ForPlayer: &DefaultPlayer{handle: player},
	})
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &DefaultVehicle{handle: vehicle},
		ForPlayer: &DefaultPlayer{handle: player},
	})
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &DefaultVehicle{handle: vehicle},
		Killer:  &DefaultPlayer{handle: killer},
	})
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &DefaultPlayer{handle: player},
		Vehicle:     &DefaultVehicle{handle: vehicle},
		IsPassenger: int(isPassenger) != 0,
	})
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &DefaultPlayer{handle: player},
		Vehicle: &DefaultVehicle{handle: vehicle},
	})
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &DefaultVehicle{handle: vehicle},
		Player:  &DefaultPlayer{handle: player},
	})
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &DefaultPlayer{handle: player},
		Vehicle:  &DefaultVehicle{handle: vehicle},
		PaintJob: int(paintJob),
	})
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehicleMod, &VehicleModEvent{
		Player:    &DefaultPlayer{handle: player},
		Vehicle:   &DefaultVehicle{handle: vehicle},
		Component: int(component),
	})
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehicleRespray, &VehicleResprayEvent{
		Player:  &DefaultPlayer{handle: player},
		Vehicle: &DefaultVehicle{handle: vehicle},
		Color:   VehicleColor{Primary: int(color1), Secondary: int(color2)},
	})
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit, interiorID C.int) {
	event.Dispatch(eventDispatcher, EventTypeEnterExitModShop, &EnterExitModShopEvent{
		Player:     &DefaultPlayer{handle: player},
		EnterExit:  int(enterexit) != 0,
		InteriorID: int(interiorID),
	})
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &DefaultVehicle{handle: vehicle},
	})
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, seat C.uchar, posX, posY, posZ, velX, velY, velZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle:  &DefaultVehicle{handle: vehicle},
		Player:   &DefaultPlayer{handle: player},
		Seat:     int(seat),
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
		Velocity: Vector3{X: float32(velX), Y: float32(velY), Z: float32(velZ)},
	})
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &DefaultPlayer{handle: player},
		Vehicle: &DefaultVehicle{handle: vehicle},
	})
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState C.uchar) {
	event.Dispatch(eventDispatcher, EventTypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &DefaultPlayer{handle: player},
		Vehicle:    &DefaultVehicle{handle: vehicle},
		SirenState: int(sirenState),
	})
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeObjectMoved, &ObjectMovedEvent{
		Object: &GlobalObject{handle: object},
	})
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectMoved, &PlayerObjectMovedEvent{
		Player: &DefaultPlayer{handle: player},
		Object: &GlobalObject{handle: object},
	})
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeObjectSelected, &ObjectSelectedEvent{
		Player:   &DefaultPlayer{handle: player},
		Object:   &GlobalObject{handle: object},
		Model:    int(model),
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
	})
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectSelected, &PlayerObjectSelectedEvent{
		Player:   &DefaultPlayer{handle: player},
		Object:   &PlayerObject{handle: object},
		Model:    int(model),
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
	})
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeObjectEdited, &ObjectEditedEvent{
		Player:   &DefaultPlayer{handle: player},
		Object:   &GlobalObject{handle: object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
	})
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectEdited, &PlayerObjectEditedEvent{
		Player:   &DefaultPlayer{handle: player},
		Object:   &PlayerObject{handle: object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
	})
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index, saved, model, bone C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, scaleX, scaleY, scaleZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerAttachedObjectEdited, &PlayerAttachedObjectEditedEvent{
		Player:   &DefaultPlayer{handle: player},
		Index:    int(index),
		Saved:    int(saved),
		Model:    int(model),
		Bone:     int(bone),
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
		Scale:    Vector3{X: float32(scaleX), Y: float32(scaleY), Z: float32(scaleZ)},
	})
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickTextDraw, &PlayerClickTextDrawEvent{
		Player:   &DefaultPlayer{handle: player},
		TextDraw: &GlobalTextDraw{handle: textdraw},
	})
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickPlayerTextDraw, &PlayerClickPlayerTextDrawEvent{
		Player:   &DefaultPlayer{handle: player},
		TextDraw: &GlobalTextDraw{handle: textdraw},
	})
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerCancelTextDrawSelection, &PlayerCancelTextDrawSelectionEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerCancelPlayerTextDrawSelection, &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

// Player model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerFinishedDownloading, &PlayerFinishedDownloadingEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(player unsafe.Pointer, _type C.int, checksum C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerRequestDownload, &PlayerRequestDownloadEvent{
		Player:   &DefaultPlayer{handle: player},
		Type:     int(_type),
		Checksum: uint(checksum),
	})
}

// Console events. TODO

//export onConsoleText
func onConsoleText(command *C.char, parameters *C.char) {
	event.Dispatch(eventDispatcher, EventTypeConsoleText, &ConsoleTextEvent{
		Command:    C.GoString(command),
		Parameters: C.GoString(parameters),
	})
}

//export onRconLoginAttempt
func onRconLoginAttempt(player unsafe.Pointer, password *C.char, success C.int) {
	event.Dispatch(eventDispatcher, EventTypeRconLoginAttempt, &RconLoginAttemptEvent{
		Player:   &DefaultPlayer{handle: player},
		Password: C.GoString(password),
		Success:  int(success) != 0,
	})
}

// Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(player, pickup unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerPickUpPickup, &PlayerPickUpPickupEvent{
		Player: &DefaultPlayer{handle: player},
		Pickup: &DefaultPickup{handle: pickup},
	})
}

// GangZone events

//export onPlayerEnterGangZone
func onPlayerEnterGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterGangZone, &PlayerEnterGangZoneEvent{
		Player:   &DefaultPlayer{handle: player},
		GangZone: &GlobalGangZone{handle: gangZone},
	})
}

//export onPlayerLeaveGangZone
func onPlayerLeaveGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveGangZone, &PlayerLeaveGangZoneEvent{
		Player:   &DefaultPlayer{handle: player},
		GangZone: &GlobalGangZone{handle: gangZone},
	})
}

//export onPlayerClickGangZone
func onPlayerClickGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickGangZone, &PlayerClickGangZoneEvent{
		Player:   &DefaultPlayer{handle: player},
		GangZone: &GlobalGangZone{handle: gangZone},
	})
}

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(player unsafe.Pointer, menuRow C.uchar) {
	event.Dispatch(eventDispatcher, EventTypePlayerSelectedMenuRow, &PlayerSelectedMenuRowEvent{
		Player:  &DefaultPlayer{handle: player},
		MenuRow: uint8(menuRow),
	})
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerExitedMenu, &PlayerExitedMenuEvent{
		Player: &DefaultPlayer{handle: player},
	})
}

//export onPlayerRequestClass
func onPlayerRequestClass(player unsafe.Pointer, classID C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerRequestClass, &PlayerRequestClassEvent{
		Player:  &DefaultPlayer{handle: player},
		ClassID: uint(classID),
	})
}
