#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct {
        Vector2 min;
        Vector2 max;
    } TurfPos;

    void* turf_create(float minX, float minY, float maxX, float maxY);
    void turf_release(void* turf);
    void turf_useCheck(void* turf, unsigned char use);
    unsigned char turf_isShownForPlayer(void* turf, void* player);
    unsigned char turf_isFlashingForPlayer(void* turf, void* player);
    void turf_showForPlayer(void* turf, void* player, uint32_t colour);
    void turf_showForAll(void* turf, uint32_t colour);
    void turf_hideForPlayer(void* turf, void* player);
    void turf_hideForAll(void* turf);
    void turf_flashForPlayer(void* turf, void* player, uint32_t colour);
    void turf_flashForAll(void* turf, uint32_t colour);
    void turf_stopFlashForPlayer(void* turf, void* player);
    void turf_stopFlashForAll(void* turf);
    TurfPos turf_getPosition(void* turf);
    void turf_setPosition(void* turf, float minX, float minY, float maxX, float maxY);
    unsigned char turf_isPlayerInside(void* turf, void* player);
    uint32_t turf_getFlashingColourForPlayer(void* turf, void* player);
    uint32_t turf_getColourForPlayer(void* turf, void* player);

#ifdef __cplusplus
}
#endif
