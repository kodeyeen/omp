#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void server_printLnU8(const char* fmt);
    void server_logLnU8(int logLevel, const char* fmt);

    void server_setModeText(String text);
    void server_setWeather(int weather);
    void server_setWorldTime(int hours);
    void server_enableStuntBonuses();

#ifdef __cplusplus
}
#endif
