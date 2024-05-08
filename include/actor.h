#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    typedef struct {
        float delta;
        unsigned char loop;
        unsigned char lockX;
        unsigned char lockY;
        unsigned char freeze;
        uint32_t time;
        String lib;
        String name;
    } CAnimationData;

    typedef struct {
        Vector3 position;
        float facingAngle;
        int skin;
    } ActorSpawnData;

    void* actor_create(int skin, float posX, float posY, float posZ, float angle);
    void actor_release(void* actor);
    void actor_setSkin(void* actor, int skin);
    int actor_getSkin(void* actor);
    void actor_applyAnimation(void* actor, float delta, unsigned char loop, unsigned char lockX, unsigned char lockY, unsigned char freeze, uint32_t time, String lib, String name);
    CAnimationData actor_getAnimation(void* actor);
    void actor_clearAnimations(void* actor);
    void actor_setHealth(void* actor, float health);
    float actor_getHealth(void* actor);
    void actor_setInvulnerable(void* actor, unsigned char invuln);
    unsigned char actor_isInvulnerable(void* actor);
    unsigned char actor_isStreamedInForPlayer(void* actor, void* player);
    ActorSpawnData actor_getSpawnData(void* actor);
    void actor_setPosition(void* actor, float posX, float posY, float posZ);
    Vector3 actor_getPosition(void* actor);
    void actor_setVirtualWorld(void* actor, int vw);
    int actor_getVirtualWorld(void* actor);
    void actor_setFacingAngle(void* actor, float angle);
    float actor_getFacingAngle(void* actor);

#ifdef __cplusplus
}
#endif
