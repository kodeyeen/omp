package omp

// #include <stdlib.h>
// #include "include/omp.h"
// #include "include/class.h"
import "C"
import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
	"unsafe"

	"github.com/kodeyeen/event"
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

var Events = event.NewDispatcher()
var Commands = newCommandManager()

func DispatchEvent[T any](_type event.Type, data T) bool {
	defer handlePanic()

	return event.Dispatch(Events, _type, data)
}

func handlePanic() {
	if r := recover(); r != nil {
		stackTrace := strings.TrimSuffix(string(debug.Stack()), "\n")

		Log(LogLevelError, fmt.Sprint(r))
		Log(LogLevelError, stackTrace)
	}
}

//export ComponentEntryPoint
func ComponentEntryPoint() unsafe.Pointer {
	cName := C.CString("OmpGo")
	defer C.free(unsafe.Pointer(cName))

	cVer := C.struct_ComponentVersion{
		major:  0,
		minor:  0,
		patch:  0,
		prerel: 0,
	}

	return C.Component_Create(C.ulonglong(uid), cName, cVer, onReady, onReset, onFree)
}

//export onGameModeInit
func onGameModeInit() C.bool {
	defer handlePanic()

	C.loadComponent()

	result := event.Dispatch(Events, EventTypeGameModeInit, &GameModeInitEvent{})

	return C.bool(result)
}

//export onGameModeExit
func onGameModeExit() C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeGameModeExit, &GameModeExitEvent{})

	return C.bool(result)
}

// Actor events

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(args *C.struct_EventArgs_onPlayerGiveDamageActor) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerGiveDamageActor, &PlayerGiveDamageActorEvent{
		Player:   &Player{handle: *args.list.player},
		Actor:    &Player{handle: *args.list.actor},
		Amount:   float32(*args.list.amount),
		Weapon:   Weapon(*args.list.weapon),
		BodyPart: BodyPart(*args.list.part),
	})

	return C.bool(result)
}

//export onActorStreamOut
func onActorStreamOut(args *C.struct_EventArgs_onActorStreamOut) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeActorStreamOut, &ActorStreamOutEvent{
		Actor:     &Player{handle: *args.list.actor},
		ForPlayer: &Player{handle: *args.list.forPlayer},
	})

	return C.bool(result)
}

//export onActorStreamIn
func onActorStreamIn(args *C.struct_EventArgs_onActorStreamIn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeActorStreamIn, &ActorStreamInEvent{
		Actor:     &Player{handle: *args.list.actor},
		ForPlayer: &Player{handle: *args.list.forPlayer},
	})

	return C.bool(result)
}

// Checkpoint events

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(args *C.struct_EventArgs_onPlayerEnterCheckpoint) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerEnterCheckpoint, &PlayerEnterCheckpointEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(args *C.struct_EventArgs_onPlayerLeaveCheckpoint) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerLeaveCheckpoint, &PlayerLeaveCheckpointEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(args *C.struct_EventArgs_onPlayerEnterRaceCheckpoint) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerEnterRaceCheckpoint, &PlayerEnterRaceCheckpointEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(args *C.struct_EventArgs_onPlayerLeaveRaceCheckpoint) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerLeaveRaceCheckpoint, &PlayerLeaveRaceCheckpointEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

// Class events

//export onPlayerRequestClass
func onPlayerRequestClass(args *C.struct_EventArgs_onPlayerRequestClass) C.bool {
	defer handlePanic()

	class := C.Class_FromID(*args.list.classId)

	result := event.Dispatch(Events, EventTypePlayerRequestClass, &PlayerRequestClassEvent{
		Player: &Player{handle: *args.list.player},
		Class:  &Class{handle: class},
	})

	return C.bool(result)
}

// Console events. TODO

//export onConsoleText
func onConsoleText(args *C.struct_EventArgs_onConsoleText) C.bool {
	defer handlePanic()

	cmd := *args.list.command
	params := *args.list.parameters

	result := event.Dispatch(Events, EventTypeConsoleText, &ConsoleTextEvent{
		Command:    C.GoStringN(cmd.data, C.int(cmd.len)),
		Parameters: C.GoStringN(params.data, C.int(params.len)),
	})

	return C.bool(result)
}

// //export onRconLoginAttempt
// func onRconLoginAttempt(args *C.struct_EventArgs_onRconLoginAttempt) C.bool {
// 	defer handlePanic()

// 	password := *args.list.password

// 	result := event.Dispatch(Events, EventTypeRconLoginAttempt, &RconLoginAttemptEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Password: C.GoStringN(password.data, C.int(password.len)),
// 		Success:  bool(*args.list.success),
// 	})

// 	return C.bool(result)
// }

// //export onTick
// func onTick(args *C.struct_EventArgs_onTick) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeTick, &TickEvent{
// 		Elapsed: time.Duration(*args.list.elapsed) * time.Microsecond,
// 		Now:     time.Unix(int64(*args.list.now), 0),
// 	})

// 	return C.bool(result)
// }

// Custom model events

//export onPlayerFinishedDownloading
func onPlayerFinishedDownloading(args *C.struct_EventArgs_onPlayerFinishedDownloading) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerFinishedDownloading, &PlayerFinishedDownloadingEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(args *C.struct_EventArgs_onPlayerRequestDownload) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerRequestDownload, &PlayerRequestDownloadEvent{
		Player:   &Player{handle: *args.list.player},
		Type:     DownloadRequestType(*args.list._type),
		Checksum: int(*args.list.checksum),
	})

	return C.bool(result)
}

// Dialog events

//export onDialogResponse
func onDialogResponse(args *C.struct_EventArgs_onDialogResponse) C.bool {
	defer handlePanic()

	player := *args.list.player
	// dialogID := *args.list.dialogId
	response := DialogResponse(*args.list.response)
	listItem := int(*args.list.listItem)
	inputText := *args.list.inputText

	eventPlayer := &Player{handle: player}
	playerID := eventPlayer.ID()

	dialog, ok := activeDialogs[playerID]
	if !ok {
		panic("active dialog is not set")
	}
	delete(activeDialogs, playerID)

	var result bool

	switch dialog := dialog.(type) {
	case *MessageDialog:
		result = event.Dispatch(dialog.Dispatcher, EventTypeDialogResponse, &MessageDialogResponseEvent{
			Player:   eventPlayer,
			Response: response,
		})

		event.Dispatch(dialog.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
			Player: eventPlayer,
		})
	case *InputDialog:
		result = event.Dispatch(dialog.Dispatcher, EventTypeDialogResponse, &InputDialogResponseEvent{
			Player:    eventPlayer,
			Response:  response,
			InputText: C.GoStringN(inputText.data, C.int(inputText.len)),
		})

		event.Dispatch(dialog.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
			Player: eventPlayer,
		})
	case *ListDialog:
		result = event.Dispatch(dialog.Dispatcher, EventTypeDialogResponse, &ListDialogResponseEvent{
			Player:     eventPlayer,
			Response:   response,
			ItemNumber: listItem,
			Item:       C.GoStringN(inputText.data, C.int(inputText.len)),
		})

		event.Dispatch(dialog.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
			Player: eventPlayer,
		})
	case *TabListDialog:
		result = event.Dispatch(dialog.Dispatcher, EventTypeDialogResponse, &TabListDialogResponseEvent{
			Player:     eventPlayer,
			Response:   response,
			ItemNumber: listItem,
			Item:       dialog.items[listItem],
		})

		event.Dispatch(dialog.Dispatcher, EventTypeDialogHide, &DialogHideEvent{
			Player: eventPlayer,
		})
	default:
		panic("unknown dialog type")
	}

	return C.bool(result)
}

// GangZone events

//export onPlayerEnterGangZone
func onPlayerEnterGangZone(args *C.struct_EventArgs_onPlayerEnterGangZone) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerEnterTurf, &PlayerEnterTurfEvent{
		Player: &Player{handle: *args.list.player},
		Turf:   &Turf{handle: *args.list.zone},
	})

	return C.bool(result)
}

// //export onPlayerEnterPlayerGangZone
// func onPlayerEnterPlayerGangZone(args *C.struct_EventArgs_onPlayerEnterPlayerGangZone) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerEnterPlayerTurf, &PlayerEnterPlayerTurfEvent{
// 		Player: eventPlayer,
// 		Turf: &PlayerTurf{
// 			handle: *args.list.zone,
// 			player: eventPlayer,
// 		},
// 	})

// 	return C.bool(result)
// }

//export onPlayerLeaveGangZone
func onPlayerLeaveGangZone(args *C.struct_EventArgs_onPlayerLeaveGangZone) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerLeaveTurf, &PlayerLeaveTurfEvent{
		Player: &Player{handle: *args.list.player},
		Turf:   &Turf{handle: *args.list.zone},
	})

	return C.bool(result)
}

// //export onPlayerLeavePlayerGangZone
// func onPlayerLeavePlayerGangZone(args *C.struct_EventArgs_onPlayerLeavePlayerGangZone) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{
// 		handle: *args.list.player,
// 	}

// 	result := event.Dispatch(Events, EventTypePlayerLeavePlayerTurf, &PlayerLeavePlayerTurfEvent{
// 		Player: eventPlayer,
// 		Turf: &PlayerTurf{
// 			handle: *args.list.zone,
// 			player: eventPlayer,
// 		},
// 	})

// 	return C.bool(result)
// }

//export onPlayerClickGangZone
func onPlayerClickGangZone(args *C.struct_EventArgs_onPlayerClickGangZone) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerClickTurf, &PlayerClickTurfEvent{
		Player: &Player{handle: *args.list.player},
		Turf:   &Turf{handle: *args.list.zone},
	})

	return C.bool(result)
}

// //export onPlayerClickPlayerGangZone
// func onPlayerClickPlayerGangZone(args *C.struct_EventArgs_onPlayerClickPlayerGangZone) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerClickPlayerTurf, &PlayerClickPlayerTurfEvent{
// 		Player: eventPlayer,
// 		Turf: &PlayerTurf{
// 			handle: *args.list.zone,
// 			player: eventPlayer,
// 		},
// 	})

// 	return C.bool(result)
// }

// Menu events

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(args *C.struct_EventArgs_onPlayerSelectedMenuRow) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerSelectedMenuRow, &PlayerSelectedMenuRowEvent{
		Player:  &Player{handle: *args.list.player},
		MenuRow: int(*args.list.row),
	})

	return C.bool(result)
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(args *C.struct_EventArgs_onPlayerExitedMenu) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerExitedMenu, &PlayerExitedMenuEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

// Object events

// //export onObjectMoved
// func onObjectMoved(args *C.struct_EventArgs_onObjectMoved) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeObjectMoved, &ObjectMovedEvent{
// 		Object: &Object{handle: *args.list.object},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerObjectMoved
// func onPlayerObjectMoved(args *C.struct_EventArgs_onPlayerObjectMoved) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerObjectMoved, &PlayerObjectMovedEvent{
// 		Player: eventPlayer,
// 		Object: &PlayerObject{
// 			handle: *args.list.object,
// 			player: eventPlayer,
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onObjectEdited
// func onObjectEdited(args *C.struct_EventArgs_onObjectEdited) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeObjectEdited, &ObjectEditedEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Object: &Object{handle: *args.list.object},
// 		Response: ObjectEditResponse(*args.list.response),
// 		Offset: Vector3{
// 			X: float32(*args.list.offsetX),
// 			Y: float32(*args.list.offsetY),
// 			Z: float32(*args.list.offsetZ),
// 		},
// 		Rotation: Vector3{
// 			X: float32(*args.list.rotationX),
// 			Y: float32(*args.list.rotationY),
// 			Z: float32(*args.list.rotationZ),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerObjectEdited
// func onPlayerObjectEdited(args *C.struct_EventArgs_onPlayerObjectEdited) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerObjectEdited, &PlayerObjectEditedEvent{
// 		Player: eventPlayer,
// 		Object: &PlayerObject{
// 			handle: *args.list.object,
// 			player: eventPlayer,
// 		},
// 		Response: ObjectEditResponse(*args.list.response),
// 		Offset: Vector3{
// 			X: float32(*args.list.offsetX),
// 			Y: float32(*args.list.offsetY),
// 			Z: float32(*args.list.offsetZ),
// 		},
// 		Rotation: Vector3{
// 			X: float32(*args.list.rotationX),
// 			Y: float32(*args.list.rotationY),
// 			Z: float32(*args.list.rotationZ),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerAttachedObjectEdited
// func onPlayerAttachedObjectEdited(args *C.struct_EventArgs_onPlayerAttachedObjectEdited) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerAttachmentEdited, &PlayerAttachmentEdited{
// 		Player: &Player{handle: *args.list.player},
// 		Index: int(*args.list.index),
// 		Saved: bool(*args.list.saved),
// 		Attachment: PlayerAttachment{
// 			ModelID: int(*args.list.model),
// 			Bone:    PlayerBone(*args.list.bone),
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Rotation: Vector3{
// 				X: float32(*args.list.rotationX),
// 				Y: float32(*args.list.rotationY),
// 				Z: float32(*args.list.rotationZ),
// 			},
// 			Scale: Vector3{
// 				X: float32(*args.list.scaleX),
// 				Y: float32(*args.list.scaleY),
// 				Z: float32(*args.list.scaleZ),
// 			},
// 			Color1: Color(*args.list.color1),
// 			Color2: Color(*args.list.color2),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onObjectSelected
// func onObjectSelected(args *C.struct_EventArgs_onObjectSelected) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeObjectSelected, &ObjectSelectedEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Object: &Object{handle: *args.list.object},
// 		Model: int(*args.list.model),
// 		Position: Vector3{
// 			X: float32(*args.list.x),
// 			Y: float32(*args.list.y),
// 			Z: float32(*args.list.z),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerObjectSelected
// func onPlayerObjectSelected(args *C.struct_EventArgs_onPlayerObjectSelected) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerObjectSelected, &PlayerObjectSelectedEvent{
// 		Player: eventPlayer,
// 		Object: &PlayerObject{
// 			handle: *args.list.object,
// 			player: eventPlayer,
// 		},
// 		Model: int(*args.list.model),
// 		Position: Vector3{
// 			X: float32(*args.list.x),
// 			Y: float32(*args.list.y),
// 			Z: float32(*args.list.z),
// 		},
// 	})

// 	return C.bool(result)
// }

// Pickup events

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(args *C.struct_EventArgs_onPlayerPickUpPickup) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerPickUpPickup, &PlayerPickUpPickupEvent{
		Player: &Player{handle: *args.list.player},
		Pickup: &Pickup{handle: *args.list.pickup},
	})

	return C.bool(result)
}

// //export onPlayerPickUpPlayerPickup
// func onPlayerPickUpPlayerPickup(args *C.struct_EventArgs_onPlayerPickUpPlayerPickup) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerPickUpPlayerPickup, &PlayerPickUpPlayerPickupEvent{
// 		Player: eventPlayer,
// 		Pickup: &PlayerPickup{
// 			handle: *args.list.pickup,
// 			player: eventPlayer,
// 		},
// 	})

// 	return C.bool(result)
// }

// Player spawn events

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(args *C.struct_EventArgs_onPlayerRequestSpawn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerRequestSpawn, &PlayerRequestSpawnEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerSpawn
func onPlayerSpawn(args *C.struct_EventArgs_onPlayerSpawn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerSpawn, &PlayerSpawnEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

// Player connect events

//export onIncomingConnection
func onIncomingConnection(args *C.struct_EventArgs_onIncomingConnection) C.bool {
	defer handlePanic()

	ipAddress := *args.list.ipAddress

	result := event.Dispatch(Events, EventTypeIncomingConnection, &IncomingConnectionEvent{
		Player:    &Player{handle: *args.list.player},
		ipAddress: C.GoStringN(ipAddress.data, C.int(ipAddress.len)),
		Port:      int(*args.list.port),
	})

	return C.bool(result)
}

//export onPlayerConnect
func onPlayerConnect(args *C.struct_EventArgs_onPlayerConnect) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerConnect, &PlayerConnectEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerDisconnect
func onPlayerDisconnect(args *C.struct_EventArgs_onPlayerDisconnect) C.bool {
	defer handlePanic()

	eventPlayer := &Player{handle: *args.list.player}

	result := event.Dispatch(Events, EventTypePlayerDisconnect, &PlayerDisconnectEvent{
		Player: eventPlayer,
		Reason: DisconnectReason(*args.list.reason),
	})

	delete(activeDialogs, eventPlayer.ID())

	return C.bool(result)
}

// //export onPlayerClientInit
// func onPlayerClientInit(args *C.struct_EventArgs_onPlayerClientInit) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerClientInit, &PlayerClientInitEvent{
// 		Player: &Player{handle: *args.list.player},
// 	})

// 	return C.bool(result)
// }

// Player stream events

//export onPlayerStreamIn
func onPlayerStreamIn(args *C.struct_EventArgs_onPlayerStreamIn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerStreamIn, &PlayerStreamInEvent{
		Player:    &Player{handle: *args.list.player},
		ForPlayer: &Player{handle: *args.list.forPlayer},
	})

	return C.bool(result)
}

//export onPlayerStreamOut
func onPlayerStreamOut(args *C.struct_EventArgs_onPlayerStreamOut) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerStreamOut, &PlayerStreamOutEvent{
		Player:    &Player{handle: *args.list.player},
		ForPlayer: &Player{handle: *args.list.forPlayer},
	})

	return C.bool(result)
}

// Player text events

//export onPlayerText
func onPlayerText(args *C.struct_EventArgs_onPlayerText) C.bool {
	defer handlePanic()

	text := *args.list.text

	result := event.Dispatch(Events, EventTypePlayerText, &PlayerTextEvent{
		Player:  &Player{handle: *args.list.player},
		Message: C.GoStringN(text.data, C.int(text.len)),
	})

	return C.bool(result)
}

//export onPlayerCommandText
func onPlayerCommandText(args *C.struct_EventArgs_onPlayerCommandText) C.bool {
	defer handlePanic()

	player := *args.list.player
	message := *args.list.command

	rawCmd := C.GoStringN(message.data, C.int(message.len))

	tmp := strings.Fields(rawCmd)
	cmdName := strings.TrimPrefix(tmp[0], "/")
	cmdArgs := tmp[1:]

	exists := Commands.Has(cmdName)
	if !exists {
		return false
	}

	Commands.run(cmdName, &Command{
		Sender:   &Player{handle: player},
		Name:     cmdName,
		Args:     cmdArgs,
		RawValue: rawCmd,
	})

	return true
}

// Player shot events

// //export onPlayerShotMissed
// func onPlayerShotMissed(args *C.struct_EventArgs_onPlayerShotMissed) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerShotMissed, &PlayerShotMissedEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Bullet: PlayerBullet{
// 			Origin: Vector3{
// 				X: float32(*args.list.originX),
// 				Y: float32(*args.list.originY),
// 				Z: float32(*args.list.originZ),
// 			},
// 			HitPosition: Vector3{
// 				X: float32(*args.list.hitPosX),
// 				Y: float32(*args.list.hitPosY),
// 				Z: float32(*args.list.hitPosZ),
// 			},
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Weapon: Weapon(*args.list.weapon),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerShotPlayer
// func onPlayerShotPlayer(args *C.struct_EventArgs_onPlayerShotPlayer) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerShotPlayer, &PlayerShotPlayerEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Target: &Player{handle: *args.list.target},
// 		Bullet: PlayerBullet{
// 			Origin: Vector3{
// 				X: float32(*args.list.originX),
// 				Y: float32(*args.list.originY),
// 				Z: float32(*args.list.originZ),
// 			},
// 			HitPosition: Vector3{
// 				X: float32(*args.list.hitPosX),
// 				Y: float32(*args.list.hitPosY),
// 				Z: float32(*args.list.hitPosZ),
// 			},
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Weapon: Weapon(*args.list.weapon),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerShotVehicle
// func onPlayerShotVehicle(args *C.struct_EventArgs_onPlayerShotVehicle) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerShotVehicle, &PlayerShotVehicleEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Target: &Vehicle{handle: *args.list.target},
// 		Bullet: PlayerBullet{
// 			Origin: Vector3{
// 				X: float32(*args.list.originX),
// 				Y: float32(*args.list.originY),
// 				Z: float32(*args.list.originZ),
// 			},
// 			HitPosition: Vector3{
// 				X: float32(*args.list.hitPosX),
// 				Y: float32(*args.list.hitPosY),
// 				Z: float32(*args.list.hitPosZ),
// 			},
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Weapon: Weapon(*args.list.weapon),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerShotObject
// func onPlayerShotObject(args *C.struct_EventArgs_onPlayerShotObject) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerShotObject, &PlayerShotObjectEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Target: &Object{handle: *args.list.target},
// 		Bullet: PlayerBullet{
// 			Origin: Vector3{
// 				X: float32(*args.list.originX),
// 				Y: float32(*args.list.originY),
// 				Z: float32(*args.list.originZ),
// 			},
// 			HitPosition: Vector3{
// 				X: float32(*args.list.hitPosX),
// 				Y: float32(*args.list.hitPosY),
// 				Z: float32(*args.list.hitPosZ),
// 			},
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Weapon: Weapon(*args.list.weapon),
// 		},
// 	})

// 	return C.bool(result)
// }

// //export onPlayerShotPlayerObject
// func onPlayerShotPlayerObject(args *C.struct_EventArgs_onPlayerShotPlayerObject) C.bool {
// 	defer handlePanic()

// 	eventPlayer := &Player{handle: *args.list.player}

// 	result := event.Dispatch(Events, EventTypePlayerShotPlayerObject, &PlayerShotPlayerObjectEvent{
// 		Player: eventPlayer,
// 		Target: &PlayerObject{
// 			handle: *args.list.target,
// 			player: eventPlayer,
// 		},
// 		Bullet: PlayerBullet{
// 			Origin: Vector3{
// 				X: float32(*args.list.originX),
// 				Y: float32(*args.list.originY),
// 				Z: float32(*args.list.originZ),
// 			},
// 			HitPosition: Vector3{
// 				X: float32(*args.list.hitPosX),
// 				Y: float32(*args.list.hitPosY),
// 				Z: float32(*args.list.hitPosZ),
// 			},
// 			Offset: Vector3{
// 				X: float32(*args.list.offsetX),
// 				Y: float32(*args.list.offsetY),
// 				Z: float32(*args.list.offsetZ),
// 			},
// 			Weapon: Weapon(*args.list.weapon),
// 		},
// 	})

// 	return C.bool(result)
// }

// Player change events

// //export onPlayerScoreChange
// func onPlayerScoreChange(args *C.struct_EventArgs_onPlayerScoreChange) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerScoreChange, &PlayerScoreChangeEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Score: int(*args.list.score),
// 	})

// 	return C.bool(result)
// }

// //export onPlayerNameChange
// func onPlayerNameChange(args *C.struct_EventArgs_onPlayerNameChange) C.bool {
// 	defer handlePanic()

// 	oldName := *args.list.oldName

// 	result := event.Dispatch(Events, EventTypePlayerNameChange, &PlayerNameChangeEvent{
// 		Player: &Player{handle: *args.list.player},
// 		OldName: C.GoStringN(oldName.data, C.int(oldName.len)),
// 	})

// 	return C.bool(result)
// }

//export onPlayerInteriorChange
func onPlayerInteriorChange(args *C.struct_EventArgs_onPlayerInteriorChange) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerInteriorChange, &PlayerInteriorChangeEvent{
		Player:      &Player{handle: *args.list.player},
		NewInterior: int(*args.list.newInterior),
		OldInterior: int(*args.list.oldInterior),
	})

	return C.bool(result)
}

//export onPlayerStateChange
func onPlayerStateChange(args *C.struct_EventArgs_onPlayerStateChange) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerStateChange, &PlayerStateChangeEvent{
		Player:   &Player{handle: *args.list.player},
		NewState: PlayerState(*args.list.newState),
		OldState: PlayerState(*args.list.oldState),
	})

	return C.bool(result)
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(args *C.struct_EventArgs_onPlayerKeyStateChange) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerKeyStateChange, &PlayerKeyStateChangeEvent{
		Player:  &Player{handle: *args.list.player},
		NewKeys: int(*args.list.newKeys),
		OldKeys: int(*args.list.oldKeys),
	})

	return C.bool(result)
}

// Player damage events

//export onPlayerDeath
func onPlayerDeath(args *C.struct_EventArgs_onPlayerDeath) C.bool {
	defer handlePanic()

	player := *args.list.player
	killer := *args.list.killer
	reason := int(*args.list.reason)

	eventKiller := &Player{handle: killer}

	if killer == nil {
		eventKiller = nil
	}

	result := event.Dispatch(Events, EventTypePlayerDeath, &PlayerDeathEvent{
		Player: &Player{handle: player},
		Killer: eventKiller,
		Reason: reason,
	})

	return C.bool(result)
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(args *C.struct_EventArgs_onPlayerTakeDamage) C.bool {
	defer handlePanic()

	player := *args.list.player
	from := *args.list.from
	amount := float32(*args.list.amount)
	weapon := Weapon(*args.list.weapon)
	bodyPart := BodyPart(*args.list.bodypart)

	eventFrom := &Player{handle: from}
	if from == nil {
		eventFrom = nil
	}

	result := event.Dispatch(Events, EventTypePlayerTakeDamage, &PlayerTakeDamageEvent{
		Player: &Player{handle: player},
		From:   eventFrom,
		Amount: amount,
		Weapon: weapon,
		Part:   bodyPart,
	})

	return C.bool(result)
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(args *C.struct_EventArgs_onPlayerGiveDamage) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerGiveDamage, &PlayerGiveDamageEvent{
		Player:   &Player{handle: *args.list.player},
		To:       &Player{handle: *args.list.to},
		Amount:   float32(*args.list.amount),
		Weapon:   Weapon(*args.list.weapon),
		BodyPart: BodyPart(*args.list.bodypart),
	})

	return C.bool(result)
}

// Player click events

//export onPlayerClickMap
func onPlayerClickMap(args *C.struct_EventArgs_onPlayerClickMap) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerClickMap, &PlayerClickMapEvent{
		Player: &Player{handle: *args.list.player},
		Position: Vector3{
			X: float32(*args.list.x),
			Y: float32(*args.list.y),
			Z: float32(*args.list.z),
		},
	})

	return C.bool(result)
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(args *C.struct_EventArgs_onPlayerClickPlayer) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerClickPlayer, &PlayerClickPlayerEvent{
		Player: &Player{handle: *args.list.player},
		Target: &Player{handle: *args.list.clicked},
		Source: PlayerClickSource(*args.list.source),
	})

	return C.bool(result)
}

// Player check events

// //export onClientCheckResponse
// func onClientCheckResponse(args *C.struct_EventArgs_onClientCheckResponse) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeClientCheckResponse, &ClientCheckResponseEvent{
// 		Player: &Player{handle: *args.list.player},
// 		ActionType: int(*args.list.actionType),
// 		Address:    int(*args.list.address),
// 		Results:    int(*args.list.result),
// 	})

// 	return C.bool(result)
// }

// Player update event

// //export onPlayerUpdate
// func onPlayerUpdate(args *C.struct_EventArgs_onPlayerUpdate) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypePlayerUpdate, &PlayerUpdateEvent{
// 		Player: &Player{handle: *args.list.player},
// 		Now: time.Unix(int64(*args.list.now), 0),
// 	})

// 	return C.bool(result)
// }

// Textdraw events

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(args *C.struct_EventArgs_onPlayerClickTextDraw) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerClickTextDraw, &PlayerClickTextDrawEvent{
		Player:   &Player{handle: *args.list.player},
		Textdraw: &Textdraw{handle: *args.list.textdraw},
	})

	return C.bool(result)
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(args *C.struct_EventArgs_onPlayerClickPlayerTextDraw) C.bool {
	defer handlePanic()

	eventPlayer := &Player{handle: *args.list.player}

	result := event.Dispatch(Events, EventTypePlayerClickPlayerTextDraw, &PlayerClickPlayerTextDrawEvent{
		Player: eventPlayer,
		Textdraw: &PlayerTextdraw{
			handle: *args.list.textdraw,
			player: eventPlayer,
		},
	})

	return C.bool(result)
}

//export onPlayerCancelTextDrawSelection
func onPlayerCancelTextDrawSelection(args *C.struct_EventArgs_onPlayerCancelTextDrawSelection) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerCancelTextDrawSelection, &PlayerCancelTextDrawSelectionEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerCancelPlayerTextDrawSelection
func onPlayerCancelPlayerTextDrawSelection(args *C.struct_EventArgs_onPlayerCancelPlayerTextDrawSelection) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerCancelPlayerTextDrawSelection, &PlayerCancelPlayerTextDrawSelectionEvent{
		Player: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

// Vehicle events

//export onVehicleStreamIn
func onVehicleStreamIn(args *C.struct_EventArgs_onVehicleStreamIn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleStreamIn, &VehicleStreamInEvent{
		Vehicle:   &Vehicle{handle: *args.list.vehicle},
		ForPlayer: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onVehicleStreamOut
func onVehicleStreamOut(args *C.struct_EventArgs_onVehicleStreamOut) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleStreamOut, &VehicleStreamOutEvent{
		Vehicle:   &Vehicle{handle: *args.list.vehicle},
		ForPlayer: &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onVehicleDeath
func onVehicleDeath(args *C.struct_EventArgs_onVehicleDeath) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleDeath, &VehicleDeathEvent{
		Vehicle: &Vehicle{handle: *args.list.vehicle},
		Killer:  &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(args *C.struct_EventArgs_onPlayerEnterVehicle) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerEnterVehicle, &PlayerEnterVehicleEvent{
		Player:      &Player{handle: *args.list.player},
		Vehicle:     &Vehicle{handle: *args.list.vehicle},
		IsPassenger: bool(*args.list.passenger),
	})

	return C.bool(result)
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(args *C.struct_EventArgs_onPlayerExitVehicle) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypePlayerExitVehicle, &PlayerExitVehicleEvent{
		Player:  &Player{handle: *args.list.player},
		Vehicle: &Vehicle{handle: *args.list.vehicle},
	})

	return C.bool(result)
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(args *C.struct_EventArgs_onVehicleDamageStatusUpdate) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleDamageStatusUpdate, &VehicleDamageStatusUpdateEvent{
		Vehicle: &Vehicle{handle: *args.list.vehicle},
		Player:  &Player{handle: *args.list.player},
	})

	return C.bool(result)
}

//export onVehiclePaintJob
func onVehiclePaintJob(args *C.struct_EventArgs_onVehiclePaintJob) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehiclePaintJob, &VehiclePaintJobEvent{
		Player:   &Player{handle: *args.list.player},
		Vehicle:  &Vehicle{handle: *args.list.vehicle},
		PaintJob: int(*args.list.paintJob),
	})

	return C.bool(result)
}

//export onVehicleMod
func onVehicleMod(args *C.struct_EventArgs_onVehicleMod) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleMod, &VehicleModEvent{
		Player:    &Player{handle: *args.list.player},
		Vehicle:   &Vehicle{handle: *args.list.vehicle},
		Component: int(*args.list.component),
	})

	return C.bool(result)
}

//export onVehicleRespray
func onVehicleRespray(args *C.struct_EventArgs_onVehicleRespray) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleRespray, &VehicleResprayEvent{
		Player:  &Player{handle: *args.list.player},
		Vehicle: &Vehicle{handle: *args.list.vehicle},
		Color: VehicleColor{
			Primary:   int(*args.list.color1),
			Secondary: int(*args.list.color2),
		},
	})

	return C.bool(result)
}

// //export onEnterExitModShop
// func onEnterExitModShop(args *C.struct_EventArgs_onEnterExitModShop) C.bool {
// 	defer handlePanic()

// 	result := event.Dispatch(Events, EventTypeEnterExitModShop, &EnterExitModShopEvent{
// 		Player: &Player{handle: *args.list.player},
// 		EnterExit:  bool(*args.list.enterexit),
// 		InteriorID: int(*args.list.interiorId),
// 	})

// 	return C.bool(result)
// }

//export onVehicleSpawn
func onVehicleSpawn(args *C.struct_EventArgs_onVehicleSpawn) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleSpawn, &VehicleSpawnEvent{
		Vehicle: &Vehicle{handle: *args.list.vehicle},
	})

	return C.bool(result)
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(args *C.struct_EventArgs_onUnoccupiedVehicleUpdate) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeUnoccupiedVehicleUpdate, &UnoccupiedVehicleUpdateEvent{
		Vehicle: &Vehicle{handle: *args.list.vehicle},
		Player:  &Player{handle: *args.list.player},
		Update: UnoccupiedVehicleUpdate{
			Seat: int(*args.list.seat),
			Position: Vector3{
				X: float32(*args.list.posX),
				Y: float32(*args.list.posY),
				Z: float32(*args.list.posZ),
			},
			Velocity: Vector3{
				X: float32(*args.list.velocityX),
				Y: float32(*args.list.velocityY),
				Z: float32(*args.list.velocityZ),
			},
		},
	})

	return C.bool(result)
}

//export onTrailerUpdate
func onTrailerUpdate(args *C.struct_EventArgs_onTrailerUpdate) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeTrailerUpdate, &TrailerUpdateEvent{
		Player:  &Player{handle: *args.list.player},
		Trailer: &Vehicle{handle: *args.list.trailer},
	})

	return C.bool(result)
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(args *C.struct_EventArgs_onVehicleSirenStateChange) C.bool {
	defer handlePanic()

	result := event.Dispatch(Events, EventTypeVehicleSirenStateChange, &VehicleSirenStateChangeEvent{
		Player:     &Player{handle: *args.list.player},
		Vehicle:    &Vehicle{handle: *args.list.vehicle},
		SirenState: int(*args.list.sirenState),
	})

	return C.bool(result)
}
