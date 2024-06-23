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

func SetGameModeTextf(format string, a ...any) {
	SetGameModeText(fmt.Sprintf(format, a...))
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

	C.server_setServerName(cName)
}

func SetServerNamef(format string, a ...any) {
	SetServerName(fmt.Sprintf(format, a...))
}

func SetMapName(name string) {
	cName := newCString(name)
	defer freeCString(cName)

	C.server_setMapName(cName)
}

func SetMapNamef(format string, a ...any) {
	SetMapName(fmt.Sprintf(format, a...))
}

func SetLanguage(language string) {
	cLanguage := newCString(language)
	defer freeCString(cLanguage)

	C.server_setLanguage(cLanguage)
}

func SetLanguagef(format string, a ...any) {
	SetLanguage(fmt.Sprintf(format, a...))
}

func SetURL(url string) {
	cUrl := newCString(url)
	defer freeCString(cUrl)

	C.server_setURL(cUrl)
}

func SetURLf(format string, a ...any) {
	SetURL(fmt.Sprintf(format, a...))
}
