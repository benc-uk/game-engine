package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	pos      *Point
	a        float64
	l        float64
	speed    float64
	inputDir Direction
}

const turnSpeed = 0.06
const fullCircle = 2 * math.Pi
const accel = 0.5
const maxSpeed = 7

func (p *Player) processInputs() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.a += turnSpeed
		if p.a >= fullCircle {
			p.a -= fullCircle
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.a -= turnSpeed
		if p.a < 0 {
			p.a += fullCircle
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) ||
		ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyE) {
		p.speed += accel
		if p.speed > maxSpeed {
			p.speed = maxSpeed
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			p.inputDir = DirectionN
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			p.inputDir = DirectionS
		}

		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			p.inputDir = DirectionW
		}

		if ebiten.IsKeyPressed(ebiten.KeyE) {
			p.inputDir = DirectionE
		}
	}

	dx := math.Sin(p.a) * p.speed
	dy := math.Cos(p.a) * p.speed

	switch p.inputDir {
	case DirectionN:
		p.pos.x += dx
		p.pos.y += dy
	case DirectionS:
		p.pos.x -= dx
		p.pos.y -= dy
	case DirectionW:
		p.pos.x += dy
		p.pos.y -= dx
	case DirectionE:
		p.pos.x -= dy
		p.pos.y += dx
	}

	if p.speed > 0 {
		p.speed -= accel * 0.5
	} else {
		p.inputDir = -1
	}
}
