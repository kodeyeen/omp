package gomp

// #include <stdlib.h>
// #include <string.h>
// #include "include/server.h"
import "C"
import "unsafe"

func SetGameModeText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.server_setModeText(C.String{
		buf:    cText,
		length: C.strlen(cText),
	})
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
