package gomp

// #include "component.h"
import "C"
import (
	"unsafe"

	"github.com/kodeyeen/gomp/event"
)

// go build -o test.dll -buildmode=c-shared

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

func main() {}

//export onGameModeInit
func onGameModeInit() {
	clibpath := C.CString("./components/Gomp.dll")
	defer C.free(unsafe.Pointer(clibpath))

	handle := C.loadLib(clibpath)

	C.initFuncs(handle)

	gm := &GameMode{}

	event.Dispatch(eventDispatcher, EventTypeGameModeInit, &GameModeInitEvent{
		GameMode: gm,
	})
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeIncomingConnection, &PlayerRequestSpawnEvent{
		Player: &Player{player},
	})
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeIncomingConnection, &PlayerSpawnEvent{
		Player: &Player{player},
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(eventDispatcher, EventTypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &Player{player},
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerConnect, &PlayerConnectEvent{
		Player: &Player{player},
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) {
	event.Dispatch(eventDispatcher, EventTypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: &Player{player},
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClientInit, &PlayerClientInitEvent{
		Player: &Player{player},
	})
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &Player{player},
		ForPlayer: &Player{forPlayer},
	})
}

//export onPlayerStreamOut
func onPlayerStreamOut(player, forPlayer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &Player{player},
		ForPlayer: &Player{forPlayer},
	})
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerText, &PlayerTextEvent{
		Player:  &Player{player},
		Message: C.GoString(message),
	})
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerCommandText, &PlayerCommandTextEvent{
		Player:  &Player{player},
		Message: C.GoString(message),
	})
}

// Player shot events

// TODO:
//
////export onPlayerWeaponShot
// func onPlayerWeaponShot(player unsafe.Pointer, weapon C.uchar, hitType C.int, hitID C.ushort) {
// 	event.Dispatch(eventDispatcher, EventTypePlayerCommandText, &PlayerCommandTextEvent{
// 		Player:  &Player{player},
// 		Message: C.GoString(message),
// 	})
// }

// Player data change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerScoreChange, &PlayerScoreChangeEvent{
		Player: &Player{player},
		Score:  int(score),
	})
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName *C.char) {
	event.Dispatch(eventDispatcher, EventTypePlayerNameChange, &PlayerNameChangeEvent{
		Player:  &Player{player},
		OldName: C.GoString(oldName),
	})
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior, oldInterior C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &Player{player},
		NewInterior: uint(newInterior),
		OldInterior: uint(oldInterior),
	})
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState, oldState C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &Player{player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys, oldKeys C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &Player{player},
		NewKeys: uint(newKeys),
		OldKeys: uint(oldKeys),
	})
}

// Player death and damage events

//export onPlayerDeath
func onPlayerDeath(player, killer unsafe.Pointer, reason C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerDeath, &PlayerDeathEvent{
		Player: &Player{player},
		Killer: &Player{killer},
		Reason: int(reason),
	})
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player, from unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &Player{player},
		From:   &Player{from},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player, to unsafe.Pointer, amount C.float, weapon C.uint, part C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerGiveDamage, &PlayerGiveDamageEvent{
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
	event.Dispatch(eventDispatcher, EventTypePlayerClickMap, &PlayerClickMapEvent{
		Player:   &Player{player},
		Position: &Position{float32(x), float32(y), float32(z)},
	})
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player, clicked unsafe.Pointer, source C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player:  &Player{player},
		Clicked: &Player{clicked},
		Source:  PlayerClickSource(source),
	})
}

// Client check event

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results C.int) {
	event.Dispatch(eventDispatcher, EventTypeClientCheckResponse, &ClientCheckResponseEvent{
		Player:     &Player{player},
		ActionType: int(actionType),
		Address:    int(address),
		Results:    int(results),
	})
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerUpdate, &PlayerUpdateEvent{
		Player: &Player{player},
		Now:    int(now),
	})
}

// Player dialog event

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID, response, listItem C.int, inputText *C.char) {
	event.Dispatch(eventDispatcher, EventTypeDialogResponse, &DialogResponseEvent{
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
	event.Dispatch(eventDispatcher, EventTypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &Player{player},
		Actor:  &Player{actor},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onActorStreamIn
func onActorStreamIn(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

//export onActorStreamOut
func onActorStreamOut(actor, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{actor},
		ForPlayer: &Player{player},
	})
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &Vehicle{vehicle},
		ForPlayer: &Player{player},
	})
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{vehicle},
		ForPlayer: &Player{player},
	})
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &Vehicle{vehicle},
		Killer:  &Player{killer},
	})
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger C.int) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &Player{player},
		Vehicle:     &Vehicle{vehicle},
		IsPassenger: int(isPassenger) != 0,
	})
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
	})
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{vehicle},
		Player:  &Player{player},
	})
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &Player{player},
		Vehicle:  &Vehicle{vehicle},
		PaintJob: int(paintJob),
	})
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehicleMod, &VehicleModEvent{
		Player:    &Player{player},
		Vehicle:   &Vehicle{vehicle},
		Component: int(component),
	})
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 C.int) {
	event.Dispatch(eventDispatcher, EventTypeVehicleRespray, &VehicleResprayEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
		Color:   VehicleColor{int(color1), int(color2)},
	})
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit, interiorID C.int) {
	event.Dispatch(eventDispatcher, EventTypeEnterExitModShop, &EnterExitModShopEvent{
		Player:     &Player{player},
		EnterExit:  int(enterexit) != 0,
		InteriorID: int(interiorID),
	})
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &Vehicle{vehicle},
	})
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, seat C.uchar, posX, posY, posZ, velX, velY, velZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle:  &Vehicle{vehicle},
		Player:   &Player{player},
		Seat:     int(seat),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
		Velocity: Vector3{float32(velX), float32(velY), float32(velZ)},
	})
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &Player{player},
		Vehicle: &Vehicle{vehicle},
	})
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState C.uchar) {
	event.Dispatch(eventDispatcher, EventTypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &Player{player},
		Vehicle:    &Vehicle{vehicle},
		SirenState: int(sirenState),
	})
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypeObjectMoved, &ObjectMovedEvent{
		Object: &Object{object, nil},
	})
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectMoved, &PlayerObjectMovedEvent{
		Player: &Player{player},
		Object: &Object{object, player},
	})
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeObjectSelected, &ObjectSelectedEvent{
		Player:   &Player{player},
		Object:   &Object{object, player},
		Model:    int(model),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
	})
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model C.int, posX, posY, posZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectSelected, &PlayerObjectSelectedEvent{
		Player:   &Player{player},
		Object:   &Object{object, player},
		Model:    int(model),
		Position: Vector3{float32(posX), float32(posY), float32(posZ)},
	})
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, EventTypeObjectEdited, &ObjectEditedEvent{
		Player:   &Player{player},
		Object:   &Object{object, player},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{float32(offsetX), float32(offsetY), float32(offsetZ)},
		Rotation: Vector3{float32(rotX), float32(rotY), float32(rotZ)},
	})
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerObjectEdited, &PlayerObjectEditedEvent{
		Player:   &Player{player},
		Object:   &Object{object, player},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{float32(offsetX), float32(offsetY), float32(offsetZ)},
		Rotation: Vector3{float32(rotX), float32(rotY), float32(rotZ)},
	})
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index, saved, model, bone C.int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, scaleX, scaleY, scaleZ C.float) {
	event.Dispatch(eventDispatcher, EventTypePlayerAttachedObjectEdited, &PlayerAttachedObjectEditedEvent{
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
	event.Dispatch(eventDispatcher, EventTypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{player},
	})
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickTextDraw, &PlayerClickTextDrawEvent{
		Player:   &Player{player},
		TextDraw: &TextDraw{textdraw, nil},
	})
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickPlayerTextDraw, &PlayerClickPlayerTextDrawEvent{
		Player:   &Player{player},
		TextDraw: &TextDraw{textdraw, nil},
	})
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerCancelTextDrawSelection, &PlayerCancelTextDrawSelectionEvent{
		Player: &Player{player},
	})
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerCancelPlayerTextDrawSelection, &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &Player{player},
	})
}

// Player model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerFinishedDownloading, &PlayerFinishedDownloadingEvent{
		Player: &Player{player},
	})
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(player unsafe.Pointer, _type C.int, checksum C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerRequestDownload, &PlayerRequestDownloadEvent{
		Player:   &Player{player},
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
		Player:   &Player{player},
		Password: C.GoString(password),
		Success:  int(success) != 0,
	})
}

// Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(player, pickup unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerPickUpPickup, &PlayerPickUpPickupEvent{
		Player: &Player{player},
		Pickup: &Pickup{pickup, nil},
	})
}

// GangZone events

//export onPlayerEnterGangZone
func onPlayerEnterGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerEnterGangZone, &PlayerEnterGangZoneEvent{
		Player:   &Player{player},
		GangZone: &GangZone{gangZone},
	})
}

//export onPlayerLeaveGangZone
func onPlayerLeaveGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerLeaveGangZone, &PlayerLeaveGangZoneEvent{
		Player:   &Player{player},
		GangZone: &GangZone{gangZone},
	})
}

//export onPlayerClickGangZone
func onPlayerClickGangZone(player, gangZone unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerClickGangZone, &PlayerClickGangZoneEvent{
		Player:   &Player{player},
		GangZone: &GangZone{gangZone},
	})
}

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(player unsafe.Pointer, menuRow C.uchar) {
	event.Dispatch(eventDispatcher, EventTypePlayerSelectedMenuRow, &PlayerSelectedMenuRowEvent{
		Player:  &Player{player},
		MenuRow: uint8(menuRow),
	})
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(player unsafe.Pointer) {
	event.Dispatch(eventDispatcher, EventTypePlayerExitedMenu, &PlayerExitedMenuEvent{
		Player: &Player{player},
	})
}

//export onPlayerRequestClass
func onPlayerRequestClass(player unsafe.Pointer, classID C.uint) {
	event.Dispatch(eventDispatcher, EventTypePlayerRequestClass, &PlayerRequestClassEvent{
		Player:  &Player{player},
		ClassID: uint(classID),
	})
}
