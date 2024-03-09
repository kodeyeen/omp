package main

// #include <stdlib.h>
// #include "component.h"
import "C"
import "unsafe"

var comp *component

// provides Go interface for open.mp component
type component struct {
	handle unsafe.Pointer
}

func getComponent() *component {
	if comp == nil {
		comp = &component{}
	}

	return comp
}

func (c *component) init(libpath string) {
	clibpath := C.CString(libpath)
	defer C.free(unsafe.Pointer(clibpath))

	c.handle = C.loadLib(clibpath)

	C.initFuncs(c.handle)
}

func (c *component) player_getName(plrHandle unsafe.Pointer) string {
	cname := C.player_getName(plrHandle)

	return C.GoString(cname)
}

func (c *component) player_sendClientMessage(plrHandle unsafe.Pointer, color int, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(plrHandle, C.int(color), cmsg)
}
