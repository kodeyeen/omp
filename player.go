package main

// #include <stdlib.h>
// #include "component.h"
import "C"
import (
	"unsafe"
)

type Player struct {
	handle unsafe.Pointer
}

func (p *Player) ID() int {
	// return int(C.player_getID(p.handle))
	return 0
}

func (p *Player) Name() string {
	cname := C.player_getName(p.handle)

	return C.GoString(cname)
}

func (p *Player) SendMessage(color int, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.player_sendClientMessage(p.handle, C.int(color), cmsg)
}
