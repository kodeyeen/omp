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

    void vehicle_release(void* vehicle)
    {
        return call<void>("vehicle_release", vehicle);
    }

    int vehicle_isStreamedInForPlayer(void* vehicle, void* player)
    {
        return call<int>("vehicle_isStreamedInForPlayer", vehicle, player);
    }

    void vehicle_setColour(void* vehicle, int col1, int col2)
    {
        return call<void>("vehicle_setColour", col1, col2);
    }

    VehicleColour vehicle_getColour(void* vehicle)
    {
        return call<VehicleColour>("vehicle_getColour", vehicle);
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

    void vehicle_addComponent(void* vehicle, int component)
    {
        return call<void>("vehicle_addComponent", vehicle, component);
    }

    int vehicle_getComponentInSlot(void* vehicle, int slot)
    {
        return call<int>("vehicle_getComponentInSlot", vehicle, slot);
    }

    void vehicle_removeComponent(void* vehicle, int component)
    {
        return call<void>("vehicle_removeComponent", vehicle, component);
    }

    void vehicle_putPlayer(void* vehicle, void* player, int seatID)
    {
        return call<void>("vehicle_putPlayer", vehicle, player, seatID);
    }

    void vehicle_setZAngle(void* vehicle, float angle)
    {
        return call<void>("vehicle_setZAngle", vehicle, angle);
    }

    float vehicle_getZAngle(void* vehicle)
    {
        return call<float>("vehicle_getZAngle", vehicle);
    }

    void vehicle_setParams(void* vehicle, VehicleParams* params)
    {
        return call<void>("vehicle_setParams", vehicle, params);
    }

    void vehicle_setParamsForPlayer(void* vehicle, void* player, VehicleParams* params)
    {
        return call<void>("vehicle_setParamsForPlayer", vehicle, player, params);
    }

    VehicleParams vehicle_getParams(void* vehicle)
    {
        return call<VehicleParams>("vehicle_getParams", vehicle);
    }

    int vehicle_isDead(void* vehicle)
    {
        return call<int>("vehicle_isDead", vehicle);
    }

    void vehicle_respawn(void* vehicle)
    {
        return call<void>("vehicle_respawn", vehicle);
    }

    long long vehicle_getRespawnDelay(void* vehicle)
    {
        return call<long long>("vehicle_getRespawnDelay", vehicle);
    }

    void vehicle_setRespawnDelay(void* vehicle, int delay)
    {
        return call<void>("vehicle_setRespawnDelay", vehicle, delay);
    }
    
    int vehicle_isRespawning(void* vehicle)
    {
        return call<int>("vehicle_isRespawning", vehicle);
    }

    void vehicle_setInterior(void* vehicle, int interiorID)
    {
        return call<void>("vehicle_setInterior", vehicle, interiorID);
    }

    int vehicle_getInterior(void* vehicle)
    {
        return call<int>("vehicle_getInterior", vehicle);
    }

    void vehicle_attachTrailer(void* vehicle, void* trailer)
    {
        return call<void>("vehicle_attachTrailer", vehicle, trailer);
    }

    void vehicle_detachTrailer(void* vehicle)
    {
        return call<void>("vehicle_detachTrailer", vehicle);
    }

    int vehicle_isTrailer(void* vehicle)
    {
        return call<int>("vehicle_isTrailer", vehicle);
    }

    void* vehicle_getTrailer(void* vehicle)
    {
        return call<void*>("vehicle_getTrailer", vehicle);
    }

    void* vehicle_getCab(void* vehicle)
    {
        return call<void*>("vehicle_getCab", vehicle);
    }

    void vehicle_repair(void* vehicle)
    {
        return call<void>("vehicle_repair", vehicle);
    }

    void vehicle_setVelocity(void* vehicle, float x, float y, float z)
    {
        return call<void>("vehicle_setVelocity", vehicle, x, y, z);
    }

    Vector3 vehicle_getVelocity(void* vehicle)
    {
        return call<Vector3>("vehicle_getVelocity", vehicle);
    }

    void vehicle_setAngularVelocity(void* vehicle, float x, float y, float z)
    {
        return call<void>("vehicle_setAngularVelocity", vehicle, x, y, z);
    }

    Vector3 vehicle_getAngularVelocity(void* vehicle)
    {
        return call<Vector3>("vector_getAngularVelocity", vehicle);
    }

    int vehicle_getModel(void* vehicle)
    {
        return call<int>("vehicle_getModel", vehicle);
    }

    uint8_t vehicle_getLandingGearState(void* vehicle)
    {
        return call<uint8_t>("vehicle_getLandingGearState", vehicle);
    }

    int vehicle_hasBeenOccupied(void* vehicle)
    {
        return call<int>("vehicle_hasBeenOccupied", vehicle);
    }

    long vehicle_getLastOccupiedTime(void* vehicle)
    {
        return call<long>("vehicle_getLastOccupiedTime", vehicle);
    }

    long vehicle_getLastSpawnTime(void* vehicle)
    {
        return call<long>("vehicle_getLastSpawnTime", vehicle);
    }

    int vehicle_isOccupied(void* vehicle)
    {
        return call<int>("vehicle_isOccupied", vehicle);
    }

    void vehicle_setSiren(void* vehicle, int status)
    {
        return call<void>("vehicle_setSiren", vehicle, status);
    }

    uint8_t vehicle_getSirenState(void* vehicle)
    {
        return call<uint8_t>("vehicle_getSirenState", vehicle);
    }

    uint32_t vehicle_getHydraThrustAngle(void* vehicle)
    {
        return call<uint32_t>("vehicle_getHydraThrustAngle", vehicle);
    }

    float vehicle_getTrainSpeed(void* vehicle)
    {
        return call<float>("vehicle_getTrainSpeed", vehicle);
    }

    int vehicle_getLastDriverPoolID(void* vehicle)
    {
        return call<int>("vehicle_getLastDriverPoolID", vehicle);
    }

    // entity

    Vector3 vehicle_getPosition(void* vehicle)
    {
        return call<Vector3>("vehicle_getPosition", vehicle);
    }

    void vehicle_setPosition(void* vehicle, float x, float y, float z)
    {
        return call<void>("vehicle_setPosition", vehicle, x, y, z);
    }

    Vector4 vehicle_getRotation(void* vehicle)
    {
        return call<Vector4>("vehicle_getRotation", vehicle);
    }

    void vehicle_setVirtualWorld(void* vehicle, int vw)
    {
        return call<void>("vehicle_setVirtualWorld", vehicle, vw);
    }

    int vehicle_getVirtualWorld(void* vehicle)
    {
        return call<int>("vehicle_getVirtualWorld", vehicle);
    }

    float vehicle_getDistanceFromPoint(void* vehicle, float pX, float pY, float pZ)
    {
        return call<float>("vehicle_getDistanceFromPoint", vehicle, pX, pY, pZ);
    }

    int vehicle_isInRangeOfPoint(void* vehicle, float range, float pX, float pY, float pZ)
    {
        return call<int>("vehicle_isInRangeOfPoint", vehicle, range, pX, pY, pZ);
    }

#ifdef __cplusplus
}
#endif
