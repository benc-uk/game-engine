package main

type Direction int

const (
	DirectionN Direction = iota
	DirectionE
	DirectionS
	DirectionW
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
	p1   Point
	p2   Point
	side bool
}
