#include "gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void config_setPlayerMarkerMode(int mode);
    void config_setNametagDrawRadius(float radius);
    void config_useEntryExitMarkers(int use);
    void config_useManualEngineAndLights(int use);
    void config_useNametags(int use);

#ifdef __cplusplus
}
#endif
