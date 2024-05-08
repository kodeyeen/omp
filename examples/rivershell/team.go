package main

type Team = int

const (
	CapturesToWin = 1

	TeamGreen Team = 1
	TeamBlue  Team = 2
)

var captures = make(map[Team]int, 2)
