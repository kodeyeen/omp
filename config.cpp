#include "include/gomp.h"
#include "include/config.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void config_setPlayerMarkerMode(int mode)
    {
        return call<void>("config_setPlayerMarkerMode", mode);
    }

    void config_setNametagDrawRadius(float radius)
    {
        return call<void>("config_setNametagDrawRadius", radius);
    }

    void config_useEntryExitMarkers(int use)
    {
        return call<void>("config_useEntryExitMarkers", use);
    }

    void config_useManualEngineAndLights(int use)
    {
        return call<void>("config_useManualEngineAndLights", use);
    }

    void config_useNametags(int use)
    {
        return call<void>("config_useNametags", use);
    }

#ifdef __cplusplus
}
#endif
