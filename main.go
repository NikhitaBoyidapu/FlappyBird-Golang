package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(screenWidth, screenHeight, "FlappyBird")
	rl.SetTargetFPS(60)
	bird_down := rl.LoadImage("assets/bird-down.png")
	bird_up := rl.LoadImage("assets/bird-up.png")
	texture := rl.LoadTextureFromImage(bird_up)
	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			y_coords += 5
		} else {
			texture = rl.LoadTextureFromImage(bird_down)
			y_coords -= 5
		}
		rl.EndDrawing()
	}
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
