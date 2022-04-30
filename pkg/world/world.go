package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var otherColor = color.RGBA{0x00, 0xff, 0x00, 0xff}

// World contains all the state of the world
type World struct {
	Height int
	Width  int
	cells  [][]bool
}

// Create generates a new world
func Create(width, height int) *World {
	c := make([][]bool, height)
	for i := range c {
		c[i] = make([]bool, width)
	}

	w := &World{
		Width:  width,
		Height: height,
		cells:  c,
	}

	w.cells[4][5] = true
	w.cells[4][4] = true
	w.cells[4][3] = true

	return w
}

// Turn executes a turn on the world state
func (w *World) Turn() {
	newWorld := Create(w.Width, w.Height)

	for i := range newWorld.cells {
		for j := range newWorld.cells[i] {
			newWorld.cells[i][j] = w.liveOrDie(i, j)
		}
	}

	w.cells = newWorld.cells
}

// liveOrDie returns if a given cell should live or dir based on the current world
func (w *World) liveOrDie(y, x int) bool {
	live := 0
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i >= w.Height {
			continue
		}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j >= w.Width {
				continue
			} else if i == y && j == x {
				continue
			}

			if w.cells[i][j] {
				live++
			}
		}
	}

	if w.cells[y][x] {
		if live == 2 || live == 3 {
			return true
		}
	} else if live == 3 {
		return true
	}

	return false
}

// Draw draws the world to a world image
func (w *World) Draw(worldImage *ebiten.Image) {
	worldImage.Fill(color.White)
	cellImage := ebiten.NewImage(1, 1)
	cellImage.Fill(color.Black)

	for j := 0; j < w.Height; j++ {
		for i := 0; i < w.Width; i++ {
			x := i
			y := j

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorM.ScaleWithColor(otherColor)

			if w.cells[y][x] {
				worldImage.DrawImage(cellImage, op)
			}

		}
	}
}
