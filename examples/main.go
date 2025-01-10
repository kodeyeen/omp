package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/kodeyeen/omp"
)

// GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o test.dll

func init() {
	omp.ListenFunc(omp.EventTypeGameModeInit, func(ctx context.Context, e omp.Event) error {
		omp.Println("GAME MODE INIT 2")

		return nil
	})
}

func init() {
	omp.ListenFunc(omp.EventTypeGameModeInit, func(ctx context.Context, e omp.Event) error {
		omp.Println("GAME MODE INIT")

		omp.EnableManualEngineAndLights()
		return nil
	})

	omp.ListenFunc(omp.EventTypePlayerConnect, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.PlayerConnectEvent)
		player := ep.Player

		player.SendClientMessage(fmt.Sprintf("Hello, %s", player.Name()), 0x00FF0000)

		omp.NewVehicle(omp.VehicleModelAlpha, omp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		omp.NewPickup(1273, 1, 0, omp.Vector3{X: 3.9905, Y: 7.0804, Z: 3.1096})

		return nil
	})

	omp.ListenFunc(omp.EventTypePlayerSpawn, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.PlayerSpawnEvent)
		player := ep.Player

		player.GiveWeapon(omp.WeaponDeagle, 100)

		return nil
	})

	omp.ListenFunc(omp.EventTypeConsoleText, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.ConsoleTextEvent)
		fmt.Println("onConsoleText", ep.Command, ep.Parameters)
		return nil
	})

	omp.ListenFunc(omp.EventTypePlayerCommandText, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.PlayerCommandTextEvent)

		switch ep.Name {
		case "rcontest":
			fmt.Println(ep)

			omp.SendRCONCommand("weather 20")
		case "msgdlg":
			dialog := omp.NewMessageDialog("Message Dialog", "Message", "Ok", "Cancel")

			// dialog.On(omp.EventTypeDialogHide, func(e *omp.DialogHideEvent) bool {
			// 	e.Player.SendClientMessage("Dialog is hiding", 0x00FFFFFF)
			// 	return true
			// })

			dialog.ShowFor(ep.Sender)
		case "inputdlg":
			dialog := omp.NewInputDialog("Input Dialog", "Enter something:", "Ok", "Cancel")

			// dialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
			// 	if e.Response == omp.DialogResponseLeft {
			// 		e.Player.SendClientMessage(fmt.Sprintf("Left button. Your input is %s", e.InputText), 0xFF0FFFFF)
			// 	} else if e.Response == omp.DialogResponseRight {
			// 		e.Player.SendClientMessage(fmt.Sprintf("Right button. Your input is %s", e.InputText), 0xFF0FFFFF)
			// 	}

			// 	return true
			// })

			dialog.ShowFor(ep.Sender)
		case "listdlg":
			dialog := omp.NewListDialog("List Dialog", "Ok", "Cancel")

			dialog.SetItems([]string{"Item 0", "Item 1", "Item 2"})

			dialog.Add("Item 3", "Item 4")

			dialog.ShowFor(ep.Sender)
		case "pwddlg":
			dialog := omp.NewPasswordDialog("Password Dialog", "Enter password:", "Ok", "Cancel")
			dialog.ShowFor(ep.Sender)
		case "hidedlg":
			time.AfterFunc(5*time.Second, func() {
				dialog, err := ep.Sender.Dialog()
				if err != nil {
					ep.Sender.SendClientMessage("You have no active dialog", 0xFF00FFFF)
					return
				}

				ep.Sender.SendClientMessage("timer worked", 0xFF00FFFF)
				dialog.HideFor(ep.Sender)
			})
		case "tablistdlg":
			dialog := omp.NewTabListDialog("Weapons Dialog", "Ok", "Cancel")

			dialog.SetItems([]omp.TabListItem{
				{"Deagle", "$5000", "100"},
				{"Sawnoff", "$5000", "100"},
				{"Pistol", "$1000", "50"},
			})

			dialog.Add(
				omp.TabListItem{"M4", "$15000", "150"},
				omp.TabListItem{"AK47", "$12000", "150"},
			)

			// dialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
			// 	e.Player.SendClientMessage(fmt.Sprintf("Response: %d, itemno: %d, item: %+v", e.Response, e.ItemNumber, e.Item), 0xFFFF00FF)
			// 	return true
			// })

			// dialog.On(omp.EventTypeDialogShow, func(e *omp.DialogShowEvent) bool {
			// 	e.Player.SendClientMessage("Dialog shown", 0xFFFF00FF)
			// 	return true
			// })

			// dialog.On(omp.EventTypeDialogHide, func(e *omp.DialogHideEvent) bool {
			// 	e.Player.SendClientMessage("Dialog hidden", 0xFFFF00FF)
			// 	return true
			// })

			dialog.ShowFor(ep.Sender)
		case "tablistheadersdlg":
			dialog := omp.NewTabListDialog("Weapons Dialog", "Ok", "Cancel")

			dialog.SetHeader(omp.TabListItem{"Weapon", "Price", "Ammo"})

			dialog.Add(omp.TabListItem{"Deagle", "$5000", "100"})
			dialog.Add(omp.TabListItem{"Sawnoff", "$5000", "100"})
			dialog.Add(omp.TabListItem{"Pistol", "$1000", "50"})

			dialog.ShowFor(ep.Sender)

			// dialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
			// 	e.Player.SendClientMessage("Dialog response triggered", 0xFFFF00FF)

			// 	return true
			// })
		case "getpos":
			ep.Sender.SendClientMessage(fmt.Sprintf("Your position is %+v", ep.Sender.Position()), 0xFFFFFFFF)
		case "createveh":
			if len(ep.Args) != 1 {
				ep.Sender.SendClientMessage("Invalid command syntax", 0xFFFFFFFF)
				return errors.New("invalid command syntax")
			}

			modelID, err := strconv.Atoi(ep.Args[0])
			if err != nil {
				ep.Sender.SendClientMessage("Invalid command syntax", 0xFFFFFFFF)
				return errors.New("invalid command syntax")
			}

			omp.NewVehicle(omp.VehicleModel(modelID), ep.Sender.Position(), 0.0)
		case "setname":
			status := ep.Sender.SetName("кириллица")
			ep.Sender.SendClientMessage(fmt.Sprintf("You changed %d your name to %s", status, ep.Sender.Name()), 0xFFFFFFFF)
		case "doors":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.AreDoorsLocked() {
				plrVeh.UnlockDoors()
				ep.Sender.SendClientMessage("Doors unlocked", 0xFFFFFFFF)
			} else {
				plrVeh.LockDoors()
				ep.Sender.SendClientMessage("Doors locked", 0xFFFFFFFF)
			}
		case "hood":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.IsHoodOpen() {
				plrVeh.CloseHood()
			} else {
				plrVeh.OpenHood()
			}
		case "trunk":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.IsTrunkOpen() {
				plrVeh.CloseTrunk()
			} else {
				plrVeh.OpenTrunk()
			}
		case "lights":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.AreLightsTurnedOn() {
				plrVeh.TurnOffLights()
			} else {
				plrVeh.TurnOnLights()
			}
		case "engine":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.IsEngineStarted() {
				plrVeh.StopEngine()
			} else {
				plrVeh.StartEngine()
			}
		case "alarm":
			plrVeh, err := ep.Sender.Vehicle()
			if err != nil {
				return errors.New("player is not in a vehicle")
			}

			if plrVeh.IsAlarmTurnedOn() {
				plrVeh.TurnOffAlarm()
			} else {
				plrVeh.TurnOnAlarm()
			}
		default:
			return errors.New("unknown command")
		}

		return nil
	})
}

func main() {}
