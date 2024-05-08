package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kodeyeen/omp"
)

// go build -buildmode=c-shared -o test.dll

func init() {
	omp.On(omp.EventTypeGameModeInit, func(evt *omp.GameModeInitEvent) bool {
		log.Println("GAME MODE INITIALIZED")
		omp.EnableManualEngineAndLights()
		return true
	})

	omp.On(omp.EventTypePlayerConnect, func(evt *omp.PlayerConnectEvent) bool {
		player := evt.Player

		player.GiveWeapon(omp.WeaponDeagle, 100)

		player.SendClientMessage(fmt.Sprintf("Hello, %s", player.Name()), 0x00FF0000)

		omp.NewVehicle(omp.VehicleModelAlpha, omp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		omp.NewPickup(1273, 1, 0, omp.Vector3{X: 3.9905, Y: 7.0804, Z: 3.1096})

		return true
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
