package omp

// #include <stdlib.h>
// #include "include/omp.h"
import "C"
import (
	"context"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
	"unsafe"
)

var baseEventEmitter EventEmitter = newEventEmitter()

func BaseEventEmitter() EventEmitter {
	return baseEventEmitter
}

var events EventEmitter = baseEventEmitter

func SetEventEmitter(em EventEmitter) {
	events = em
}

func Emit(ctx context.Context, e any) error {
	return events.Emit(ctx, e)
}

func On[T any](f func(context.Context, T) error) {
	t := reflect.TypeFor[T]()

	events.addListener(t, func(ctx context.Context, e any) error {
		return f(ctx, e.(T))
	})
}

func recoverPanic() {
	if r := recover(); r != nil {
		stackTrace := strings.TrimSuffix(string(debug.Stack()), "\n")

		Log(LogLevelError, "%s", fmt.Sprint(r))
		Log(LogLevelError, "%s", stackTrace)
	}
}

//export onGameModeInit
func onGameModeInit() C.bool {
	defer recoverPanic()

	C.loadComponent()

	err := events.Emit(context.Background(), &GameModeInitEvent{})

	return err == nil
}

//export onGameModeExit
func onGameModeExit() C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &GameModeExitEvent{})

	return err == nil
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(player, actor unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerGiveDamageActorEvent{
		Player: &Player{handle: player},
		Actor:  &Player{handle: actor},
		Amount: amount,
		Weapon: weapon,
		Part:   BodyPart(part),
	})

	return err == nil
}

//export onActorStreamOut
func onActorStreamOut(actor, forPlayer unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ActorStreamOutEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: forPlayer},
	})

	return err == nil
}

//export onActorStreamIn
func onActorStreamIn(actor, forPlayer unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ActorStreamInEvent{
		Actor:     &Player{handle: actor},
		ForPlayer: &Player{handle: forPlayer},
	})

	return err == nil
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerEnterCheckpointEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerLeaveCheckpointEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerEnterRaceCheckpointEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

// Class events

//export onPlayerRequestClass
func onPlayerRequestClass(player, class unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerRequestClassEvent{
		Player: &Player{handle: player},
		Class:  &Class{handle: class},
	})

	return err == nil
}

// Console events. TODO

//export onConsoleText
func onConsoleText(command C.String, parameters C.String) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ConsoleTextEvent{
		Command:    C.GoStringN(command.buf, C.int(command.length)),
		Parameters: C.GoStringN(parameters.buf, C.int(parameters.length)),
	})

	return err == nil
}

//export onRconLoginAttempt
func onRconLoginAttempt(player unsafe.Pointer, password C.String, success bool) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &RconLoginAttemptEvent{
		Player:   &Player{handle: player},
		Password: C.GoStringN(password.buf, C.int(password.length)),
		Success:  success,
	})

	return err == nil
}

// Custom model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerFinishedDownloadingEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(player unsafe.Pointer, _type uint8, checksum uint32) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerRequestDownloadEvent{
		Player:   &Player{handle: player},
		Type:     int(_type),
		Checksum: int(checksum),
	})

	return err == nil
}

// Dialog events

//export onDialogResponse
func onDialogResponse(player unsafe.Pointer, dialogID, response, listItem int, inputText C.String) C.bool {
	defer recoverPanic()

	ctx := context.Background()
	eventPlayer := &Player{handle: player}
	playerID := eventPlayer.ID()

	d, ok := activeDialogs[playerID]
	if !ok {
		panic("active dialog is not set")
	}
	delete(activeDialogs, playerID)

	var err error

	switch d := d.(type) {
	case *MessageDialog:
		if d.onResponse != nil {
			err = d.onResponse(ctx, &MessageDialogResponseEvent{
				Player:   eventPlayer,
				Response: DialogResponse(response),
			})
		}
	case *InputDialog:
		if d.onResponse != nil {
			err = d.onResponse(ctx, &InputDialogResponseEvent{
				Player:    eventPlayer,
				Response:  DialogResponse(response),
				InputText: C.GoStringN(inputText.buf, C.int(inputText.length)),
			})
		}
	case *ListDialog:
		if d.onResponse != nil {
			err = d.onResponse(ctx, &ListDialogResponseEvent{
				Player:     eventPlayer,
				Response:   DialogResponse(response),
				ItemNumber: listItem,
				Item:       C.GoStringN(inputText.buf, C.int(inputText.length)),
			})
		}
	case *TabListDialog:
		if d.onResponse != nil {
			err = d.onResponse(ctx, &TabListDialogResponseEvent{
				Player:     eventPlayer,
				Response:   DialogResponse(response),
				ItemNumber: listItem,
				Item:       d.items[listItem],
			})
		}
	}

	return err == nil
}

// GangZone events

//export onPlayerEnterGangZone
func onPlayerEnterGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerEnterTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})

	return err == nil
}

//export onPlayerEnterPlayerGangZone
func onPlayerEnterPlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerEnterPlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})

	return err == nil
}

//export onPlayerLeaveGangZone
func onPlayerLeaveGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerLeaveTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})

	return err == nil
}

//export onPlayerLeavePlayerGangZone
func onPlayerLeavePlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerLeavePlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})

	return err == nil
}

//export onPlayerClickGangZone
func onPlayerClickGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickTurfEvent{
		Player: &Player{handle: player},
		Turf:   &Turf{handle: gangzone},
	})

	return err == nil
}

//export onPlayerClickPlayerGangZone
func onPlayerClickPlayerGangZone(player, gangzone unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickPlayerTurfEvent{
		Player: &Player{handle: player},
		Turf:   &PlayerTurf{handle: gangzone},
	})

	return err == nil
}

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(player unsafe.Pointer, menuRow uint8) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerSelectedMenuRowEvent{
		Player:  &Player{handle: player},
		MenuRow: menuRow,
	})

	return err == nil
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerExitedMenuEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

// Object events

//export onObjectMoved
func onObjectMoved(object unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ObjectMovedEvent{
		Object: &Object{handle: object},
	})

	return err == nil
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(player, object unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerObjectMovedEvent{
		Player: &Player{handle: player},
		Object: &PlayerObject{handle: object},
	})

	return err == nil
}

//export onObjectSelected
func onObjectSelected(player, object unsafe.Pointer, model int, pos C.Vector3) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ObjectSelectedEvent{
		Player: &Player{handle: player},
		Object: &Object{handle: object},
		Model:  model,
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})

	return err == nil
}

//export onPlayerObjectSelected
func onPlayerObjectSelected(player, object unsafe.Pointer, model int, pos C.Vector3) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerObjectSelectedEvent{
		Player: &Player{handle: player},
		Object: &PlayerObject{handle: object},
		Model:  model,
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})

	return err == nil
}

//export onObjectEdited
func onObjectEdited(player, object unsafe.Pointer, response int, offset, rot C.Vector3) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ObjectEditedEvent{
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

	return err == nil
}

//export onPlayerObjectEdited
func onPlayerObjectEdited(player, object unsafe.Pointer, response int, offset, rot C.Vector3) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerObjectEditedEvent{
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

	return err == nil
}

//export onPlayerAttachedObjectEdited
func onPlayerAttachedObjectEdited(player unsafe.Pointer, index int, saved bool, data C.PlayerAttachedObject) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerAttachmentEditedEvent{
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
			Color1: uint(data.colour1),
			Color2: uint(data.colour2),
		},
	})

	return err == nil
}

// // Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(player, pickup unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerPickUpPickupEvent{
		Player: &Player{handle: player},
		Pickup: &Pickup{handle: pickup},
	})

	return err == nil
}

//export onPlayerPickUpPlayerPickup
func onPlayerPickUpPlayerPickup(player, pickup unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerPickUpPlayerPickupEvent{
		Player: &Player{handle: player},
		Pickup: &PlayerPickup{handle: pickup},
	})

	return err == nil
}

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerRequestSpawnEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerSpawn
func onPlayerSpawn(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerSpawnEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(player unsafe.Pointer, ipAddress C.String, port C.ushort) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &IncomingConnectionEvent{
		Player:    &Player{handle: player},
		IPAddress: C.GoStringN(ipAddress.buf, C.int(ipAddress.length)),
		Port:      int(port),
	})

	return err == nil
}

//export onPlayerConnect
func onPlayerConnect(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerConnectEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerDisconnect
func onPlayerDisconnect(player unsafe.Pointer, reason int) C.bool {
	defer recoverPanic()

	eventPlayer := &Player{handle: player}

	err := events.Emit(context.Background(), &PlayerDisconnectEvent{
		Player: eventPlayer,
		Reason: DisconnectReason(reason),
	})

	delete(activeDialogs, eventPlayer.ID())

	return err == nil
}

//export onPlayerClientInit
func onPlayerClientInit(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClientInitEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(player, forPlayer unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerStreamInEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})

	return err == nil
}

//export onPlayerStreamOut
func onPlayerStreamOut(player, forPlayer unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerStreamOutEvent{
		Player:    &Player{handle: player},
		ForPlayer: &Player{handle: forPlayer},
	})

	return err == nil
}

// Player text events

//export onPlayerText
func onPlayerText(player unsafe.Pointer, message *C.char) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerTextEvent{
		Player:  &Player{handle: player},
		Message: C.GoString(message),
	})

	return err == nil
}

//export onPlayerCommandText
func onPlayerCommandText(player unsafe.Pointer, message C.String) C.bool {
	defer recoverPanic()

	rawVal := C.GoStringN(message.buf, C.int(message.length))

	tmp := strings.Fields(rawVal)
	name := strings.TrimPrefix(tmp[0], "/")
	args := tmp[1:]

	err := events.Emit(context.Background(), &PlayerCommandTextEvent{
		Sender:   &Player{handle: player},
		Name:     name,
		Args:     args,
		RawValue: rawVal,
	})

	return err == nil
}

// Player shot events

//export onPlayerShotMissed
func onPlayerShotMissed(player unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerShotMissedEvent{
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

	return err == nil
}

//export onPlayerShotPlayer
func onPlayerShotPlayer(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerShotPlayerEvent{
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

	return err == nil
}

//export onPlayerShotVehicle
func onPlayerShotVehicle(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerShotVehicleEvent{
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

	return err == nil
}

//export onPlayerShotObject
func onPlayerShotObject(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerShotObjectEvent{
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

	return err == nil
}

//export onPlayerShotPlayerObject
func onPlayerShotPlayerObject(player, target unsafe.Pointer, bulletData C.PlayerBulletData) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerShotPlayerObjectEvent{
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

	return err == nil
}

// Player change events

//export onPlayerScoreChange
func onPlayerScoreChange(player unsafe.Pointer, score int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerScoreChangeEvent{
		Player: &Player{handle: player},
		Score:  score,
	})

	return err == nil
}

//export onPlayerNameChange
func onPlayerNameChange(player unsafe.Pointer, oldName C.String) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerNameChangeEvent{
		Player:  &Player{handle: player},
		OldName: C.GoStringN(oldName.buf, C.int(oldName.length)),
	})

	return err == nil
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(player unsafe.Pointer, newInterior, oldInterior uint) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerInteriorChangeEvent{
		Player:      &Player{handle: player},
		NewInterior: newInterior,
		OldInterior: oldInterior,
	})

	return err == nil
}

//export onPlayerStateChange
func onPlayerStateChange(player unsafe.Pointer, newState, oldState int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerStateChangeEvent{
		Player:   &Player{handle: player},
		NewState: PlayerState(newState),
		OldState: PlayerState(oldState),
	})

	return err == nil
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(player unsafe.Pointer, newKeys, oldKeys uint) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerKeyStateChangeEvent{
		Player:  &Player{handle: player},
		NewKeys: newKeys,
		OldKeys: oldKeys,
	})

	return err == nil
}

// Player damage events

//export onPlayerDeath
func onPlayerDeath(player, killer unsafe.Pointer, reason int) C.bool {
	defer recoverPanic()

	eventKiller := &Player{handle: killer}
	if killer == nil {
		eventKiller = nil
	}

	err := events.Emit(context.Background(), &PlayerDeathEvent{
		Player: &Player{handle: player},
		Killer: eventKiller,
		Reason: reason,
	})

	return err == nil
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(player, from unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer recoverPanic()

	eventFrom := &Player{handle: from}
	if from == nil {
		eventFrom = nil
	}

	err := events.Emit(context.Background(), &PlayerTakeDamageEvent{
		Player: &Player{handle: player},
		From:   eventFrom,
		Amount: amount,
		Weapon: Weapon(weapon),
		Part:   BodyPart(part),
	})

	return err == nil
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(player, to unsafe.Pointer, amount float32, weapon uint, part int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerGiveDamageEvent{
		Player: &Player{handle: player},
		To:     &Player{handle: to},
		Amount: amount,
		Weapon: Weapon(weapon),
		Part:   BodyPart(part),
	})

	return err == nil
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(player unsafe.Pointer, pos C.Vector3) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickMapEvent{
		Player: &Player{handle: player},
		Position: Vector3{
			X: float32(pos.x),
			Y: float32(pos.y),
			Z: float32(pos.z),
		},
	})

	return err == nil
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(player, clicked unsafe.Pointer, source int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickPlayerEvent{
		Player:  &Player{handle: player},
		Clicked: &Player{handle: clicked},
		Source:  PlayerClickSource(source),
	})

	return err == nil
}

// Player check events

//export onClientCheckResponse
func onClientCheckResponse(player unsafe.Pointer, actionType, address, results int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &ClientCheckResponseEvent{
		Player:     &Player{handle: player},
		ActionType: actionType,
		Address:    address,
		Results:    results,
	})

	return err == nil
}

// Player update event

//export onPlayerUpdate
func onPlayerUpdate(player unsafe.Pointer, now C.longlong) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerUpdateEvent{
		Player: &Player{handle: player},
		Now:    time.Unix(0, int64(now)*int64(time.Millisecond)),
	})

	return err == nil
}

// Textdraw events

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(player, textdraw unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &Textdraw{handle: textdraw},
	})

	return err == nil
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(player, textdraw unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerClickPlayerTextDrawEvent{
		Player:   &Player{handle: player},
		Textdraw: &PlayerTextdraw{handle: textdraw},
	})

	return err == nil
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerCancelTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &Player{handle: player},
	})

	return err == nil
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(vehicle, player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleStreamInEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})

	return err == nil
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicle, player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{handle: vehicle},
		ForPlayer: &Player{handle: player},
	})

	return err == nil
}

//export onVehicleDeath
func onVehicleDeath(vehicle, killer unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleDeathEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Killer:  &Player{handle: killer},
	})

	return err == nil
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(player, vehicle unsafe.Pointer, isPassenger int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerEnterVehicleEvent{
		Player:      &Player{handle: player},
		Vehicle:     &Vehicle{handle: vehicle},
		IsPassenger: isPassenger != 0,
	})

	return err == nil
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(player, vehicle unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &PlayerExitVehicleEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})

	return err == nil
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicle, player unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{handle: vehicle},
		Player:  &Player{handle: player},
	})

	return err == nil
}

//export onVehiclePaintJob
func onVehiclePaintJob(player, vehicle unsafe.Pointer, paintJob int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehiclePaintJobEvent{
		Player:   &Player{handle: player},
		Vehicle:  &Vehicle{handle: vehicle},
		PaintJob: paintJob,
	})

	return err == nil
}

//export onVehicleMod
func onVehicleMod(player, vehicle unsafe.Pointer, component int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleModEvent{
		Player:    &Player{handle: player},
		Vehicle:   &Vehicle{handle: vehicle},
		Component: component,
	})

	return err == nil
}

//export onVehicleRespray
func onVehicleRespray(player, vehicle unsafe.Pointer, color1, color2 int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleResprayEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
		Color:   VehicleColor{Primary: color1, Secondary: color2},
	})

	return err == nil
}

//export onEnterExitModShop
func onEnterExitModShop(player unsafe.Pointer, enterexit bool, interiorID int) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &EnterExitModShopEvent{
		Player:     &Player{handle: player},
		EnterExit:  enterexit,
		InteriorID: interiorID,
	})

	return err == nil
}

//export onVehicleSpawn
func onVehicleSpawn(vehicle unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleSpawnEvent{
		Vehicle: &Vehicle{handle: vehicle},
	})

	return err == nil
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicle, player unsafe.Pointer, updateData C.UnoccupiedVehicleUpdate) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &UnoccupiedVehicleUpdateEvent{
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

	return err == nil
}

//export onTrailerUpdate
func onTrailerUpdate(player, vehicle unsafe.Pointer) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &TrailerUpdateEvent{
		Player:  &Player{handle: player},
		Vehicle: &Vehicle{handle: vehicle},
	})

	return err == nil
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(player, vehicle unsafe.Pointer, sirenState uint8) C.bool {
	defer recoverPanic()

	err := events.Emit(context.Background(), &VehicleSirenStateChangeEvent{
		Player:     &Player{handle: player},
		Vehicle:    &Vehicle{handle: vehicle},
		SirenState: int(sirenState),
	})

	return err == nil
}
