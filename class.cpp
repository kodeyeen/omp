#include "include/class.h"

extern "C" {
    void* class_create(ClassData* data) {
        return call<void*>("class_create", data);
    }

    void class_release(void* _class) {
        return call<void>("class_release", _class);
    }

    int class_getID(void* class_) {
        return call<int>("class_getID", class_);
    }

    void class_setClass(void* class_, ClassData* data) {
        return call<void>("class_setClass", class_, data);
    }

    ClassData class_getClass(void* class_) {
        return call<ClassData>("class_getClass", class_);
    }

    void* Class_FromID(int classid) {
        return call<void*>("Class_FromID", classid);
    }
}
