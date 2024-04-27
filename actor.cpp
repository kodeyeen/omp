#include "include/gomp.h"
#include "include/actor.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* actor_create(int skin, float posX, float posY, float posZ, float angle)
    {
        return call<void*>("actor_create", skin, posX, posY, posZ, angle);
    }

    void actor_release(void* actor)
    {
        return call<void>("actor_release", actor);
    }

    void actor_setSkin(void* actor, int skin)
    {
        return call<void>("actor_setSkin", actor, skin);
    }

    int actor_getSkin(void* actor)
    {
        return call<int>("actor_getSkin", actor);
    }

    void actor_applyAnimation(void* actor, float delta, unsigned char loop, unsigned char lockX, unsigned char lockY, unsigned char freeze, uint32_t time, String lib, String name)
    {
        return call<void>("actor_getSkin", actor, delta, loop, lockX, lockY, freeze, time, lib, name);
    }

    CAnimationData actor_getAnimation(void* actor)
    {
        return call<CAnimationData>("actor_getAnimation", actor);
    }

    void actor_clearAnimations(void* actor)
    {
        return call<void>("actor_clearAnimations", actor);
    }

    void actor_setHealth(void* actor, float health)
    {
        return call<void>("actor_setHealth", actor, health);
    }

    float actor_getHealth(void* actor)
    {
        return call<float>("actor_getHealth", actor);
    }

    void actor_setInvulnerable(void* actor, unsigned char invuln)
    {
        return call<void>("actor_setInvulnerable", actor, invuln);
    }

    unsigned char actor_isInvulnerable(void* actor)
    {
        return call<unsigned char>("actor_isInvulnerable", actor);
    }

    unsigned char actor_isStreamedInForPlayer(void* actor, void* player)
    {
        return call<unsigned char>("actor_isStreamedInForPlayer", actor, player);
    }

    ActorSpawnData actor_getSpawnData(void* actor)
    {
        return call<ActorSpawnData>("actor_getSpawnData", actor);
    }

#ifdef __cplusplus
}
#endif
