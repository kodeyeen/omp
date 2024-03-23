#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

void* vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, int addSiren);
int vehicle_isStreamedInForPlayer(void* vehicle, void* player);
void vehicle_setHealth(void* vehicle, float health);
float vehicle_getHealth(void* vehicle);
void* vehicle_getDriver(void* vehicle);
Array* vehicle_getPassengers(void* vehicle);
void vehicle_setPlate(void* vehicle, String plate);
String vehicle_getPlate(void* vehicle);
void vehicle_setDamageStatus(void* vehicle, int PanelStatus, int DoorStatus, uint8_t LightStatus, uint8_t TyreStatus, void* vehicleUpdater);
void vehicle_getDamageStatus(void* vehicle, int* PanelStatus, int* DoorStatus, int* LightStatus, int* TyreStatus);
void vehicle_setPaintjob(void* vehicle, int paintjob);
int vehicle_getPaintjob(void* vehicle);

void vehicle_setColour(void* vehicle, int col1, int col2);
VehicleColour vehicle_getColour(void* vehicle);
void vehicle_setParams(void* vehicle, VehicleParams *params);
VehicleParams vehicle_getParams(void* vehicle);

#ifdef __cplusplus
}
#endif
