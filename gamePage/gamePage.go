package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Apple struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)
	rl.InitAudioDevice()
	eat_noise := rl.LoadSound("../sound/eat.wav")
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")
	rl.SetTargetFPS(60)
	bird_down := rl.LoadImage("../assets/bird-down.png")
	bird_up := rl.LoadImage("../assets/bird-up.png")
	texture := rl.LoadTextureFromImage(bird_up)
	rand.Seed(time.Now().UnixNano())
	var apple_loc int = rand.Intn(450-2+1) - 2
	Apples := []Apple{}
	current_apple := Apple{screenWidth, int32(apple_loc), 25, 25, rl.Red}
	Apples = append(Apples, current_apple)
	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40
	var score int = 0
	var lives int = 3 // Number of lives
	var name string

	// Receive name from start page
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "Player" // Default name if not provided
	}

	for !rl.WindowShouldClose() && lives > 0 {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// Draw the bird
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		// rl.DrawText("Current Score: "+strconv.Itoa(score), 0, 0, 30, rl.Black)
		// rl.DrawText("Lives: "+strconv.Itoa(lives), 0, 40, 30, rl.Black)
		rl.DrawText("Current Score: "+strconv.Itoa(score), 10, 0, 30, rl.Black)
		rl.DrawText("Lives: "+strconv.Itoa(lives), 310, 0, 30, rl.Black)
		rl.DrawText("Name: "+name, 450, 0, 30, rl.Black)
		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			y_coords -= 5
		} else {
			texture = rl.LoadTextureFromImage(bird_down)
			y_coords += 5
		}

		for io, current_apple := range Apples {
			rl.DrawRectangle(current_apple.posX, current_apple.posY, current_apple.width, current_apple.height, current_apple.Color)
			Apples[io].posX = Apples[io].posX - 5
			if current_apple.posX < 0 {
				Apples[io].posX = 800
				// Apples[io].posY = int32(rand.Intn(580-2+1) - 2)
				Apples[io].posY = int32(rand.Intn(460-2+1) - 2) // Adjusted range to 460 to avoid overlap with name text
				score--
			}
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords), float32(34), float32(24)), rl.NewRectangle(float32(current_apple.posX), float32(current_apple.posY), float32(current_apple.width), float32(current_apple.height))) {
				Apples[io].posX = 800
				//Apples[io].posY = int32(rand.Intn(580-2+1) - 2)
				Apples[io].posY = int32(rand.Intn(460-2+1) - 2) // Adjusted range to 460 to avoid overlap with name text
				score++
				rl.PlaySound(eat_noise)
			}
		}

		// Reduce lives if bird goes below screen
		if y_coords > 600 {
			lives--
			if lives > 0 {
				y_coords = screenHeight/2 - texture.Height/2 - 40
			}
			if lives == 0 {
				rl.UnloadTexture(texture)
				Apples = nil
				rl.DrawText("Your final score is: "+strconv.Itoa(score), 30, 40, 30, rl.Red)
			}
		}

		rl.EndDrawing()
		time.Sleep(50000000)
	}

	rl.UnloadSound(eat_noise)
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
