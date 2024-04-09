package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	textboxWidth  = 200
	textboxHeight = 40
)

var (
	nameTextbox string
)

func update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Handle mouse click
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Handle enter key press
	}

	// Clear the screen
	screen.Fill(color.White)
	drawTextbox(screen, 100, 100, "Flappy Bird")
	// Draw the textbox
	drawTextbox(screen, 100, 100, "Enter your name:")

	return nil
}

func drawTextbox(screen *ebiten.Image, x, y int, text string) {
	// Draw the textbox background
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(textboxWidth), float64(textboxHeight), color.RGBA{0, 0, 0, 255})

	// Draw the text inside the textbox
	ebitenutil.DebugPrintAt(screen, text, x+10, y+textboxHeight/2)
}

func main() {
	// Create a new game window
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Start Page"); err != nil {
		panic(err)
	}
}
