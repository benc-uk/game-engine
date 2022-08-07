package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	width   int
	height  int
	width2  float64
	height2 float64
	player  *Player
}

const baseWidth = 160
const baseHeight = 120
const fovScale = 100

var wall = &Wall{
	p1:     Point2{x: 50, y: 250},
	p2:     Point2{x: 150, y: 270},
	height: 40,
}

func NewGame(res int) *Game {
	return &Game{
		width:   baseWidth * res,
		height:  baseHeight * res,
		width2:  float64(baseWidth*res) / 2.0,
		height2: float64(baseHeight*res) / 2.0,

		player: &Player{
			pos:      NewPoint3(100, 100, 15),
			angle:    0,
			angleCos: math.Cos(0),
			angleSin: math.Sin(0),
			look:     1,
			speed:    0.0,
		},
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.processInputs()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := color.RGBA{255, 255, 255, 255}

	s1 := g.pointToScreen(Point3{x: wall.p1.x, y: wall.p1.y, z: wall.z})
	s2 := g.pointToScreen(Point3{x: wall.p2.x, y: wall.p2.y, z: wall.z})
	s3 := g.pointToScreen(Point3{x: wall.p1.x, y: wall.p1.y, z: wall.z + wall.height})
	s4 := g.pointToScreen(Point3{x: wall.p2.x, y: wall.p2.y, z: wall.z + wall.height})

	ebitenutil.DrawLine(screen, s1.x, float64(g.height)-s1.y, s2.x, float64(g.height)-s2.y, c)
	ebitenutil.DrawLine(screen, s3.x, float64(g.height)-s3.y, s4.x, float64(g.height)-s4.y, c)

	g.drawDebug(screen)
}

func (g *Game) pointToScreen(p Point3) Point2 {
	x := p.x - g.player.pos.x
	y := p.y - g.player.pos.y

	// worldP position of wall
	worldP := Point3{
		x: x*g.player.angleCos - y*g.player.angleSin,
		y: y*g.player.angleCos + x*g.player.angleSin,
	}

	// Calc Z needs the world Y coord so we do it separately
	worldP.z = p.z - g.player.pos.z + ((g.player.look * worldP.y) / 32)

	// screen position of wall points
	return Point2{
		x: ((worldP.x * fovScale) / worldP.y) + g.width2,
		y: ((worldP.z * fovScale) / worldP.y) + g.height2,
	}
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.player.pos.x-2, g.player.pos.y-2, 4, 4, color.RGBA{255, 255, 255, 255})
	dx := math.Sin(g.player.angle) * 40
	dy := math.Cos(g.player.angle) * 40
	ebitenutil.DrawLine(screen, g.player.pos.x, g.player.pos.y,
		g.player.pos.x+dx, g.player.pos.y+dy, color.RGBA{255, 0, 0, 255})

	_ = ebitenutil.DebugPrint(screen, fmt.Sprintf("speed: %.2f pos:%.1f,%.1f ang:%.2f",
		g.player.speed, g.player.pos.x, g.player.pos.y, g.player.angle))

	ebitenutil.DrawLine(screen, wall.p1.x, wall.p1.y, wall.p2.x, wall.p2.y, color.RGBA{255, 255, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
