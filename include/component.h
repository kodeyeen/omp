#include "omp.h"

#ifdef __cplusplus
extern "C" {
#endif

	void* Component_Create(uint64_t uid, const char* name, struct ComponentVersion version, void* onReady, void* onReset, void* onFree);

#ifdef __cplusplus
}
#endif
