package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kodeyeen/omp"
)

// GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o test.dll

func init() {
	omp.On(omp.EventTypeGameModeInit, func(evt *omp.GameModeInitEvent) bool {
		omp.Println("GAME MODE INIT")

		omp.EnableManualEngineAndLights()
		return true
	})

	omp.On(omp.EventTypePlayerConnect, func(evt *omp.PlayerConnectEvent) bool {
		player := evt.Player

		player.SendClientMessage(fmt.Sprintf("Hello, %s", player.Name()), 0x00FF0000)

		omp.NewVehicle(omp.VehicleModelAlpha, omp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		omp.NewPickup(1273, 1, 0, omp.Vector3{X: 3.9905, Y: 7.0804, Z: 3.1096})

		return true
	})

	omp.On(omp.EventTypePlayerSpawn, func(evt *omp.PlayerSpawnEvent) bool {
		player := evt.Player

		player.GiveWeapon(omp.WeaponDeagle, 100)

		return true
	})

	omp.AddCommand("msgdlg", func(cmd *omp.Command) {
		dialog := omp.NewMessageDialog("Message Dialog", "Message", "Ok", "Cancel")

		dialog.On(omp.EventTypeDialogHide, func(evt *omp.DialogHideEvent) bool {
			evt.Player.SendClientMessage("Dialog is hiding", 0x00FFFFFF)
			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.AddCommand("inputdlg", func(cmd *omp.Command) {
		dialog := omp.NewInputDialog("Input Dialog", "Enter something:", "Ok", "Cancel")

		dialog.On(omp.EventTypeDialogResponse, func(evt *omp.InputDialogResponseEvent) bool {
			if evt.Response == omp.DialogResponseLeft {
				evt.Player.SendClientMessage(fmt.Sprintf("Left button. Your input is %s", evt.InputText), 0xFF0FFFFF)
			} else if evt.Response == omp.DialogResponseRight {
				evt.Player.SendClientMessage(fmt.Sprintf("Right button. Your input is %s", evt.InputText), 0xFF0FFFFF)
			}

			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.AddCommand("listdlg", func(cmd *omp.Command) {
		dialog := omp.NewListDialog("List Dialog", "Ok", "Cancel")

		dialog.SetItems([]string{"Item 0", "Item 1", "Item 2"})

		dialog.Add("Item 3", "Item 4")

		dialog.ShowFor(cmd.Sender)
	})

	omp.AddCommand("pwddlg", func(cmd *omp.Command) {
		dialog := omp.NewPasswordDialog("Password Dialog", "Enter password:", "Ok", "Cancel")
		dialog.ShowFor(cmd.Sender)
	})

	omp.AddCommand("hidedlg", func(cmd *omp.Command) {
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

	omp.AddCommand("tablistdlg", func(cmd *omp.Command) {
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

		dialog.On(omp.EventTypeDialogResponse, func(event *omp.TabListDialogResponseEvent) bool {
			event.Player.SendClientMessage(fmt.Sprintf("Response: %d, itemno: %d, item: %+v", event.Response, event.ItemNumber, event.Item), 0xFFFF00FF)
			return true
		})

		dialog.On(omp.EventTypeDialogShow, func(event *omp.DialogShowEvent) bool {
			event.Player.SendClientMessage("Dialog shown", 0xFFFF00FF)
			return true
		})

		dialog.On(omp.EventTypeDialogHide, func(event *omp.DialogHideEvent) bool {
			event.Player.SendClientMessage("Dialog hidden", 0xFFFF00FF)
			return true
		})

		dialog.ShowFor(cmd.Sender)
	})

	omp.AddCommand("tablistheadersdlg", func(cmd *omp.Command) {
		dialog := omp.NewTabListDialog("Weapons Dialog", "Ok", "Cancel")

		dialog.SetHeader(omp.TabListItem{"Weapon", "Price", "Ammo"})

		dialog.Add(omp.TabListItem{"Deagle", "$5000", "100"})
		dialog.Add(omp.TabListItem{"Sawnoff", "$5000", "100"})
		dialog.Add(omp.TabListItem{"Pistol", "$1000", "50"})

		dialog.ShowFor(cmd.Sender)

		dialog.On(omp.EventTypeDialogResponse, func(evt *omp.TabListDialogResponseEvent) bool {
			evt.Player.SendClientMessage("Dialog response triggered", 0xFFFF00FF)

			return true
		})
	})

	omp.AddCommand("getpos", func(cmd *omp.Command) {
		cmd.Sender.SendClientMessage(fmt.Sprintf("Your position is %+v", cmd.Sender.Position()), 0xFFFFFFFF)
	})

	omp.AddCommand("createveh", func(cmd *omp.Command) {
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

	omp.AddCommand("setname", func(cmd *omp.Command) {
		status := cmd.Sender.SetName("кириллица")
		cmd.Sender.SendClientMessage(fmt.Sprintf("You changed %d your name to %s", status, cmd.Sender.Name()), 0xFFFFFFFF)
	})

	omp.AddCommand("doors", func(cmd *omp.Command) {
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

	omp.AddCommand("hood", func(cmd *omp.Command) {
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

	omp.AddCommand("trunk", func(cmd *omp.Command) {
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

	omp.AddCommand("lights", func(cmd *omp.Command) {
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

	omp.AddCommand("engine", func(cmd *omp.Command) {
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

	omp.AddCommand("alarm", func(cmd *omp.Command) {
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
