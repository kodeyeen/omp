#include "include/gomp.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* pickup_create(int modelId, unsigned char type, float x, float y, float z, unsigned int virtualWorld, int isStatic, void* player)
    {
        return call<void*>("pickup_create", modelId, type, x, y, z, virtualWorld, isStatic, player);
    }

#ifdef __cplusplus
}
#endif
