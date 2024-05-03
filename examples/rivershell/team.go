package main

const (
	CapturesToWin = 5
)

var (
	TeamGreen = Team{ID: 1}
	TeamBlue  = Team{ID: 2}
)

type Team struct {
	ID, VehicleCapturedCount int
}
