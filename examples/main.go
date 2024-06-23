package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kodeyeen/omp"
)

// GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o test.dll

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, func(e *omp.GameModeInitEvent) bool {
		omp.Println("GAME MODE INIT 2")

		return true
	})
}

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, func(e *omp.GameModeInitEvent) bool {
		omp.Println("GAME MODE INIT")

		omp.EnableManualEngineAndLights()
		return true
	})

	omp.Events.Listen(omp.EventTypePlayerConnect, func(e *omp.PlayerConnectEvent) bool {
		player := e.Player

		player.SendClientMessage(fmt.Sprintf("Hello, %s", player.Name()), 0x00FF0000)

		omp.NewVehicle(omp.VehicleModelAlpha, omp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		omp.NewPickup(1273, 1, 0, omp.Vector3{X: 3.9905, Y: 7.0804, Z: 3.1096})

		return true
	})

	omp.Events.Listen(omp.EventTypePlayerSpawn, func(e *omp.PlayerSpawnEvent) bool {
		player := e.Player

		player.GiveWeapon(omp.WeaponDeagle, 100)

		return true
	})

	omp.Events.Listen(omp.EventTypeConsoleText, func(e *omp.ConsoleTextEvent) bool {
		fmt.Println("onConsoleText", e.Command, e.Parameters)
		return false
	})

	omp.Commands.Add("rcontest", func(cmd *omp.Command) {
		fmt.Println(cmd)

		omp.SendRCONCommand("weather 20")
	})

	omp.Commands.Add("msgdlg", func(cmd *omp.Command) {
		dialog := omp.NewMessageDialog("Message Dialog", "Message", "Ok", "Cancel")

		dialog.On(omp.EventTypeDialogHide, func(e *omp.DialogHideEvent) bool {
			e.Player.SendClientMessage("Dialog is hiding", 0x00FFFFFF)
			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.Commands.Add("inputdlg", func(cmd *omp.Command) {
		dialog := omp.NewInputDialog("Input Dialog", "Enter something:", "Ok", "Cancel")

		dialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
			if e.Response == omp.DialogResponseLeft {
				e.Player.SendClientMessage(fmt.Sprintf("Left button. Your input is %s", e.InputText), 0xFF0FFFFF)
			} else if e.Response == omp.DialogResponseRight {
				e.Player.SendClientMessage(fmt.Sprintf("Right button. Your input is %s", e.InputText), 0xFF0FFFFF)
			}

			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.Commands.Add("listdlg", func(cmd *omp.Command) {
		dialog := omp.NewListDialog("List Dialog", "Ok", "Cancel")

		dialog.SetItems([]string{"Item 0", "Item 1", "Item 2"})

		dialog.Add("Item 3", "Item 4")

		dialog.ShowFor(cmd.Sender)
	})

	omp.Commands.Add("pwddlg", func(cmd *omp.Command) {
		dialog := omp.NewPasswordDialog("Password Dialog", "Enter password:", "Ok", "Cancel")
		dialog.ShowFor(cmd.Sender)
	})

	omp.Commands.Add("hidedlg", func(cmd *omp.Command) {
		time.AfterFunc(5*time.Second, func() {
			dialog, err := cmd.Sender.Dialog()
			if err != nil {
				cmd.Sender.SendClientMessage("You have no active dialog", 0xFF00FFFF)
				return
			}

			cmd.Sender.SendClientMessage("timer worked", 0xFF00FFFF)
			dialog.HideFor(cmd.Sender)
		})
	})

	omp.Commands.Add("tablistdlg", func(cmd *omp.Command) {
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

		dialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
			e.Player.SendClientMessage(fmt.Sprintf("Response: %d, itemno: %d, item: %+v", e.Response, e.ItemNumber, e.Item), 0xFFFF00FF)
			return true
		})

		dialog.On(omp.EventTypeDialogShow, func(e *omp.DialogShowEvent) bool {
			e.Player.SendClientMessage("Dialog shown", 0xFFFF00FF)
			return true
		})

		dialog.On(omp.EventTypeDialogHide, func(e *omp.DialogHideEvent) bool {
			e.Player.SendClientMessage("Dialog hidden", 0xFFFF00FF)
			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.Commands.Add("tablistheadersdlg", func(cmd *omp.Command) {
		dialog := omp.NewTabListDialog("Weapons Dialog", "Ok", "Cancel")

		dialog.SetHeader(omp.TabListItem{"Weapon", "Price", "Ammo"})

		dialog.Add(omp.TabListItem{"Deagle", "$5000", "100"})
		dialog.Add(omp.TabListItem{"Sawnoff", "$5000", "100"})
		dialog.Add(omp.TabListItem{"Pistol", "$1000", "50"})

		dialog.ShowFor(cmd.Sender)

		dialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
			e.Player.SendClientMessage("Dialog response triggered", 0xFFFF00FF)

			return true
		})
	})

	omp.Commands.Add("getpos", func(cmd *omp.Command) {
		cmd.Sender.SendClientMessage(fmt.Sprintf("Your position is %+v", cmd.Sender.Position()), 0xFFFFFFFF)
	})

	omp.Commands.Add("createveh", func(cmd *omp.Command) {
		if len(cmd.Args) != 1 {
			cmd.Sender.SendClientMessage("Invalid command syntax", 0xFFFFFFFF)
			return
		}

		modelID, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			cmd.Sender.SendClientMessage("Invalid command syntax", 0xFFFFFFFF)
			return
		}

		omp.NewVehicle(omp.VehicleModel(modelID), cmd.Sender.Position(), 0.0)
	})

	omp.Commands.Add("setname", func(cmd *omp.Command) {
		status := cmd.Sender.SetName("кириллица")
		cmd.Sender.SendClientMessage(fmt.Sprintf("You changed %d your name to %s", status, cmd.Sender.Name()), 0xFFFFFFFF)
	})

	omp.Commands.Add("doors", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.AreDoorsLocked() {
			plrVeh.UnlockDoors()
			cmd.Sender.SendClientMessage("Doors unlocked", 0xFFFFFFFF)
		} else {
			plrVeh.LockDoors()
			cmd.Sender.SendClientMessage("Doors locked", 0xFFFFFFFF)
		}
	})

	omp.Commands.Add("hood", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.IsHoodOpen() {
			plrVeh.CloseHood()
		} else {
			plrVeh.OpenHood()
		}
	})

	omp.Commands.Add("trunk", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.IsTrunkOpen() {
			plrVeh.CloseTrunk()
		} else {
			plrVeh.OpenTrunk()
		}
	})

	omp.Commands.Add("lights", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.AreLightsTurnedOn() {
			plrVeh.TurnOffLights()
		} else {
			plrVeh.TurnOnLights()
		}
	})

	omp.Commands.Add("engine", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.IsEngineStarted() {
			plrVeh.StopEngine()
		} else {
			plrVeh.StartEngine()
		}
	})

	omp.Commands.Add("alarm", func(cmd *omp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.IsAlarmTurnedOn() {
			plrVeh.TurnOffAlarm()
		} else {
			plrVeh.TurnOnAlarm()
		}
	})
}

func main() {}
