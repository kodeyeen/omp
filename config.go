package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/config.h"
import "C"
import (
	"errors"
	"unsafe"
)

const (
	PlayerMarkerModeOff = iota
	PlayerMarkerModeGlobal
	PlayerMarkerModeStreamed
)

var (
	ErrInvalidValueType   = errors.New("invalid value type")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUnknownOption      = errors.New("unknown option")
	ErrUnsupportedOption  = errors.New("unsupported option")
)

func SetConfig(key string, value any) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cKeyStr := C.String{
		buf:    cKey,
		length: C.strlen(cKey),
	}

	cType := C.config_getType(cKeyStr)

	switch cType {
	case -1:
		return ErrUnknownOption
	case 0:
		v, ok := value.(int)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setInt(cKeyStr, C.int(v))
	case 1:
		return ErrUnsupportedOption
	case 2:
		v, ok := value.(float64)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setFloat(cKeyStr, C.float(v))
	case 3:
		return ErrUnsupportedOption
	case 4:
		v, ok := value.(bool)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setBool(cKeyStr, boolToCUchar(v))
	}

	return ErrSomethingWentWrong
}

func Config(key string) (any, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cKeyStr := C.String{
		buf:    cKey,
		length: C.strlen(cKey),
	}

	cType := C.config_getType(cKeyStr)

	switch cType {
	case -1:
		return nil, ErrUnknownOption
	case 0:
		return int(C.config_getInt(cKeyStr)), nil
	case 1:
		return nil, ErrUnsupportedOption
	case 2:
		return float64(C.config_getFloat(cKeyStr)), nil
	case 3:
		return nil, ErrUnsupportedOption
	case 4:
		return C.config_getBool(cKeyStr) != 0, nil
	}

	return nil, ErrSomethingWentWrong
}
