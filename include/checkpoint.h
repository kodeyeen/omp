#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void* checkpoint_setPosition(void* checkpoint, float posX, float posY, float posZ);
    Vector3 checkpoint_getPosition(void* checkpoint);
    void checkpoint_setRadius(void* checkpoint, float radius);
    float checkpoint_getRadius(void* checkpoint);
    unsigned char checkpoint_isPlayerInside(void* checkpoint);
    void checkpoint_enable(void* checkpoint);
    void checkpoint_disable(void* checkpoint);
    unsigned char checkpoint_isEnabled(void* checkpoint);

    // race checkpoint

    void* raceCheckpoint_setPosition(void* checkpoint, float posX, float posY, float posZ);
    Vector3 raceCheckpoint_getPosition(void* checkpoint);
    void raceCheckpoint_setRadius(void* checkpoint, float radius);
    float raceCheckpoint_getRadius(void* checkpoint);
    unsigned char raceCheckpoint_isPlayerInside(void* checkpoint);
    void raceCheckpoint_enable(void* checkpoint);
    void raceCheckpoint_disable(void* checkpoint);
    unsigned char raceCheckpoint_isEnabled(void* checkpoint);
    void raceCheckpoint_setType(void* checkpoint, int type);
    int raceCheckpoint_getType(void* checkpoint);
    void raceCheckpoint_setNextPosition(void* checkpoint, float posX, float posY, float posZ);
    Vector3 raceCheckpoint_getNextPosition(void* checkpoint);

#ifdef __cplusplus
}
#endif
