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

//const baseAspectRatio = 1.3333333333

func NewGame(res int) *Game {
	return &Game{
		width:  baseWidth * res,
		height: baseHeight * res,
		player: &Player{
			pos:   NewPoint(10, 10, 20),
			a:     0,
			speed: 0.0,
		},
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.processInputs()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// for i := 0; i < g.width; i++ {
	// 	for j := 0; j < g.height; j++ {
	// 		screen.Set(i, j, color.RGBA{uint8(i), uint8(j), 0, 255})
	// 	}
	// }
	screen.Fill(color.RGBA{0, 0, 50, 255})
	ebitenutil.DrawRect(screen, g.player.pos.x, g.player.pos.y, 5, 5, color.RGBA{255, 255, 255, 255})
	dx := math.Sin(g.player.a) * 40
	dy := math.Cos(g.player.a) * 40
	ebitenutil.DrawLine(screen, g.player.pos.x, g.player.pos.y,
		g.player.pos.x+dx, g.player.pos.y+dy, color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("speed: %f pos:%+v ang:%f", g.player.speed, g.player.pos, g.player.a))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
