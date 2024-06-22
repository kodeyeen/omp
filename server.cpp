#include "include/server.h"

#ifdef __cplusplus
extern "C" {
#endif

    void server_printLnU8(const char* fmt) {
        return call<void>("server_printLnU8", fmt);
    }

    void server_logLnU8(int logLevel, const char* fmt) {
        return call<void>("server_logLnU8", logLevel, fmt);
    }

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

    void server_setServerName(String name) {
        return call<void>("server_setServerName", name);
    }

    void server_setMapName(String name) {
        return call<void>("server_setMapName", name);
    }

    void server_setLanguage(String language) {
        return call<void>("server_setLanguage", language);
    }

    void server_setURL(String url) {
        return call<void>("server_setURL", url);
    }

#ifdef __cplusplus
}
#endif
