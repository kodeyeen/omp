#include "include/gomp.h"
#include "include/model.h"

#ifdef __cplusplus
extern "C"
{
#endif

	unsigned char model_add(uint8_t type, int32_t id, int32_t baseId, String dffName, String txdName, int32_t virtualWorld, uint8_t timeOn, uint8_t timeOff)
    {
        return call<unsigned char>("model_add", type, id, baseId, dffName, txdName, virtualWorld, timeOn, timeOff);
    }

    String model_getNameFromCheckSum(uint32_t crc)
    {
        return call<String>("model_getNameFromCheckSum", crc);
    }

    unsigned char model_getPath(int32_t modelId, String* dffPath, String* txdPath)
    {
        return call<unsigned char>("model_getPath", modelId, dffPath, txdPath);
    }

#ifdef __cplusplus
}
#endif
