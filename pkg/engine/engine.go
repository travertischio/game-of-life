package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/travertischio/game-of-life/pkg/world"
)

var backgroundColor = color.RGBA{0x30, 0xD5, 0xC8, 0xff}

// Game manages the world and world image
type Game struct {
	startScreen bool
	world       *world.World
	WorldImage  *ebiten.Image
}

// NewGame generates a new Game object
func NewGame() *Game {
	return &Game{
		startScreen: false,
		world:       world.Create(100, 100),
	}
}

// Update updates the current game state
func (g *Game) Update() error {
	if !g.startScreen {
		g.world.Turn()
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
	sw, sh := screen.Size()
	x := (sw - g.world.Width) / 2
	y := (sh - g.world.Height) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.WorldImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 120, 120
}
