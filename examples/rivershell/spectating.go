package main

type SpectateState int

const (
	SpectateStateNone SpectateState = iota
	SpectateStatePlayer
	SpectateStateFixed
)
