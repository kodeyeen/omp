package main

import (
	"time"

	"github.com/kodeyeen/omp"
)

const (
	RespawnDelay = 20 * time.Second

	ColorObjective omp.Color = 0xE2C063FF
	ColorGreen     omp.Color = 0x77CC77FF
	ColorBlue      omp.Color = 0x7777DDFF
)

var (
	chars              = make(map[int]*Character, 1000)
	objectiveVehGreen  omp.Vehicle
	objectiveVehBlue   omp.Vehicle
	isObjectiveReached bool
	objectiveGreenChar Character
	objectiveBlueChar  Character
)

func onGameModeInit(e *omp.GameModeInitEvent) bool {
	omp.SetGameModeText("Rivershell")

	omp.SetPlayerMarkerMode(omp.PlayerMarkerModeOff)
	omp.EnableNametags()
	omp.SetWorldTime(17)
	omp.SetWeather(11)
	omp.EnablePlayerPedAnims()
	omp.EnableVehicleFriendlyFire()
	omp.SetNametagDrawRadius(110.0)
	omp.DisableInteriorEnterExits()

	// Green classes
	omp.NewClass(0, 162, omp.Vector3{X: 2117.0129, Y: -224.4389, Z: 8.15}, 0.0, omp.WeaponM4, 100, omp.WeaponMP5, 200, omp.WeaponSniper, 10)
	omp.NewClass(0, 157, omp.Vector3{X: 2148.6606, Y: -224.3336, Z: 8.15}, 347.1396, omp.WeaponM4, 100, omp.WeaponMP5, 200, omp.WeaponSniper, 10)

	// Blue classes
	omp.NewClass(0, 154, omp.Vector3{X: 2352.9873, Y: 580.3051, Z: 7.7813}, 178.1424, omp.WeaponM4, 100, omp.WeaponMP5, 200, omp.WeaponSniper, 10)
	omp.NewClass(0, 138, omp.Vector3{X: 2281.1504, Y: 567.6248, Z: 7.7813}, 163.7289, omp.WeaponM4, 100, omp.WeaponMP5, 200, omp.WeaponSniper, 10)

	// Objective vehicles
	greenReefer, _ := NewVehicle(omp.VehicleModelReefer, omp.Vector3{X: 2184.7156, Y: -188.5401, Z: -0.0239}, 0.0000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second) // gr reefer
	blueReefer, _ := NewVehicle(omp.VehicleModelReefer, omp.Vector3{X: 2380.0542, Y: 535.2582, Z: -0.0272}, 178.4999, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)  // bl reefer

	objectiveVehGreen = *blueReefer
	objectiveVehBlue = *greenReefer

	// Green Dhingys
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2096.0833, Y: -168.7771, Z: 0.3528}, 4.5000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2103.2510, Y: -168.7598, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2099.4966, Y: -168.8216, Z: 0.3528}, 2.8200, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2107.1143, Y: -168.7798, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2111.0674, Y: -168.7609, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2114.8933, Y: -168.7898, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2167.2217, Y: -169.0570, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2170.4294, Y: -168.9724, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2173.7952, Y: -168.9217, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2177.0386, Y: -168.9767, Z: 0.3528}, 3.1800, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2161.5786, Y: -191.9538, Z: 0.3528}, 89.1000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2161.6394, Y: -187.2925, Z: 0.3528}, 89.1000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2161.7610, Y: -183.0225, Z: 0.3528}, 89.1000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2162.0283, Y: -178.5106, Z: 0.3528}, 89.1000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	// Green Mavericks
	NewVehicle(omp.VehicleModelMaverick, omp.Vector3{X: 2088.7905, Y: -227.9593, Z: 8.3662}, 0.0000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)
	NewVehicle(omp.VehicleModelMaverick, omp.Vector3{X: 2204.5991, Y: -225.3703, Z: 8.2400}, 0.0000, omp.VehicleColor{Primary: 114, Secondary: 1}, 100*time.Second)

	// Blue Dhingys
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2370.3198, Y: 518.3151, Z: 0.1240}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2362.6484, Y: 518.3978, Z: 0.0598}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2358.6550, Y: 518.2167, Z: 0.2730}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2366.5544, Y: 518.2680, Z: 0.1080}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2354.6321, Y: 518.1960, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2350.7449, Y: 518.1929, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2298.8977, Y: 518.4470, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2295.6118, Y: 518.3963, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2292.3237, Y: 518.4249, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2289.0901, Y: 518.4363, Z: 0.3597}, 180.3600, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2304.8232, Y: 539.7859, Z: 0.3597}, 270.5998, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2304.6936, Y: 535.0454, Z: 0.3597}, 270.5998, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2304.8245, Y: 530.3308, Z: 0.3597}, 270.5998, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelDinghy, omp.Vector3{X: 2304.8142, Y: 525.7471, Z: 0.3597}, 270.5998, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)

	// Blue Mavericks
	NewVehicle(omp.VehicleModelMaverick, omp.Vector3{X: 2260.2637, Y: 578.5220, Z: 8.1223}, 182.3401, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)
	NewVehicle(omp.VehicleModelMaverick, omp.Vector3{X: 2379.9792, Y: 580.0323, Z: 8.0178}, 177.9601, omp.VehicleColor{Primary: 79, Secondary: 7}, 100*time.Second)

	// Green Base Section
	omp.NewObject(9090, 500.0, omp.Vector3{X: 2148.64, Y: -222.88, Z: -20.60}, omp.Vector3{X: 0.00, Y: 0.00, Z: 179.70})
	// Green resupply hut
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, omp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})

	// Blue Base Section
	omp.NewObject(9090, 500.0, omp.Vector3{X: 2317.09, Y: 572.27, Z: -20.97}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	// Blue resupply hut
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, omp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})

	// General mapping
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, omp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	omp.NewObject(19300, 500.0, omp.Vector3{X: 2137.33, Y: -237.17, Z: 46.61}, omp.Vector3{X: 0.00, Y: 0.00, Z: 180.00})
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, omp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	omp.NewObject(19300, 500.0, omp.Vector3{X: 2325.41, Y: 587.93, Z: 47.37}, omp.Vector3{X: 0.00, Y: 0.00, Z: 180.00})
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, omp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, omp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, omp.Vector3{X: 0.00, Y: 0.00, Z: -89.94})
	omp.NewObject(12991, 500.0, omp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, omp.Vector3{X: 0.00, Y: 0.00, Z: 89.88})
	omp.NewObject(18228, 500.0, omp.Vector3{X: 1887.93, Y: -59.78, Z: -2.14}, omp.Vector3{X: 0.00, Y: 0.00, Z: 20.34})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 1990.19, Y: 541.37, Z: -22.32}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(18227, 500.0, omp.Vector3{X: 2000.82, Y: 494.15, Z: -7.53}, omp.Vector3{X: 11.70, Y: -25.74, Z: 154.38})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 1992.35, Y: 539.80, Z: -2.97}, omp.Vector3{X: 9.12, Y: 30.66, Z: 0.00})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 1991.88, Y: 483.77, Z: -0.66}, omp.Vector3{X: -2.94, Y: -5.22, Z: 12.78})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2070.57, Y: -235.87, Z: -6.05}, omp.Vector3{X: -7.20, Y: 4.08, Z: 114.30})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2056.50, Y: -228.77, Z: -19.67}, omp.Vector3{X: 14.16, Y: 19.68, Z: 106.56})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2074.00, Y: -205.33, Z: -18.60}, omp.Vector3{X: 16.02, Y: 60.60, Z: 118.86})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2230.39, Y: -242.59, Z: -11.41}, omp.Vector3{X: 5.94, Y: 7.56, Z: 471.24})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2252.53, Y: -213.17, Z: -20.81}, omp.Vector3{X: 18.90, Y: -6.30, Z: -202.38})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2233.04, Y: -234.08, Z: -19.00}, omp.Vector3{X: 21.84, Y: -8.88, Z: -252.06})
	omp.NewObject(17027, 500.0, omp.Vector3{X: 2235.05, Y: -201.49, Z: -11.90}, omp.Vector3{X: -11.94, Y: -4.08, Z: 136.32})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2226.11, Y: -237.07, Z: -2.45}, omp.Vector3{X: 8.46, Y: 2.10, Z: 471.24})
	omp.NewObject(4368, 500.0, omp.Vector3{X: 2433.79, Y: 446.26, Z: 4.67}, omp.Vector3{X: -8.04, Y: -9.30, Z: 61.02})
	omp.NewObject(4368, 500.0, omp.Vector3{X: 2031.23, Y: 489.92, Z: -13.20}, omp.Vector3{X: -8.04, Y: -9.30, Z: -108.18})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2458.36, Y: 551.10, Z: -6.95}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2465.37, Y: 511.35, Z: -7.70}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2474.80, Y: 457.71, Z: -5.17}, omp.Vector3{X: 0.00, Y: 0.00, Z: 172.74})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2466.03, Y: 426.28, Z: -5.17}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2310.45, Y: -229.38, Z: 7.41}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2294.00, Y: -180.15, Z: 7.41}, omp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2017.50, Y: -305.30, Z: 7.29}, omp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2106.45, Y: -279.86, Z: 20.05}, omp.Vector3{X: 0.00, Y: 0.00, Z: 60.90})
	omp.NewObject(706, 500.0, omp.Vector3{X: 2159.13, Y: -263.71, Z: 19.22}, omp.Vector3{X: 356.86, Y: 0.00, Z: -17.18})
	omp.NewObject(706, 500.0, omp.Vector3{X: 2055.75, Y: -291.53, Z: 13.98}, omp.Vector3{X: 356.86, Y: 0.00, Z: -66.50})
	omp.NewObject(791, 500.0, omp.Vector3{X: 1932.65, Y: -315.88, Z: 6.77}, omp.Vector3{X: 0.00, Y: 0.00, Z: -35.76})
	omp.NewObject(790, 500.0, omp.Vector3{X: 2429.40, Y: 575.79, Z: 10.42}, omp.Vector3{X: 0.00, Y: 0.00, Z: 3.14})
	omp.NewObject(790, 500.0, omp.Vector3{X: 2403.40, Y: 581.56, Z: 10.42}, omp.Vector3{X: 0.00, Y: 0.00, Z: 29.48})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2083.44, Y: 365.48, Z: 13.19}, omp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2040.15, Y: 406.02, Z: 13.33}, omp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	omp.NewObject(791, 500.0, omp.Vector3{X: 1995.36, Y: 588.10, Z: 7.50}, omp.Vector3{X: 356.86, Y: 0.00, Z: -1.95})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2126.11, Y: 595.15, Z: 5.99}, omp.Vector3{X: 0.00, Y: 0.00, Z: -35.82})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2188.35, Y: 588.90, Z: 6.04}, omp.Vector3{X: 0.00, Y: 0.00, Z: 0.00})
	omp.NewObject(791, 500.0, omp.Vector3{X: 2068.56, Y: 595.58, Z: 5.99}, omp.Vector3{X: 0.00, Y: 0.00, Z: 52.62})
	omp.NewObject(698, 500.0, omp.Vector3{X: 2385.32, Y: 606.16, Z: 9.79}, omp.Vector3{X: 0.00, Y: 0.00, Z: 34.62})
	omp.NewObject(698, 500.0, omp.Vector3{X: 2309.29, Y: 606.92, Z: 9.79}, omp.Vector3{X: 0.00, Y: 0.00, Z: -54.54})
	omp.NewObject(790, 500.0, omp.Vector3{X: 2347.14, Y: 619.77, Z: 9.94}, omp.Vector3{X: 0.00, Y: 0.00, Z: 3.14})
	omp.NewObject(698, 500.0, omp.Vector3{X: 2255.28, Y: 606.94, Z: 9.79}, omp.Vector3{X: 0.00, Y: 0.00, Z: -92.76})
	omp.NewObject(4298, 500.0, omp.Vector3{X: 2121.37, Y: 544.12, Z: -5.74}, omp.Vector3{X: -10.86, Y: 6.66, Z: 3.90})
	omp.NewObject(4368, 500.0, omp.Vector3{X: 2273.18, Y: 475.02, Z: -15.30}, omp.Vector3{X: 4.80, Y: 8.10, Z: 266.34})
	omp.NewObject(18227, 500.0, omp.Vector3{X: 2232.38, Y: 451.61, Z: -30.71}, omp.Vector3{X: -18.54, Y: -6.06, Z: 154.38})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2228.15, Y: 518.87, Z: -16.51}, omp.Vector3{X: 13.14, Y: -1.32, Z: -20.10})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2230.42, Y: 558.52, Z: -18.38}, omp.Vector3{X: -2.94, Y: -5.22, Z: 12.78})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2228.97, Y: 573.62, Z: 5.17}, omp.Vector3{X: 17.94, Y: -15.60, Z: -4.08})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2116.67, Y: -87.71, Z: -2.31}, omp.Vector3{X: 5.94, Y: 7.56, Z: 215.22})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2078.66, Y: -83.87, Z: -27.30}, omp.Vector3{X: 13.02, Y: -53.94, Z: -0.30})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2044.80, Y: -36.91, Z: -9.26}, omp.Vector3{X: -13.74, Y: 27.90, Z: 293.76})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2242.41, Y: 426.16, Z: -15.43}, omp.Vector3{X: -21.54, Y: 22.26, Z: 154.80})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2220.06, Y: 450.07, Z: -34.78}, omp.Vector3{X: -1.32, Y: 10.20, Z: -45.84})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2252.49, Y: 439.08, Z: -19.47}, omp.Vector3{X: -41.40, Y: 20.16, Z: 331.86})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2241.41, Y: 431.93, Z: -5.62}, omp.Vector3{X: -2.22, Y: -4.80, Z: 53.64})
	omp.NewObject(17029, 500.0, omp.Vector3{X: 2141.10, Y: -81.30, Z: -2.41}, omp.Vector3{X: 5.94, Y: 7.56, Z: 39.54})
	omp.NewObject(17031, 500.0, omp.Vector3{X: 2277.07, Y: 399.31, Z: -1.65}, omp.Vector3{X: -2.22, Y: -4.80, Z: -121.74})
	omp.NewObject(17026, 500.0, omp.Vector3{X: 2072.75, Y: -224.40, Z: -5.25}, omp.Vector3{X: 0.00, Y: 0.00, Z: -41.22})

	// Ramps
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2131.97, Y: 110.24, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 153.72})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2124.59, Y: 113.69, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 157.56})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2116.31, Y: 116.44, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 160.08})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2113.22, Y: 108.48, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 340.20})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2121.21, Y: 105.21, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 340.20})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2127.84, Y: 102.06, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 334.68})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2090.09, Y: 40.90, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2098.73, Y: 39.12, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2107.17, Y: 37.94, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2115.88, Y: 36.47, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 348.36})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2117.46, Y: 45.86, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 529.20})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2108.98, Y: 46.95, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 529.20})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2100.42, Y: 48.11, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 526.68})
	omp.NewObject(1632, 500.0, omp.Vector3{X: 2091.63, Y: 50.02, Z: 0.00}, omp.Vector3{X: 0.00, Y: 0.00, Z: 526.80})

	return true
}

func onPlayerConnect(e *omp.PlayerConnectEvent) bool {
	char := &Character{
		Player: e.Player,
	}

	chars[char.ID()] = char

	char.SetColor(0x888888FF)
	char.ShowGameText("~r~SA-MP: ~w~Rivershell", 2*time.Second, 5)
	char.RemoveNeededBuildings()

	return true
}

func onPlayerDisconnect(e *omp.PlayerDisconnectEvent) bool {
	delete(chars, e.Player.ID())
	return true
}

func onPlayerRequestClass(e *omp.PlayerRequestClassEvent) bool {
	char := chars[e.Player.ID()]

	char.SetupForClassSelection()
	char.SetTeamFromClass(e.Class)

	clsID := e.Class.ID()

	if clsID == 0 || clsID == 1 {
		char.ShowGameText("~g~GREEN ~w~TEAM", 1*time.Second, 5)
	} else if clsID == 2 || clsID == 3 {
		char.ShowGameText("~b~BLUE ~w~TEAM", 1*time.Second, 5)
	}

	return true
}

func onPlayerSpawn(e *omp.PlayerSpawnEvent) bool {
	char := chars[e.Player.ID()]

	if !char.LastDiedAt.IsZero() && time.Since(char.LastDiedAt) < RespawnDelay {
		char.SendClientMessage("Waiting to respawn....", 0xFFAAEEEE)
		char.EnableSpectating()

		if char.LastKiller == nil {
			return true
		}

		if char.LastKiller.State() == omp.PlayerStateOnFoot ||
			char.LastKiller.State() == omp.PlayerStateDriver ||
			char.LastKiller.State() == omp.PlayerStatePassenger {

			char.SpectateCharacter(char.LastKiller)
			char.SpectateState = SpectateStatePlayer
		}

		return true
	}

	char.SetColorFromTeam()

	if char.Team() == TeamGreen {
		char.ShowGameText("Defend the ~g~GREEN ~w~team's ~y~Reefer~n~~w~Capture the ~b~BLUE ~w~team's ~y~Reefer", 6*time.Second, 5)
	} else if char.Team() == TeamBlue {
		char.ShowGameText("Defend the ~b~BLUE ~w~team's ~y~Reefer~n~~w~Capture the ~g~GREEN ~w~team's ~y~Reefer", 6*time.Second, 5)
	}

	char.SetHealth(100.0)
	char.SetArmor(100.0)
	char.SetWorldBounds(omp.Vector4{X: 2500.0, Y: 1850.0, Z: 631.2963, W: -454.9898})

	char.SpectateState = SpectateStateNone

	return true
}

func onPlayerEnterCheckpoint(e *omp.PlayerEnterCheckpointEvent) bool {
	if isObjectiveReached {
		return true
	}

	char := chars[e.Player.ID()]

	charVeh, err := char.Vehicle()
	if err != nil {
		return true
	}

	if *charVeh == objectiveVehGreen && char.Team() == TeamGreen {
		// Green OBJECTIVE REACHED.
		captures[TeamGreen]++
		char.SetScore(char.Score() + 5)

		if captures[TeamGreen] == CapturesToWin {
			omp.ShowGameTextForAll("~g~GREEN ~w~team wins!", 3*time.Second, 5)
			isObjectiveReached = true
			PlaySoundForAll(1185, omp.Vector3{})
			time.AfterFunc(6*time.Second, ExitTheGameMode) // Set up a timer to exit this mode.
		} else {
			omp.ShowGameTextForAll("~g~GREEN ~w~team captured the ~y~boat!", 3*time.Second, 5)
			objectiveVehGreen.Respawn()
		}
	} else if *charVeh == objectiveVehBlue && char.Team() == TeamBlue {
		// Blue OBJECTIVE REACHED.
		captures[TeamBlue]++
		char.SetScore(char.Score() + 5)

		if captures[TeamBlue] == CapturesToWin {
			omp.ShowGameTextForAll("~b~BLUE ~w~team wins!", 3*time.Second, 5)
			isObjectiveReached = true
			PlaySoundForAll(1185, omp.Vector3{})
			time.AfterFunc(6*time.Second, ExitTheGameMode) // Set up a timer to exit this mode.
		} else {
			omp.ShowGameTextForAll("~b~BLUE ~w~team captured the ~y~boat!", 3*time.Second, 5)
			objectiveVehBlue.Respawn()
		}
	}

	return true
}

func onPlayerDeath(e *omp.PlayerDeathEvent) bool {
	char := chars[e.Player.ID()]
	var killer *Character

	if e.Killer == nil {
		omp.SendDeathMessage(nil, char.Player, e.Reason)
	} else {
		killer = chars[e.Killer.ID()]

		if killer.Team() != char.Team() {
			// Valid kill
			omp.SendDeathMessage(killer.Player, char.Player, e.Reason)
			killer.SetScore(killer.Score() + 1)
		} else {
			// Team kill
			omp.SendDeathMessage(killer.Player, char.Player, e.Reason)
		}
	}

	char.LastDiedAt = time.Now()
	char.LastKiller = killer

	return true
}

func onVehicleStreamIn(e *omp.VehicleStreamInEvent) bool {
	// As the vehicle streams in, player team dependant params are applied. They can't be
	// applied to vehicles that don't exist in the player's world.
	char := chars[e.ForPlayer.ID()]

	if *e.Vehicle == objectiveVehBlue {
		if char.Team() == TeamGreen {
			objectiveVehBlue.EnableObjectiveFor(char.Player)
			objectiveVehBlue.LockDoorsFor(char.Player)
		} else if char.Team() == TeamBlue {
			objectiveVehBlue.EnableObjectiveFor(char.Player)
			objectiveVehBlue.UnlockDoorsFor(char.Player)
		}
	} else if *e.Vehicle == objectiveVehGreen {
		if char.Team() == TeamBlue {
			objectiveVehGreen.EnableObjectiveFor(char.Player)
			objectiveVehGreen.LockDoorsFor(char.Player)
		} else if char.Team() == TeamGreen {
			objectiveVehGreen.EnableObjectiveFor(char.Player)
			objectiveVehGreen.UnlockDoorsFor(char.Player)
		}
	}

	return true
}

func onPlayerUpdate(e *omp.PlayerUpdateEvent) bool {
	char := chars[e.Player.ID()]

	if char.IsBot() {
		return true
	}

	if char.State() == omp.PlayerStateSpectating {
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
	if char.State() == omp.PlayerStateOnFoot {
		if char.IsInRangeOf(omp.Vector3{X: 2140.83, Y: -235.13, Z: 7.13}, 2.5) ||
			char.IsInRangeOf(omp.Vector3{X: 2318.73, Y: 590.96, Z: 6.75}, 2.5) {

			char.DoResupply()
		}
	}

	return true
}

func onPlayerStateChange(e *omp.PlayerStateChangeEvent) bool {
	char := chars[e.Player.ID()]

	if e.NewState == omp.PlayerStateDriver {
		veh, err := char.Vehicle()
		if err != nil {
			return true
		}

		if char.Team() == TeamGreen && *veh == objectiveVehGreen {
			// It's the objective vehicle
			char.SetColor(ColorObjective)
			char.ShowGameText("~w~Take the ~y~boat ~w~back to the ~r~spawn!", 3*time.Second, 5)

			cp := char.DefaultCheckpoint()
			cp.SetPosition(omp.Vector3{X: 2135.7368, Y: -179.8811, Z: -0.5323})
			cp.SetRadius(10.0)
			cp.Enable()

			objectiveGreenChar = *char
		} else if char.Team() == TeamBlue && *veh == objectiveVehBlue {
			// It's the objective vehicle
			char.SetColor(ColorObjective)
			char.ShowGameText("~w~Take the ~y~boat ~w~back to the ~r~spawn!", 3*time.Second, 5)

			cp := char.DefaultCheckpoint()
			cp.SetPosition(omp.Vector3{X: 2329.4226, Y: 532.7426, Z: 0.5862})
			cp.SetRadius(10.0)
			cp.Enable()

			objectiveBlueChar = *char
		}
	} else if e.NewState == omp.PlayerStateOnFoot {
		if *char == objectiveGreenChar {
			objectiveGreenChar = Character{}
			char.SetColorFromTeam()
			char.DefaultCheckpoint().Disable()
		} else if *char == objectiveBlueChar {
			objectiveBlueChar = Character{}
			char.SetColorFromTeam()
			char.DefaultCheckpoint().Disable()
		}
	}

	return true
}

func NewVehicle(model omp.VehicleModel, pos omp.Vector3, angle float32, color omp.VehicleColor, respawnDelay time.Duration) (*omp.Vehicle, error) {
	veh, err := omp.NewVehicle(model, pos, angle)
	if err != nil {
		return nil, err
	}

	veh.SetColor(color)
	veh.SetRespawnDelay(respawnDelay)

	return veh, nil
}

func PlaySoundForAll(sound int, pos omp.Vector3) {
	for _, player := range omp.Players() {
		player.PlaySound(sound, pos)
	}
}

func ExitTheGameMode() {
	PlaySoundForAll(1186, omp.Vector3{})
	// omp.ExitGameMode()
}

func main() {}

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.Events.Listen(omp.EventTypePlayerDisconnect, onPlayerDisconnect)
	omp.Events.Listen(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.Events.Listen(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.Events.Listen(omp.EventTypePlayerDeath, onPlayerDeath)
	omp.Events.Listen(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
	omp.Events.Listen(omp.EventTypePlayerUpdate, onPlayerUpdate)
	omp.Events.Listen(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.Events.Listen(omp.EventTypeVehicleStreamIn, onVehicleStreamIn)
}
