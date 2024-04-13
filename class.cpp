#include "include/gomp.h"
#include "include/class.h"

#ifdef __cplusplus
extern "C"
{
#endif

    void* class_create(ClassData* data)
    {
        return call<void*>("class_create", data);
    }

    int class_getID(void* class_)
    {
        return call<int>("class_getID", class_);
    }

    void class_setClass(void* class_, ClassData* data)
    {
        return call<void>("class_setClass", class_, data);
    }

    ClassData class_getClass(void* class_)
    {
        return call<ClassData>("class_getClass", class_);
    }

#ifdef __cplusplus
}
#endif

