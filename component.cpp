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
        funcs.emplace("freeArray", findFunc(handle, "freeArray"));

        funcs.emplace("useManualEngineAndLights", findFunc(handle, "useManualEngineAndLights"));

        funcs.emplace("player_getID", findFunc(handle, "player_getID"));
        funcs.emplace("player_getName", findFunc(handle, "player_getName"));
        funcs.emplace("player_setName", findFunc(handle, "player_setName"));
        funcs.emplace("player_getPosition", findFunc(handle, "player_getPosition"));
        funcs.emplace("player_sendClientMessage", findFunc(handle, "player_sendClientMessage"));
        funcs.emplace("player_getVehicle", findFunc(handle, "player_getVehicle"));

        funcs.emplace("vehicle_create", findFunc(handle, "vehicle_create"));
        funcs.emplace("vehicle_isStreamedInForPlayer", findFunc(handle, "vehicle_isStreamedInForPlayer"));
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

        funcs.emplace("vehicle_setColour", findFunc(handle, "vehicle_setColour"));
        funcs.emplace("vehicle_getColour", findFunc(handle, "vehicle_getColour"));
        funcs.emplace("vehicle_setParams", findFunc(handle, "vehicle_setParams"));
        funcs.emplace("vehicle_getParams", findFunc(handle, "vehicle_getParams"));

        funcs.emplace("pickup_create", findFunc(handle, "pickup_create"));
    }

    void freeArray(Array* arr)
    {
        return call<void>("freeArray", arr);
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

    String player_getName(void* player)
    {
        return call<String>("player_getName", player);
    }

    int player_setName(void* player, String name)
    {
        return call<int>("player_setName", player, name);
    }

    Vector3 player_getPosition(void* player)
    {
        return call<Vector3>("player_getPosition", player);
    }

    void player_sendClientMessage(void* player, int colour, String message)
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

    int vehicle_isStreamedInForPlayer(void* vehicle, void* player)
    {
        return call<int>("vehicle_isStreamedInForPlayer", vehicle, player);
    }

    void vehicle_setHealth(void* vehicle, float health)
    {
        return call<void>("vehicle_setHealth", vehicle, health);
    }

    float vehicle_getHealth(void* vehicle)
    {
        return call<float>("vehicle_getHealth", vehicle);
    }

    void* vehicle_getDriver(void* vehicle)
    {
        return call<void*>("vehicle_getDriver", vehicle);
    }

    Array* vehicle_getPassengers(void* vehicle)
    {
        return call<Array*>("vehicle_getPassengers", vehicle);
    }

    void vehicle_setPlate(void* vehicle, String plate)
    {
        return call<void>("vehicle_setPlate", vehicle, plate);
    }

    String vehicle_getPlate(void* vehicle)
    {
        return call<String>("vehicle_getPlate", vehicle);
    }

    void vehicle_setDamageStatus(void* vehicle, int PanelStatus, int DoorStatus, uint8_t LightStatus, uint8_t TyreStatus, void* vehicleUpdater)
    {
        return call<void>("vehicle_setDamageStatus", vehicle, PanelStatus, DoorStatus, LightStatus, TyreStatus, vehicleUpdater);
    }

    void vehicle_getDamageStatus(void* vehicle, int* PanelStatus, int* DoorStatus, int* LightStatus, int* TyreStatus)
    {
        return call<void>("vehicle_getDamageStatus", vehicle, PanelStatus, DoorStatus, LightStatus, TyreStatus);
    }

    void vehicle_setPaintjob(void* vehicle, int paintjob)
    {
        return call<void>("vehicle_setPaintjob", vehicle, paintjob);
    }

    int vehicle_getPaintjob(void* vehicle)
    {
        return call<int>("vehicle_getPaintjob", vehicle);
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

    VehicleParams vehicle_getParams(void* vehicle)
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
