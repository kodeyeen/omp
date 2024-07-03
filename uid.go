package omp

import (
	"crypto/rand"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func newUID() uint64 {
	token := make([]byte, 8)
	rand.Read(token)

	sb := &strings.Builder{}

	for _, t := range token {
		sb.WriteString(fmt.Sprintf("%02X", t))
	}

	uid, err := strconv.ParseUint(sb.String(), 16, 64)
	if err != nil {
		log.Panicf("NewUID: %s\n", err.Error())
	}

	return uid
}
