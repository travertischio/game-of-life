package engine

import (
	"time"

	"github.com/travertischio/game-of-life/pkg/world"
)

func RunGame() {
	world.Create()
	for {
		world.Print()
		world.Turn()
		time.Sleep(time.Millisecond * time.Duration(50))
	}
}
