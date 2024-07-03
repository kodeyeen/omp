package omp

// #include <stdlib.h>
// #include "include/component.h"
import "C"
import "unsafe"

type ComponentVersion struct {
	major, minor, patch uint8
	prerel              uint16
}

func newComponent(uid uint64, name string, ver ComponentVersion, onReady, onReset, onFree unsafe.Pointer) unsafe.Pointer {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cVer := C.struct_ComponentVersion{
		major:  C.uchar(ver.major),
		minor:  C.uchar(ver.minor),
		patch:  C.uchar(ver.patch),
		prerel: C.ushort(ver.prerel),
	}

	return C.Component_Create(C.ulonglong(uid), cName, cVer, onReady, onReset, onFree)
}
