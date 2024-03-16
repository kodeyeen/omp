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

    return (void*)(func);
}

void initFuncs(void* handle)
{
    funcs.emplace("player_getName", findFunc(handle, "player_getName"));
    funcs.emplace("player_sendClientMessage", findFunc(handle, "player_sendClientMessage"));
    funcs.emplace("vehicle_create", findFunc(handle, "vehicle_create"));
    funcs.emplace("pickup_create", findFunc(handle, "pickup_create"));
}

const char* player_getName(void* player)
{
    return call<const char*>("player_getName", player);
}

void player_sendClientMessage(void* player, int colour, const char* message)
{
    return call<void>("player_sendClientMessage", player, colour, message);
}

void* vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren)
{
    return call<void*>("vehicle_create", isStatic, modelId, x, y, z, angle, colour1, colour2, respawnDelay, addSiren);
}

void* pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void* player)
{
    return call<void*>("pickup_create", modelId, type, x, y, z, virtualWorld, isStatic, player);
}

#ifdef __cplusplus
}
#endif
