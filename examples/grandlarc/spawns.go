package main

import "github.com/kodeyeen/gomp"

type Spawn struct {
	pos   gomp.Vector3
	angle float32
}

func losSantosSpawns() []Spawn {
	return []Spawn{
		{gomp.Vector3{X: 1751.1097, Y: -2106.4529, Z: 13.5469}, 183.1979}, // El-Corona - Outside random house
		{gomp.Vector3{X: 2652.6418, Y: -1989.9175, Z: 13.9988}, 182.7107}, // Random house in willowfield - near playa de seville and stadium
		{gomp.Vector3{X: 2489.5225, Y: -1957.9258, Z: 13.5881}, 2.3440},   // Hotel in willowfield - near cluckin bell
		{gomp.Vector3{X: 2689.5203, Y: -1695.9354, Z: 10.0517}, 39.5312},  // Outside stadium - lots of cars
		{gomp.Vector3{X: 2770.5393, Y: -1628.3069, Z: 12.1775}, 4.9637},   // South in east beach - north of stadium - carparks nearby
		{gomp.Vector3{X: 2807.9282, Y: -1176.8883, Z: 25.3805}, 173.6018}, // North in east beach - near apartments
		{gomp.Vector3{X: 2552.5417, Y: -958.0850, Z: 82.6345}, 280.2542},  // Random house north of Las Colinas
		{gomp.Vector3{X: 2232.1309, Y: -1159.5679, Z: 25.8906}, 103.2939}, // Jefferson motel
		{gomp.Vector3{X: 2388.1003, Y: -1279.8933, Z: 25.1291}, 94.3321},  // House south of pig pen
		{gomp.Vector3{X: 2481.1885, Y: -1536.7186, Z: 24.1467}, 273.4944}, // East LS - near clucking bell and car wash
		{gomp.Vector3{X: 2495.0720, Y: -1687.5278, Z: 13.5150}, 359.6696}, // Outside CJ's house - lots of cars nearby
		{gomp.Vector3{X: 2306.8252, Y: -1675.4340, Z: 13.9221}, 2.6271},   // House in ganton - lots of cars nearby
		{gomp.Vector3{X: 2191.8403, Y: -1455.8251, Z: 25.5391}, 267.9925}, // House in south jefferson - lots of cars nearby
		{gomp.Vector3{X: 1830.1359, Y: -1092.1849, Z: 23.8656}, 94.0113},  // Mulholland intersection carpark
		{gomp.Vector3{X: 2015.3630, Y: -1717.2535, Z: 13.5547}, 93.3655},  // Idlewood house
		{gomp.Vector3{X: 1654.7091, Y: -1656.8516, Z: 22.5156}, 177.9729}, // Right next to PD
		{gomp.Vector3{X: 1219.0851, Y: -1812.8058, Z: 16.5938}, 190.0045}, // Conference Center
		{gomp.Vector3{X: 1508.6849, Y: -1059.0846, Z: 25.0625}, 1.8058},   // Across the street of BANK - lots of cars in intersection carpark
		{gomp.Vector3{X: 1421.0819, Y: -885.3383, Z: 50.6531}, 3.6516},    // Outside house in vinewood
		{gomp.Vector3{X: 1133.8237, Y: -1272.1558, Z: 13.5469}, 192.4113}, // Near hospital
		{gomp.Vector3{X: 1235.2196, Y: -1608.6111, Z: 13.5469}, 181.2655}, // Backalley west of mainstreet
		{gomp.Vector3{X: 590.4648, Y: -1252.2269, Z: 18.2116}, 25.0473},   // Outside "BAnk of San Andreas"
		{gomp.Vector3{X: 842.5260, Y: -1007.7679, Z: 28.4185}, 213.9953},  // North of Graveyard
		{gomp.Vector3{X: 911.9332, Y: -1232.6490, Z: 16.9766}, 5.2999},    // LS Film Studio
		{gomp.Vector3{X: 477.6021, Y: -1496.6207, Z: 20.4345}, 266.9252},  // Rodeo Place
		{gomp.Vector3{X: 255.4621, Y: -1366.3256, Z: 53.1094}, 312.0852},  // Outside propery in richman
		{gomp.Vector3{X: 281.5446, Y: -1261.4562, Z: 73.9319}, 305.0017},  // Another richman property
		{gomp.Vector3{X: 790.1918, Y: -839.8533, Z: 60.6328}, 191.9514},   // Mulholland house
		{gomp.Vector3{X: 1299.1859, Y: -801.4249, Z: 84.1406}, 269.5274},  // Maddoggs
		{gomp.Vector3{X: 1240.3170, Y: -2036.6886, Z: 59.9575}, 276.4659}, // Verdant Bluffs
		{gomp.Vector3{X: 2215.5181, Y: -2627.8174, Z: 13.5469}, 273.7786}, // Ocean docks 1
		{gomp.Vector3{X: 2509.4346, Y: -2637.6543, Z: 13.6453}, 358.3565}, // Ocean Docks spawn 2
	}
}

func sanFierroSpawns() []Spawn {
	return []Spawn{
		{gomp.Vector3{X: 1435.8024, Y: 2662.3647, Z: 11.3926}, 1.1650},   //  Northern train station
		{gomp.Vector3{X: 1457.4762, Y: 2773.4868, Z: 10.8203}, 272.2754}, //  Northern golf club
		{gomp.Vector3{X: 1739.6390, Y: 2803.0569, Z: 14.2735}, 285.3929}, //  Northern housing estate 1
		{gomp.Vector3{X: 1870.3096, Y: 2785.2471, Z: 14.2734}, 42.3102},  //  Northern housing estate 2
		{gomp.Vector3{X: 1959.7142, Y: 2754.6863, Z: 10.8203}, 181.4731}, //  Northern house 1
		{gomp.Vector3{X: 2314.2556, Y: 2759.4504, Z: 10.8203}, 93.2711},  //  Northern industrial estate 1
		{gomp.Vector3{X: 2216.5674, Y: 2715.0334, Z: 10.8130}, 267.6540}, //  Northern industrial estate 2
		{gomp.Vector3{X: 2101.4192, Y: 2678.7874, Z: 10.8130}, 92.0607},  //  Northern near railway line
		{gomp.Vector3{X: 1951.1090, Y: 2660.3877, Z: 10.8203}, 180.8461}, //  Northern house 2
		{gomp.Vector3{X: 1666.6949, Y: 2604.9861, Z: 10.8203}, 179.8495}, //  Northern house 3
		{gomp.Vector3{X: 2808.3367, Y: 2421.5107, Z: 11.0625}, 136.2060}, //  Northern shopping centre
		{gomp.Vector3{X: 2633.3203, Y: 2349.7061, Z: 10.6719}, 178.7175}, //  V-Rock
		{gomp.Vector3{X: 2606.6348, Y: 2161.7490, Z: 10.8203}, 88.7508},  //  South V-Rock
		{gomp.Vector3{X: 2616.5286, Y: 2100.6226, Z: 10.8158}, 177.7834}, //  North Ammunation 1
		{gomp.Vector3{X: 2491.8816, Y: 2397.9370, Z: 10.8203}, 266.6003}, //  North carpark 1
		{gomp.Vector3{X: 2531.7891, Y: 2530.3223, Z: 21.8750}, 91.6686},  //  North carpark 2
		{gomp.Vector3{X: 2340.6677, Y: 2530.4324, Z: 10.8203}, 177.8630}, //  North Pizza Stack
		{gomp.Vector3{X: 2097.6855, Y: 2491.3313, Z: 14.8390}, 181.8117}, //  Emerald Isle
		{gomp.Vector3{X: 1893.1000, Y: 2423.2412, Z: 11.1782}, 269.4385}, //  Souvenir shop
		{gomp.Vector3{X: 1698.9330, Y: 2241.8320, Z: 10.8203}, 357.8584}, //  Northern casino
		{gomp.Vector3{X: 1479.4559, Y: 2249.0769, Z: 11.0234}, 306.3790}, //  Baseball stadium 1
		{gomp.Vector3{X: 1298.1548, Y: 2083.4016, Z: 10.8127}, 256.7034}, //  Baseball stadium 2
		{gomp.Vector3{X: 1117.8785, Y: 2304.1514, Z: 10.8203}, 81.5490},  //  North carparks
		{gomp.Vector3{X: 1108.9878, Y: 1705.8639, Z: 10.8203}, 0.6785},   //  Dirtring racing 1
		{gomp.Vector3{X: 1423.9780, Y: 1034.4188, Z: 10.8203}, 90.9590},  //  Sumo
		{gomp.Vector3{X: 1537.4377, Y: 752.0641, Z: 11.0234}, 271.6893},  //  Church
		{gomp.Vector3{X: 1917.9590, Y: 702.6984, Z: 11.1328}, 359.2682},  //  Southern housing estate
		{gomp.Vector3{X: 2089.4785, Y: 658.0414, Z: 11.2707}, 357.3572},  //  Southern house 1
		{gomp.Vector3{X: 2489.8286, Y: 928.3251, Z: 10.8280}, 67.2245},   //  Wedding chapel
		{gomp.Vector3{X: 2697.4717, Y: 856.4916, Z: 9.8360}, 267.0983},   //  Southern construction site
		{gomp.Vector3{X: 2845.6104, Y: 1288.1444, Z: 11.3906}, 3.6506},   //  Southern train station
		{gomp.Vector3{X: 2437.9370, Y: 1293.1442, Z: 10.8203}, 86.3830},  //  Wedding chapel (near Pyramid)
		{gomp.Vector3{X: 2299.5430, Y: 1451.4177, Z: 10.8203}, 269.1287}, //  Carpark (near Pyramid)
		{gomp.Vector3{X: 2214.3008, Y: 2041.9165, Z: 10.8203}, 268.7626}, //  Central parking lot
		{gomp.Vector3{X: 2005.9174, Y: 2152.0835, Z: 10.8203}, 270.1372}, //  Central motel
		{gomp.Vector3{X: 2222.1042, Y: 1837.4220, Z: 10.8203}, 88.6461},  //  Clowns Pocket
		{gomp.Vector3{X: 2025.6753, Y: 1916.4363, Z: 12.3382}, 272.5852}, //  The Visage
		{gomp.Vector3{X: 2087.9902, Y: 1516.5336, Z: 10.8203}, 48.9300},  //  Royal Casino
		{gomp.Vector3{X: 2172.1624, Y: 1398.7496, Z: 11.0625}, 91.3783},  //  Auto Bahn
		{gomp.Vector3{X: 2139.1841, Y: 987.7975, Z: 10.8203}, 0.2315},    //  Come-a-lot
		{gomp.Vector3{X: 1860.9672, Y: 1030.2910, Z: 10.8203}, 271.6988}, //  Behind 4 Dragons
		{gomp.Vector3{X: 1673.2345, Y: 1316.1067, Z: 10.8203}, 177.7294}, //  Airport carpark
		{gomp.Vector3{X: 1412.6187, Y: 2000.0596, Z: 14.7396}, 271.3568}, //  South baseball stadium houses
	}
}

func lasVenturasSpawns() []Spawn {
	return []Spawn{
		{gomp.Vector3{X: -2723.4639, Y: -314.8138, Z: 7.1839}, 43.5562},   // golf course spawn
		{gomp.Vector3{X: -2694.5344, Y: 64.5550, Z: 4.3359}, 95.0190},     // in front of a house
		{gomp.Vector3{X: -2458.2000, Y: 134.5419, Z: 35.1719}, 303.9446},  // hotel
		{gomp.Vector3{X: -2796.6589, Y: 219.5733, Z: 7.1875}, 88.8288},    // house
		{gomp.Vector3{X: -2706.5261, Y: 397.7129, Z: 4.3672}, 179.8611},   // park
		{gomp.Vector3{X: -2866.7683, Y: 691.9363, Z: 23.4989}, 286.3060},  // house
		{gomp.Vector3{X: -2764.9543, Y: 785.6434, Z: 52.7813}, 357.6817},  // donut shop
		{gomp.Vector3{X: -2660.9402, Y: 883.2115, Z: 79.7738}, 357.4440},  // house
		{gomp.Vector3{X: -2861.0796, Y: 1047.7109, Z: 33.6068}, 188.2750}, //  parking lot
		{gomp.Vector3{X: -2629.2009, Y: 1383.1367, Z: 7.1833}, 179.7006},  // parking lot at the bridge
		{gomp.Vector3{X: -2079.6802, Y: 1430.0189, Z: 7.1016}, 177.6486},  // pier
		{gomp.Vector3{X: -1660.2294, Y: 1382.6698, Z: 9.8047}, 136.2952},  //  pier 69
		{gomp.Vector3{X: -1674.1964, Y: 430.3246, Z: 7.1797}, 226.1357},   // gas station]
		{gomp.Vector3{X: -1954.9982, Y: 141.8080, Z: 27.1747}, 277.7342},  // train station
		{gomp.Vector3{X: -1956.1447, Y: 287.1091, Z: 35.4688}, 90.4465},   // car shop
		{gomp.Vector3{X: -1888.1117, Y: 615.7245, Z: 35.1719}, 128.4498},  // random
		{gomp.Vector3{X: -1922.5566, Y: 886.8939, Z: 35.3359}, 272.1293},  // random
		{gomp.Vector3{X: -1983.3458, Y: 1117.0645, Z: 53.1243}, 271.2390}, // church
		{gomp.Vector3{X: -2417.6458, Y: 970.1491, Z: 45.2969}, 269.3676},  // gas station
		{gomp.Vector3{X: -2108.0171, Y: 902.8030, Z: 76.5792}, 5.7139},    // house
		{gomp.Vector3{X: -2097.5664, Y: 658.0771, Z: 52.3672}, 270.4487},  // random
		{gomp.Vector3{X: -2263.6650, Y: 393.7423, Z: 34.7708}, 136.4152},  // random
		{gomp.Vector3{X: -2287.5027, Y: 149.1875, Z: 35.3125}, 266.3989},  // baseball parking lot
		{gomp.Vector3{X: -2039.3571, Y: -97.7205, Z: 35.1641}, 7.4744},    // driving school
		{gomp.Vector3{X: -1867.5022, Y: -141.9203, Z: 11.8984}, 22.4499},  // factory
		{gomp.Vector3{X: -1537.8992, Y: 116.0441, Z: 17.3226}, 120.8537},  // docks ship
		{gomp.Vector3{X: -1708.4763, Y: 7.0187, Z: 3.5489}, 319.3260},     // docks hangar
		{gomp.Vector3{X: -1427.0858, Y: -288.9430, Z: 14.1484}, 137.0812}, // airport
		{gomp.Vector3{X: -2173.0654, Y: -392.7444, Z: 35.3359}, 237.0159}, // stadium
		{gomp.Vector3{X: -2320.5286, Y: -180.3870, Z: 35.3135}, 179.6980}, // burger shot
		{gomp.Vector3{X: -2930.0049, Y: 487.2518, Z: 4.9141}, 3.8258},     // harbor
	}
}
