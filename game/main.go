package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const windowScale = 2

func main() {
	rand.Seed(time.Now().UnixNano())

	g := NewGame(4)

	ebiten.SetWindowSize(g.width*windowScale, g.height*windowScale)
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
