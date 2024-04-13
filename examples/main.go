package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kodeyeen/gomp"
)

// go build -o test.dll -buildmode=c-shared

func init() {
	gomp.On(gomp.EventTypeGameModeInit, func(evt *gomp.GameModeInitEvent) {
		log.Println("GAME MODE INITIALIZED")
		gomp.EnableManualEngineAndLights()
	})

	gomp.On(gomp.EventTypePlayerConnect, func(evt *gomp.PlayerConnectEvent) {
		plr := evt.Player

		plr.GiveWeapon(gomp.WeaponDeagle, 100)

		plr.SendMessage(fmt.Sprintf("Hello, %s", plr.Name()), 0x00FF0000)

		gomp.NewVehicle(gomp.VehicleModelAlpha, gomp.Vector3{X: 2161.8389, Y: -1143.7473, Z: 24.6501}, 266.9070)

		gomp.NewPickup(1273, 2, gomp.Vector3{X: 2216.0325, Y: -1161.7224, Z: 25.7266}, 0)
	})

	var veh *gomp.Vehicle

	gomp.On(gomp.EventTypePlayerCommandText, func(evt *gomp.PlayerCommandTextEvent) {
		plr := evt.Player

		tmp := strings.Fields(evt.Command)
		cmdName := strings.TrimPrefix(tmp[0], "/")
		cmdArgs := tmp[1:]

		plr.SendMessage(cmdName, 0xFFFFFFFF)
		log.Println("COMMAND", cmdName)

		switch cmdName {
		// case "createcp":
		// 	plr.NewDefaultCheckpoint(5.0, gomp.Vector3{X: -38.9655, Y: 30.4141, Z: 3.1172})
		case "getpos":
			plr.SendMessage(fmt.Sprintf("Your position is %+v", plr.Position()), 0xFFFFFFFF)
		case "setname":
			status := plr.SetName("кириллица")
			plr.SendMessage(fmt.Sprintf("You changed %d your name to %s", status, plr.Name()), 0xFFFFFFFF)
		case "createveh":
			modelID, err := strconv.Atoi(cmdArgs[0])
			if err != nil {
				plr.SendMessage("Invalid command syntax", 0xFFFFFFFF)
				return
			}

			veh, _ = gomp.NewVehicle(gomp.VehicleModel(modelID), plr.Position(), 0.0)
		case "doors":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.AreDoorsLocked() {
				plrVeh.UnlockDoors()
				plr.SendMessage("Doors unlocked", 0xFFFFFFFF)
			} else {
				plrVeh.LockDoors()
				plr.SendMessage("Doors locked", 0xFFFFFFFF)
			}
		case "hood":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.IsHoodOpen() {
				plrVeh.CloseHood()
			} else {
				plrVeh.OpenHood()
			}
		case "trunk":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.IsTrunkOpen() {
				plrVeh.CloseTrunk()
			} else {
				plrVeh.OpenTrunk()
			}
		case "lights":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.AreLightsTurnedOn() {
				plrVeh.TurnOffLights()
			} else {
				plrVeh.TurnOnLights()
			}
		case "engine":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.IsEngineStarted() {
				plrVeh.StopEngine()
			} else {
				plrVeh.StartEngine()
			}
		case "alarm":
			plrVeh, err := plr.Vehicle()
			if err != nil {
				return
			}

			if plrVeh.IsAlarmTurnedOn() {
				plrVeh.TurnOffAlarm()
			} else {
				plrVeh.TurnOnAlarm()
			}
		case "passengers":
			// plrVeh, err := plr.Vehicle()
			// if err != nil {
			// 	log.Println("NO VEH")
			// 	return
			// }

			passengers := veh.Passengers()
			log.Println("PASSES", passengers)
		}
	})
}

func main() {}
