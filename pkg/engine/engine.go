package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/travertischio/game-of-life/pkg/world"
)

const (
	screenWidth  = 120
	screenHeight = 120
	worldWidth   = 100
	worldHeight  = 100
)

var backgroundColor = color.RGBA{0x30, 0xD5, 0xC8, 0xff}

// Game manages the world and world image
type Game struct {
	startScreen bool
	world       *world.World
	WorldImage  *ebiten.Image

	xMargin int
	yMargin int
}

// NewGame generates a new Game object
func NewGame() *Game {
	x := (screenWidth - worldWidth) / 2
	y := (screenHeight - worldHeight) / 2

	return &Game{
		startScreen: true,
		world:       world.Create(worldWidth, worldHeight),
		xMargin:     x,
		yMargin:     y,
	}
}

// Update updates the current game state
func (g *Game) Update() error {
	if g.startScreen {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()

			if x > g.xMargin && x < g.world.Width+g.xMargin && y > g.yMargin && y < g.world.Height+g.yMargin {
				cellX := x - g.xMargin
				cellY := y - g.yMargin

				g.world.Update(cellX, cellY)
				// The start button is the entire bottom of the screen
			} else if y > (g.world.Height + g.yMargin) {
				g.startScreen = false
			}
		}
	} else {
		g.world.Turn()

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			_, y := ebiten.CursorPosition()

			// The start button is the entire bottom of the screen
			if y > (g.world.Height + g.yMargin) {
				g.startScreen = true
			}
		}
	}

	return nil
}

// Draw draws the current world to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	if g.WorldImage == nil {
		g.WorldImage = ebiten.NewImage(g.world.Width, g.world.Height)
	}

	screen.Fill(backgroundColor)
	g.world.Draw(g.WorldImage)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.xMargin), float64(g.yMargin))
	screen.DrawImage(g.WorldImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 120, 120
}
