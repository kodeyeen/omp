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

    void server_setWeather(int weather) {
        return call<void>("server_setWeather", weather);
    }

    void server_setWorldTime(int hours) {
        return call<void>("server_setWorldTime", hours);
    }

    void server_enableStuntBonuses() {
        return call<void>("server_enableStuntBonuses");
    }

    void server_setData(int type, String data) {
        return call<void>("server_setData", type, data);
    }

#ifdef __cplusplus
}
#endif
