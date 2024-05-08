#include "include/checkpoint.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* checkpoint_setPosition(void* checkpoint, float posX, float posY, float posZ) {
        return call<void*>("checkpoint_setPosition", checkpoint, posX, posY, posZ);
    }

    Vector3 checkpoint_getPosition(void* checkpoint) {
        return call<Vector3>("checkpoint_getPosition", checkpoint);
    }

    void checkpoint_setRadius(void* checkpoint, float radius) {
        return call<void>("checkpoint_setRadius", checkpoint, radius);
    }

    float checkpoint_getRadius(void* checkpoint) {
        return call<float>("checkpoint_getRadius", checkpoint);
    }

    unsigned char checkpoint_isPlayerInside(void* checkpoint) {
        return call<unsigned char>("checkpoint_isPlayerInside", checkpoint);
    }

    void checkpoint_enable(void* checkpoint) {
        return call<void>("checkpoint_enable", checkpoint);
    }

    void checkpoint_disable(void* checkpoint) {
        return call<void>("checkpoint_disable", checkpoint);
    }

    unsigned char checkpoint_isEnabled(void* checkpoint) {
        return call<unsigned char>("checkpoint_isEnabled", checkpoint);
    }

    void* raceCheckpoint_setPosition(void* checkpoint, float posX, float posY, float posZ) {
        return call<void*>("raceCheckpoint_setPosition", checkpoint, posX, posY, posZ);
    }

    Vector3 raceCheckpoint_getPosition(void* checkpoint) {
        return call<Vector3>("raceCheckpoint_getPosition", checkpoint);
    }

    void raceCheckpoint_setRadius(void* checkpoint, float radius) {
        return call<void>("raceCheckpoint_setRadius", checkpoint, radius);
    }

    float raceCheckpoint_getRadius(void* checkpoint) {
        return call<float>("raceCheckpoint_getRadius", checkpoint);
    }

    unsigned char raceCheckpoint_isPlayerInside(void* checkpoint) {
        return call<unsigned char>("raceCheckpoint_isPlayerInside", checkpoint);
    }

    void raceCheckpoint_enable(void* checkpoint) {
        return call<void>("raceCheckpoint_enable", checkpoint);
    }

    void raceCheckpoint_disable(void* checkpoint) {
        return call<void>("raceCheckpoint_disable", checkpoint);
    }

    unsigned char raceCheckpoint_isEnabled(void* checkpoint) {
        return call<unsigned char>("raceCheckpoint_isEnabled", checkpoint);
    }

    void raceCheckpoint_setType(void* checkpoint, int type) {
        return call<void>("raceCheckpoint_setType", checkpoint, type);
    }

    int raceCheckpoint_getType(void* checkpoint) {
        return call<int>("raceCheckpoint_getType", checkpoint);
    }

    void raceCheckpoint_setNextPosition(void* checkpoint, float posX, float posY, float posZ) {
        return call<void>("raceCheckpoint_setNextPosition", checkpoint, posX, posY, posZ);
    }

    Vector3 raceCheckpoint_getNextPosition(void* checkpoint) {
        return call<Vector3>("raceCheckpoint_getNextPosition", checkpoint);
    }

#ifdef __cplusplus
}
#endif
