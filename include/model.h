#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

	unsigned char model_add(uint8_t type, int32_t id, int32_t baseId, String dffName, String txdName, int32_t virtualWorld, uint8_t timeOn, uint8_t timeOff);
    String model_getNameFromCheckSum(uint32_t crc);
    unsigned char model_getPath(int32_t modelId, String* dffPath, String* txdPath);

#ifdef __cplusplus
}
#endif
