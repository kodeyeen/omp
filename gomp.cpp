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
        funcs.emplace("player_getName", findFunc(handle, "player_getName"));
        funcs.emplace("player_setName", findFunc(handle, "player_setName"));
        funcs.emplace("player_getPosition", findFunc(handle, "player_getPosition"));
        funcs.emplace("player_sendClientMessage", findFunc(handle, "player_sendClientMessage"));
        funcs.emplace("player_getVehicle", findFunc(handle, "player_getVehicle"));

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
