package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kodeyeen/omp"
)

const (
	ColorWhite omp.Color = 0xFFFFFFFF
)

type City int

const (
	CityLosSantos City = iota
	CitySanFierro
	CityLasVenturas
)

type Character struct {
	*omp.Player
	citySelection      City
	hasCitySelected    bool
	lastCitySelectedAt time.Time
}

var chars = make(map[int]*Character, 1000)

var lsTd *omp.Textdraw
var sfTd *omp.Textdraw
var lvTd *omp.Textdraw
var classSelHelperTd *omp.Textdraw

var vehFiles = []string{
	"bone.txt", "flint.txt", "ls_airport.txt", "ls_gen_inner.txt", "ls_gen_outer.txt",
	"ls_law.txt", "lv_airport.txt", "lv_gen.txt", "lv_law.txt", "pilots.txt",
	"red_county.txt", "sf_airport.txt", "sf_gen.txt", "sf_law.txt", "sf_train.txt",
	"tierra.txt", "trains_platform.txt", "trains.txt", "whetstone.txt",
}

var lsSpawns = losSantosSpawns()
var sfSpawns = sanFierroSpawns()
var lvSpawns = lasVenturasSpawns()

func onGameModeInit(event *omp.GameModeInitEvent) bool {
	omp.SetGameModeText("Grand Larceny")
	// omp.SetPlayerMarkerMode(omp.PlayerMarkerModeGlobal)
	// omp.EnableNametags()
	// omp.SetNametagDrawRadius(40.0)
	// omp.EnableStuntBonuses()
	// omp.DisableEntryExitMarkers()
	omp.SetWeather(2)
	omp.SetWorldTime(11)

	lsTd, _ = NewCityNameTextdraw("Los Santos")
	sfTd, _ = NewCityNameTextdraw("San Fierro")
	lvTd, _ = NewCityNameTextdraw("Las Venturas")

	classSelHelperTd, _ = omp.NewTextdraw("Press ~b~~k~~GO_LEFT~ ~w~or ~b~~k~~GO_RIGHT~ ~w~to switch cities.~n~ Press ~r~~k~~PED_FIREWEAPON~ ~w~to select.", omp.Vector2{X: 10.0, Y: 415.0})
	classSelHelperTd.EnableBox()
	classSelHelperTd.SetBoxColor(0x222222BB)
	classSelHelperTd.SetLetterSize(omp.Vector2{X: 0.3, Y: 1.0})
	classSelHelperTd.SetTextSize(omp.Vector2{X: 400.0, Y: 40.0})
	classSelHelperTd.SetStyle(omp.TextdrawStyle2)
	classSelHelperTd.SetShadow(0)
	classSelHelperTd.SetOutline(1)
	classSelHelperTd.SetBackgroundColor(0x000000FF)
	classSelHelperTd.SetColor(0xFFFFFFFF)

	omp.NewClass(0, 298, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 299, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 300, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 301, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 302, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 303, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 304, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 305, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 280, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 281, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 282, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 283, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 284, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 285, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 286, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 287, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 288, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 289, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 265, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 266, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 267, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 268, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 269, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 270, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 1, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 2, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 3, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 4, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 5, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 6, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 8, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 42, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 65, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	//omp.NewClass(0, 74, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 86, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 119, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 149, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 208, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 273, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 289, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)

	omp.NewClass(0, 47, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 48, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 49, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 50, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 51, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 52, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 53, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 54, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 55, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 56, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 57, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 58, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 68, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 69, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 70, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 71, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 72, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 73, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 75, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 76, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 78, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 79, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 80, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 81, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 82, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 83, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 84, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 85, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 87, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 88, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 89, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 91, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 92, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 93, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 95, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 96, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 97, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 98, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)
	omp.NewClass(0, 99, omp.Vector3{X: 1759.0189, Y: -1898.1260, Z: 13.5622}, 266.4503, -1, -1, -1, -1, -1, -1)

	var vehCnt int
	for _, vehFile := range vehFiles {
		filename := filepath.Join("scriptfiles/vehicles", vehFile)

		cnt, err := LoadStaticVehiclesFromFile(filename)
		if err != nil {
			fmt.Printf("Failed to load vehicles from %s: %s\n", filename, err.Error())
		}

		vehCnt += cnt
	}

	fmt.Printf("Total vehicles from files: %d\n", vehCnt)

	return true
}

func onPlayerConnect(event *omp.PlayerConnectEvent) bool {
	char := &Character{
		Player:             event.Player,
		citySelection:      -1,
		hasCitySelected:    false,
		lastCitySelectedAt: time.Now(),
	}

	chars[char.ID()] = char

	char.ShowGameText("~w~Grand Larceny", 3*time.Second, 4)
	char.SendClientMessage("Welcome to {88AA88}G{FFFFFF}rand {88AA88}L{FFFFFF}arceny", ColorWhite)

	return true
}

func onPlayerSpawn(event *omp.PlayerSpawnEvent) bool {
	char := chars[event.Player.ID()]

	if char.IsBot() {
		return true
	}

	char.SetInterior(0)
	char.HideClock()
	char.ResetMoney()
	char.GiveMoney(30000)

	var randSpawn Spawn

	if char.citySelection == CityLosSantos {
		randN := rand.IntN(len(lsSpawns))
		randSpawn = lsSpawns[randN]
	} else if char.citySelection == CitySanFierro {
		randN := rand.IntN(len(sfSpawns))
		randSpawn = sfSpawns[randN]
	} else if char.citySelection == CityLasVenturas {
		randN := rand.IntN(len(lvSpawns))
		randSpawn = lvSpawns[randN]
	}

	char.SetPosition(omp.Vector3{X: randSpawn.pos.X, Y: randSpawn.pos.Y, Z: randSpawn.pos.Z})
	char.SetFacingAngle(randSpawn.angle)

	char.GiveWeapon(omp.WeaponColt45, 100)
	char.HideClock()

	return true
}

func onPlayerRequestClass(event *omp.PlayerRequestClassEvent) bool {
	char := chars[event.Player.ID()]

	if char.IsBot() {
		return true
	}

	if char.hasCitySelected {
		setupCharSelection(char)
		return true
	}

	if char.State() != omp.PlayerStateSpectating {
		char.EnableSpectating()
		classSelHelperTd.ShowFor(char.Player)
		char.citySelection = -1
	}

	return false
}

func onPlayerUpdate(event *omp.PlayerUpdateEvent) bool {
	char := chars[event.Player.ID()]

	if char.IsBot() {
		return true
	}

	if !char.hasCitySelected && char.State() == omp.PlayerStateSpectating {
		handleCitySelection(char)
		return true
	}

	if char.ArmedWeapon() == omp.WeaponMinigun {
		char.Kick()
		return false
	}

	return true
}

func onPlayerDeath(event *omp.PlayerDeathEvent) bool {
	char := chars[event.Player.ID()]

	char.hasCitySelected = false

	var killer *Character
	if event.Killer != nil {
		killer = chars[event.Killer.ID()]
	}

	if killer == nil {
		char.ResetMoney()
		return true
	}

	if char.Money() > 0 {
		killer.GiveMoney(char.Money())
		char.ResetMoney()
	}

	return true
}

func NewCityNameTextdraw(cityName string) (*omp.Textdraw, error) {
	td, err := omp.NewTextdraw(cityName, omp.Vector2{X: 10.0, Y: 380.0})
	if err != nil {
		return nil, err
	}

	td.DisableBox()
	td.SetLetterSize(omp.Vector2{X: 1.25, Y: 3.0})
	td.SetStyle(omp.TextdrawStyle0)
	td.SetShadow(0)
	td.SetOutline(1)
	td.SetColor(0xEEEEEEFF)
	td.SetBackgroundColor(0x000000FF)

	return td, nil
}

func LoadStaticVehiclesFromFile(filename string) (int, error) {
	inf, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer inf.Close()

	r := bufio.NewReader(inf)

	const delim = '\n'
	var cnt int
	var eof bool

	for !eof {
		line, err := r.ReadString(delim)
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				return cnt, err
			}
		}

		line = strings.TrimSuffix(line, string(delim))

		split := strings.Split(line, ",")

		model, err := strconv.ParseInt(split[0], 10, 0)
		if err != nil {
			continue
		}

		spawnX, err := strconv.ParseFloat(split[1], 32)
		if err != nil {
			continue
		}

		spawnY, err := strconv.ParseFloat(split[2], 32)
		if err != nil {
			continue
		}

		spawnZ, err := strconv.ParseFloat(split[3], 32)
		if err != nil {
			continue
		}

		rot, err := strconv.ParseFloat(split[4], 32)
		if err != nil {
			continue
		}

		primaryColor, err := strconv.Atoi(split[5])
		if err != nil {
			continue
		}

		secColAndName := strings.Split(split[6], ";")

		secondaryColor, err := strconv.Atoi(strings.TrimSpace(secColAndName[0]))
		if err != nil {
			continue
		}

		veh, err := omp.NewStaticVehicle(omp.VehicleModel(model), omp.Vector3{X: float32(spawnX), Y: float32(spawnY), Z: float32(spawnZ)}, float32(rot))
		if err != nil {
			continue
		}

		veh.SetColor(omp.VehicleColor{
			Primary:   primaryColor,
			Secondary: secondaryColor,
		})

		cnt++
	}

	fmt.Printf("Loaded %d vehicles from: %s\n", cnt, filename)

	return cnt, nil
}

func handleCitySelection(char *Character) {
	if char.citySelection == -1 {
		switchToNextCity(char)
		return
	}

	if time.Since(char.lastCitySelectedAt) < 500*time.Millisecond {
		return
	}

	keyData := char.KeyData()

	if keyData.Keys&omp.PlayerKeyFire != 0 {
		char.hasCitySelected = true
		lsTd.HideFor(char.Player)
		sfTd.HideFor(char.Player)
		lvTd.HideFor(char.Player)
		classSelHelperTd.HideFor(char.Player)
		char.DisableSpectating()
		return
	}

	if keyData.LeftRight > 0 {
		switchToNextCity(char)
	} else if keyData.LeftRight < 0 {
		switchToPrevCity(char)
	}
}

func switchToNextCity(char *Character) {
	char.citySelection++
	if char.citySelection > CityLasVenturas {
		char.citySelection = CityLosSantos
	}

	char.PlaySound(1052, omp.Vector3{})
	char.lastCitySelectedAt = time.Now()
	setupSelectedCity(char)
}

func switchToPrevCity(char *Character) {
	char.citySelection--
	if char.citySelection < CityLosSantos {
		char.citySelection = CityLasVenturas
	}

	char.PlaySound(1053, omp.Vector3{})
	char.lastCitySelectedAt = time.Now()
	setupSelectedCity(char)
}

func setupSelectedCity(char *Character) {
	if char.citySelection == -1 {
		char.citySelection = CityLosSantos
	}

	char.SetInterior(0)

	if char.citySelection == CityLosSantos {
		char.SetCameraPosition(omp.Vector3{X: 1630.6136, Y: -2286.0298, Z: 110.0})
		char.SetCameraLookAt(omp.Vector3{X: 1887.6034, Y: -1682.1442, Z: 47.6167}, omp.PlayerCameraCutTypeCut)

		lsTd.ShowFor(char.Player)
		sfTd.HideFor(char.Player)
		lvTd.HideFor(char.Player)
	} else if char.citySelection == CitySanFierro {
		char.SetCameraPosition(omp.Vector3{X: -1300.8754, Y: 68.0546, Z: 129.4823})
		char.SetCameraLookAt(omp.Vector3{X: -1817.9412, Y: 769.3878, Z: 132.6589}, omp.PlayerCameraCutTypeCut)

		lsTd.HideFor(char.Player)
		sfTd.ShowFor(char.Player)
		lvTd.HideFor(char.Player)
	} else if char.citySelection == CityLasVenturas {
		char.SetCameraPosition(omp.Vector3{X: 1310.6155, Y: 1675.9182, Z: 110.7390})
		char.SetCameraLookAt(omp.Vector3{X: 2285.2944, Y: 1919.3756, Z: 68.2275}, omp.PlayerCameraCutTypeCut)

		lsTd.HideFor(char.Player)
		sfTd.HideFor(char.Player)
		lvTd.ShowFor(char.Player)
	}
}

func setupCharSelection(char *Character) {
	if char.citySelection == CityLosSantos {
		char.SetInterior(11)
		char.SetPosition(omp.Vector3{X: 508.7362, Y: -87.4335, Z: 998.9609})
		char.SetFacingAngle(0.0)
		char.SetCameraPosition(omp.Vector3{X: 508.7362, Y: -83.4335, Z: 998.9609})
		char.SetCameraLookAt(omp.Vector3{X: 508.7362, Y: -87.4335, Z: 998.9609}, omp.PlayerCameraCutTypeCut)
	} else if char.citySelection == CitySanFierro {
		char.SetInterior(3)
		char.SetPosition(omp.Vector3{X: -2673.8381, Y: 1399.7424, Z: 918.3516})
		char.SetFacingAngle(181.0)
		char.SetCameraPosition(omp.Vector3{X: -2673.2776, Y: 1394.3859, Z: 918.3516})
		char.SetCameraLookAt(omp.Vector3{X: -2673.8381, Y: 1399.7424, Z: 918.3516}, omp.PlayerCameraCutTypeCut)
	} else if char.citySelection == CityLasVenturas {
		char.SetInterior(3)
		char.SetPosition(omp.Vector3{X: 349.0453, Y: 193.2271, Z: 1014.1797})
		char.SetFacingAngle(286.25)
		char.SetCameraPosition(omp.Vector3{X: 352.9164, Y: 194.5702, Z: 1014.1875})
		char.SetCameraLookAt(omp.Vector3{X: 349.0453, Y: 193.2271, Z: 1014.1797}, omp.PlayerCameraCutTypeCut)
	}
}

func main() {}

func init() {
	omp.On(omp.EventTypeGameModeInit, onGameModeInit)
	omp.On(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.On(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.On(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.On(omp.EventTypePlayerUpdate, onPlayerUpdate)
	omp.On(omp.EventTypePlayerDeath, onPlayerDeath)
}
