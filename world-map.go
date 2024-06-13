package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Location struct {
	X, Y         int
	Name         string
	IsMission    bool
	MissionIndex int
}

var worldLocations []Location

func initializeWorldLocations() {
	worldLocations = []Location{
		{100, 100, "City Center", true, 0},
		{200, 200, "Suburbs", false, -1},
		{300, 100, "Industrial Zone", true, 1},
		{400, 200, "Tech Park", false, -1},
		{500, 100, "Corporate HQ", true, 2},
	}
}

func handleWorldMapInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
		return
	}

	// Check for interactions with map nodes
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		for i, location := range worldLocations {
			if cx > location.X && cx < location.X+40 && cy > location.Y && cy < location.Y+40 {
				travelToLocation(i)
				break
			}
		}
	}
}

func drawWorldMap(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x10, 0x10, 0x10, 0xff})
	for _, location := range worldLocations {
		col := color.RGBA{0x00, 0xFF, 0x00, 0xFF} // Green for normal locations
		if location.IsMission {
			col = color.RGBA{0xFF, 0x00, 0x00, 0xFF} // Red for mission locations
		}
		drawLocation(screen, location.X, location.Y, col)
	}
	drawCenteredText(screen, "World Map\nClick on a location to travel", 30, fontFace)
}

func drawLocation(screen *ebiten.Image, x, y int, col color.Color) {
	rect := image.Rect(x, y, x+40, y+40)
	draw.Draw(screen, rect, &image.Uniform{col}, image.Point{}, draw.Src)
}

func travelToLocation(index int) {
	location := &worldLocations[index]
	currentMessage = fmt.Sprintf("Traveled to %s!", location.Name)
	if location.IsMission {
		startMission(location.MissionIndex)
	}
}
