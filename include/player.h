#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    typedef struct {
        Vector3 camFrontVector;
        Vector3 camPos;
        float aimZ;
        float camZoom;
        float aspectRatio;
        int8_t weaponState;
        uint8_t camMode;
    } PlayerAimData;

    typedef struct {
        uint8_t id;
        uint32_t ammo;
    } WeaponSlotData;

    typedef struct {
        int hours;
        int minutes;
    } Time;

    typedef struct {
        uint16_t ID;
        uint16_t flags;
    } PlayerAnimationData;

    typedef struct {
        uint32_t keys;
        int16_t upDown;
        int16_t leftRight;
    } PlayerKeyData;

    typedef struct {
        unsigned char spectating;
        int spectateID;
        int type;
    } CPlayerSpectateData;

    typedef enum {
        PlayerAnimationSyncType_NoSync,
        PlayerAnimationSyncType_Sync,
        PlayerAnimationSyncType_SyncOthers
    } PlayerAnimationSyncType;

    void* player_getByID(int id);
    Array player_getAll();
    void player_sendDeathMessageToAll(void* killer, void* killee, int weapon);
    void player_sendEmptyDeathMessageToAll();
    void player_sendGameTextToAll(String msg, int time, int style);

    int player_getID(void* player);

    void player_kick(void* player);
    void player_ban(void* player, String reason);
    int player_isBot(void* player);
    unsigned player_getPing(void* player);
    void player_spawn(void* player);
    int player_isSpawned(void* player);
    uint8_t player_getClientVersion(void* player);
    String player_getClientVersionName(void* player);
    void player_setPositionFindZ(void* player, float x, float y, float z);
    void player_setCameraPosition(void* player, float x, float y, float z);
    Vector3 player_getCameraPosition(void* player);
    void player_setCameraLookAt(void* player, float x, float y, float z, int cutType);
    Vector3 player_getCameraLookAt(void* player);
    void player_setCameraBehind(void* player);
    void player_interpolateCameraPosition(void* player, float fromX, float fromY, float fromZ, float toX, float toY, float toZ, int time, int cutType);
    void player_interpolateCameraLookAt(void* player, float fromX, float fromY, float fromZ, float toX, float toY, float toZ, int time, int cutType);
    void player_attachCameraToObject(void* player, void* object);
    int player_setName(void* player, String name);
    String player_getName(void* player);
    String player_getSerial(void* player);
    void player_giveWeapon(void* player, uint8_t id, uint32_t ammo);
    void player_removeWeapon(void* player, uint8_t weapon);
    void player_setWeaponAmmo(void* player, uint8_t id, uint32_t ammo);
    Array* player_getWeapons(void* player);
    WeaponSlotData player_getWeaponSlot(void* player, int slot);
    void player_resetWeapons(void* player);
    void player_setArmedWeapon(void* player, uint32_t weapon);
    uint32_t player_getArmedWeapon(void* player);
    uint32_t player_getArmedWeaponAmmo(void* player);
    void player_setShopName(void* player, String name);
    String player_getShopName(void* player);
    void player_setDrunkLevel(void* player, int level);
    int player_getDrunkLevel(void* player);
    void player_setColour(void* player, uint32_t colour);
    uint32_t player_getColour(void* player);
    void player_setOtherColour(void* player, void* other, uint32_t colour);
    int player_getOtherColour(void* player, void* other, uint32_t* colour);
    void player_setControllable(void* player, int controllable);
    int player_getControllable(void* player);
    void player_setSpectating(void* player, int spectating);
    void player_setWantedLevel(void* player, unsigned level);
    unsigned player_getWantedLevel(void* player);
    void player_playSound(void* player, uint32_t sound, float posX, float posY, float posZ);
    uint32_t player_lastPlayedSound(void* player);
    void player_playAudio(void* player, String url, unsigned char usePos, float posX, float posY, float posZ, float distance);
    int player_playerCrimeReport(void* player, void* suspect, int crime);
    void player_stopAudio(void* player);
    String player_lastPlayedAudio(void* player);
    void player_createExplosion(void* player, float vecX, float vecY, float vecZ, int type, float radius);
    void player_sendDeathMessage(void* player, void* killee, void* killer, int weapon);
	void player_sendEmptyDeathMessage(void* player);
    void player_removeDefaultObjects(void* player, unsigned model, float posX, float posY, float posZ, float radius);
    void player_forceClassSelection(void* player);
    void player_setMoney(void* player, int money);
    void player_giveMoney(void* player, int money);
    void player_resetMoney(void* player);
    int player_getMoney(void* player);
    void player_setMapIcon(void* player, int id, float posX, float posY, float posZ, int type, uint32_t colour, int style);
    void player_unsetMapIcon(void* player, int id);
    void player_useStuntBonuses(void* player, int enable);
    void player_toggleOtherNameTag(void* player, void* other, int toggle);
    void player_setTime(void* player, int hr, int min);
    Time player_getTime(void* player);
    void player_useClock(void* player, int enable);
    int player_hasClock(void* player);
    void player_useWidescreen(void* player, int enable);
    int player_hasWidescreen(void* player);
    void player_setHealth(void* player, float health);
    float player_getHealth(void* player);
    void player_setScore(void* player, int score);
    int player_getScore(void* player);
    void player_setArmour(void* player, float armour);
    float player_getArmour(void* player);
    void player_setGravity(void* player, float gravity);
    float player_getGravity(void* player);
    void player_setWorldTime(void* player, int time);
    void player_applyAnimation(void* player, float delta, unsigned char loop, unsigned char lockX, unsigned char lockY, unsigned char freeze, uint32_t time, String lib, String name, int syncType);
    void player_clearAnimations(void* player, int syncType);
    PlayerAnimationData player_getAnimationData(void* player);
    int player_isStreamedInForPlayer(void* player, void* other);
    int player_getState(void* player);
    void player_setTeam(void* player, int team);
    int player_getTeam(void* player);
    void player_setSkin(void* player, int skin, int send);
    int player_getSkin(void* player);
    void player_setChatBubble(void* player, String text, uint32_t colour, float drawDist, int expire);
    void player_sendClientMessage(void* player, uint32_t colour, String message);
    void player_sendChatMessage(void* player, void* sender, String message);
    void player_sendGameText(void* player, String message, int time, int style);
    void player_hideGameText(void* player, int style);
    int player_hasGameText(void* player, int style);
    int player_getGameText(void* player, int style, String* message, int* time, int* remaining);
    void player_setWeather(void* player, int weatherID);
    int player_getWeather(void* player);
    void player_setWorldBounds(void* player, float x, float y, float z, float w);
    Vector4 player_getWorldBounds(void* player);
    void player_setFightingStyle(void* player, int style);
    int player_getFightingStyle(void* player);
    void player_setSkillLevel(void* player, int skill, int level);
    void player_setAction(void* player, int action);
    int player_getAction(void* player);
    void player_setVelocity(void* player, float velX, float velY, float velZ);
    Vector3 player_getVelocity(void* player);
    void player_setInterior(void* player, unsigned interior);
    unsigned player_getInterior(void* player);
    PlayerKeyData player_getKeyData(void* player);
    const PlayerAimData* player_getAimData(void* player);
    const PlayerBulletData* player_getBulletData(void* player);
    void player_useCameraTargeting(void* player, int enable);
    unsigned char player_hasCameraTargeting(void* player);
    void player_removeFromVehicle(void* player, unsigned char force);
    void* player_getCameraTargetPlayer(void* player);
    void* player_getCameraTargetVehicle(void* player);
    void* player_getCameraTargetObject(void* player);
    void* player_getCameraTargetActor(void* player);
    void* player_getTargetPlayer(void* player);
    void* player_getTargetActor(void* player);
    void player_setRemoteVehicleCollisions(void* player, int collide);
    void player_spectatePlayer(void* player, void* target, int mode);
    void player_spectateVehicle(void* player, void* target, int mode);
    CPlayerSpectateData player_getSpectateData(void* player);
    void player_sendClientCheck(void* player, int actionType, int address, int offset, int count);
    void player_toggleGhostMode(void* player, int toggle);
    int player_isGhostModeEnabled(void* player);
    int player_getDefaultObjectsRemoved(void* player);
    void player_clearTasks(void* player, PlayerAnimationSyncType syncType);
    void player_allowWeapons(void* player, int allow);
    int player_areWeaponsAllowed(void* player);
    void player_allowTeleport(void* player, int allow);
    int player_isTeleportAllowed(void* player);
    int player_isUsingOfficialClient(void* player);

    // entity
    void player_setPosition(void* player, float posX, float posY, float posZ);
    Vector3 player_getPosition(void* player);
    Vector4 player_getRotation(void* player);
    void player_setVirtualWorld(void* player, int vw);
    int player_getVirtualWorld(void* player);

    // console data
    void player_setConsoleAccessibility(void* player, int set);
    int player_hasConsoleAccess(void* player);

    // checkpoint data
    void* player_getCheckpoint(void* player);
    void* player_getRaceCheckpoint(void* player);

    // custom models data
    int player_getCustomSkin(void* player);
    unsigned char player_redirectDownload(void* player, String url);

    // dialog data
    void player_showDialog(void* player, int id, int style, String title, String body, String button1, String button2);
    void player_hideDialog(void* player);

    // network data
    String player_getIp(void* player);
    uint32_t player_getRawIp(void* player);

    // vehicle data
    void* player_getVehicle(void* player);
    int player_getSeat(void* player);
    int player_isInModShop(void* player);

    // object data
    void player_beginObjectEditing(void* player, void* object);
    void player_endObjectEditing(void* player);
    unsigned char player_isEditingObject(void* player);
    void player_beginObjectSelecting(void* player);
    unsigned char player_isSelectingObject(void* player);
    void player_setAttachedObject(void* player, int index, int modelId, int bone, float offsetX, float offsetY, float offsetZ, float rotX, float rotY, float rotZ, float scaleX, float scaleY, float scaleZ, uint32_t colour1, uint32_t colour2);
    PlayerAttachedObject player_getAttachedObject(void* player, int index);
    void player_removeAttachedObject(void* player, int index);
    void player_editAttachedObject(void* player, int index);
    unsigned char player_hasAttachedObject(void* player, int index);

    // misc

    float player_getDistanceFromPoint(void* player, float pX, float pY, float pZ);
    void player_setFacingAngle(void* player, float angle);
    float player_getFacingAngle(void* player);
    unsigned char player_isInRangeOfPoint(void* player, float range, float pX, float pY, float pZ);

#ifdef __cplusplus
}
#endif
