package main

import (
	"math/rand"
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
	screenHeight := int32(450)
	rl.InitAudioDevice()
	eat_noise := rl.LoadSound("sound/eat.wav")
	defer rl.UnloadSound(eat_noise)

	rand.Seed(time.Now().UnixNano())
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")
	rl.SetTargetFPS(60)

	bird_down := rl.LoadImage("assets/bird-down.png")
	defer rl.UnloadImage(bird_down)
	bird_up := rl.LoadImage("assets/bird-up.png")
	defer rl.UnloadImage(bird_up)

	texture := rl.LoadTextureFromImage(bird_up)
	defer rl.UnloadTexture(texture)

	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40
	var score int = 0

	Apples := []Apple{}
	var apple_loc int32

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		rl.DrawText("Current Score: "+strconv.Itoa(score), 10, 10, 30, rl.LightGray)

		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			y_coords -= 5
		} else {
			texture = rl.LoadTextureFromImage(bird_down)
			y_coords += 5
		}

		if len(Apples) == 0 {
			apple_loc = int32(rand.Intn(int(screenHeight-25)))
			current_apple := Apple{screenWidth, apple_loc, 25, 25, rl.Red}
			Apples = append(Apples, current_apple)
		}

		for io, current_apple := range Apples {
			rl.DrawRectangle(current_apple.posX, current_apple.posY, current_apple.width, current_apple.height, current_apple.Color)
			Apples[io].posX = Apples[io].posX - 5

			if current_apple.posX < 0 {
				Apples = append(Apples[:io], Apples[io+1:]...)
				score--
			}

			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords), 34, 24), rl.NewRectangle(float32(current_apple.posX), float32(current_apple.posY), float32(current_apple.width), float32(current_apple.height)) {
				Apples = append(Apples[:io], Apples[io+1:]...)
				score++
				rl.PlaySoundMulti(eat_noise)
			}
		}

		if y_coords > screenHeight {
			rl.DrawText("Your final score is: "+strconv.Itoa(score), 30, 40, 30, rl.Red)
		}

		rl.EndDrawing()
		time.Sleep(16 * time.Millisecond)
	}

	rl.StopSoundMulti()
	rl.CloseWindow()
}
