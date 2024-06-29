package omp

// #include "include/config.h"
import "C"

const (
	PlayerMarkerModeOff = iota
	PlayerMarkerModeGlobal
	PlayerMarkerModeStreamed
)

type ConfigOptionValue interface {
	float64 | int | bool
}

func SetConfigOption[T ConfigOptionValue](key string, value T) {
	cKey := newCString(key)
	defer freeCString(cKey)

	switch v := any(value).(type) {
	case float64:
		C.config_setFloat(cKey, C.float(v))
	case int:
		C.config_setInt(cKey, C.int(v))
	case bool:
		C.config_setBool(cKey, newCUchar(v))
	}
}

func ConfigOption[T ConfigOptionValue](key string) T {
	cKey := newCString(key)
	defer freeCString(cKey)

	var t T
	var result any

	switch any(t).(type) {
	case float64:
		result = float64(C.config_getFloat(cKey))
	case int:
		result = int(C.config_getInt(cKey))
	case bool:
		result = bool(C.config_getBool(cKey) != 0)
	}

	return result.(T)
}

func AllowAdminTeleport() {
	SetConfigOption("rcon.allow_teleport", true)
}

func DisallowAdminTeleport() {
	SetConfigOption("rcon.allow_teleport", false)
}

func IsAdminTeleportAllowed() bool {
	return ConfigOption[bool]("rcon.allow_teleport")
}

func AreInteriorWeaponsAllowed() bool {
	return ConfigOption[bool]("game.allow_interior_weapons")
}

func EnableInteriorEnterExits() {
	SetConfigOption("game.use_entry_exit_markers", true)
}

func DisableInteriorEnterExits() {
	SetConfigOption("game.use_entry_exit_markers", false)
}

func AreInteriorEnterExitsEnabled() bool {
	return ConfigOption[bool]("game.use_entry_exit_markers")
}

func EnableNametagLOS() {
	SetConfigOption("game.use_nametag_los", true)
}

func DisableNametagLOS() {
	SetConfigOption("game.use_nametag_los", false)
}

func IsNametagLOSEnabled() bool {
	return ConfigOption[bool]("game.use_nametag_los")
}

func EnableAllAnimations() {
	SetConfigOption("game.use_all_animations", true)
}

func DisableAllAnimations() {
	SetConfigOption("game.use_all_animations", false)
}

func AreAllAnimationsEnabled() bool {
	return ConfigOption[bool]("game.use_all_animations")
}

func EnableVehicleFriendlyFire() {
	SetConfigOption("game.use_vehicle_friendly_fire", true)
}

func DisableVehicleFriendlyFire() {
	SetConfigOption("game.use_vehicle_friendly_fire", false)
}

func IsVehicleFriendlyFireEnabled() bool {
	return ConfigOption[bool]("game.use_vehicle_friendly_fire")
}

func EnableZoneNames() {
	SetConfigOption("game.use_zone_names", true)
}

func DisableZoneNames() {
	SetConfigOption("game.use_zone_names", false)
}

func AreZoneNamesEnabled() bool {
	return ConfigOption[bool]("game.use_zone_names")
}

func MaxPlayers() int {
	return ConfigOption[int]("max_players")
}

func Weather() int {
	return ConfigOption[int]("game.weather")
}

func WorldTime() int {
	return ConfigOption[int]("game.time")
}

func IsIPBanned(IP string) bool {
	cIP := newCString(IP)
	defer freeCString(cIP)

	return C.config_isBanned(cIP) != 0
}

func LimitGlobalChatRadius(radius float64) {
	SetConfigOption("game.use_chat_radius", true)
	SetConfigOption("game.chat_radius", radius)
}

func LimitPlayerMarkerRadius(radius float64) {
	SetConfigOption("game.use_player_marker_draw_radius", true)
	SetConfigOption("game.player_marker_draw_radius", radius)
}

func EnableManualEngineAndLights() {
	SetConfigOption("game.use_manual_engine_and_lights", true)
}

func DisableManualEngineAndLights() {
	SetConfigOption("game.use_manual_engine_and_lights", true)
}

func IsManualEngineAndLightsEnabled() bool {
	return ConfigOption[bool]("game.use_manual_engine_and_lights")
}

func SetNametagDrawRadius(radius float64) {
	SetConfigOption("game.nametag_draw_radius", radius)
}

func EnableNametags() {
	SetConfigOption("game.use_nametags", true)
}

func DisableNametags() {
	SetConfigOption("game.use_nametags", true)
}

func IsNametagsEnabled() bool {
	return ConfigOption[bool]("game.use_nametags")
}

func SetPlayerMarkerMode(mode int) {
	SetConfigOption("game.player_marker_mode", mode)
}

func PlayerMarkerMode() int {
	return ConfigOption[int]("game.player_marker_mode")
}

func EnableChatInputFilter() {
	SetConfigOption("chat_input_filter", true)
}

func DisableChatInputFilter() {
	SetConfigOption("chat_input_filter", false)
}

func IsChatInputFilterEnabled() bool {
	return ConfigOption[bool]("chat_input_filter")
}

func EnablePlayerPedAnims() {
	SetConfigOption("game.use_player_ped_anims", true)
}

func DisablePlayerPedAnims() {
	SetConfigOption("game.use_player_ped_anims", false)
}

func ArePlayerPedAnimsEnabled() bool {
	return ConfigOption[bool]("game.use_player_ped_anims")
}
