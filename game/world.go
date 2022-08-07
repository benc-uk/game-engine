package main

type Direction int

const (
	DirectionN Direction = iota
	DirectionE
	DirectionS
	DirectionW
)

const (
	DirectionFwd Direction = iota
	DirectionRight
	DirectionBack
	DirectionLeft
)

type Point struct {
	x float64
	y float64
	z float64
}

func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

type Wall struct {
	points []Point
}
