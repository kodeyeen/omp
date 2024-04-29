package gomp

// #include <string.h>
// #include "include/gomp.h"
import "C"

func boolToCUchar(val bool) C.uchar {
	if val {
		return 1
	}

	return 0
}

func stringToCString(str string) C.String {
	cKey := C.CString(str)

	return C.String{
		buf:    cKey,
		length: C.strlen(cKey),
	}
}
