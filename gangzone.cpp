#include "include/gomp.h"
#include "include/gangzone.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* gangZone_create(float minX, float minY, float maxX, float maxY)
    {
        return call<void*>("gangZone_create", minX, minY, maxX, maxY);
    }

    void gangZone_release(void* gangZone)
    {
        return call<void>("gangZone_release", gangZone);
    }

    void gangZone_useCheck(void* gangZone, unsigned char use)
    {
        return call<void>("gangZone_useCheck", gangZone, use);
    }

    unsigned char gangZone_isShownForPlayer(void* gangZone, void* player)
    {
        return call<unsigned char>("gangZone_isShownForPlayer", gangZone, player);
    }

    unsigned char gangZone_isFlashingForPlayer(void* gangZone, void* player)
    {
        return call<unsigned char>("gangZone_isFlashingForPlayer", gangZone, player);
    }

    void gangZone_showForPlayer(void* gangZone, void* player, uint32_t colour)
    {
        return call<void>("gangZone_showForPlayer", gangZone, player, colour);
    }

    void gangZone_hideForPlayer(void* gangZone, void* player)
    {
        return call<void>("gangZone_hideForPlayer", gangZone, player);
    }

    void gangZone_flashForPlayer(void* gangZone, void* player, uint32_t colour)
    {
        return call<void>("gangZone_flashForPlayer", gangZone, player, colour);
    }

    void gangZone_stopFlashForPlayer(void* gangZone, void* player)
    {
        return call<void>("gangZone_stopFlashForPlayer", gangZone, player);
    }

    GangZonePos gangZone_getPosition(void* gangZone)
    {
        return call<GangZonePos>("gangZone_getPosition", gangZone);
    }

    void gangZone_setPosition(void* gangZone, float minX, float minY, float maxX, float maxY)
    {
        return call<void>("gangZone_setPosition", gangZone, minX, minY, maxX, maxY);
    }

    unsigned char gangZone_isPlayerInside(void* gangZone, void* player)
    {
        return call<unsigned char>("gangZone_isPlayerInside", gangZone, player);
    }

    uint32_t gangZone_getFlashingColourForPlayer(void* gangZone, void* player)
    {
        return call<uint32_t>("gangZone_getFlashingColourForPlayer", gangZone, player);
    }

    uint32_t gangZone_getColourForPlayer(void* gangZone, void* player)
    {
        return call<uint32_t>("gangZone_getColourForPlayer", gangZone, player);
    }

#ifdef __cplusplus
}
#endif
