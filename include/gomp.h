#ifndef GOMP_H
#define GOMP_H

#include <stdint.h>

typedef struct
{
    const char* buf;
    size_t length;
} String;

typedef struct
{
	void** buf;
	size_t length;
} Array;

typedef struct
{
    float x;
    float y;
    float z;
    float w;
} Vector4;

typedef struct
{
    float x;
    float y;
    float z;
} Vector3;

typedef struct
{
    float x;
    float y;
} Vector2;

typedef struct
{
    int primary;
    int secondary;
} VehicleColour;

#ifdef __cplusplus

#include <string>
#include <unordered_map>

extern std::unordered_map<std::string, void*> funcs;

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


extern "C"
{
#endif

    void* openLib(const char* path);
    void* findFunc(void* handle, const char* name);
    void initFuncs(void* handle);
    void freeArray(Array arr);

#ifdef __cplusplus
}
#endif

#endif // GOMP_H
