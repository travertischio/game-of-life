package world

import (
	"fmt"
)

const (
	height = 10
	width  = 10
)

var world [height][width]bool

func Create() {
	world[4][5] = true
	world[4][4] = true
	world[4][3] = true
}

func Print() {
	worldString := ""
	for i := range world {
		for _, val := range world[i] {
			box := " "
			if val {
				box = "X"
			}
			worldString += box
		}
		worldString += "\n"
	}

	fmt.Println(worldString)
}

func Turn() {
	var newWorld [height][width]bool

	for i := range newWorld {
		for j := range newWorld[i] {
			newWorld[i][j] = liveOrDie(i, j)
		}
	}

	world = newWorld
}

func liveOrDie(y, x int) bool {
	live := 0
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i >= height {
			continue
		}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j >= width {
				continue
			} else if i == y && j == x {
				continue
			}

			if world[i][j] {
				live++
			}
		}
	}

	if world[y][x] {
		if live == 2 || live == 3 {
			return true
		}
	} else if live == 3 {
		return true
	}

	return false
}
