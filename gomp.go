package gomp

// #include <stdlib.h>
// #include "include/gomp.h"
import "C"
import (
	"strings"
	"time"
	"unsafe"

	"github.com/kodeyeen/event"
)

// go build -o test.dll -buildmode=c-shared

type Animation struct {
	Lib, Name                  string
	Delta                      float32
	Loop, LockX, LockY, Freeze bool
	Duration                   time.Duration
}

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

type Color uint

var evtDispatcher = event.NewDispatcher()
var cmdManager = newCommandManager()

func On(evtType event.Type, handler any) {
	evtDispatcher.On(evtType, handler)
}

func Once(evtType event.Type, handler any) {
	evtDispatcher.Once(evtType, handler)
}

func Off(evtType event.Type, handler any) {
	evtDispatcher.Off(evtType, handler)
}

func Dispatch[T any](evtType event.Type, evt T) {
	event.Dispatch(evtDispatcher, evtType, evt)
}

func AddCommand(name string, handler CommandHandler) {
	cmdManager.add(name, handler)
}

//export onGameModeInit
func onGameModeInit() {
	clibpath := C.CString("./components/Gomponent.dll")
	defer C.free(unsafe.Pointer(clibpath))

	handle := C.openLib(clibpath)

	C.initFuncs(handle)

	event.Dispatch(evtDispatcher, EventTypeGameModeInit, &GameModeInitEvent{})
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerRequesSpawn, &PlayerRequestSpawnEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerSpawn, &PlayerSpawnEvent{
		Player: &Player{handle: player},
	})
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress *C.char, port C.ushort) {
	event.Dispatch(evtDispatcher, EventTypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &Player{handle: player},
		IPAddress: C.GoString(ipAddress),
		Port:      int(port),
	})
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerConnect, &PlayerConnectEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) {
	event.Dispatch(evtDispatcher, EventTypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: &Player{handle: player},
		Reason: DisconnectReason(reason),
	})
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerClientInit, &PlayerClientInitEvent{
		Player: &Player{handle: player},
	})
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player, forPlayer unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})
}

//export onPlayerStreamOut
func onPlayerStreamOut(player, forPlayer unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) {
	event.Dispatch(evtDispatcher, EventTypePlayerText, &PlayerTextEvent{
		Player:  &Player{handle: player},
		Message: C.GoString(message),
	})
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message C.String) bool {
	rawCmd := C.GoStringN(message.buf, C.int(message.length))

	tmp := strings.Fields(rawCmd)
	name := strings.TrimPrefix(tmp[0], "/")
	args := tmp[1:]

	exists := cmdManager.has(name)
	if !exists {
		return false
	}

	cmdManager.run(name, &Command{
		Sender:   &Player{handle: player},
		Name:     name,
		Args:     args,
		RawValue: rawCmd,
	})

	return true
}

// Player shot events

// TODO:
//
////export onPlayerWeaponShot
// func onPlayerWeaponShot(player unsafe.Pointer, weapon C.uchar, hitType int, hitID C.ushort) {
// 	event.Dispatch(evtDispatcher, EventTypePlayerCommandText, &PlayerCommandTextEvent{
// 		Player:  &Player{handle: player},
// 		Message: C.GoString(message),
// 	})
// }

// Player data change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score int) {
	event.Dispatch(evtDispatcher, EventTypePlayerScoreChange, &PlayerScoreChangeEvent{
		Player: &Player{handle: player},
		Score:  score,
	})
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName *C.char) {
	event.Dispatch(evtDispatcher, EventTypePlayerNameChange, &PlayerNameChangeEvent{
		Player:  &Player{handle: player},
		OldName: C.GoString(oldName),
	})
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior, oldInterior uint) {
	event.Dispatch(evtDispatcher, EventTypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &Player{handle: player},
		NewInterior: uint(newInterior),
		OldInterior: uint(oldInterior),
	})
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState, oldState int) {
	event.Dispatch(evtDispatcher, EventTypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &Player{handle: player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys, oldKeys C.uint) {
	event.Dispatch(evtDispatcher, EventTypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &Player{handle: player},
		NewKeys: uint(newKeys),
		OldKeys: uint(oldKeys),
	})
}

// Player death and damage events

//export onPlayerDeath
func onPlayerDeath(player, killer unsafe.Pointer, reason int) {
	evtKiller := &Player{handle: killer}
	if killer == nil {
		evtKiller = nil
	}

	event.Dispatch(evtDispatcher, EventTypePlayerDeath, &PlayerDeathEvent{
		Player: &Player{handle: player},
		Killer: evtKiller,
		Reason: reason,
	})
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player, from unsafe.Pointer, amount C.float, weapon C.uint, part int) {
	event.Dispatch(evtDispatcher, EventTypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &Player{handle: player},
		From:   &Player{handle: from},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player, to unsafe.Pointer, amount C.float, weapon C.uint, part int) {
	event.Dispatch(evtDispatcher, EventTypePlayerGiveDamage, &PlayerGiveDamageEvent{
		Player: &Player{handle: player},
		To:     &Player{handle: to},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(player unsafe.Pointer, x, y, z C.float) {
	event.Dispatch(evtDispatcher, EventTypePlayerClickMap, &PlayerClickMapEvent{
		Player:   &Player{handle: player},
		Position: Vector3{X: float32(x), Y: float32(y), Z: float32(z)},
	})
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player, clicked unsafe.Pointer, source int) {
	event.Dispatch(evtDispatcher, EventTypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player:  &Player{handle: player},
		Clicked: &Player{handle: clicked},
		Source:  PlayerClickSource(source),
	})
}

// Client check event

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results int) {
	event.Dispatch(evtDispatcher, EventTypeClientCheckResponse, &ClientCheckResponseEvent{
		Player:     &Player{handle: player},
		ActionType: actionType,
		Address:    address,
		Results:    results,
	})
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.long) bool {
	return event.Dispatch(evtDispatcher, EventTypePlayerUpdate, &PlayerUpdateEvent{
		Player: &Player{handle: player},
		Now:    time.Unix(0, int64(now)*int64(time.Millisecond)),
	})
}

// Player dialog event

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID, response, listItem int, inputText *C.char) {
	event.Dispatch(evtDispatcher, EventTypeDialogResponse, &DialogResponseEvent{
		Player:    &Player{handle: player},
		DialogID:  dialogID,
		Response:  DialogResponse(response),
		ListItem:  listItem,
		InputText: C.GoString(inputText),
	})
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(player, actor unsafe.Pointer, amount C.float, weapon C.uint, part int) {
	event.Dispatch(evtDispatcher, EventTypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &Player{handle: player},
		Actor:  &Player{handle: actor},
		Amount: float32(amount),
		Weapon: uint(weapon),
		Part:   BodyPart(part),
	})
}

//export onActorStreamIn
func onActorStreamIn(actor, player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: player},
	})
}

//export onActorStreamOut
func onActorStreamOut(actor, player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: player},
	})
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Killer:  &Player{handle: killer},
	})
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger int) {
	event.Dispatch(evtDispatcher, EventTypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &Player{handle: player},
		Vehicle:     &Vehicle{handle: vehicle},
		IsPassenger: isPassenger != 0,
	})
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Player:  &Player{handle: player},
	})
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob int) {
	event.Dispatch(evtDispatcher, EventTypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &Player{handle: player},
		Vehicle:  &Vehicle{handle: vehicle},
		PaintJob: paintJob,
	})
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component int) {
	event.Dispatch(evtDispatcher, EventTypeVehicleMod, &VehicleModEvent{
		Player:    &Player{handle: player},
		Vehicle:   &Vehicle{handle: vehicle},
		Component: component,
	})
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 int) {
	event.Dispatch(evtDispatcher, EventTypeVehicleRespray, &VehicleResprayEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
		Color:   VehicleColor{Primary: Color(color1), Secondary: Color(color2)},
	})
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit, interiorID int) {
	event.Dispatch(evtDispatcher, EventTypeEnterExitModShop, &EnterExitModShopEvent{
		Player:     &Player{handle: player},
		EnterExit:  enterexit != 0,
		InteriorID: interiorID,
	})
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &Vehicle{handle: vehicle},
	})
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, seat C.uchar, posX, posY, posZ, velX, velY, velZ C.float) {
	event.Dispatch(evtDispatcher, EventTypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle:  &Vehicle{handle: vehicle},
		Player:   &Player{handle: player},
		Seat:     int(seat),
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
		Velocity: Vector3{X: float32(velX), Y: float32(velY), Z: float32(velZ)},
	})
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState C.uchar) {
	event.Dispatch(evtDispatcher, EventTypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &Player{handle: player},
		Vehicle:    &Vehicle{handle: vehicle},
		SirenState: int(sirenState),
	})
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypeObjectMoved, &ObjectMovedEvent{
		Object: &Object{handle: object},
	})
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerObjectMoved, &PlayerObjectMovedEvent{
		Player: &Player{handle: player},
		Object: &Object{handle: object},
	})
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model int, posX, posY, posZ C.float) {
	event.Dispatch(evtDispatcher, EventTypeObjectSelected, &ObjectSelectedEvent{
		Player:   &Player{handle: player},
		Object:   &Object{handle: object},
		Model:    model,
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
	})
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model int, posX, posY, posZ C.float) {
	event.Dispatch(evtDispatcher, EventTypePlayerObjectSelected, &PlayerObjectSelectedEvent{
		Player:   &Player{handle: player},
		Object:   &Object{handle: object},
		Model:    model,
		Position: Vector3{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
	})
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(evtDispatcher, EventTypeObjectEdited, &ObjectEditedEvent{
		Player:   &Player{handle: player},
		Object:   &Object{handle: object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
	})
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ C.float) {
	event.Dispatch(evtDispatcher, EventTypePlayerObjectEdited, &PlayerObjectEditedEvent{
		Player:   &Player{handle: player},
		Object:   &Object{handle: object},
		Response: ObjectEditResponse(response),
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
	})
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index, saved, model, bone int, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, scaleX, scaleY, scaleZ C.float) {
	event.Dispatch(evtDispatcher, EventTypePlayerAttachedObjectEdited, &PlayerAttachedObjectEditedEvent{
		Player:   &Player{handle: player},
		Index:    index,
		Saved:    saved,
		Model:    model,
		Bone:     bone,
		Offset:   Vector3{X: float32(offsetX), Y: float32(offsetY), Z: float32(offsetZ)},
		Rotation: Vector3{X: float32(rotX), Y: float32(rotY), Z: float32(rotZ)},
		Scale:    Vector3{X: float32(scaleX), Y: float32(scaleY), Z: float32(scaleZ)},
	})
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerClickTextDraw, &PlayerClickTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &Textdraw{handle: textdraw},
	})
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(player, textdraw unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerClickPlayerTextDraw, &PlayerClickPlayerTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &Textdraw{handle: textdraw},
	})
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerCancelTextDrawSelection, &PlayerCancelTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerCancelPlayerTextDrawSelection, &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})
}

// Player model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerFinishedDownloading, &PlayerFinishedDownloadingEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(player unsafe.Pointer, _type int, checksum C.uint) {
	event.Dispatch(evtDispatcher, EventTypePlayerRequestDownload, &PlayerRequestDownloadEvent{
		Player:   &Player{handle: player},
		Type:     _type,
		Checksum: uint(checksum),
	})
}

// Console events. TODO

//export onConsoleText
func onConsoleText(command *C.char, parameters *C.char) {
	event.Dispatch(evtDispatcher, EventTypeConsoleText, &ConsoleTextEvent{
		Command:    C.GoString(command),
		Parameters: C.GoString(parameters),
	})
}

//export onRconLoginAttempt
func onRconLoginAttempt(player unsafe.Pointer, password *C.char, success int) {
	event.Dispatch(evtDispatcher, EventTypeRconLoginAttempt, &RconLoginAttemptEvent{
		Player:   &Player{handle: player},
		Password: C.GoString(password),
		Success:  success != 0,
	})
}

// Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(player, pickup unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerPickUpPickup, &PlayerPickUpPickupEvent{
		Player: &Player{handle: player},
		Pickup: &Pickup{handle: pickup},
	})
}

// Turf events

//export onPlayerEnterTurf
func onPlayerEnterTurf(player, turf unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerEnterTurf, &PlayerEnterTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: turf},
	})
}

//export onPlayerLeaveTurf
func onPlayerLeaveTurf(player, turf unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerLeaveTurf, &PlayerLeaveTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: turf},
	})
}

//export onPlayerClickTurf
func onPlayerClickTurf(player, turf unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerClickTurf, &PlayerClickTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: turf},
	})
}

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(player unsafe.Pointer, menuRow C.uchar) {
	event.Dispatch(evtDispatcher, EventTypePlayerSelectedMenuRow, &PlayerSelectedMenuRowEvent{
		Player:  &Player{handle: player},
		MenuRow: uint8(menuRow),
	})
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(player unsafe.Pointer) {
	event.Dispatch(evtDispatcher, EventTypePlayerExitedMenu, &PlayerExitedMenuEvent{
		Player: &Player{handle: player},
	})
}

//export onPlayerRequestClass
func onPlayerRequestClass(player, class unsafe.Pointer) bool {
	return event.Dispatch(evtDispatcher, EventTypePlayerRequestClass, &PlayerRequestClassEvent{
		Player: &Player{handle: player},
		Class:  &Class{handle: class},
	})
}
