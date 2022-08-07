package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	width  int
	height int
	player *Player
}

const baseWidth = 160
const baseHeight = 120
const fovScale = 100

var wall = &Wall{
	points: []Point{
		{x: 50, y: 250, z: 0},
		{x: 150, y: 250, z: 0},
		{x: 50, y: 250, z: 30},
		{x: 150, y: 250, z: 30},
	},
}

func NewGame(res int) *Game {
	return &Game{
		width:  baseWidth * res,
		height: baseHeight * res,

		player: &Player{
			pos:   NewPoint(100, 100, 15),
			angle: 0,
			look:  1,
			speed: 0.0,
		},
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.processInputs()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w2 := float64(g.width / 2.0)
	h2 := float64(g.height / 2.0)

	// Precompute the sin and cos of the player's angle
	cs := math.Cos(g.player.angle)
	sn := math.Sin(g.player.angle)

	// Render all wall points
	for _, wallP := range wall.points {
		// Translate to player position
		x := wallP.x - g.player.pos.x
		y := wallP.y - g.player.pos.y

		// worldP position of wall
		worldP := Point{
			x: x*cs - y*sn,
			y: y*cs + x*sn,
		}

		// Calc Z needs the world Y coord so we do it separately
		worldP.z = wallP.z - g.player.pos.z + ((g.player.look * worldP.y) / 32)

		// screen position of wall points
		scrP := Point{
			x: ((worldP.x * fovScale) / worldP.y) + w2,
			y: ((worldP.z * fovScale) / worldP.y) + h2,
		}

		c := color.RGBA{255, 255, 255, 255}
		if wallP.z > 0 {
			c = color.RGBA{255, 0, 255, 255}
		}

		// Flip the Y coord for drawing
		ebitenutil.DrawRect(screen, scrP.x, float64(g.height)-scrP.y, 3, 3, c)
	}

	g.drawDebug(screen)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.player.pos.x-2, g.player.pos.y-2, 4, 4, color.RGBA{255, 255, 255, 255})
	dx := math.Sin(g.player.angle) * 40
	dy := math.Cos(g.player.angle) * 40
	ebitenutil.DrawLine(screen, g.player.pos.x, g.player.pos.y,
		g.player.pos.x+dx, g.player.pos.y+dy, color.RGBA{255, 0, 0, 255})

	_ = ebitenutil.DebugPrint(screen, fmt.Sprintf("speed: %.2f pos:%.1f,%.1f ang:%.2f",
		g.player.speed, g.player.pos.x, g.player.pos.y, g.player.angle))

	ebitenutil.DrawLine(screen, wall.points[0].x, wall.points[0].y, wall.points[1].x, wall.points[1].y, color.RGBA{255, 255, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
