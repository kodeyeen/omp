#include "include/player.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* player_getByID(int id) {
        return call<void*>("player_getByID", id);
    }

    Array player_getAll() {
        return call<Array>("player_getAll");
    }

    void player_sendDeathMessageToAll(void* killer, void* killee, int weapon) {
        return call<void>("player_sendDeathMessageToAll", killer, killee, weapon);
    }

    void player_sendEmptyDeathMessageToAll() {
        return call<void>("player_sendEmptyDeathMessageToAll");
    }

    void player_sendGameTextToAll(String msg, int time, int style) {
        return call<void>("player_sendGameTextToAll", msg, time, style);
    }

    int player_getID(void* player) {
        return call<int>("player_getID", player);
    }

    void player_kick(void* player) {
        return call<void>("player_kick", player);
    }

    void player_ban(void* player, String reason) {
        return call<void>("player_ban", player, reason);
    }

    int player_isBot(void* player) {
        return call<int>("player_isBot", player);
    }

    unsigned player_getPing(void* player) {
        return call<unsigned>("player_getPing", player);
    }

    void player_spawn(void* player) {
        return call<void>("player_spawn", player);
    }

    int player_isSpawned(void* player) {
        return call<int>("player_isSpawned", player);
    }

    uint8_t player_getClientVersion(void* player) {
        return call<uint8_t>("player_getClientVersion", player);
    }

    String player_getClientVersionName(void* player) {
        return call<String>("player_getClientVersionName", player);
    }

    void player_setPositionFindZ(void* player, float x, float y, float z) {
        return call<void>("player_setPositionFindZ", player, x, y, z);
    }

    void player_setCameraPosition(void* player, float x, float y, float z) {
        return call<void>("player_setCameraPosition", player, x, y, z);
    }

    Vector3 player_getCameraPosition(void* player) {
        return call<Vector3>("player_getCameraPosition", player);
    }

    void player_setCameraLookAt(void* player, float x, float y, float z, int cutType) {
        return call<void>("player_setCameraLookAt", player, x, y, z, cutType);
    }

    Vector3 player_getCameraLookAt(void* player) {
        return call<Vector3>("player_getCameraLookAt", player);
    }

    void player_setCameraBehind(void* player) {
        return call<void>("player_setCameraBehind", player);
    }

    void player_interpolateCameraPosition(void* player, float fromX, float fromY, float fromZ, float toX, float toY, float toZ, int time, int cutType) {
        return call<void>("player_interpolateCameraPosition", player, fromX, fromY, fromZ, toX, toY, toZ, time, cutType);
    }

    void player_interpolateCameraLookAt(void* player, float fromX, float fromY, float fromZ, float toX, float toY, float toZ, int time, int cutType) {
        return call<void>("player_interpolateCameraLookAt", player, fromX, fromY, fromZ, toX, toY, toZ, time, cutType);
    }

    void player_attachCameraToObject(void* player, void* object) {
        return call<void>("player_attachCameraToObject", player, object);
    }

    int player_setName(void* player, String name) {
        return call<int>("player_setName", player, name);
    }

    String player_getName(void* player) {
        return call<String>("player_getName", player);
    }

    String player_getSerial(void* player) {
        return call<String>("player_getSerial", player);
    }

    void player_giveWeapon(void* player, uint8_t id, uint32_t ammo) {
        return call<void>("player_giveWeapon", player, id, ammo);
    }

    void player_removeWeapon(void* player, uint8_t weapon) {
        return call<void>("player_removeWeapon", player, weapon);
    }

    void player_setWeaponAmmo(void* player, uint8_t id, uint32_t ammo) {
        return call<void>("player_setWeaponAmmo", player, id, ammo);
    }

    Array* player_getWeapons(void* player) {
        return call<Array*>("player_getWeapons", player);
    }

    WeaponSlotData player_getWeaponSlot(void* player, int slot) {
        return call<WeaponSlotData>("player_getWeaponSlot", player, slot);
    }

    void player_resetWeapons(void* player) {
        return call<void>("player_resetWeapons", player);
    }

    void player_setArmedWeapon(void* player, uint32_t weapon) {
        return call<void>("player_setArmedWeapon", player, weapon);
    }

    uint32_t player_getArmedWeapon(void* player) {
        return call<uint32_t>("player_getArmedWeapon", player);
    }

    uint32_t player_getArmedWeaponAmmo(void* player) {
        return call<uint32_t>("player_getArmedWeaponAmmo", player);
    }

    void player_setShopName(void* player, String name) {
        return call<void>("player_setShopName", player, name);
    }

    String player_getShopName(void* player) {
        return call<String>("player_getShopName", player);
    }

    void player_setDrunkLevel(void* player, int level) {
        return call<void>("player_setDrunkLevel", player, level);
    }

    int player_getDrunkLevel(void* player) {
        return call<int>("player_getDrunkLevel", player);
    }

    void player_setColour(void* player, uint32_t colour) {
        return call<void>("player_setColour", player, colour);
    }

    uint32_t player_getColour(void* player) {
        return call<uint32_t>("player_getColour", player);
    }

    void player_setOtherColour(void* player, void* other, uint32_t colour) {
        return call<void>("player_setOtherColour", player, other, colour);
    }

    int player_getOtherColour(void* player, void* other, uint32_t* colour) {
        return call<int>("player_getOtherColour", player, other, colour);
    }

    void player_setControllable(void* player, int controllable) {
        return call<void>("player_setControllable", player, controllable);
    }

    int player_getControllable(void* player) {
        return call<int>("player_getControllable", player);
    }

    void player_setSpectating(void* player, int spectating) {
        return call<void>("player_setSpectating", player, spectating);
    }

    void player_setWantedLevel(void* player, unsigned level) {
        return call<void>("player_setWantedLevel", player, level);
    }

    unsigned player_getWantedLevel(void* player) {
        return call<unsigned>("player_getWantedLevel", player);
    }

    void player_playSound(void* player, uint32_t sound, float posX, float posY, float posZ) {
        return call<void>("player_playSound", player, sound, posX, posY, posZ);
    }

    uint32_t player_lastPlayedSound(void* player) {
        return call<uint32_t>("player_lastPlayedSound", player);
    }

    void player_playAudio(void* player, String url, unsigned char usePos, float posX, float posY, float posZ, float distance) {
        return call<void>("player_playAudio", player, url, usePos, posX, posY, posZ, distance);
    }

    int player_playerCrimeReport(void* player, void* suspect, int crime) {
        return call<int>("player_playerCrimeReport", player, suspect, crime);
    }

    void player_stopAudio(void* player) {
        return call<void>("player_stopAudio", player);
    }

    String player_lastPlayedAudio(void* player) {
        return call<String>("player_lastPlayedAudio", player);
    }

    void player_createExplosion(void* player, float vecX, float vecY, float vecZ, int type, float radius) {
        return call<void>("player_createExplosion", player, vecX, vecY, vecZ, type, radius);
    }

    void player_sendDeathMessage(void* player, void* killee, void* killer, int weapon) {
        return call<void>("player_sendDeathMessage", player, killee, killer, weapon);
    }

    void player_sendEmptyDeathMessage(void* player) {
        return call<void>("player_sendEmptyDeathMessage", player);
    }

    void player_removeDefaultObjects(void* player, unsigned model, float posX, float posY, float posZ, float radius) {
        return call<void>("player_removeDefaultObjects", player, model, posX, posY, posZ, radius);
    }

    void player_forceClassSelection(void* player) {
        return call<void>("player_forceClassSelection", player);
    }

    void player_setMoney(void* player, int money) {
        return call<void>("player_setMoney", player, money);
    }

    void player_giveMoney(void* player, int money) {
        return call<void>("player_giveMoney", player, money);
    }

    void player_resetMoney(void* player) {
        return call<void>("player_resetMoney", player);
    }

    int player_getMoney(void* player) {
        return call<int>("player_getMoney", player);
    }

    void player_setMapIcon(void* player, int id, float posX, float posY, float posZ, int type, uint32_t colour, int style) {
        return call<void>("player_setMapIcon", player, id, posX, posY, posZ, type, colour, style);
    }

    void player_unsetMapIcon(void* player, int id) {
        return call<void>("player_unsetMapIcon", player, id);
    }

    void player_useStuntBonuses(void* player, int enable) {
        return call<void>("player_useStuntBonuses", player, enable);
    }

    void player_toggleOtherNameTag(void* player, void* other, int toggle) {
        return call<void>("player_toggleOtherNameTag", player, other, toggle);
    }

    void player_setTime(void* player, int hr, int min) {
        return call<void>("player_setTime", player, hr, min);
    }

    Time player_getTime(void* player) {
        return call<Time>("player_getTime", player);
    }

    void player_useClock(void* player, int enable) {
        return call<void>("player_useClock", player, enable);
    }

    int player_hasClock(void* player) {
        return call<int>("player_hasClock", player);
    }

    void player_useWidescreen(void* player, int enable) {
        return call<void>("player_useWidescreen", player, enable);
    }

    int player_hasWidescreen(void* player) {
        return call<int>("player_hasWidescreen", player);
    }

    void player_setHealth(void* player, float health) {
        return call<void>("player_setHealth", player, health);
    }

    float player_getHealth(void* player) {
        return call<float>("player_getHealth", player);
    }

    void player_setScore(void* player, int score) {
        return call<void>("player_setScore", player, score);
    }

    int player_getScore(void* player) {
        return call<int>("player_getScore", player);
    }

    void player_setArmour(void* player, float armour) {
        return call<void>("player_setArmour", player, armour);
    }

    float player_getArmour(void* player) {
        return call<float>("player_getArmour", player);
    }

    void player_setGravity(void* player, float gravity) {
        return call<void>("player_setGravity", player, gravity);
    }

    float player_getGravity(void* player) {
        return call<float>("player_getGravity", player);
    }

    void player_setWorldTime(void* player, int time) {
        return call<void>("player_setWorldTime", player, time);
    }

    void player_applyAnimation(void* player, float delta, unsigned char loop, unsigned char lockX, unsigned char lockY, unsigned char freeze, uint32_t time, String lib, String name, int syncType) {
        return call<void>("player_applyAnimation", player, delta, loop, lockX, lockY, freeze, time, lib, name, syncType);
    }

    void player_clearAnimations(void* player, int syncType) {
        return call<void>("player_clearAnimations", player, syncType);
    }

    PlayerAnimationData player_getAnimationData(void* player) {
        return call<PlayerAnimationData>("player_getAnimationData", player);
    }

    int player_isStreamedInForPlayer(void* player, void* other) {
        return call<int>("player_isStreamedInForPlayer", player, other);
    }

    int player_getState(void* player) {
        return call<int>("player_getState", player);
    }

    void player_setTeam(void* player, int team) {
        return call<void>("player_setTeam", player, team);
    }

    int player_getTeam(void* player) {
        return call<int>("player_getTeam", player);
    }

    void player_setSkin(void* player, int skin, int send) {
        return call<void>("player_setSkin", player, skin, send);
    }

    int player_getSkin(void* player) {
        return call<int>("player_getSkin", player);
    }

    void player_setChatBubble(void* player, String text, uint32_t colour, float drawDist, int expire) {
        return call<void>("player_setChatBubble", player, text, colour, drawDist, expire);
    }

    void player_sendClientMessage(void* player, uint32_t colour, String message) {
        return call<void>("player_sendClientMessage", player, colour, message);
    }

    void player_sendChatMessage(void* player, void* sender, String message) {
        return call<void>("player_sendChatMessage", player, sender, message);
    }

    void player_sendGameText(void* player, String message, int time, int style) {
        return call<void>("player_sendGameText", player, message, time, style);
    }

    void player_hideGameText(void* player, int style) {
        return call<void>("player_hideGameText", player, style);
    }

    int player_hasGameText(void* player, int style) {
        return call<int>("player_hasGameText", player, style);
    }

    int player_getGameText(void* player, int style, String* message, int* time, int* remaining) {
        return call<int>("player_getGameText", player, style, message, time, remaining);
    }

    void player_setWeather(void* player, int weatherID) {
        return call<void>("player_setWeather", player, weatherID);
    }

    int player_getWeather(void* player) {
        return call<int>("player_getWeather", player);
    }

    void player_setWorldBounds(void* player, float x, float y, float z, float w) {
        return call<void>("player_setWorldBounds", player, x, y, z, w);
    }

    Vector4 player_getWorldBounds(void* player) {
        return call<Vector4>("player_getWorldBounds", player);
    }

    void player_setFightingStyle(void* player, int style) {
        return call<void>("player_setFightingStyle", player, style);
    }

    int player_getFightingStyle(void* player) {
        return call<int>("player_getFightingStyle", player);
    }

    void player_setSkillLevel(void* player, int skill, int level) {
        return call<void>("player_setSkillLevel", player, skill, level);
    }

    void player_setAction(void* player, int action) {
        return call<void>("player_setAction", player, action);
    }

    int player_getAction(void* player) {
        return call<int>("player_getAction", player);
    }

    void player_setVelocity(void* player, float velX, float velY, float velZ) {
        return call<void>("player_setVelocity", player, velX, velY, velZ);
    }

    Vector3 player_getVelocity(void* player) {
        return call<Vector3>("player_getVelocity", player);
    }

    void player_setInterior(void* player, unsigned interior) {
        return call<void>("player_setInterior", player, interior);
    }

    unsigned player_getInterior(void* player) {
        return call<unsigned>("player_getInterior", player);
    }

    PlayerKeyData player_getKeyData(void* player) {
        return call<PlayerKeyData>("player_getKeyData", player);
    }

    const PlayerAimData* player_getAimData(void* player) {
        return call<const PlayerAimData*>("player_getAimData", player);
    }

    const PlayerBulletData* player_getBulletData(void* player) {
        return call<const PlayerBulletData*>("player_getBulletData", player);
    }

    void player_useCameraTargeting(void* player, int enable) {
        return call<void>("player_useCameraTargeting", player, enable);
    }

    unsigned char player_hasCameraTargeting(void* player) {
        return call<unsigned char>("player_hasCameraTargeting", player);
    }

    void player_removeFromVehicle(void* player, unsigned char force) {
        return call<void>("player_removeFromVehicle", player, force);
    }

    void* player_getCameraTargetPlayer(void* player) {
        return call<void*>("player_getCameraTargetPlayer", player);
    }

    void* player_getCameraTargetVehicle(void* player) {
        return call<void*>("player_getCameraTargetVehicle", player);
    }

    void* player_getCameraTargetObject(void* player) {
        return call<void*>("player_getCameraTargetObject", player);
    }

    void* player_getCameraTargetActor(void* player) {
        return call<void*>("player_getCameraTargetActor", player);
    }

    void* player_getTargetPlayer(void* player) {
        return call<void*>("player_getTargetPlayer", player);
    }

    void* player_getTargetActor(void* player) {
        return call<void*>("player_getTargetActor", player);
    }

    void player_setRemoteVehicleCollisions(void* player, int collide) {
        return call<void>("player_setRemoteVehicleCollisions", player, collide);
    }

    void player_spectatePlayer(void* player, void* target, int mode) {
        return call<void>("player_spectatePlayer", player, target, mode);
    }

    void player_spectateVehicle(void* player, void* target, int mode) {
        return call<void>("player_spectateVehicle", player, target, mode);
    }

    CPlayerSpectateData player_getSpectateData(void* player) {
        return call<CPlayerSpectateData>("player_getSpectateData", player);
    }

    void player_sendClientCheck(void* player, int actionType, int address, int offset, int count) {
        return call<void>("player_sendClientCheck", player, actionType, address, offset, count);
    }

    void player_toggleGhostMode(void* player, int toggle) {
        return call<void>("player_toggleGhostMode", player, toggle);
    }

    int player_isGhostModeEnabled(void* player) {
        return call<int>("player_isGhostModeEnabled", player);
    }

    int player_getDefaultObjectsRemoved(void* player) {
        return call<int>("player_getDefaultObjectsRemoved", player);
    }

    void player_clearTasks(void* player, PlayerAnimationSyncType syncType) {
        return call<void>("player_clearTasks", player, syncType);
    }

    void player_allowWeapons(void* player, int allow) {
        return call<void>("player_allowWeapons", player, allow);
    }

    int player_areWeaponsAllowed(void* player) {
        return call<int>("player_areWeaponsAllowed", player);
    }

    void player_allowTeleport(void* player, int allow) {
        return call<void>("player_allowTeleport", player, allow);
    }

    int player_isTeleportAllowed(void* player) {
        return call<int>("player_isTeleportAllowed", player);
    }

    int player_isUsingOfficialClient(void* player) {
        return call<int>("player_isUsingOfficialClient", player);
    }

    void player_setPosition(void* player, float posX, float posY, float posZ) {
        return call<void>("player_setPosition", player, posX, posY, posZ);
    }

    Vector3 player_getPosition(void* player) {
        return call<Vector3>("player_getPosition", player);
    }

    Vector4 player_getRotation(void* player) {
        return call<Vector4>("player_getRotation", player);
    }

    void player_setVirtualWorld(void* player, int vw) {
        return call<void>("player_setVirtualWorld", player, vw);
    }

    int player_getVirtualWorld(void* player) {
        return call<int>("player_getVirtualWorld", player);
    }

    void player_setConsoleAccessibility(void* player, int set) {
        return call<void>("player_setConsoleAccessibility", player, set);
    }

    int player_hasConsoleAccess(void* player) {
        return call<int>("player_hasConsoleAccess", player);
    }

    void* player_getCheckpoint(void* player) {
        return call<void*>("player_getCheckpoint", player);
    }

    void* player_getRaceCheckpoint(void* player) {
        return call<void*>("player_getRaceCheckpoint", player);
    }

    int player_getCustomSkin(void* player) {
        return call<int>("player_getCustomSkin", player);
    }

    unsigned char player_redirectDownload(void* player, String url) {
        return call<unsigned char>("player_redirectDownload", player, url);
    }

    void player_showDialog(void* player, int id, int style, String title, String body, String button1, String button2) {
        return call<void>("player_showDialog", player, id, style, title, body, button1, button2);
    }

    void player_hideDialog(void* player) {
        return call<void>("player_hideDialog", player);
    }

    String player_getIp(void* player) {
        return call<String>("player_getIp", player);
    }

    uint32_t player_getRawIp(void* player) {
        return call<uint32_t>("player_getRawIp", player);
    }

    void* player_getVehicle(void* player) {
        return call<void*>("player_getVehicle", player);
    }

    int player_getSeat(void* player) {
        return call<int>("player_getSeat", player);
    }

    int player_isInModShop(void* player) {
        return call<int>("player_isInModShop", player);
    }

    void player_beginObjectEditing(void* player, void* object) {
        return call<void>("player_beginObjectEditing", player, object);
    }

    void player_endObjectEditing(void* player) {
        return call<void>("player_endObjectEditing", player);
    }

    unsigned char player_isEditingObject(void* player) {
        return call<unsigned char>("player_isEditingObject", player);
    }

    void player_beginObjectSelecting(void* player) {
        return call<void>("player_beginObjectSelecting", player);
    }

    unsigned char player_isSelectingObject(void* player) {
        return call<unsigned char>("player_isSelectingObject", player);
    }

    void player_setAttachedObject(void* player, int index, int modelId, int bone, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ, float scaleX, float scaleY, float scaleZ, uint32_t colour1, uint32_t colour2) {
        return call<void>("player_setAttachedObject", player, index, modelId, bone, offsetX, offsetY, offsetZ, rotX, rotY, rotZ, scaleX, scaleY, scaleZ, colour1, colour2);
    }

    PlayerAttachedObject player_getAttachedObject(void* player, int index) {
        return call<PlayerAttachedObject>("player_getAttachedObject", player, index);
    }

    void player_removeAttachedObject(void* player, int index) {
        return call<void>("player_removeAttachedObject", player, index);
    }

    void player_editAttachedObject(void* player, int index) {
        return call<void>("player_editAttachedObject", player, index);
    }

    unsigned char player_hasAttachedObject(void* player, int index) {
        return call<unsigned char>("player_hasAttachedObject", player, index);
    }

    float player_getDistanceFromPoint(void* player, float pX, float pY, float pZ) {
        return call<float>("player_getDistanceFromPoint", player, pX, pY, pZ);
    }

    void player_setFacingAngle(void* player, float angle) {
        return call<void>("player_setFacingAngle", player, angle);
    }

    float player_getFacingAngle(void* player) {
        return call<float>("player_getFacingAngle", player);
    }

    unsigned char player_isInRangeOfPoint(void* player, float range, float pX, float pY, float pZ) {
        return call<unsigned char>("player_isInRangeOfPoint", player, range, pX, pY, pZ);
    }

#ifdef __cplusplus
}
#endif
