#include <Windows.h>

#include "include/gomp.h"

std::unordered_map<std::string, void*> funcs;

#ifdef __cplusplus
extern "C"
{
#endif

    void* openLib(const char* path)
    {
        return LoadLibrary((LPCTSTR)path);
    }

    void* findFunc(void* handle, const char* name)
    {
        FARPROC func = GetProcAddress((HMODULE)handle, name);

        return (void*)func;
    }

    void initFuncs(void* handle)
    {
        funcs.emplace("freeArray", findFunc(handle, "freeArray"));

        funcs.emplace("useManualEngineAndLights", findFunc(handle, "useManualEngineAndLights"));

        funcs.emplace("player_getID", findFunc(handle, "player_getID"));
        funcs.emplace("player_kick", findFunc(handle, "player_kick"));
        funcs.emplace("player_ban", findFunc(handle, "player_ban"));
        funcs.emplace("player_isBot", findFunc(handle, "player_isBot"));
        funcs.emplace("player_getPing", findFunc(handle, "player_getPing"));
        funcs.emplace("player_spawn", findFunc(handle, "player_spawn"));
        funcs.emplace("player_isSpawned", findFunc(handle, "player_isSpawned"));
        funcs.emplace("player_getClientVersion", findFunc(handle, "player_getClientVersion"));
        funcs.emplace("player_getClientVersionName", findFunc(handle, "player_getClientVersionName"));
        funcs.emplace("player_setPositionFindZ", findFunc(handle, "player_setPositionFindZ"));
        funcs.emplace("player_setCameraPosition", findFunc(handle, "player_setCameraPosition"));
        funcs.emplace("player_getCameraPosition", findFunc(handle, "player_getCameraPosition"));
        funcs.emplace("player_setCameraLookAt", findFunc(handle, "player_setCameraLookAt"));
        funcs.emplace("player_getCameraLookAt", findFunc(handle, "player_getCameraLookAt"));
        funcs.emplace("player_setCameraBehind", findFunc(handle, "player_setCameraBehind"));
        funcs.emplace("player_interpolateCameraPosition", findFunc(handle, "player_interpolateCameraPosition"));
        funcs.emplace("player_interpolateCameraLookAt", findFunc(handle, "player_interpolateCameraLookAt"));
        funcs.emplace("player_attachCameraToObject", findFunc(handle, "player_attachCameraToObject"));
        funcs.emplace("player_setName", findFunc(handle, "player_setName"));
        funcs.emplace("player_getName", findFunc(handle, "player_getName"));
        funcs.emplace("player_getSerial", findFunc(handle, "player_getSerial"));
        funcs.emplace("player_giveWeapon", findFunc(handle, "player_giveWeapon"));
        funcs.emplace("player_removeWeapon", findFunc(handle, "player_removeWeapon"));
        funcs.emplace("player_setWeaponAmmo", findFunc(handle, "player_setWeaponAmmo"));
        funcs.emplace("player_getWeapons", findFunc(handle, "player_getWeapons"));
        funcs.emplace("player_getWeaponSlot", findFunc(handle, "player_getWeaponSlot"));
        funcs.emplace("player_resetWeapons", findFunc(handle, "player_resetWeapons"));
        funcs.emplace("player_setArmedWeapon", findFunc(handle, "player_setArmedWeapon"));
        funcs.emplace("player_getArmedWeapon", findFunc(handle, "player_getArmedWeapon"));
        funcs.emplace("player_getArmedWeaponAmmo", findFunc(handle, "player_getArmedWeaponAmmo"));
        funcs.emplace("player_setShopName", findFunc(handle, "player_setShopName"));
        funcs.emplace("player_getShopName", findFunc(handle, "player_getShopName"));
        funcs.emplace("player_setDrunkLevel", findFunc(handle, "player_setDrunkLevel"));
        funcs.emplace("player_getDrunkLevel", findFunc(handle, "player_getDrunkLevel"));
        funcs.emplace("player_setColour", findFunc(handle, "player_setColour"));
        funcs.emplace("player_getColour", findFunc(handle, "player_getColour"));
        funcs.emplace("player_setOtherColour", findFunc(handle, "player_setOtherColour"));
        funcs.emplace("player_getOtherColour", findFunc(handle, "player_getOtherColour"));
        funcs.emplace("player_setControllable", findFunc(handle, "player_setControllable"));
        funcs.emplace("player_getControllable", findFunc(handle, "player_getControllable"));
        funcs.emplace("player_setSpectating", findFunc(handle, "player_setSpectating"));
        funcs.emplace("player_setWantedLevel", findFunc(handle, "player_setWantedLevel"));
        funcs.emplace("player_getWantedLevel", findFunc(handle, "player_getWantedLevel"));
        funcs.emplace("player_playSound", findFunc(handle, "player_playSound"));
        funcs.emplace("player_lastPlayedSound", findFunc(handle, "player_lastPlayedSound"));
        funcs.emplace("player_playAudio", findFunc(handle, "player_playAudio"));
        funcs.emplace("player_playerCrimeReport", findFunc(handle, "player_playerCrimeReport"));
        funcs.emplace("player_stopAudio", findFunc(handle, "player_stopAudio"));
        funcs.emplace("player_lastPlayedAudio", findFunc(handle, "player_lastPlayedAudio"));
        funcs.emplace("player_createExplosion", findFunc(handle, "player_createExplosion"));
        funcs.emplace("player_sendDeathMessage", findFunc(handle, "player_sendDeathMessage"));
        funcs.emplace("player_sendEmptyDeathMessage", findFunc(handle, "player_sendEmptyDeathMessage"));
        funcs.emplace("player_removeDefaultObjects", findFunc(handle, "player_removeDefaultObjects"));
        funcs.emplace("player_forceClassSelection", findFunc(handle, "player_forceClassSelection"));
        funcs.emplace("player_setMoney", findFunc(handle, "player_setMoney"));
        funcs.emplace("player_giveMoney", findFunc(handle, "player_giveMoney"));
        funcs.emplace("player_resetMoney", findFunc(handle, "player_resetMoney"));
        funcs.emplace("player_getMoney", findFunc(handle, "player_getMoney"));
        funcs.emplace("player_setMapIcon", findFunc(handle, "player_setMapIcon"));
        funcs.emplace("player_unsetMapIcon", findFunc(handle, "player_unsetMapIcon"));
        funcs.emplace("player_useStuntBonuses", findFunc(handle, "player_useStuntBonuses"));
        funcs.emplace("player_toggleOtherNameTag", findFunc(handle, "player_toggleOtherNameTag"));
        funcs.emplace("player_setTime", findFunc(handle, "player_setTime"));
        funcs.emplace("player_getTime", findFunc(handle, "player_getTime"));
        funcs.emplace("player_useClock", findFunc(handle, "player_useClock"));
        funcs.emplace("player_useWidescreen", findFunc(handle, "player_useWidescreen"));
        funcs.emplace("player_hasWidescreen", findFunc(handle, "player_hasWidescreen"));
        funcs.emplace("player_setHealth", findFunc(handle, "player_setHealth"));
        funcs.emplace("player_getHealth", findFunc(handle, "player_getHealth"));
        funcs.emplace("player_setScore", findFunc(handle, "player_setScore"));
        funcs.emplace("player_getScore", findFunc(handle, "player_getScore"));
        funcs.emplace("player_setArmour", findFunc(handle, "player_setArmour"));
        funcs.emplace("player_getArmour", findFunc(handle, "player_getArmour"));
        funcs.emplace("player_setGravity", findFunc(handle, "player_setGravity"));
        funcs.emplace("player_getGravity", findFunc(handle, "player_getGravity"));
        funcs.emplace("player_setWorldTime", findFunc(handle, "player_setWorldTime"));
        funcs.emplace("player_applyAnimation", findFunc(handle, "player_applyAnimation"));
        funcs.emplace("player_clearAnimations", findFunc(handle, "player_clearAnimations"));
        funcs.emplace("player_getAnimationData", findFunc(handle, "player_getAnimationData"));
        funcs.emplace("player_isStreamedInForPlayer", findFunc(handle, "player_isStreamedInForPlayer"));
        funcs.emplace("player_getState", findFunc(handle, "player_getState"));
        funcs.emplace("player_setTeam", findFunc(handle, "player_setTeam"));
        funcs.emplace("player_getTeam", findFunc(handle, "player_getTeam"));
        funcs.emplace("player_setSkin", findFunc(handle, "player_setSkin"));
        funcs.emplace("player_getSkin", findFunc(handle, "player_getSkin"));
        funcs.emplace("player_setChatBubble", findFunc(handle, "player_setChatBubble"));
        funcs.emplace("player_sendClientMessage", findFunc(handle, "player_sendClientMessage"));
        funcs.emplace("player_sendChatMessage", findFunc(handle, "player_sendChatMessage"));
        funcs.emplace("player_sendCommand", findFunc(handle, "player_sendCommand"));
        funcs.emplace("player_sendGameText", findFunc(handle, "player_sendGameText"));
        funcs.emplace("player_hideGameText", findFunc(handle, "player_hideGameText"));
        funcs.emplace("player_hasGameText", findFunc(handle, "player_hasGameText"));
        funcs.emplace("player_getGameText", findFunc(handle, "player_getGameText"));
        funcs.emplace("player_setWeather", findFunc(handle, "player_setWeather"));
        funcs.emplace("player_getWeather", findFunc(handle, "player_getWeather"));
        funcs.emplace("player_setWorldBounds", findFunc(handle, "player_setWorldBounds"));
        funcs.emplace("player_getWorldBounds", findFunc(handle, "player_getWorldBounds"));
        funcs.emplace("player_setFightingStyle", findFunc(handle, "player_setFightingStyle"));
        funcs.emplace("player_getFightingStyle", findFunc(handle, "player_getFightingStyle"));
        funcs.emplace("player_setSkillLevel", findFunc(handle, "player_setSkillLevel"));
        funcs.emplace("player_setAction", findFunc(handle, "player_setAction"));
        funcs.emplace("player_getAction", findFunc(handle, "player_getAction"));
        funcs.emplace("player_setVelocity", findFunc(handle, "player_setVelocity"));
        funcs.emplace("player_getVelocity", findFunc(handle, "player_getVelocity"));
        funcs.emplace("player_setInterior", findFunc(handle, "player_setInterior"));
        funcs.emplace("player_getInterior", findFunc(handle, "player_getInterior"));
        funcs.emplace("player_getKeyData", findFunc(handle, "player_getKeyData"));
        funcs.emplace("player_getAimData", findFunc(handle, "player_getAimData"));
        funcs.emplace("player_getBulletData", findFunc(handle, "player_getBulletData"));
        funcs.emplace("player_useCameraTargetting", findFunc(handle, "player_useCameraTargetting"));
        funcs.emplace("player_hasCameraTargetting", findFunc(handle, "player_hasCameraTargetting"));
        funcs.emplace("player_removeFromVehicle", findFunc(handle, "player_removeFromVehicle"));
        funcs.emplace("player_getCameraTargetPlayer", findFunc(handle, "player_getCameraTargetPlayer"));
        funcs.emplace("player_getCameraTargetVehicle", findFunc(handle, "player_getCameraTargetVehicle"));
        funcs.emplace("player_getCameraTargetObject", findFunc(handle, "player_getCameraTargetObject"));
        funcs.emplace("player_getCameraTargetActor", findFunc(handle, "player_getCameraTargetActor"));
        funcs.emplace("player_getTargetPlayer", findFunc(handle, "player_getTargetPlayer"));
        funcs.emplace("player_getTargetActor", findFunc(handle, "player_getTargetActor"));
        funcs.emplace("player_setRemoteVehicleCollisions", findFunc(handle, "player_setRemoteVehicleCollisions"));
        funcs.emplace("player_spectatePlayer", findFunc(handle, "player_spectatePlayer"));
        funcs.emplace("player_spectateVehicle", findFunc(handle, "player_spectateVehicle"));
        funcs.emplace("player_getSpectateData", findFunc(handle, "player_getSpectateData"));
        funcs.emplace("player_sendClientCheck", findFunc(handle, "player_sendClientCheck"));
        funcs.emplace("player_toggleGhostMode", findFunc(handle, "player_toggleGhostMode"));
        funcs.emplace("player_isGhostModeEnabled", findFunc(handle, "player_isGhostModeEnabled"));
        funcs.emplace("player_getDefaultObjectsRemoved", findFunc(handle, "player_getDefaultObjectsRemoved"));
        funcs.emplace("player_getKickStatus", findFunc(handle, "player_getKickStatus"));
        funcs.emplace("player_clearTasks", findFunc(handle, "player_clearTasks"));
        funcs.emplace("player_allowWeapons", findFunc(handle, "player_allowWeapons"));
        funcs.emplace("player_areWeaponsAllowed", findFunc(handle, "player_areWeaponsAllowed"));
        funcs.emplace("player_allowTeleport", findFunc(handle, "player_allowTeleport"));
        funcs.emplace("player_isTeleportAllowed", findFunc(handle, "player_isTeleportAllowed"));
        funcs.emplace("player_isUsingOfficialClient", findFunc(handle, "player_isUsingOfficialClient"));
        funcs.emplace("player_setPosition", findFunc(handle, "player_setPosition"));
        funcs.emplace("player_getPosition", findFunc(handle, "player_getPosition"));
        funcs.emplace("player_setVirtualWorld", findFunc(handle, "player_setVirtualWorld"));
        funcs.emplace("player_getVirtualWorld", findFunc(handle, "player_getVirtualWorld"));

        funcs.emplace("player_setCheckpoint", findFunc(handle, "player_setCheckpoint"));

        funcs.emplace("player_setConsoleAccessibility", findFunc(handle, "player_setConsoleAccessibility"));
        funcs.emplace("player_hasConsoleAccess", findFunc(handle, "player_hasConsoleAccess"));
        funcs.emplace("player_getCustomSkin", findFunc(handle, "player_getCustomSkin"));
        funcs.emplace("player_getIp", findFunc(handle, "player_getIp"));
        funcs.emplace("player_getRawIp", findFunc(handle, "player_getRawIp"));
        funcs.emplace("player_getVehicle", findFunc(handle, "player_getVehicle"));
        funcs.emplace("player_getSeat", findFunc(handle, "player_getSeat"));
        funcs.emplace("player_isInModShop", findFunc(handle, "player_isInModShop"));
        funcs.emplace("player_isInDriveByMode", findFunc(handle, "player_isInDriveByMode"));
        funcs.emplace("player_isCuffed", findFunc(handle, "player_isCuffed"));
        funcs.emplace("player_getDistanceFromPoint", findFunc(handle, "player_getDistanceFromPoint"));
        funcs.emplace("player_setFacingAngle", findFunc(handle, "player_setFacingAngle"));
        funcs.emplace("player_getFacingAngle", findFunc(handle, "player_getFacingAngle"));
        funcs.emplace("player_getRotationQuat", findFunc(handle, "player_getRotationQuat"));
        funcs.emplace("player_isInRangeOfPoint", findFunc(handle, "player_isInRangeOfPoint"));

        funcs.emplace("vehicle_create", findFunc(handle, "vehicle_create"));
        funcs.emplace("vehicle_isStreamedInForPlayer", findFunc(handle, "vehicle_isStreamedInForPlayer"));
        funcs.emplace("vehicle_setColour", findFunc(handle, "vehicle_setColour"));
        funcs.emplace("vehicle_getColour", findFunc(handle, "vehicle_getColour"));
        funcs.emplace("vehicle_setHealth", findFunc(handle, "vehicle_setHealth"));
        funcs.emplace("vehicle_getHealth", findFunc(handle, "vehicle_getHealth"));
        funcs.emplace("vehicle_getDriver", findFunc(handle, "vehicle_getDriver"));
        funcs.emplace("vehicle_getPassengers", findFunc(handle, "vehicle_getPassengers"));
        funcs.emplace("vehicle_setPlate", findFunc(handle, "vehicle_setPlate"));
        funcs.emplace("vehicle_getPlate", findFunc(handle, "vehicle_getPlate"));
        funcs.emplace("vehicle_setDamageStatus", findFunc(handle, "vehicle_setDamageStatus"));
        funcs.emplace("vehicle_getDamageStatus", findFunc(handle, "vehicle_getDamageStatus"));
        funcs.emplace("vehicle_setPaintjob", findFunc(handle, "vehicle_setPaintjob"));
        funcs.emplace("vehicle_getPaintjob", findFunc(handle, "vehicle_getPaintjob"));
        funcs.emplace("vehicle_addComponent", findFunc(handle, "vehicle_addComponent"));
        funcs.emplace("vehicle_getComponentInSlot", findFunc(handle, "vehicle_getComponentInSlot"));
        funcs.emplace("vehicle_removeComponent", findFunc(handle, "vehicle_removeComponent"));
        funcs.emplace("vehicle_putPlayer", findFunc(handle, "vehicle_putPlayer"));
        funcs.emplace("vehicle_setZAngle", findFunc(handle, "vehicle_setZAngle"));
        funcs.emplace("vehicle_getZAngle", findFunc(handle, "vehicle_getZAngle"));
        funcs.emplace("vehicle_setParams", findFunc(handle, "vehicle_setParams"));
        funcs.emplace("vehicle_setParamsForPlayer", findFunc(handle, "vehicle_setParamsForPlayer"));
        funcs.emplace("vehicle_getParams", findFunc(handle, "vehicle_getParams"));
        funcs.emplace("vehicle_isDead", findFunc(handle, "vehicle_isDead"));
        funcs.emplace("vehicle_respawn", findFunc(handle, "vehicle_respawn"));
        funcs.emplace("vehicle_getRespawnDelay", findFunc(handle, "vehicle_getRespawnDelay"));
        funcs.emplace("vehicle_setRespawnDelay", findFunc(handle, "vehicle_setRespawnDelay"));
        funcs.emplace("vehicle_isRespawning", findFunc(handle, "vehicle_isRespawning"));
        funcs.emplace("vehicle_setInterior", findFunc(handle, "vehicle_setInterior"));
        funcs.emplace("vehicle_getInterior", findFunc(handle, "vehicle_getInterior"));
        funcs.emplace("vehicle_attachTrailer", findFunc(handle, "vehicle_attachTrailer"));
        funcs.emplace("vehicle_detachTrailer", findFunc(handle, "vehicle_detachTrailer"));
        funcs.emplace("vehicle_isTrailer", findFunc(handle, "vehicle_isTrailer"));
        funcs.emplace("vehicle_getTrailer", findFunc(handle, "vehicle_getTrailer"));
        funcs.emplace("vehicle_getCab", findFunc(handle, "vehicle_getCab"));
        funcs.emplace("vehicle_repair", findFunc(handle, "vehicle_repair"));
        funcs.emplace("vehicle_setVelocity", findFunc(handle, "vehicle_setVelocity"));
        funcs.emplace("vehicle_getVelocity", findFunc(handle, "vehicle_getVelocity"));
        funcs.emplace("vehicle_setAngularVelocity", findFunc(handle, "vehicle_setAngularVelocity"));
        funcs.emplace("vehicle_getAngularVelocity", findFunc(handle, "vehicle_getAngularVelocity"));
        funcs.emplace("vehicle_getModel", findFunc(handle, "vehicle_getModel"));
        funcs.emplace("vehicle_getLandingGearState", findFunc(handle, "vehicle_getLandingGearState"));
        funcs.emplace("vehicle_hasBeenOccupied", findFunc(handle, "vehicle_hasBeenOccupied"));
        funcs.emplace("vehicle_isOccupied", findFunc(handle, "vehicle_isOccupied"));
        funcs.emplace("vehicle_setSiren", findFunc(handle, "vehicle_setSiren"));
        funcs.emplace("vehicle_getSirenState", findFunc(handle, "vehicle_getSirenState"));
        funcs.emplace("vehicle_getHydraThrustAngle", findFunc(handle, "vehicle_getHydraThrustAngle"));
        funcs.emplace("vehicle_getTrainSpeed", findFunc(handle, "vehicle_getTrainSpeed"));

        funcs.emplace("pickup_create", findFunc(handle, "pickup_create"));
    }

    void freeArray(Array* arr)
    {
        return call<void>("freeArray", arr);
    }

#ifdef __cplusplus
}
#endif
