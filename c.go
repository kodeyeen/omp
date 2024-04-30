package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/gomp.h"
import "C"
import "unsafe"

func newCUchar(goBool bool) C.uchar {
	if goBool {
		return 1
	}

	return 0
}

func newCString(goStr string) C.String {
	cStr := C.CString(goStr)

	return C.String{
		buf:    cStr,
		length: C.strlen(cStr),
	}
}

func freeCString(cStr C.String) {
	C.free(unsafe.Pointer(cStr.buf))
}
