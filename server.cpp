#include "include/server.h"

#ifdef __cplusplus
extern "C" {
#endif

    void server_setModeText(String text) {
        return call<void>("server_setModeText", text);
    }

    void server_setWeather(int weather) {
        return call<void>("server_setWeather", weather);
    }

    void server_setWorldTime(int hours) {
        return call<void>("server_setWorldTime", hours);
    }

    void server_enableStuntBonuses() {
        return call<void>("server_enableStuntBonuses");
    }

#ifdef __cplusplus
}
#endif
