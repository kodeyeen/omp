#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct {
        Vector2 min;
        Vector2 max;
    } GangZonePos;

    void* gangZone_create(float minX, float minY, float maxX, float maxY);
    void gangZone_release(void* gangZone);
    unsigned char gangZone_isShownForPlayer(void* gangZone, void* player);
    unsigned char gangZone_isFlashingForPlayer(void* gangZone, void* player);
    void gangZone_showForPlayer(void* gangZone, void* player, uint32_t colour);
    void gangZone_hideForPlayer(void* gangZone, void* player);
    void gangZone_flashForPlayer(void* gangZone, void* player, uint32_t colour);
    void gangZone_stopFlashForPlayer(void* gangZone, void* player);
    GangZonePos gangZone_getPosition(void* gangZone);
    void gangZone_setPosition(void* gangZone, float minX, float minY, float maxX, float maxY);
    unsigned char gangZone_isPlayerInside(void* gangZone, void* player);
    uint32_t gangZone_getFlashingColourForPlayer(void* gangZone, void* player);
    uint32_t gangZone_getColourForPlayer(void* gangZone, void* player);

#ifdef __cplusplus
}
#endif
