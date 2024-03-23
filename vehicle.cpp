#include "include/gomp.h"
#include "include/vehicle.h"

#ifdef __cplusplus
extern "C"
{
#endif

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

#ifdef __cplusplus
}
#endif
