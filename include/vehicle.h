#include "omp.h"

#ifdef __cplusplus
extern "C"
{
#endif

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

    void* vehicle_create(int isStatic, int modelId, float x, float y, float z, float angle, int colour1, int colour2, int respawnDelay, unsigned char addSiren);
    void vehicle_release(void* vehicle);
    void* vehicle_getByID(int id);
    int vehicle_isStreamedInForPlayer(void* vehicle, void* player);
    void vehicle_setColour(void* vehicle, int col1, int col2);
    VehicleColour vehicle_getColour(void* vehicle);
    void vehicle_setHealth(void* vehicle, float health);
    float vehicle_getHealth(void* vehicle);
    void* vehicle_getDriver(void* vehicle);
    Array vehicle_getPassengers(void* vehicle);
    void vehicle_setPlate(void* vehicle, String plate);
    String vehicle_getPlate(void* vehicle);
    void vehicle_setDamageStatus(void* vehicle, int PanelStatus, int DoorStatus, uint8_t LightStatus, uint8_t TyreStatus, void* vehicleUpdater);
    void vehicle_getDamageStatus(void* vehicle, int* PanelStatus, int* DoorStatus, int* LightStatus, int* TyreStatus);
    void vehicle_setPaintjob(void* vehicle, int paintjob);
    int vehicle_getPaintjob(void* vehicle);
    void vehicle_addComponent(void* vehicle, int component);
    int vehicle_getComponentInSlot(void* vehicle, int slot);
    void vehicle_removeComponent(void* vehicle, int component);
    void vehicle_putPlayer(void* vehicle, void* player, int seatID);
    void vehicle_setZAngle(void* vehicle, float angle);
    float vehicle_getZAngle(void* vehicle);
    void vehicle_setParams(void* vehicle, VehicleParams* params);
    void vehicle_setParamsForPlayer(void* vehicle, void* player, VehicleParams* params);
    VehicleParams vehicle_getParams(void* vehicle);
    int vehicle_isDead(void* vehicle);
    void vehicle_respawn(void* vehicle);
    void vehicle_setRespawnDelay(void* vehicle, int delay);
    long long vehicle_getRespawnDelay(void* vehicle);
    int vehicle_isRespawning(void* vehicle);
    void vehicle_setInterior(void* vehicle, int interiorID);
    int vehicle_getInterior(void* vehicle);
    void vehicle_attachTrailer(void* vehicle, void* trailer);
    void vehicle_detachTrailer(void* vehicle);
    int vehicle_isTrailer(void* vehicle);
    void* vehicle_getTrailer(void* vehicle);
    void* vehicle_getCab(void* vehicle);
    void vehicle_repair(void* vehicle);
    void vehicle_setVelocity(void* vehicle, float x, float y, float z);
    Vector3 vehicle_getVelocity(void* vehicle);
    void vehicle_setAngularVelocity(void* vehicle, float x, float y, float z);
    Vector3 vehicle_getAngularVelocity(void* vehicle);
    int vehicle_getModel(void* vehicle);
    uint8_t vehicle_getLandingGearState(void* vehicle);
    int vehicle_hasBeenOccupied(void* vehicle);
    long vehicle_getLastOccupiedTime(void* vehicle);
    long vehicle_getLastSpawnTime(void* vehicle);
    int vehicle_isOccupied(void* vehicle);
    void vehicle_setSiren(void* vehicle, int status);
    uint8_t vehicle_getSirenState(void* vehicle);
    uint32_t vehicle_getHydraThrustAngle(void* vehicle);
    float vehicle_getTrainSpeed(void* vehicle);
    int vehicle_getLastDriverPoolID(void* vehicle);

    Vector3 vehicle_getPosition(void* vehicle);
    void vehicle_setPosition(void* vehicle, float x, float y, float z);
    Vector4 vehicle_getRotation(void* vehicle);
    void vehicle_setVirtualWorld(void* vehicle, int vw);
    int vehicle_getVirtualWorld(void* vehicle);

    float vehicle_getDistanceFromPoint(void* vehicle, float pX, float pY, float pZ);
    int vehicle_isInRangeOfPoint(void* vehicle, float range, float pX, float pY, float pZ);

#ifdef __cplusplus
}
#endif
