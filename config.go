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
	cKey := stringToCString(key)
	defer C.free(unsafe.Pointer(cKey.buf))

	cType := C.config_getType(cKey)

	switch cType {
	case -1:
		return ErrUnknownOption
	case 0:
		v, ok := value.(int)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setInt(cKey, C.int(v))
	case 1:
		return ErrUnsupportedOption
	case 2:
		v, ok := value.(float64)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setFloat(cKey, C.float(v))
	case 3:
		return ErrUnsupportedOption
	case 4:
		v, ok := value.(bool)
		if !ok {
			return ErrInvalidValueType
		}

		C.config_setBool(cKey, boolToCUchar(v))
	}

	return ErrSomethingWentWrong
}

func Config(key string) (any, error) {
	cKey := stringToCString(key)
	defer C.free(unsafe.Pointer(cKey.buf))

	cType := C.config_getType(cKey)

	switch cType {
	case -1:
		return nil, ErrUnknownOption
	case 0:
		return int(C.config_getInt(cKey)), nil
	case 1:
		return nil, ErrUnsupportedOption
	case 2:
		return float64(C.config_getFloat(cKey)), nil
	case 3:
		return nil, ErrUnsupportedOption
	case 4:
		return C.config_getBool(cKey) != 0, nil
	}

	return nil, ErrSomethingWentWrong
}
