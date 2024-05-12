#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
	#include <Windows.h>
#else
	#include <dlfcn.h>
#endif

#include "include/omp.h"

void* libHandle;
std::unordered_map<std::string, void*> funcs;

#ifdef __cplusplus
extern "C" {
#endif

    void init(const char* libPath) {
        libHandle = openLib(libPath);
    }

    void* openLib(const char* path) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        return LoadLibrary((LPCTSTR)path);
#else
        return dlopen(path, RTLD_GLOBAL | RTLD_NOW);
#endif
    }

    void closeLib(void* handle) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        FreeLibrary((HMODULE)handle);
#else
        dlclose(handle);
#endif
    }

    void* findFunc(void* handle, const char* name) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        return (void*)GetProcAddress((HMODULE)handle, name);
#else
        return dlsym(handle, name);
#endif
    }

    void freeArray(Array arr) {
        return call<void>("freeArray", arr);
    }

    uint8_t getWeaponSlotIndex(uint8_t weapon) {
        return call<uint8_t>("getWeaponSlotIndex", weapon);
    }

    unsigned char getVehicleModelInfo(int model, int type, Vector3* out) {
        return call<unsigned char>("getVehicleModelInfo", model, type, out);
    }

#ifdef __cplusplus
}
#endif
