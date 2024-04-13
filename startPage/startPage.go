package main

import (
	"os/exec"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TextInput represents a text input box.
type TextInput struct {
	rect       rl.Rectangle
	text       string
	active     bool
	fontSize   int32
	fontColor  rl.Color
	normalText string // Text to display when not active
	label      string // Label for the text input box
}

// NewTextInput creates a new text input box.
func NewTextInput(rect rl.Rectangle, label string, normalText string, fontSize int32, fontColor rl.Color) TextInput {
	return TextInput{
		rect:       rect,
		label:      label,
		normalText: normalText,
		fontSize:   fontSize,
		fontColor:  fontColor,
	}
}

// Draw draws the text input box.
func (t *TextInput) Draw() {
	// Draw label
	rl.DrawText(t.label, int32(t.rect.X), int32(t.rect.Y-25), t.fontSize, t.fontColor)

	if t.active {
		rl.DrawRectangleRec(t.rect, rl.White) // Background color
		rl.DrawRectangleLinesEx(t.rect, 1, rl.Black)
		rl.DrawText(t.text, int32(t.rect.X+5), int32(t.rect.Y+10), t.fontSize, t.fontColor)
	} else {
		rl.DrawRectangleRec(t.rect, rl.LightGray)
		rl.DrawRectangleLinesEx(t.rect, 1, rl.Gray)
		rl.DrawText(t.normalText, int32(t.rect.X+5), int32(t.rect.Y+10), t.fontSize, rl.Gray)
	}
}

// Update handles text input for the text input box.
func (t *TextInput) Update() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), t.rect) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		t.active = true
	} else if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		t.active = false
	}

	if t.active {
		key := rl.GetKeyPressed()
		if key >= 32 && key <= 125 && len(t.text) < 16 {
			t.text += string(key)
		} else if key == rl.KeyBackspace && len(t.text) > 0 {
			t.text = t.text[:len(t.text)-1]
		}
	}
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)

	rl.InitWindow(screenWidth, screenHeight, "Window with Texture and Button")

	// Load texture
	texture := rl.LoadTexture("Icon.png")
	if texture.ID == 0 {
		rl.TraceLog(rl.LogError, "Failed to load texture")
		return
	}

	// Padding for the texture from the left edge
	padding := int32(50)

	// Button position and size
	buttonWidth := int32(200)
	buttonHeight := int32(50)
	buttonX := screenWidth - buttonWidth - 50 // 50 pixels padding from the right edge
	buttonY := screenHeight/2 - buttonHeight/2

	// Padding between button and text box
	paddingY := int32(20)

	// Username text box
	usernameBox := NewTextInput(rl.NewRectangle(float32(buttonX), float32(buttonY-buttonHeight-paddingY), float32(buttonWidth), float32(buttonHeight)), "Username:", "", 15, rl.Black)

	// Main loop
	for !rl.WindowShouldClose() {
		// Check if the button is clicked
		mousePosition := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePosition, rl.NewRectangle(float32(buttonX), float32(buttonY), float32(buttonWidth), float32(buttonHeight))) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if usernameBox.text == "" {
				usernameBox.text = "user"
			}
			cmd := exec.Command("go", "run", "../gamePage/gamePage.go", usernameBox.text)
			err := cmd.Run()
			if err != nil {
				rl.TraceLog(rl.LogError, "Failed to open gamePage.go:", err)
			}
			usernameBox.text = "" // Reset the username after starting the game
		}

		// Update username text box
		usernameBox.Update()

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(getBiscuitColor())

		// Draw texture with left padding
		textureX := padding
		textureY := screenHeight/2 - texture.Height/2
		rl.DrawTexture(texture, textureX, textureY, rl.White)

		// Draw button on the right half
		rl.DrawRectangle(buttonX, buttonY, buttonWidth, buttonHeight, rl.Green)
		rl.DrawText("New Game", buttonX+25, buttonY+15, 20, rl.White)

		// Draw username text box
		usernameBox.Draw()

		rl.EndDrawing()
	}

	// Unload texture
	rl.UnloadTexture(texture)

	rl.CloseWindow()
}

// getBiscuitColor returns the "biscuit" color (RGB: 255, 228, 196).
func getBiscuitColor() rl.Color {
	return rl.NewColor(255, 228, 196, 255)
}
