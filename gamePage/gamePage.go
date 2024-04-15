package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUsername = "sql5698581"
	DBPassword = "9JsLU26Zye"
	DBHost     = "sql5.freemysqlhosting.net"
	DBPort     = "3306"
	DBName     = "sql5698581"
)

type Apple struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

type User struct {
	ID    int
	Name  string
	Lives int
	Score int
}

func main() {
	screenWidth := int32(1200)
	screenHeight := int32(800)
	rl.InitAudioDevice()
	eatNoise := rl.LoadSound("../sound/eat.wav")
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")
	rl.SetTargetFPS(60)
	birdDown := rl.LoadImage("../assets/bird-down.png")
	birdUp := rl.LoadImage("../assets/bird-up.png")
	texture := rl.LoadTextureFromImage(birdUp)
	rand.Seed(time.Now().UnixNano())
	var appleLoc int = rand.Intn(450-2+1) - 2
	Apples := []Apple{}
	currentApple := Apple{screenWidth, int32(appleLoc), 25, 25, rl.Red}
	Apples = append(Apples, currentApple)
	var xCoords int32 = screenWidth/2 - texture.Width/2
	var yCoords int32 = screenHeight/2 - texture.Height/2 - 40
	var score int = 0
	var lives int = 3 // Number of lives
	var name string
	saveClicked := false

	// Receive name from start page
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "Player" // Default name if not provided
	}

	saveButton := rl.NewRectangle(float32(screenWidth-120), 10, 110, 40)

	// Fetch high score from the database
	highScore := getHighScore()

	for !rl.WindowShouldClose() && lives > 0 {
		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(255, 228, 196, 255))

		// Draw the bird
		rl.DrawTexture(texture, xCoords, yCoords, rl.White)
		rl.DrawText("Score: "+strconv.Itoa(score), 10, 0, 30, rl.Black)
		rl.DrawText("Lives: "+strconv.Itoa(lives), 180, 0, 30, rl.Black)
		rl.DrawText("Name: "+name, 330, 0, 30, rl.Black)
		rl.DrawText("High Score: "+strconv.Itoa(highScore), 680, 0, 30, rl.Black)

		// Draw the save button
		rl.DrawRectangleRec(saveButton, rl.Green)
		rl.DrawText("Save", int32(saveButton.X)+25, int32(saveButton.Y)+10, 20, rl.White)

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), saveButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			saveClicked = true
		}

		if saveClicked {
			saveGameDetails(name, score, lives)
			rl.CloseWindow()
		}

		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(birdUp)
			yCoords -= 5
		} else {
			texture = rl.LoadTextureFromImage(birdDown)
			yCoords += 5
		}

		for io, currentApple := range Apples {
			rl.DrawRectangle(currentApple.posX, currentApple.posY, currentApple.width, currentApple.height, currentApple.Color)
			Apples[io].posX = Apples[io].posX - 5
			if currentApple.posX < 0 {
				Apples[io].posX = 800
				Apples[io].posY = int32(rand.Intn(580-2+1) - 2)
				score--
			}
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(xCoords), float32(yCoords), float32(34), float32(24)), rl.NewRectangle(float32(currentApple.posX), float32(currentApple.posY), float32(currentApple.width), float32(currentApple.height))) {
				Apples[io].posX = 800
				Apples[io].posY = int32(rand.Intn(580-2+1) - 2)
				score++
				rl.PlaySound(eatNoise)
			}
		}

		// Reduce lives if bird goes below screen
		if yCoords > 600 {
			lives--
			if lives > 0 {
				yCoords = screenHeight/2 - texture.Height/2 - 40
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

	rl.UnloadSound(eatNoise)
	rl.UnloadTexture(texture)
}

func saveGameDetails(name string, score, lives int) {
	// Connect to the database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Insert new user record
	newUser := User{Name: name, Lives: lives, Score: score}
	err = insertUser(db, newUser)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
	}
}

func insertUser(db *sql.DB, user User) error {
	query := `
        INSERT INTO User (username, lives, score)
        VALUES (?, ?, ?)
        ON DUPLICATE KEY UPDATE
        lives = VALUES(lives),
        score = VALUES(score)
    `
	_, err := db.Exec(query, user.Name, user.Lives, user.Score)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
	}
	return err
}

func getHighScore() int {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	var highScore int
	err = db.QueryRow("SELECT MAX(score) FROM User").Scan(&highScore)
	if err != nil {
		log.Fatal("Failed to get high score:", err)
	}

	return highScore
}
