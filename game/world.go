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

type Point3 struct {
	x float64
	y float64
	z float64
}

type Point2 struct {
	x float64
	y float64
}

func NewPoint3(x, y, z float64) *Point3 {
	return &Point3{x, y, z}
}

func NewGame2(x, y float64) *Point2 {
	return &Point2{x, y}
}

type Wall struct {
	p1     Point2
	p2     Point2
	height float64
	z      float64
}
