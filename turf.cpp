#include "include/gomp.h"
#include "include/turf.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* turf_create(float minX, float minY, float maxX, float maxY)
    {
        return call<void*>("gangZone_create", minX, minY, maxX, maxY);
    }

    void turf_release(void* turf)
    {
        return call<void>("gangZone_release", turf);
    }

    void turf_useCheck(void* turf, unsigned char use)
    {
        return call<void>("gangZone_useCheck", turf, use);
    }

    unsigned char turf_isShownForPlayer(void* turf, void* player)
    {
        return call<unsigned char>("gangZone_isShownForPlayer", turf, player);
    }

    unsigned char turf_isFlashingForPlayer(void* turf, void* player)
    {
        return call<unsigned char>("gangZone_isFlashingForPlayer", turf, player);
    }

    void turf_showForPlayer(void* turf, void* player, uint32_t colour)
    {
        return call<void>("gangZone_showForPlayer", turf, player, colour);
    }

    void turf_showForAll(void* turf, uint32_t colour)
    {
        return call<void>("gangZone_showForAll", turf, colour);
    }

    void turf_hideForPlayer(void* turf, void* player)
    {
        return call<void>("gangZone_hideForPlayer", turf, player);
    }

    void turf_hideForAll(void* turf)
    {
        return call<void>("gangZone_hideForAll", turf);
    }

    void turf_flashForPlayer(void* turf, void* player, uint32_t colour)
    {
        return call<void>("gangZone_flashForPlayer", turf, player, colour);
    }

    void turf_flashForAll(void* turf, uint32_t colour)
    {
        return call<void>("gangZone_flashForAll", turf, colour);
    }

    void turf_stopFlashForPlayer(void* turf, void* player)
    {
        return call<void>("gangZone_stopFlashForPlayer", turf, player);
    }

    void turf_stopFlashForAll(void* turf)
    {
        return call<void>("gangZone_stopFlashForAll", turf);
    }

    TurfPos turf_getPosition(void* turf)
    {
        return call<TurfPos>("gangZone_getPosition", turf);
    }

    void turf_setPosition(void* turf, float minX, float minY, float maxX, float maxY)
    {
        return call<void>("gangZone_setPosition", turf, minX, minY, maxX, maxY);
    }

    unsigned char turf_isPlayerInside(void* turf, void* player)
    {
        return call<unsigned char>("gangZone_isPlayerInside", turf, player);
    }

    uint32_t turf_getFlashingColourForPlayer(void* turf, void* player)
    {
        return call<uint32_t>("gangZone_getFlashingColourForPlayer", turf, player);
    }

    uint32_t turf_getColourForPlayer(void* turf, void* player)
    {
        return call<uint32_t>("gangZone_getColourForPlayer", turf, player);
    }

    // Player Turf

    void* playerTurf_create(void* player, float minX, float minY, float maxX, float maxY)
    {
        return call<void*>("playerGangZone_create", player, minX, minY, maxX, maxY);
    }

    void playerTurf_release(void* turf, void* player)
    {
        return call<void>("playerGangZone_release", turf, player);
    }

#ifdef __cplusplus
}
#endif
