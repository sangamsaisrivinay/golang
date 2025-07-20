package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	x, y float64
}

func (g *Game) Update() error {
	// Move the player square using arrow keys
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Define black and white colors using RGBA values
	black := color.RGBA{0, 0, 0, 255}  // Black
	white := color.RGBA{255, 255, 255, 255}  // White

	// Clear the screen to black
	screen.Fill(black)

	// Draw the player square (white)
	ebitenutil.DrawRect(screen, g.x, g.y, 20, 20, white)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the dimensions of the game screen
	return screenWidth, screenHeight
}

func main() {
	// Create a new Game instance
	game := &Game{x: screenWidth / 2, y: screenHeight / 2}

	// Start the game and handle errors
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
