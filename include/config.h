#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

    void config_setFloat(String key, float value);
    void config_setInt(String key, int value);
    void config_setBool(String key, unsigned char value);
    int config_getType(String key);
    float config_getFloat(String key);
    int config_getInt(String key);
    unsigned char config_getBool(String key);
    unsigned char config_isBanned(String ip);

#ifdef __cplusplus
}
#endif
