package gomp

// #include <stdlib.h>
// #include "include/model.h"
import "C"
import (
	"errors"
	"unsafe"
)

type modelPath struct {
	DffPath, TxdPath string
}

func AddCharModel(baseID, newID int, dff, txd string) error {
	cDff := stringToCString(dff)
	defer C.free(unsafe.Pointer(cDff.buf))

	cTxd := stringToCString(txd)
	defer C.free(unsafe.Pointer(cTxd.buf))

	cOk := C.model_add(1, C.int(newID), C.int(baseID), cDff, cTxd, -1, 0, 0)
	if cOk == 0 {
		return errors.New("failed to add character model")
	}

	return nil
}

func AddSimpleModel(vw, baseID, newID int, dff, txd string) error {
	cDff := stringToCString(dff)
	defer C.free(unsafe.Pointer(cDff.buf))

	cTxd := stringToCString(txd)
	defer C.free(unsafe.Pointer(cTxd.buf))

	cOk := C.model_add(2, C.int(newID), C.int(baseID), cDff, cTxd, C.int(vw), 0, 0)
	if cOk == 0 {
		return errors.New("failed to add simple model")
	}

	return nil
}

func AddSimpleModelTimed(vw, baseID, newID int, dff, txd string, timeOn, timeOff int) error {
	cDff := stringToCString(dff)
	defer C.free(unsafe.Pointer(cDff.buf))

	cTxd := stringToCString(txd)
	defer C.free(unsafe.Pointer(cTxd.buf))

	cOk := C.model_add(2, C.int(newID), C.int(baseID), cDff, cTxd, C.int(vw), C.uchar(timeOn), C.uchar(timeOff))
	if cOk == 0 {
		return errors.New("failed to add timed simple model")
	}

	return nil
}

func ModelNameFromCRC(crc int) string {
	cName := C.model_getNameFromCheckSum(C.uint(crc))

	return C.GoStringN(cName.buf, C.int(cName.length))
}

func ModelPath(modelID int) (modelPath, error) {
	var cDffPath C.String
	var cTxdPath C.String

	cOk := C.model_getPath(C.uint(modelID), &cDffPath, &cTxdPath)
	if cOk == 0 {
		return modelPath{}, errors.New("unable to get model path")
	}

	path := modelPath{
		DffPath: C.GoStringN(cDffPath.buf, C.int(cDffPath.length)),
		TxdPath: C.GoStringN(cTxdPath.buf, C.int(cTxdPath.length)),
	}

	return path, nil
}
