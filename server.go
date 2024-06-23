package omp

// #include <stdlib.h>
// #include "include/server.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelMessage
	LogLevelWarning
	LogLevelError
)

type settableCoreDataType int

const (
	settableCoreDataTypeServerName settableCoreDataType = iota
	settableCoreDataTypeModeText
	settableCoreDataTypeMapName
	settableCoreDataTypeLanguage
	settableCoreDataTypeURL
	settableCoreDataTypePassword
	settableCoreDataTypeAdminPassword
)

func Println(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)

	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))

	C.server_printLnU8(cMsg)
}

func Log(level LogLevel, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)

	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))

	C.server_logLnU8(C.int(level), cMsg)
}

func SetWeather(weather int) {
	C.server_setWeather(C.int(weather))
}

func SetWorldTime(hours int) {
	C.server_setWorldTime(C.int(hours))
}

func EnableStuntBonuses() {
	C.server_enableStuntBonuses()
}

func SetServerName(name string) {
	cName := newCString(name)
	defer freeCString(cName)

	C.server_setData(C.int(settableCoreDataTypeServerName), cName)
}

func SetGameModeText(text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.server_setData(C.int(settableCoreDataTypeModeText), cText)
}

func SetMapName(name string) {
	cName := newCString(name)
	defer freeCString(cName)

	C.server_setData(C.int(settableCoreDataTypeMapName), cName)
}

func SetLanguage(language string) {
	cLanguage := newCString(language)
	defer freeCString(cLanguage)

	C.server_setData(C.int(settableCoreDataTypeLanguage), cLanguage)
}

func SetURL(url string) {
	cURL := newCString(url)
	defer freeCString(cURL)

	C.server_setData(C.int(settableCoreDataTypeURL), cURL)
}
