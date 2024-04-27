package gomp

import "C"

func boolToCUchar(val bool) C.uchar {
	if val {
		return 1
	}

	return 0
}
