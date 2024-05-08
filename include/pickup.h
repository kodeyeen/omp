#include <stdint.h>
#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* pickup_create(int modelId, uint8_t type, float posX, float posY, float posZ, uint32_t virtualWorld, int isStatic);
    void pickup_release(void* pickup);
    int pickup_getID(void* pickup);
    void pickup_setType(void* pickup, uint8_t type);
    uint8_t pickup_getType(void* pickup);
    void pickup_setModel(void* pickup, int model);
    int pickup_getModel(void* pickup);
    unsigned char pickup_isStreamedInForPlayer(void* pickup, void* player);
    void pickup_setPickupHiddenForPlayer(void* pickup, void* player, unsigned char hidden);
    unsigned char pickup_isPickupHiddenForPlayer(void* pickup, void* player);

    // entity
    void pickup_setPosition(void* pickup, float posX, float posY, float posZ);
    Vector3 pickup_getPosition(void* pickup);
    void pickup_setVirtualWorld(void* pickup, int vw);
    int pickup_getVirtualWorld(void* pickup);

    void* playerPickup_create(void* player, int modelId, uint8_t type, float posX, float posY, float posZ, uint32_t virtualWorld, int isStatic);
    void playerPickup_release(void* pickup, void* player);

#ifdef __cplusplus
}
#endif
