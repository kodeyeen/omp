#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void server_printLnU8(const char* fmt);
    void server_logLnU8(int logLevel, const char* fmt);

    void server_setWeather(int weather);
    void server_setWorldTime(int hours);
    void server_enableStuntBonuses();

    void server_setData(int type, String data);

#ifdef __cplusplus
}
#endif
