#include <string>
#include <unordered_map>
#include <Windows.h>

#include "component.h"

std::unordered_map<std::string, void*> funcs;

template <typename R, typename... Args>
R call(const std::string& funcName, Args... args)
{
    auto it = funcs.find(funcName);

    // if (it == funcs.end())
    // {
    //     return;
    // }

    typedef R (* FuncType)(Args...);

    FuncType func = (FuncType)it->second;

    return (*func)(std::forward<Args>(args)...);
}

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
    funcs.emplace("useManualEngineAndLights", findFunc(handle, "useManualEngineAndLights"));

    funcs.emplace("player_getID", findFunc(handle, "player_getID"));
    funcs.emplace("player_getName", findFunc(handle, "player_getName"));
    funcs.emplace("player_setName", findFunc(handle, "player_setName"));
    funcs.emplace("player_getPosition", findFunc(handle, "player_getPosition"));
    funcs.emplace("player_sendClientMessage", findFunc(handle, "player_sendClientMessage"));
    funcs.emplace("player_getVehicle", findFunc(handle, "player_getVehicle"));

    funcs.emplace("vehicle_create", findFunc(handle, "vehicle_create"));
    funcs.emplace("vehicle_setColour", findFunc(handle, "vehicle_setColour"));
    funcs.emplace("vehicle_getColour", findFunc(handle, "vehicle_getColour"));
    funcs.emplace("vehicle_setParams", findFunc(handle, "vehicle_setParams"));
    funcs.emplace("vehicle_getParams", findFunc(handle, "vehicle_getParams"));

    funcs.emplace("pickup_create", findFunc(handle, "pickup_create"));
}

// Game

void useManualEngineAndLights()
{
    return call<void>("useManualEngineAndLights");
}

// Player

int player_getID(void* player)
{
    return call<int>("player_getID", player);
}

const char* player_getName(void* player)
{
    return call<const char*>("player_getName", player);
}

int player_setName(void* player, const char* name)
{
    return call<int>("player_setName", player);
}

Vector3 player_getPosition(void* player)
{
    return call<Vector3>("player_getPosition", player);
}

void player_sendClientMessage(void* player, int colour, const char* message)
{
    return call<void>("player_sendClientMessage", player, colour, message);
}

void* player_getVehicle(void* player)
{
    return call<void*>("player_getVehicle", player);
}

// Vehicle

void* vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren)
{
    return call<void*>("vehicle_create", isStatic, modelId, x, y, z, angle, colour1, colour2, respawnDelay, addSiren);
}

void vehicle_setColour(void* vehicle, int col1, int col2)
{
    return call<void>("vehicle_setColour", col1, col2);
}

VehicleColour vehicle_getColour(void* vehicle)
{
    return call<VehicleColour>("vehicle_getColour", vehicle);
}

void vehicle_setParams(void* vehicle, VehicleParams* params)
{
    return call<void>("vehicle_setParams", vehicle, params);
}

VehicleParams vehicle_getParams(void *vehicle)
{
    return call<VehicleParams>("vehicle_getParams", vehicle);
}

// Pickup

void* pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void* player)
{
    return call<void*>("pickup_create", modelId, type, x, y, z, virtualWorld, isStatic, player);
}

#ifdef __cplusplus
}
#endif
