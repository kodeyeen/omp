package gomp

// #include "include/config.h"
import "C"

func SetPlayerMarkerMode(mode PlayerMarkerMode) {
	C.config_setPlayerMarkerMode(C.int(mode))
}

func SetNametagDrawRadius(radius float32) {
	C.config_setNametagDrawRadius(C.float(radius))
}

func EnableEntryExitMarkers() {
	C.config_useEntryExitMarkers(1)
}

func DisableEntryExitMarkers() {
	C.config_useEntryExitMarkers(0)
}

func EnableManualEngineAndLights() {
	C.config_useManualEngineAndLights(1)
}

func DisableManualEngineAndLights() {
	C.config_useManualEngineAndLights(0)
}

func EnableNametags() {
	C.config_useNametags(1)
}

func DisableNametags() {
	C.config_useNametags(0)
}
