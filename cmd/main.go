package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/travertischio/game-of-life/pkg/engine"
)

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Game of Life")
	ebiten.SetMaxTPS(5)

	g := engine.NewGame()
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
