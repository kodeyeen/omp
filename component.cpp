#include "include/component.h"

extern "C" {
    void* Component_Create(uint64_t uid, const char* name, struct ComponentVersion version, void* onReady, void* onReset, void* onFree) {
        return call<void*>("Component_Create", uid, name, version, onReady, onReset, onFree);
    }
}
