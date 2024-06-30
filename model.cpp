#include "include/model.h"

extern "C" {
    unsigned char model_add(uint8_t type, int32_t id, int32_t baseId, String dffName, String txdName, int32_t virtualWorld, uint8_t timeOn, uint8_t timeOff) {
        return call<unsigned char>("customModel_add", type, id, baseId, dffName, txdName, virtualWorld, timeOn, timeOff);
    }

    String model_getNameFromCheckSum(uint32_t checksum) {
        return call<String>("customModel_getNameFromCheckSum", checksum);
    }

    unsigned char model_isValid(int32_t modelId) {
        return call<unsigned char>("customModel_isValid", modelId);
    }

    unsigned char model_getPath(int32_t modelId, String* dffPath, String* txdPath) {
        return call<unsigned char>("customModel_getPath", modelId, dffPath, txdPath);
    }
}
