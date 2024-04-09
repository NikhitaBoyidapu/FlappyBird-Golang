// working startPage
package main

import (
	"image/color"
	"log"
	"os"
	"os/exec"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	buttonWidth   = 200
	buttonHeight  = 60
	buttonPadding = 20
)

var (
	iconImage   *ebiten.Image
	buttonX     = (screenWidth - buttonWidth) / 2
	buttonY     = (screenHeight - buttonHeight) / 2
	imageWidth  = 200 // Adjust as per requirement
	imageHeight = 200 // Adjust as per requirement
)

func update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX >= buttonX && mouseX <= buttonX+buttonWidth && mouseY >= buttonY && mouseY <= buttonY+buttonHeight {
			// Handle button click
			log.Println("Button clicked!")
			err := exec.Command("go", "run", "../gamePage/gamePage.go").Run()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0) // Exit after opening gamePage.go
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Handle enter key press
	}

	// Clear the screen
	screen.Fill(color.White)

	// Draw the icon image
	if iconImage != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(screenWidth/2-imageWidth/2), float64(screenHeight/4-imageHeight/2))
		screen.DrawImage(iconImage, opts)
	}

	// Draw the "New Game" button
	drawButton(screen, buttonX, buttonY, "New Game")

	return nil
}

func drawButton(screen *ebiten.Image, x, y int, text string) {
	// Draw the button background
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(buttonWidth), float64(buttonHeight), color.RGBA{100, 100, 100, 255})

	// Draw the text inside the button
	textWidth := ebitenutil.MeasureText(text)
	textX := x + (buttonWidth-textWidth)/2
	textY := y + (buttonHeight-ebitenutil.DefaultFont.Metrics().Height.Ceil())/2
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

func main() {
	// Load the icon image
	var err error
	iconImage, _, err = ebitenutil.NewImageFromFile("Icon.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new game window
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Start Page"); err != nil {
		log.Fatal(err)
	}
}
