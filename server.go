package omp

// #include "include/server.h"
import "C"

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
