package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const windowScale = 2

var Version = "0.0.0"

func main() {
	fmt.Println("Starting, version", Version)
	rand.Seed(time.Now().UnixNano())

	g := NewGame(3)

	ebiten.SetWindowSize(g.width*windowScale, g.height*windowScale)
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
