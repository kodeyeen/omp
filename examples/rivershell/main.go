package main

import (
	"time"

	"github.com/kodeyeen/gomp"
)

const (
	RespawnDelay = 20 * time.Second

	ColorObjective gomp.Color = 0xE2C063FF
	ColorGreen     gomp.Color = 0x77CC77FF
	ColorBlue      gomp.Color = 0x7777DDFF
)

var (
	chars              = make(map[int]*Character, 1000)
	teams              = make(map[int]*Team, 2)
	objectiveVehGreen  *gomp.Vehicle
	objectiveVehBlue   *gomp.Vehicle
	isObjectiveReached bool
	objectiveGreenChar *Character
	objectiveBlueChar  *Character
)

func onGameModeInit(evt *gomp.GameModeInitEvent) bool {
	gomp.SetGameModeText("Rivershell")

	gomp.SetPlayerMarkerMode(gomp.PlayerMarkerModeOff)
	gomp.EnableNametags()
	gomp.SetWorldTime(17)
	gomp.SetWeather(11)
	gomp.EnablePlayerPedAnims()
	gomp.EnableVehicleFriendlyFire()
	gomp.SetNametagDrawRadius(110.0)
	gomp.DisableInteriorEnterExits()

	// Green classes
	gomp.NewClass(0, 162, gomp.Vector3{X: 2117.0129, Y: -224.4389, Z: 8.15}, 0.0, gomp.WeaponM4, 100, gomp.WeaponMP5, 200, gomp.WeaponSniper, 10)
	gomp.NewClass(0, 157, gomp.Vector3{X: 2148.6606, Y: -224.3336, Z: 8.15}, 347.1396, gomp.WeaponM4, 100, gomp.WeaponMP5, 200, gomp.WeaponSniper, 10)

	// Blue classes
	gomp.NewClass(0, 154, gomp.Vector3{X: 2352.9873, Y: 580.3051, Z: 7.7813}, 178.1424, gomp.WeaponM4, 100, gomp.WeaponMP5, 200, gomp.WeaponSniper, 10)
	gomp.NewClass(0, 138, gomp.Vector3{X: 2281.1504, Y: 567.6248, Z: 7.7813}, 163.7289, gomp.WeaponM4, 100, gomp.WeaponMP5, 200, gomp.WeaponSniper, 10)

	// Objective vehicles
	objectiveVehGreen, _ = NewVehicle(gomp.VehicleModelReefer, gomp.Vector3{X: 2184.7156, Y: -188.5401, Z: -0.0239}, 0.0000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second) // gr reefer
	objectiveVehBlue, _ = NewVehicle(gomp.VehicleModelReefer, gomp.Vector3{X: 2380.0542, Y: 535.2582, Z: -0.0272}, 178.4999, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)  // bl reefer

	// Green Dhingys
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2096.0833, Y: -168.7771, Z: 0.3528}, 4.5000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2103.2510, Y: -168.7598, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2099.4966, Y: -168.8216, Z: 0.3528}, 2.8200, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2107.1143, Y: -168.7798, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2111.0674, Y: -168.7609, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2114.8933, Y: -168.7898, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2167.2217, Y: -169.0570, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2170.4294, Y: -168.9724, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2173.7952, Y: -168.9217, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2177.0386, Y: -168.9767, Z: 0.3528}, 3.1800, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2161.5786, Y: -191.9538, Z: 0.3528}, 89.1000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2161.6394, Y: -187.2925, Z: 0.3528}, 89.1000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2161.7610, Y: -183.0225, Z: 0.3528}, 89.1000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2162.0283, Y: -178.5106, Z: 0.3528}, 89.1000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	// Green Mavericks
	NewVehicle(gomp.VehicleModelMaverick, gomp.Vector3{X: 2088.7905, Y: -227.9593, Z: 8.3662}, 0.0000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(gomp.VehicleModelMaverick, gomp.Vector3{X: 2204.5991, Y: -225.3703, Z: 8.2400}, 0.0000, gomp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)

	// Blue Dhingys
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2370.3198, Y: 518.3151, Z: 0.1240}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2362.6484, Y: 518.3978, Z: 0.0598}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2358.6550, Y: 518.2167, Z: 0.2730}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2366.5544, Y: 518.2680, Z: 0.1080}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2354.6321, Y: 518.1960, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2350.7449, Y: 518.1929, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2298.8977, Y: 518.4470, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2295.6118, Y: 518.3963, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2292.3237, Y: 518.4249, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2289.0901, Y: 518.4363, Z: 0.3597}, 180.3600, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2304.8232, Y: 539.7859, Z: 0.3597}, 270.5998, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2304.6936, Y: 535.0454, Z: 0.3597}, 270.5998, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2304.8245, Y: 530.3308, Z: 0.3597}, 270.5998, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelDinghy, gomp.Vector3{X: 2304.8142, Y: 525.7471, Z: 0.3597}, 270.5998, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)

	// Blue Mavericks
	NewVehicle(gomp.VehicleModelMaverick, gomp.Vector3{X: 2260.2637, Y: 578.5220, Z: 8.1223}, 182.3401, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(gomp.VehicleModelMaverick, gomp.Vector3{X: 2379.9792, Y: 580.0323, Z: 8.0178}, 177.9601, gomp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)

	// Green Base Section
	gomp.NewObject(9090, 500.0, gomp.Vector3{X: 2148.64, Y: -222.88, Z: -20.60}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 179.70})
	// Green resupply hut
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})

	// Blue Base Section
	gomp.NewObject(9090, 500.0, gomp.Vector3{X: 2317.09, Y: 572.27, Z: -20.97}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	// Blue resupply hut
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})

	// General mapping
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	gomp.NewObject(19300, 500.0, gomp.Vector3{X: 2137.33, Y: -237.17, Z: 46.61}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 180.00})
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	gomp.NewObject(19300, 500.0, gomp.Vector3{X: 2325.41, Y: 587.93, Z: 47.37}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 180.00})
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	gomp.NewObject(12991, 500.0, gomp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	gomp.NewObject(18228, 500.0, gomp.Vector3{X: 1887.93, Y: -59.78, Z: -2.14}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 20.34})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 1990.19, Y: 541.37, Z: -22.32}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(18227, 500.0, gomp.Vector3{X: 2000.82, Y: 494.15, Z: -7.53}, gomp.Vector3{X: 11.70, Y: -25.74, Z: 154.38})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 1992.35, Y: 539.80, Z: -2.97}, gomp.Vector3{X: 9.12, Y: 30.66, Z: 0.00})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 1991.88, Y: 483.77, Z: -0.66}, gomp.Vector3{X: -2.94, Y: -5.22, Z: 12.78})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2070.57, Y: -235.87, Z: -6.05}, gomp.Vector3{X: -7.20, Y: 4.08, Z: 114.30})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2056.50, Y: -228.77, Z: -19.67}, gomp.Vector3{X: 14.16, Y: 19.68, Z: 106.56})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2074.00, Y: -205.33, Z: -18.60}, gomp.Vector3{X: 16.02, Y: 60.60, Z: 118.86})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2230.39, Y: -242.59, Z: -11.41}, gomp.Vector3{X: 5.94, Y: 7.56, Z: 471.24})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2252.53, Y: -213.17, Z: -20.81}, gomp.Vector3{X: 18.90, Y: -6.30, Z: -202.38})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2233.04, Y: -234.08, Z: -19.00}, gomp.Vector3{X: 21.84, Y: -8.88, Z: -252.06})
	gomp.NewObject(17027, 500.0, gomp.Vector3{X: 2235.05, Y: -201.49, Z: -11.90}, gomp.Vector3{X: -11.94, Y: -4.08, Z: 136.32})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2226.11, Y: -237.07, Z: -2.45}, gomp.Vector3{X: 8.46, Y: 2.10, Z: 471.24})
	gomp.NewObject(4368, 500.0, gomp.Vector3{X: 2433.79, Y: 446.26, Z: 4.67}, gomp.Vector3{X: -8.04, Y: -9.30, Z: 61.02})
	gomp.NewObject(4368, 500.0, gomp.Vector3{X: 2031.23, Y: 489.92, Z: -13.20}, gomp.Vector3{X: -8.04, Y: -9.30, Z: -108.18})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2458.36, Y: 551.10, Z: -6.95}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2465.37, Y: 511.35, Z: -7.70}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2474.80, Y: 457.71, Z: -5.17}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 172.74})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2466.03, Y: 426.28, Z: -5.17}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2310.45, Y: -229.38, Z: 7.41}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2294.00, Y: -180.15, Z: 7.41}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2017.50, Y: -305.30, Z: 7.29}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2106.45, Y: -279.86, Z: 20.05}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	gomp.NewObject(706, 500.0, gomp.Vector3{X: 2159.13, Y: -263.71, Z: 19.22}, gomp.Vector3{X: 356.86, Y: 0.00, Z: -17.18})
	gomp.NewObject(706, 500.0, gomp.Vector3{X: 2055.75, Y: -291.53, Z: 13.98}, gomp.Vector3{X: 356.86, Y: 0.00, Z: -66.50})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 1932.65, Y: -315.88, Z: 6.77}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -35.76})
	gomp.NewObject(790, 500.0, gomp.Vector3{X: 2429.40, Y: 575.79, Z: 10.42}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 3.14})
	gomp.NewObject(790, 500.0, gomp.Vector3{X: 2403.40, Y: 581.56, Z: 10.42}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 29.48})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2083.44, Y: 365.48, Z: 13.19}, gomp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2040.15, Y: 406.02, Z: 13.33}, gomp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 1995.36, Y: 588.10, Z: 7.50}, gomp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2126.11, Y: 595.15, Z: 5.99}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -35.82})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2188.35, Y: 588.90, Z: 6.04}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	gomp.NewObject(791, 500.0, gomp.Vector3{X: 2068.56, Y: 595.58, Z: 5.99}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 52.62})
	gomp.NewObject(698, 500.0, gomp.Vector3{X: 2385.32, Y: 606.16, Z: 9.79}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 34.62})
	gomp.NewObject(698, 500.0, gomp.Vector3{X: 2309.29, Y: 606.92, Z: 9.79}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -54.54})
	gomp.NewObject(790, 500.0, gomp.Vector3{X: 2347.14, Y: 619.77, Z: 9.94}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 3.14})
	gomp.NewObject(698, 500.0, gomp.Vector3{X: 2255.28, Y: 606.94, Z: 9.79}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -92.76})
	gomp.NewObject(4298, 500.0, gomp.Vector3{X: 2121.37, Y: 544.12, Z: -5.74}, gomp.Vector3{X: -10.86, Y: 6.66, Z: 3.90})
	gomp.NewObject(4368, 500.0, gomp.Vector3{X: 2273.18, Y: 475.02, Z: -15.30}, gomp.Vector3{X: 4.80, Y: 8.10, Z: 266.34})
	gomp.NewObject(18227, 500.0, gomp.Vector3{X: 2232.38, Y: 451.61, Z: -30.71}, gomp.Vector3{X: -18.54, Y: -6.06, Z: 154.38})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2228.15, Y: 518.87, Z: -16.51}, gomp.Vector3{X: 13.14, Y: -1.32, Z: -20.10})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2230.42, Y: 558.52, Z: -18.38}, gomp.Vector3{X: -2.94, Y: -5.22, Z: 12.78})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2228.97, Y: 573.62, Z: 5.17}, gomp.Vector3{X: 17.94, Y: -15.60, Z: -4.08})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2116.67, Y: -87.71, Z: -2.31}, gomp.Vector3{X: 5.94, Y: 7.56, Z: 215.22})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2078.66, Y: -83.87, Z: -27.30}, gomp.Vector3{X: 13.02, Y: -53.94, Z: -0.30})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2044.80, Y: -36.91, Z: -9.26}, gomp.Vector3{X: -13.74, Y: 27.90, Z: 293.76})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2242.41, Y: 426.16, Z: -15.43}, gomp.Vector3{X: -21.54, Y: 22.26, Z: 154.80})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2220.06, Y: 450.07, Z: -34.78}, gomp.Vector3{X: -1.32, Y: 10.20, Z: -45.84})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2252.49, Y: 439.08, Z: -19.47}, gomp.Vector3{X: -41.40, Y: 20.16, Z: 331.86})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2241.41, Y: 431.93, Z: -5.62}, gomp.Vector3{X: -2.22, Y: -4.80, Z: 53.64})
	gomp.NewObject(17029, 500.0, gomp.Vector3{X: 2141.10, Y: -81.30, Z: -2.41}, gomp.Vector3{X: 5.94, Y: 7.56, Z: 39.54})
	gomp.NewObject(17031, 500.0, gomp.Vector3{X: 2277.07, Y: 399.31, Z: -1.65}, gomp.Vector3{X: -2.22, Y: -4.80, Z: -121.74})
	gomp.NewObject(17026, 500.0, gomp.Vector3{X: 2072.75, Y: -224.40, Z: -5.25}, gomp.Vector3{X: 0.00, Y: 0.00, Z: -41.22})

	// Ramps
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2131.97, Y: 110.24, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 153.72})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2124.59, Y: 113.69, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 157.56})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2116.31, Y: 116.44, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 160.08})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2113.22, Y: 108.48, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 340.20})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2121.21, Y: 105.21, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 340.20})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2127.84, Y: 102.06, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 334.68})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2090.09, Y: 40.90, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2098.73, Y: 39.12, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2107.17, Y: 37.94, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2115.88, Y: 36.47, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2117.46, Y: 45.86, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 529.20})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2108.98, Y: 46.95, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 529.20})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2100.42, Y: 48.11, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 526.68})
	gomp.NewObject(1632, 500.0, gomp.Vector3{X: 2091.63, Y: 50.02, Z: 0.00}, gomp.Vector3{X: 0.00, Y: 0.00, Z: 526.80})

	return true
}

func onPlayerConnect(evt *gomp.PlayerConnectEvent) bool {
	char := &Character{
		Player: evt.Player,
	}

	chars[char.ID()] = char

	char.SetColor(0x888888FF)
	char.ShowGameText("~r~SA-MP: ~w~Rivershell", 2*time.Second, 5)
	char.RemoveNeededBuildings()

	return true
}

func onPlayerDisconnect(evt *gomp.PlayerDisconnectEvent) bool {
	delete(chars, evt.Player.ID())
	return true
}

func onPlayerRequestClass(evt *gomp.PlayerRequestClassEvent) bool {
	char := chars[evt.Player.ID()]

	char.SetupForClassSelection()
	char.SetTeamFromClass(evt.Class)

	clsID := evt.Class.ID()

	if clsID == 0 || clsID == 1 {
		char.ShowGameText("~g~GREEN ~w~TEAM", 1*time.Second, 5)
	} else if clsID == 2 || clsID == 3 {
		char.ShowGameText("~b~BLUE ~w~TEAM", 1*time.Second, 5)
	}

	return true
}

func onPlayerSpawn(evt *gomp.PlayerSpawnEvent) bool {
	char := chars[evt.Player.ID()]

	if !char.LastDiedAt.IsZero() && time.Since(char.LastDiedAt) < RespawnDelay {
		char.SendMessage("Waiting to respawn....", 0xFFAAEEEE)
		char.EnableSpectating()

		if char.LastKiller.State() == gomp.PlayerStateOnFoot ||
			char.LastKiller.State() == gomp.PlayerStateDriver ||
			char.LastKiller.State() == gomp.PlayerStatePassenger {

			char.SpectateCharacter(char.LastKiller)
			char.SpectateState = SpectateStatePlayer
		}

		return true
	}

	char.SetColorFromTeam()

	if char.Team() == TeamGreen.ID {
		char.ShowGameText("Defend the ~g~GREEN ~w~team's ~y~Reefer~n~~w~Capture the ~b~BLUE ~w~team's ~y~Reefer", 6*time.Second, 5)
	} else if char.Team() == TeamBlue.ID {
		char.ShowGameText("Defend the ~b~BLUE ~w~team's ~y~Reefer~n~~w~Capture the ~g~GREEN ~w~team's ~y~Reefer", 6*time.Second, 5)
	}

	char.SetHealth(100.0)
	char.SetArmor(100.0)
	char.SetWorldBounds(gomp.Vector4{X: 2500.0, Y: 1850.0, Z: 631.2963, W: -454.9898})

	char.SpectateState = SpectateStateNone

	return true
}

func onPlayerEnterCheckpoint(evt *gomp.PlayerEnterCheckpointEvent) bool {
	if isObjectiveReached {
		return true
	}

	char := chars[evt.Player.ID()]

	charVeh, err := char.Vehicle()
	if err != nil {
		return true
	}

	if charVeh == objectiveVehGreen && char.Team() == TeamGreen.ID {
		// Green OBJECTIVE REACHED.
		teams[TeamGreen.ID].VehicleCapturedCount++
		char.SetScore(char.Score() + 5)

		if TeamGreen.VehicleCapturedCount == CapturesToWin {
			// gomp.ShowGameTextForAll("~g~GREEN ~w~team wins!", 3*time.Second, 5)
			isObjectiveReached = true
			PlaySoundForAll(1185, gomp.Vector3{})
			time.AfterFunc(6*time.Second, ExitTheGameMode) // Set up a timer to exit this mode.
		} else {
			// gomp.ShowGameTextForAll("~g~GREEN ~w~team captured the ~y~boat!", 3*time.Second, 5)
			objectiveVehGreen.Respawn()
		}
	} else if charVeh == objectiveVehBlue && char.Team() == TeamBlue.ID {
		// Blue OBJECTIVE REACHED.
		teams[TeamBlue.ID].VehicleCapturedCount++
		char.SetScore(char.Score() + 5)

		if TeamBlue.VehicleCapturedCount == CapturesToWin {
			// gomp.ShowGameTextForAll("~b~BLUE ~w~team wins!", 3*time.Second, 5)
			isObjectiveReached = true
			PlaySoundForAll(1185, gomp.Vector3{})
			time.AfterFunc(6*time.Second, ExitTheGameMode) // Set up a timer to exit this mode.
		} else {
			// gomp.ShowGameTextForAll("~b~BLUE ~w~team captured the ~y~boat!", 3*time.Second, 5)
			objectiveVehBlue.Respawn()
		}
	}

	return true
}

func onPlayerDeath(evt *gomp.PlayerDeathEvent) bool {
	char := chars[evt.Player.ID()]
	var killer *Character

	if evt.Killer == nil {
		gomp.SendDeathMessage(nil, char.Player, evt.Reason)
	} else {
		killer = chars[evt.Killer.ID()]

		if killer.Team() != char.Team() {
			// Valid kill
			gomp.SendDeathMessage(killer.Player, char.Player, evt.Reason)
			killer.SetScore(killer.Score() + 1)
		} else {
			// Team kill
			gomp.SendDeathMessage(killer.Player, char.Player, evt.Reason)
		}
	}

	char.LastDiedAt = time.Now()
	char.LastKiller = killer

	return true
}

func onVehicleStreamIn(evt *gomp.VehicleStreamInEvent) bool {
	// As the vehicle streams in, player team dependant params are applied. They can't be
	// applied to vehicles that don't exist in the player's world.
	char := chars[evt.ForPlayer.ID()]

	if evt.Vehicle == objectiveVehBlue {
		if char.Team() == TeamGreen.ID {
			objectiveVehBlue.EnableObjectiveFor(char.Player)
			objectiveVehBlue.LockDoorsFor(char.Player)
		} else if char.Team() == TeamBlue.ID {
			objectiveVehBlue.EnableObjectiveFor(char.Player)
			objectiveVehBlue.UnlockDoorsFor(char.Player)
		}
	} else if evt.Vehicle == objectiveVehGreen {
		if char.Team() == TeamBlue.ID {
			objectiveVehGreen.EnableObjectiveFor(char.Player)
			objectiveVehGreen.LockDoorsFor(char.Player)
		} else if char.Team() == TeamGreen.ID {
			objectiveVehGreen.EnableObjectiveFor(char.Player)
			objectiveVehGreen.UnlockDoorsFor(char.Player)
		}
	}

	return true
}

func onPlayerUpdate(evt *gomp.PlayerUpdateEvent) bool {
	char := chars[evt.Player.ID()]

	if char.IsBot() {
		return true
	}

	if char.State() == gomp.PlayerStateSpectating {
		if char.LastDiedAt.IsZero() {
			char.DisableSpectating()
			return true
		}

		// Allow respawn after an arbitrary time has passed
		if time.Since(char.LastDiedAt) > RespawnDelay {
			char.DisableSpectating()
			return true
		}

		char.HandleSpectating()
		return true
	}

	// Check the resupply huts
	if char.State() == gomp.PlayerStateOnFoot {
		if char.IsInRangeOf(gomp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, 2.5) ||
			char.IsInRangeOf(gomp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, 2.5) {

			char.DoResupply()
		}
	}

	return true
}

func onPlayerStateChange(evt *gomp.PlayerStateChangeEvent) bool {
	char := chars[evt.Player.ID()]

	if evt.NewState == gomp.PlayerStateDriver {
		veh, err := char.Vehicle()
		if err != nil {
			return true
		}

		if char.Team() == TeamGreen.ID && veh == objectiveVehGreen {
			// It's the objective vehicle
			char.SetColor(ColorObjective)
			char.ShowGameText("~w~Take the ~y~boat ~w~back to the ~r~spawn!", 3*time.Second, 5)

			cp := char.DefaultCheckpoint()
			cp.SetPosition(gomp.Vector3{X: 2135.7368, Y: -179.8811, Z: -0.5323})
			cp.SetRadius(10.0)
			cp.Enable()

			objectiveGreenChar = char
		} else if char.Team() == TeamBlue.ID && veh == objectiveVehBlue {
			// It's the objective vehicle
			char.SetColor(ColorObjective)
			char.ShowGameText("~w~Take the ~y~boat ~w~back to the ~r~spawn!", 3*time.Second, 5)

			cp := char.DefaultCheckpoint()
			cp.SetPosition(gomp.Vector3{X: 2329.4226, Y: 532.7426, Z: 0.5862})
			cp.SetRadius(10.0)
			cp.Enable()

			objectiveBlueChar = char
		}
	} else if evt.NewState == gomp.PlayerStateOnFoot {
		if char == objectiveGreenChar {
			objectiveGreenChar = nil
			char.SetColorFromTeam()
			char.DefaultCheckpoint().Disable()
		} else if char == objectiveBlueChar {
			objectiveBlueChar = nil
			char.SetColorFromTeam()
			char.DefaultCheckpoint().Disable()
		}
	}

	return true
}

func NewVehicle(model gomp.VehicleModel, pos gomp.Vector3, angle float32, color gomp.VehicleColor, respawnDelay time.Duration) (*gomp.Vehicle, error) {
	veh, err := gomp.NewVehicle(model, pos, angle)
	if err != nil {
		return nil, err
	}

	veh.SetColor(color)
	veh.SetRespawnDelay(respawnDelay)

	return veh, nil
}

func PlaySoundForAll(sound int, pos gomp.Vector3) {
	for _, player := range gomp.Players() {
		player.PlaySound(sound, pos)
	}
}

func ExitTheGameMode() {
	PlaySoundForAll(1186, gomp.Vector3{})
	// gomp.ExitGameMode()
}

func main() {}

func init() {
	gomp.On(gomp.EventTypeGameModeInit, onGameModeInit)
	gomp.On(gomp.EventTypePlayerConnect, onPlayerConnect)
	gomp.On(gomp.EventTypePlayerDisconnect, onPlayerDisconnect)
	gomp.On(gomp.EventTypePlayerRequestClass, onPlayerRequestClass)
	gomp.On(gomp.EventTypePlayerSpawn, onPlayerSpawn)
	gomp.On(gomp.EventTypePlayerDeath, onPlayerDeath)
	gomp.On(gomp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
	gomp.On(gomp.EventTypePlayerUpdate, onPlayerUpdate)
	gomp.On(gomp.EventTypePlayerStateChange, onPlayerStateChange)
	gomp.On(gomp.EventTypeVehicleStreamIn, onVehicleStreamIn)
}
