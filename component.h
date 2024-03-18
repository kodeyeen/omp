#ifndef COMPONENT_H
#define COMPONENT_H

#include <stdint.h>

#ifdef __cplusplus
extern "C"
{
#endif

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

void* openLib(const char *path);
void* findFunc(void *handle, const char *name);
void initFuncs(void *handle);

void useManualEngineAndLights();

int player_getID(void *player);
const char *player_getName(void *player);
int player_setName(void* player, const char *name);
Vector3 player_getPosition(void* player);
void player_sendClientMessage(void *player, int colour, const char *message);
void* player_getVehicle(void *player);

void *vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren);
void vehicle_setColour(void *vehicle, int col1, int col2);
VehicleColour vehicle_getColour(void *vehicle);
void vehicle_setParams(void *vehicle, VehicleParams *params);
VehicleParams vehicle_getParams(void *vehicle);

void *pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void *player);

#ifdef __cplusplus
}
#endif

#endif // COMPONENT_H
