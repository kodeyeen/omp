#include "include/pickup.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* pickup_create(int modelId, uint8_t type, float posX, float posY, float posZ, uint32_t virtualWorld, int isStatic) {
        return call<void*>("pickup_create", modelId, type, posX, posY, posZ, virtualWorld, isStatic);
    }

    void pickup_release(void* pickup) {
        return call<void>("pickup_release", pickup);
    }

    int pickup_getID(void* pickup) {
        return call<int>("pickup_getID", pickup);
    }

    void pickup_setType(void* pickup, uint8_t type) {
        return call<void>("pickup_setType", pickup, type);
    }

    uint8_t pickup_getType(void* pickup) {
        return call<uint8_t>("pickup_getType", pickup);
    }

    void pickup_setModel(void* pickup, int model) {
        return call<void>("pickup_setModel", pickup, model);
    }

    int pickup_getModel(void* pickup) {
        return call<int>("pickup_getModel", pickup);
    }

    unsigned char pickup_isStreamedInForPlayer(void* pickup, void* player) {
        return call<unsigned char>("pickup_isStreamedInForPlayer", pickup, player);
    }

    void pickup_setPickupHiddenForPlayer(void* pickup, void* player, unsigned char hidden) {
        return call<void>("pickup_setPickupHiddenForPlayer", pickup, player, hidden);
    }

    unsigned char pickup_isPickupHiddenForPlayer(void* pickup, void* player) {
        return call<unsigned char>("pickup_isPickupHiddenForPlayer", pickup, player);
    }

    void pickup_setPosition(void* pickup, float posX, float posY, float posZ) {
        return call<void>("pickup_setPosition", pickup, posX, posY, posZ);
    }

    Vector3 pickup_getPosition(void* pickup) {
        return call<Vector3>("pickup_getPosition", pickup);
    }

    void pickup_setVirtualWorld(void* pickup, int vw) {
        return call<void>("pickup_setVirtualWorld", pickup, vw);
    }

    int pickup_getVirtualWorld(void* pickup) {
        return call<int>("pickup_getVirtualWorld", pickup);
    }

    void* playerPickup_create(void* player, int modelId, uint8_t type, float posX, float posY, float posZ, uint32_t virtualWorld, int isStatic) {
        return call<void*>("playerPickup_create", player, modelId, type, posX, posY, posZ, virtualWorld, isStatic);
    }

    void playerPickup_release(void* pickup, void* player) {
        return call<void>("playerPickup_release", pickup, player);
    }

#ifdef __cplusplus
}
#endif
