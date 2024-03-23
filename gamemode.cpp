#include "include/gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void useManualEngineAndLights()
    {
        return call<void>("useManualEngineAndLights");
    }

#ifdef __cplusplus
}
#endif
