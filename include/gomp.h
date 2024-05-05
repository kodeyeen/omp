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

typedef struct
{
    int model;
    int bone;
    Vector3 offset;
    Vector3 rotation;
    Vector3 scale;
    uint32_t colour1;
    uint32_t colour2;
} PlayerAttachedObject;

typedef struct
{
	Vector3 origin;
	Vector3 hitPos;
	Vector3 offset;
	uint8_t weapon;
	uint8_t hitType;
	uint16_t hitID;
} PlayerBulletData;

typedef struct
{
	uint8_t seat;
	Vector3 position;
	Vector3 velocity;
} UnoccupiedVehicleUpdate;

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
