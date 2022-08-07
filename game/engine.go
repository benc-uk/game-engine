package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	width  int
	height int
}

const baseWidth = 160
const baseHeight = 120

//const baseAspectRatio = 1.3333333333

func NewGame(res int) *Game {
	return &Game{
		width:  baseWidth * res,
		height: baseHeight * res,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	//g.updateFirePixels()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			screen.Set(i, j, color.RGBA{uint8(i), uint8(j), 0, 255})
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
