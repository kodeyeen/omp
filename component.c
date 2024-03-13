#include "component.h"

player_getNamePtr player_getNameFunc;
player_sendClientMessagePtr player_sendClientMessageFunc;
vehicle_createPtr vehicle_createFunc;
pickup_createPtr pickup_createFunc;

const char *player_getName(void *player)
{
    return player_getNameFunc(player);
}

void player_sendClientMessage(void *player, int colour, const char *message)
{
    return player_sendClientMessageFunc(player, colour, message);
}

void *vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren)
{
    return vehicle_createFunc(isStatic, modelId, x, y, z, angle, colour1, colour2, respawnDelay, addSiren);
}

void *pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void *player)
{
    return pickup_createFunc(modelId, type, x, y, z, virtualWorld, isStatic, player);
}

// 

void *loadLib(const char *name)
{
    return LoadLibrary((LPCTSTR)name);
}

void unloadLib(void *handle)
{
    FreeLibrary((HMODULE)handle);
}

void *findFunc(void *handle, const char *name)
{
    return GetProcAddress((HMODULE)handle, (LPCSTR)name);
}

void initFuncs(void *handle)
{
    player_getNameFunc = (player_getNamePtr)findFunc(handle, "player_getName");
    player_sendClientMessageFunc = (player_sendClientMessagePtr)findFunc(handle, "player_sendClientMessage");
    vehicle_createFunc = (vehicle_createPtr)findFunc(handle, "vehicle_create");
    pickup_createFunc = (pickup_createPtr)findFunc(handle, "pickup_create");
}