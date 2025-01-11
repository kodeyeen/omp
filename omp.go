package omp

// #include <stdlib.h>
// #include "include/omp.h"
import "C"
import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"
	"unsafe"
)

type Animation struct {
	Lib, Name                  string
	Delta                      float32
	Loop, LockX, LockY, Freeze bool
	Duration                   time.Duration
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Vector3 struct {
	X, Y, Z float32
}

type Vector2 struct {
	X, Y float32
}

type Color uint

var DefaultDispatcher = NewDispatcher()
var eventListener Listener

func EventListener() Listener {
	if eventListener == nil {
		eventListener = DefaultDispatcher
	}

	return eventListener
}

func SetEventListener(listener Listener) {
	eventListener = listener
}

func Listen(_type EventType, listener Listener) {
	DefaultDispatcher.Listen(_type, listener)
}

func ListenFunc(_type EventType, listener func(context.Context, Event) error) {
	DefaultDispatcher.ListenFunc(_type, listener)
}

func handlePanic() {
	if r := recover(); r != nil {
		stackTrace := strings.TrimSuffix(string(debug.Stack()), "\n")

		Log(LogLevelError, "%s", fmt.Sprint(r))
		Log(LogLevelError, "%s", stackTrace)
	}
}

//export onGameModeInit
func onGameModeInit() C.bool {
	defer handlePanic()

	C.loadComponent()

	evt := NewEvent(EventTypeGameModeInit, nil)
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onGameModeExit
func onGameModeExit() C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeGameModeExit, nil)
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(player, actor unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player: &Player{handle: player},
		Actor:  &Player{handle: actor},
		Amount: amount,
		Weapon: weapon,
		Part:   BodyPart(part),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onActorStreamOut
func onActorStreamOut(actor, forPlayer unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: forPlayer},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onActorStreamIn
func onActorStreamIn(actor, forPlayer unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: forPlayer},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Class events

//export onPlayerRequestClass
func onPlayerRequestClass(player, class unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerRequestClass, &PlayerRequestClassEvent{
		Player: &Player{handle: player},
		Class:  &Class{handle: class},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Console events. TODO

//export onConsoleText
func onConsoleText(command C.String, parameters C.String) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeConsoleText, &ConsoleTextEvent{
		Command:    C.GoStringN(command.buf, C.int(command.length)),
		Parameters: C.GoStringN(parameters.buf, C.int(parameters.length)),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onRconLoginAttempt
func onRconLoginAttempt(player unsafe.Pointer, password C.String, success bool) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeRconLoginAttempt, &RconLoginAttemptEvent{
		Player:   &Player{handle: player},
		Password: C.GoStringN(password.buf, C.int(password.length)),
		Success:  success,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Custom model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerFinishedDownloading, &PlayerFinishedDownloadingEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(player unsafe.Pointer, _type uint8, checksum uint32) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerRequestDownload, &PlayerRequestDownloadEvent{
		Player:   &Player{handle: player},
		Type:     int(_type),
		Checksum: int(checksum),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Dialog events

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID, response, listItem int, inputText C.String) C.bool {
	defer handlePanic()

	ctx := context.Background()
	eventPlayer := &Player{handle: player}
	playerID := eventPlayer.ID()

	dialog, ok := activeDialogs[playerID]
	if !ok {
		panic("active dialog is not set")
	}
	delete(activeDialogs, playerID)

	var err error

	switch d := dialog.(type) {
	case *MessageDialog:
		evt := NewEvent(EventTypeDialogResponse, &MessageDialogResponseEvent{
			Player:   eventPlayer,
			Response: DialogResponse(response),
		})
		err = d.Events.HandleEvent(ctx, evt)
	case *InputDialog:
		evt := NewEvent(EventTypeDialogResponse, &InputDialogResponseEvent{
			Player:    eventPlayer,
			Response:  DialogResponse(response),
			InputText: C.GoStringN(inputText.buf, C.int(inputText.length)),
		})
		err = d.Events.HandleEvent(ctx, evt)
	case *ListDialog:
		evt := NewEvent(EventTypeDialogResponse, &ListDialogResponseEvent{
			Player:     eventPlayer,
			Response:   DialogResponse(response),
			ItemNumber: listItem,
			Item:       C.GoStringN(inputText.buf, C.int(inputText.length)),
		})
		err = d.Events.HandleEvent(ctx, evt)
	case *TabListDialog:
		evt := NewEvent(EventTypeDialogResponse, &TabListDialogResponseEvent{
			Player:     eventPlayer,
			Response:   DialogResponse(response),
			ItemNumber: listItem,
			Item:       d.items[listItem],
		})
		err = d.Events.HandleEvent(ctx, evt)
	default:
		panic("unknown dialog type")
	}

	return err == nil
}

// GangZone events

//export onPlayerEnterGangZone
func onPlayerEnterGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerEnterTurf, &PlayerEnterTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerEnterPlayerGangZone
func onPlayerEnterPlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerEnterPlayerTurf, &PlayerEnterPlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerLeaveGangZone
func onPlayerLeaveGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerLeaveTurf, &PlayerLeaveTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerLeavePlayerGangZone
func onPlayerLeavePlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerLeavePlayerTurf, &PlayerLeavePlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerClickGangZone
func onPlayerClickGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickTurf, &PlayerClickTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerClickPlayerGangZone
func onPlayerClickPlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickPlayerTurf, &PlayerClickPlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(player unsafe.Pointer, menuRow uint8) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerSelectedMenuRow, &PlayerSelectedMenuRowEvent{
		Player:  &Player{handle: player},
		MenuRow: menuRow,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerExitedMenu, &PlayerExitedMenuEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeObjectMoved, &ObjectMovedEvent{
		Object: &Object{handle: object},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerObjectMoved, &PlayerObjectMovedEvent{
		Player: &Player{handle: player},
		Object: &PlayerObject{handle: object},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model int, pos C.Vector3) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeObjectSelected, &ObjectSelectedEvent{
		Player: &Player{handle: player},
		Object: &Object{handle: object},
		Model:  model,
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model int, pos C.Vector3) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerObjectSelected, &PlayerObjectSelectedEvent{
		Player: &Player{handle: player},
		Object: &PlayerObject{handle: object},
		Model:  model,
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response int, offset, rot C.Vector3) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeObjectEdited, &ObjectEditedEvent{
		Player:   &Player{handle: player},
		Object:   &Object{handle: object},
		Response: ObjectEditResponse(response),
		Offset: Vector3{
			X: float32(offset.x),
			Y: float32(offset.y),
			Z: float32(offset.z),
		},
		Rotation: Vector3{
			X: float32(rot.x),
			Y: float32(rot.y),
			Z: float32(rot.z),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response int, offset, rot C.Vector3) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerObjectEdited, &PlayerObjectEditedEvent{
		Player:   &Player{handle: player},
		Object:   &PlayerObject{handle: object},
		Response: ObjectEditResponse(response),
		Offset: Vector3{
			X: float32(offset.x),
			Y: float32(offset.y),
			Z: float32(offset.z),
		},
		Rotation: Vector3{
			X: float32(rot.x),
			Y: float32(rot.y),
			Z: float32(rot.z),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index int, saved bool, data C.PlayerAttachedObject) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerAttachmentEdited, &PlayerAttachmentEdited{
		Player: &Player{handle: player},
		Index:  index,
		Saved:  saved,
		Attachment: PlayerAttachment{
			ModelID: int(data.model),
			Bone:    PlayerBone(data.bone),
			Offset: Vector3{
				X: float32(data.offset.x),
				Y: float32(data.offset.y),
				Z: float32(data.offset.z),
			},
			Rot: Vector3{
				X: float32(data.rotation.x),
				Y: float32(data.rotation.y),
				Z: float32(data.rotation.z),
			},
			Scale: Vector3{
				X: float32(data.scale.x),
				Y: float32(data.scale.y),
				Z: float32(data.scale.z),
			},
			Color1: Color(data.colour1),
			Color2: Color(data.colour2),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// // Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(player, pickup unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerPickUpPickup, &PlayerPickUpPickupEvent{
		Player: &Player{handle: player},
		Pickup: &Pickup{handle: pickup},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerPickUpPlayerPickup
func onPlayerPickUpPlayerPickup(player, pickup unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerPickUpPlayerPickup, &PlayerPickUpPlayerPickupEvent{
		Player: &Player{handle: player},
		Pickup: &PlayerPickup{handle: pickup},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerRequestSpawn, &PlayerRequestSpawnEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerSpawn, &PlayerSpawnEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress C.String, port C.ushort) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &Player{handle: player},
		IPAddress: C.GoStringN(ipAddress.buf, C.int(ipAddress.length)),
		Port:      int(port),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerConnect, &PlayerConnectEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) C.bool {
	defer handlePanic()

	eventPlayer := &Player{handle: player}

	evt := NewEvent(EventTypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: eventPlayer,
		Reason: DisconnectReason(reason),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	delete(activeDialogs, eventPlayer.ID())

	return err == nil
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClientInit, &PlayerClientInitEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player, forPlayer unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerStreamOut
func onPlayerStreamOut(player, forPlayer unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerText, &PlayerTextEvent{
		Player:  &Player{handle: player},
		Message: C.GoString(message),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message C.String) C.bool {
	defer handlePanic()

	rawVal := C.GoStringN(message.buf, C.int(message.length))

	tmp := strings.Fields(rawVal)
	name := strings.TrimPrefix(tmp[0], "/")
	args := tmp[1:]

	evt := NewEvent(EventTypePlayerText, &PlayerCommandTextEvent{
		Sender:   &Player{handle: player},
		Name:     name,
		Args:     args,
		RawValue: rawVal,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player shot events

//export onPlayerShotMissed
func onPlayerShotMissed(player unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerShotMissed, &PlayerShotMissedEvent{
		Player: &Player{handle: player},
		Bullet: PlayerBullet{
			Origin: Vector3{
				X: float32(bulletData.origin.x),
				Y: float32(bulletData.origin.y),
				Z: float32(bulletData.origin.z),
			},
			HitPos: Vector3{
				X: float32(bulletData.hitPos.x),
				Y: float32(bulletData.hitPos.y),
				Z: float32(bulletData.hitPos.z),
			},
			Offset: Vector3{
				X: float32(bulletData.offset.x),
				Y: float32(bulletData.offset.y),
				Z: float32(bulletData.offset.z),
			},
			Weapon: Weapon(bulletData.weapon),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerShotPlayer
func onPlayerShotPlayer(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerShotPlayer, &PlayerShotPlayerEvent{
		Player: &Player{handle: player},
		Target: &Player{handle: target},
		Bullet: PlayerBullet{
			Origin: Vector3{
				X: float32(bulletData.origin.x),
				Y: float32(bulletData.origin.y),
				Z: float32(bulletData.origin.z),
			},
			HitPos: Vector3{
				X: float32(bulletData.hitPos.x),
				Y: float32(bulletData.hitPos.y),
				Z: float32(bulletData.hitPos.z),
			},
			Offset: Vector3{
				X: float32(bulletData.offset.x),
				Y: float32(bulletData.offset.y),
				Z: float32(bulletData.offset.z),
			},
			Weapon: Weapon(bulletData.weapon),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerShotVehicle
func onPlayerShotVehicle(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerShotVehicle, &PlayerShotVehicleEvent{
		Player: &Player{handle: player},
		Target: &Vehicle{handle: target},
		Bullet: PlayerBullet{
			Origin: Vector3{
				X: float32(bulletData.origin.x),
				Y: float32(bulletData.origin.y),
				Z: float32(bulletData.origin.z),
			},
			HitPos: Vector3{
				X: float32(bulletData.hitPos.x),
				Y: float32(bulletData.hitPos.y),
				Z: float32(bulletData.hitPos.z),
			},
			Offset: Vector3{
				X: float32(bulletData.offset.x),
				Y: float32(bulletData.offset.y),
				Z: float32(bulletData.offset.z),
			},
			Weapon: Weapon(bulletData.weapon),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerShotObject
func onPlayerShotObject(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerShotObject, &PlayerShotObjectEvent{
		Player: &Player{handle: player},
		Target: &Object{handle: target},
		Bullet: PlayerBullet{
			Origin: Vector3{
				X: float32(bulletData.origin.x),
				Y: float32(bulletData.origin.y),
				Z: float32(bulletData.origin.z),
			},
			HitPos: Vector3{
				X: float32(bulletData.hitPos.x),
				Y: float32(bulletData.hitPos.y),
				Z: float32(bulletData.hitPos.z),
			},
			Offset: Vector3{
				X: float32(bulletData.offset.x),
				Y: float32(bulletData.offset.y),
				Z: float32(bulletData.offset.z),
			},
			Weapon: Weapon(bulletData.weapon),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerShotPlayerObject
func onPlayerShotPlayerObject(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerShotPlayerObject, &PlayerShotPlayerObjectEvent{
		Player: &Player{handle: player},
		Target: &PlayerObject{handle: target},
		Bullet: PlayerBullet{
			Origin: Vector3{
				X: float32(bulletData.origin.x),
				Y: float32(bulletData.origin.y),
				Z: float32(bulletData.origin.z),
			},
			HitPos: Vector3{
				X: float32(bulletData.hitPos.x),
				Y: float32(bulletData.hitPos.y),
				Z: float32(bulletData.hitPos.z),
			},
			Offset: Vector3{
				X: float32(bulletData.offset.x),
				Y: float32(bulletData.offset.y),
				Z: float32(bulletData.offset.z),
			},
			Weapon: Weapon(bulletData.weapon),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerScoreChange, &PlayerScoreChangeEvent{
		Player: &Player{handle: player},
		Score:  score,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName C.String) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerNameChange, &PlayerNameChangeEvent{
		Player:  &Player{handle: player},
		OldName: C.GoStringN(oldName.buf, C.int(oldName.length)),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior, oldInterior uint) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &Player{handle: player},
		NewInterior: newInterior,
		OldInterior: oldInterior,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState, oldState int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &Player{handle: player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys, oldKeys uint) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &Player{handle: player},
		NewKeys: newKeys,
		OldKeys: oldKeys,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player damage events

//export onPlayerDeath
func onPlayerDeath(player, killer unsafe.Pointer, reason int) C.bool {
	defer handlePanic()

	eventKiller := &Player{handle: killer}
	if killer == nil {
		eventKiller = nil
	}

	evt := NewEvent(EventTypePlayerDeath, &PlayerDeathEvent{
		Player: &Player{handle: player},
		Killer: eventKiller,
		Reason: reason,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player, from unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer handlePanic()

	eventFrom := &Player{handle: from}
	if from == nil {
		eventFrom = nil
	}

	evt := NewEvent(EventTypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &Player{handle: player},
		From:   eventFrom,
		Amount: amount,
		Weapon: Weapon(weapon),
		Part:   BodyPart(part),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player, to unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerGiveDamage, &PlayerGiveDamageEvent{
		Player: &Player{handle: player},
		To:     &Player{handle: to},
		Amount: amount,
		Weapon: Weapon(weapon),
		Part:   BodyPart(part),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(player unsafe.Pointer, pos C.Vector3) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickMap, &PlayerClickMapEvent{
		Player: &Player{handle: player},
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player, clicked unsafe.Pointer, source int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player:  &Player{handle: player},
		Clicked: &Player{handle: clicked},
		Source:  PlayerClickSource(source),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player check events

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeClientCheckResponse, &ClientCheckResponseEvent{
		Player:     &Player{handle: player},
		ActionType: actionType,
		Address:    address,
		Results:    results,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.longlong) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerUpdate, &PlayerUpdateEvent{
		Player: &Player{handle: player},
		Now:    time.Unix(0, int64(now)*int64(time.Millisecond)),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Textdraw events

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(player, textdraw unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickTextDraw, &PlayerClickTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &Textdraw{handle: textdraw},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(player, textdraw unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerClickPlayerTextDraw, &PlayerClickPlayerTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &PlayerTextdraw{handle: textdraw},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerCancelTextDrawSelection, &PlayerCancelTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerCancelPlayerTextDrawSelection, &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Killer:  &Player{handle: killer},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &Player{handle: player},
		Vehicle:     &Vehicle{handle: vehicle},
		IsPassenger: isPassenger != 0,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Player:  &Player{handle: player},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &Player{handle: player},
		Vehicle:  &Vehicle{handle: vehicle},
		PaintJob: paintJob,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleMod, &VehicleModEvent{
		Player:    &Player{handle: player},
		Vehicle:   &Vehicle{handle: vehicle},
		Component: component,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleRespray, &VehicleResprayEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
		Color:   VehicleColor{Primary: color1, Secondary: color2},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit bool, interiorID int) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeEnterExitModShop, &EnterExitModShopEvent{
		Player:     &Player{handle: player},
		EnterExit:  enterexit,
		InteriorID: interiorID,
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &Vehicle{handle: vehicle},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, updateData C.UnoccupiedVehicleUpdate) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Player:  &Player{handle: player},
		Update: UnoccupiedVehicleUpdate{
			Seat: int(updateData.seat),
			Position: Vector3{
				X: float32(updateData.position.x),
				Y: float32(updateData.position.y),
				Z: float32(updateData.position.z),
			},
			Velocity: Vector3{
				X: float32(updateData.velocity.x),
				Y: float32(updateData.velocity.y),
				Z: float32(updateData.velocity.z),
			},
		},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState uint8) C.bool {
	defer handlePanic()

	evt := NewEvent(EventTypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &Player{handle: player},
		Vehicle:    &Vehicle{handle: vehicle},
		SirenState: int(sirenState),
	})
	err := EventListener().HandleEvent(context.Background(), evt)

	return err == nil
}
