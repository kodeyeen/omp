#ifndef OMP_H
#define OMP_H

#include "capi/include/ompcapi.h"

#ifdef __cplusplus
extern "C" {
#endif

    bool onGameModeInit();
    bool onGameModeExit();
    bool onPlayerConnect(struct EventArgs_onPlayerConnect* args);

#ifdef __cplusplus
}
#endif

#endif // OMP_H
