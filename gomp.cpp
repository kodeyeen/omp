#include <Windows.h>

#include "gomp.h"

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

#ifdef __cplusplus
}
#endif
