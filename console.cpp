#include "include/console.h"

#ifdef __cplusplus
extern "C" {
#endif

    void console_send(String command) {
        return call<void>("console_send", command);
    }

#ifdef __cplusplus
}
#endif
