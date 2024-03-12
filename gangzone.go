package main

import "unsafe"

type GangZone struct {
	handle unsafe.Pointer
}

func NewGangZone() *GangZone {
	return &GangZone{}
}

func NewPlayerGangZone() *GangZone {
	return &GangZone{}
}
