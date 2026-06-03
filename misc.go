package omp

import "time"

type Animation struct {
	Lib, Name                  string
	Delta                      float32
	Loop, LockX, LockY, Freeze bool
	Duration                   time.Duration
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Vector3 struct {
	X, Y, Z float32
}

type Vector2 struct {
	X, Y float32
}
