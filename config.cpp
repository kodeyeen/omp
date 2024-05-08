#include "include/config.h"

#ifdef __cplusplus
extern "C" {
#endif

    void config_setFloat(String key, float value) {
        return call<void>("config_setFloat", key, value);
    }

    void config_setInt(String key, int value) {
        return call<void>("config_setInt", key, value);
    }

    void config_setBool(String key, unsigned char value) {
        return call<void>("config_setBool", key, value);
    }

    int config_getType(String key) {
        return call<int>("config_getType", key);
    }

    float config_getFloat(String key) {
        return call<float>("config_getFloat", key);
    }

    int config_getInt(String key) {
        return call<int>("config_getInt", key);
    }

    unsigned char config_getBool(String key) {
        return call<unsigned char>("config_getBool", key);
    }

    unsigned char config_isBanned(String ip) {
        return call<unsigned char>("config_isBanned", ip);
    }

#ifdef __cplusplus
}
#endif
