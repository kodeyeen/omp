package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kodeyeen/gomp"
)

// go build -buildmode=c-shared -o test.dll

func init() {
	gomp.On(gomp.EventTypeGameModeInit, func(evt *gomp.GameModeInitEvent) bool {
		log.Println("GAME MODE INITIALIZED")
		gomp.EnableManualEngineAndLights()
		return true
	})

	gomp.On(gomp.EventTypePlayerConnect, func(evt *gomp.PlayerConnectEvent) bool {
		plr := evt.Player

		plr.GiveWeapon(gomp.WeaponDeagle, 100)

		plr.SendMessage(fmt.Sprintf("Hello, %s", plr.Name()), 0x00FF0000)

		gomp.NewVehicle(gomp.VehicleModelAlpha, gomp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		gomp.NewPickup(1273, 1, 0, gomp.Vector3{X: 3.9905, Y: 7.0804, Z: 3.1096})

		return true
	})

	gomp.AddCommand("getpos", func(cmd *gomp.Command) {
		cmd.Sender.SendMessage(fmt.Sprintf("Your position is %+v", cmd.Sender.Position()), 0xFFFFFFFF)
	})

	gomp.AddCommand("createveh", func(cmd *gomp.Command) {
		if len(cmd.Args) != 1 {
			cmd.Sender.SendMessage("Invalid command syntax", 0xFFFFFFFF)
			return
		}

		modelID, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			cmd.Sender.SendMessage("Invalid command syntax", 0xFFFFFFFF)
			return
		}

		gomp.NewVehicle(gomp.VehicleModel(modelID), cmd.Sender.Position(), 0.0)
	})

	gomp.AddCommand("setname", func(cmd *gomp.Command) {
		status := cmd.Sender.SetName("кириллица")
		cmd.Sender.SendMessage(fmt.Sprintf("You changed %d your name to %s", status, cmd.Sender.Name()), 0xFFFFFFFF)
	})

	gomp.AddCommand("doors", func(cmd *gomp.Command) {
		plrVeh, err := cmd.Sender.Vehicle()
		if err != nil {
			return
		}

		if plrVeh.AreDoorsLocked() {
			plrVeh.UnlockDoors()
			cmd.Sender.SendMessage("Doors unlocked", 0xFFFFFFFF)
		} else {
			plrVeh.LockDoors()
			cmd.Sender.SendMessage("Doors locked", 0xFFFFFFFF)
		}
	})

	gomp.AddCommand("hood", func(cmd *gomp.Command) {
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

	gomp.AddCommand("trunk", func(cmd *gomp.Command) {
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

	gomp.AddCommand("lights", func(cmd *gomp.Command) {
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

	gomp.AddCommand("engine", func(cmd *gomp.Command) {
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

	gomp.AddCommand("alarm", func(cmd *gomp.Command) {
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
