package omp

// #include "include/console.h"
import "C"

func SendRCONCommand(command string) {
	cCommand := newCString(command)
	defer freeCString(cCommand)

	C.console_send(cCommand)
}
