package omp

// #include "include/config.h"
import "C"

const (
	PlayerMarkerModeOff = iota
	PlayerMarkerModeGlobal
	PlayerMarkerModeStreamed
)

type configOptionValue interface {
	float64 | int | bool
}

func setConfigOption[T configOptionValue](key string, value T) {
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

func configOption[T configOptionValue](key string) T {
	cKey := newCString(key)
	defer freeCString(cKey)

	t := any(new(T))
	var result any

	switch t.(type) {
	case float64:
		result = C.config_getFloat(cKey)
	case int:
		result = C.config_getInt(cKey)
	case bool:
		result = C.config_getBool(cKey)
	}

	return result.(T)
}

func AllowAdminTeleport() {
	setConfigOption("rcon.allow_teleport", true)
}

func DisallowAdminTeleport() {
	setConfigOption("rcon.allow_teleport", false)
}

func IsAdminTeleportAllowed() bool {
	return configOption[bool]("rcon.allow_teleport")
}

func AreInteriorWeaponsAllowed() bool {
	return configOption[bool]("game.allow_interior_weapons")
}

func EnableInteriorEnterExits() {
	setConfigOption("game.use_entry_exit_markers", true)
}

func DisableInteriorEnterExits() {
	setConfigOption("game.use_entry_exit_markers", false)
}

func AreInteriorEnterExitsEnabled() bool {
	return configOption[bool]("game.use_entry_exit_markers")
}

func EnableNametagLOS() {
	setConfigOption("game.use_nametag_los", true)
}

func DisableNametagLOS() {
	setConfigOption("game.use_nametag_los", false)
}

func IsNametagLOSEnabled() bool {
	return configOption[bool]("game.use_nametag_los")
}

func EnableAllAnimations() {
	setConfigOption("game.use_all_animations", true)
}

func DisableAllAnimations() {
	setConfigOption("game.use_all_animations", false)
}

func AreAllAnimationsEnabled() bool {
	return configOption[bool]("game.use_all_animations")
}

func EnableVehicleFriendlyFire() {
	setConfigOption("game.use_vehicle_friendly_fire", true)
}

func DisableVehicleFriendlyFire() {
	setConfigOption("game.use_vehicle_friendly_fire", false)
}

func IsVehicleFriendlyFireEnabled() bool {
	return configOption[bool]("game.use_vehicle_friendly_fire")
}

func EnableZoneNames() {
	setConfigOption("game.use_zone_names", true)
}

func DisableZoneNames() {
	setConfigOption("game.use_zone_names", false)
}

func AreZoneNamesEnabled() bool {
	return configOption[bool]("game.use_zone_names")
}

func MaxPlayers() int {
	return configOption[int]("max_players")
}

func Weather() int {
	return configOption[int]("game.weather")
}

func WorldTime() int {
	return configOption[int]("game.time")
}

func IsIPBanned(IP string) bool {
	cIP := newCString(IP)
	defer freeCString(cIP)

	return C.config_isBanned(cIP) != 0
}

func LimitGlobalChatRadius(radius float64) {
	setConfigOption("game.use_chat_radius", true)
	setConfigOption("game.chat_radius", radius)
}

func LimitPlayerMarkerRadius(radius float64) {
	setConfigOption("game.use_player_marker_draw_radius", true)
	setConfigOption("game.player_marker_draw_radius", radius)
}

func EnableManualEngineAndLights() {
	setConfigOption("game.use_manual_engine_and_lights", true)
}

func DisableManualEngineAndLights() {
	setConfigOption("game.use_manual_engine_and_lights", true)
}

func IsManualEngineAndLightsEnabled() bool {
	return configOption[bool]("game.use_manual_engine_and_lights")
}

func SetNametagDrawRadius(radius float64) {
	setConfigOption("game.nametag_draw_radius", radius)
}

func EnableNametags() {
	setConfigOption("game.use_nametags", true)
}

func DisableNametags() {
	setConfigOption("game.use_nametags", true)
}

func IsNametagsEnabled() bool {
	return configOption[bool]("game.use_nametags")
}

func SetPlayerMarkerMode(mode int) {
	setConfigOption("game.player_marker_mode", mode)
}

func PlayerMarkerMode() int {
	return configOption[int]("game.player_marker_mode")
}

func EnableChatInputFilter() {
	setConfigOption("chat_input_filter", true)
}

func DisableChatInputFilter() {
	setConfigOption("chat_input_filter", false)
}

func IsChatInputFilterEnabled() bool {
	return configOption[bool]("chat_input_filter")
}

func EnablePlayerPedAnims() {
	setConfigOption("game.use_player_ped_anims", true)
}

func DisablePlayerPedAnims() {
	setConfigOption("game.use_player_ped_anims", false)
}

func ArePlayerPedAnimsEnabled() bool {
	return configOption[bool]("game.use_player_ped_anims")
}
