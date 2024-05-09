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

func SetGameModeText(text string) {
	cText := newCString(text)
	defer freeCString(cText)

	C.server_setModeText(cText)
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
