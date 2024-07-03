#include "omp.h"

struct ComponentVersion {
	uint8_t major;
	uint8_t minor;
	uint8_t patch;
	uint16_t prerel;
};

#ifdef __cplusplus
extern "C" {
#endif

	void* Component_Create(uint64_t uid, const char* name, struct ComponentVersion version, void* onReady, void* onReset, void* onFree);

#ifdef __cplusplus
}
#endif
