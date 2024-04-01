#ifndef API_H
#define API_H

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
    int8_t engine;
	int8_t lights;
	int8_t alarm;
	int8_t doors;
	int8_t bonnet;
	int8_t boot;
	int8_t objective;
	int8_t siren;
	int8_t doorDriver;
	int8_t doorPassenger;
	int8_t doorBackLeft;
	int8_t doorBackRight;
	int8_t windowDriver;
	int8_t windowPassenger;
	int8_t windowBackLeft;
	int8_t windowBackRight;
} VehicleParams;

typedef struct
{
    int primary;
    int secondary;
} VehicleColour;

typedef struct
{
	uint8_t id;
	uint32_t ammo;
} WeaponSlotData;

typedef struct
{
	int hours;
	int minutes;
} Time;

typedef struct
{
	uint16_t ID;
	uint16_t flags;
} PlayerAnimationData;

typedef struct
{
	uint32_t keys;
	int16_t upDown;
	int16_t leftRight;
} PlayerKeyData;

typedef struct
{
	Vector3 camFrontVector;
	Vector3 camPos;
	float aimZ;
	float camZoom;
	float aspectRatio;
	int8_t weaponState;
	uint8_t camMode;
} PlayerAimData;

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
	int spectating;
	int spectateID;
	int type;
} PlayerSpectateData;

typedef enum
{
	PlayerAnimationSyncType_NoSync,
	PlayerAnimationSyncType_Sync,
	PlayerAnimationSyncType_SyncOthers
} PlayerAnimationSyncType;

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

    void freeArray(Array* arr);

#ifdef __cplusplus
}
#endif

#endif // API_H
