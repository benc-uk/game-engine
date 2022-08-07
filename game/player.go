package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	pos      *Point3
	angle    float64
	angleCos float64
	angleSin float64
	look     float64
	speed    float64
	inputDir Direction
}

const turnSpeed = 0.06
const fullCircle = 2 * math.Pi
const accel = 0.5
const maxSpeed = 7

func (p *Player) processInputs() {
	// Turn the player left
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.angle += turnSpeed
		if p.angle >= fullCircle {
			p.angle -= fullCircle
		}

		p.angleCos = math.Cos(p.angle)
		p.angleSin = math.Sin(p.angle)
	}

	// Turn the player right
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.angle -= turnSpeed
		if p.angle < 0 {
			p.angle += fullCircle
		}

		p.angleCos = math.Cos(p.angle)
		p.angleSin = math.Sin(p.angle)
	}

	// Move the player forward, backwards, or strafe
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) ||
		ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyE) {
		p.speed += accel
		if p.speed > maxSpeed {
			p.speed = maxSpeed
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			p.inputDir = DirectionFwd
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			p.inputDir = DirectionBack
		}

		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			p.inputDir = DirectionLeft
		}

		if ebiten.IsKeyPressed(ebiten.KeyE) {
			p.inputDir = DirectionRight
		}
	}

	// Slight optimization: if the player is not moving, don't update the position
	if p.inputDir >= 0 {
		// Work out how far to move the player
		dx := math.Sin(p.angle) * p.speed
		dy := math.Cos(p.angle) * p.speed

		// Move player x,y based on current input
		switch p.inputDir {
		case DirectionFwd:
			p.pos.x += dx
			p.pos.y += dy
		case DirectionBack:
			p.pos.x -= dx
			p.pos.y -= dy
		case DirectionLeft:
			p.pos.x += dy
			p.pos.y -= dx
		case DirectionRight:
			p.pos.x -= dy
			p.pos.y += dx
		}
	}

	// Deceleration logic
	if p.speed > 0 {
		p.speed -= accel * 0.5
	} else {
		p.inputDir = -1
	}

	// Float up
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		p.pos.z += 1
	}

	// Float down
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.pos.z -= 1
	}

	// Look up
	if ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		p.look += 1
	}

	// Look down
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		p.look -= 1
	}
}
