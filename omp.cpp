#include <Windows.h>

#include "include/omp.h"

void* libHandle;
std::unordered_map<std::string, void*> funcs;

#ifdef __cplusplus
extern "C"
{
#endif

    void init(const char* libPath)
    {
        libHandle = openLib(libPath);
    }

    void* openLib(const char* path)
    {
        return LoadLibrary((LPCTSTR)path);
    }

    void closeLib(void* handle)
    {
        FreeLibrary((HMODULE)handle);
    }

    void* findFunc(void* handle, const char* name)
    {
        FARPROC func = GetProcAddress((HMODULE)handle, name);

        return (void*)func;
    }

    void freeArray(Array arr)
    {
        return call<void>("freeArray", arr);
    }

#ifdef __cplusplus
}
#endif
