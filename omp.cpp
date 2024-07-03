#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
	#include <Windows.h>
#else
	#include <dlfcn.h>
#endif

#include "include/omp.h"

void* libHandle;
std::unordered_map<std::string, void*> funcs;

extern "C" {
    void loadComponent() {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        libHandle = LoadLibrary("./components/Go.dll");
#else
        libHandle = dlopen("./components/Go.so", RTLD_GLOBAL | RTLD_NOW);
#endif
    }

    void unloadComponent() {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        FreeLibrary((HMODULE)libHandle);
#else
        dlclose(libHandle);
#endif
    }

    void* findFunc(const char* name) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        return (void*)GetProcAddress((HMODULE)libHandle, name);
#else
        return dlsym(libHandle, name);
#endif
    }

    bool Event_AddHandler(const char* name, int priority, void* callback) {
        return call<bool>("Event_AddHandler", name, priority, callback);
    }

    bool Event_RemoveHandler(const char* name, int priority, void* callback) {
        return call<bool>("Event_RemoveHandler", name, priority, callback);
    }

    void onReady() {
        Event_AddHandler("onPlayerConnect", 0, (void*)onPlayerConnect);

        onGameModeInit();
    }

    void onFree() {
        Event_RemoveHandler("onPlayerConnect", 0, (void*)onPlayerConnect);

        onGameModeExit();
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
}
