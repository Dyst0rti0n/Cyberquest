package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	volume            = 50
	screenWidth       = 800
	screenHeight      = 600
	controlScheme     = "WASD"
	settingsOptions   = []string{"Volume", "Screen Resolution", "Controls"}
	screenResolutions = []struct {
		width, height int
	}{
		{800, 600},
		{1024, 768},
		{1280, 720},
		{1920, 1080},
	}
	currentResolutionIndex = 0
	selectedOption         = 0
)

func saveSettings() {
	file, err := os.Create("settings.txt")
	if err != nil {
		log.Println("Error saving settings:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "Volume: %d\n", volume)
	fmt.Fprintf(file, "ScreenWidth: %d\n", screenWidth)
	fmt.Fprintf(file, "ScreenHeight: %d\n", screenHeight)
	fmt.Fprintf(file, "ControlScheme: %s\n", controlScheme)
}

func loadSettings() {
	file, err := os.Open("settings.txt")
	if err != nil {
		log.Println("Error loading settings:", err)
		return
	}
	defer file.Close()

	fmt.Fscanf(file, "Volume: %d\n", &volume)
	fmt.Fscanf(file, "ScreenWidth: %d\n", &screenWidth)
	fmt.Fscanf(file, "ScreenHeight: %d\n", &screenHeight)
	fmt.Fscanf(file, "ControlScheme: %s\n", &controlScheme)
	for i, res := range screenResolutions {
		if res.width == screenWidth && res.height == screenHeight {
			currentResolutionIndex = i
			break
		}
	}
}

func handleSettingsInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		saveSettings()
		transitionToState(StateMainMenu)
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		selectedOption = (selectedOption - 1 + len(settingsOptions)) % len(settingsOptions)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		selectedOption = (selectedOption + 1) % len(settingsOptions)
	}

	switch selectedOption {
	case 0: // Volume
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			volume = max(0, volume-10)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			volume = min(100, volume+10)
		}
	case 1: // Screen Resolution
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			currentResolutionIndex = (currentResolutionIndex - 1 + len(screenResolutions)) % len(screenResolutions)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			currentResolutionIndex = (currentResolutionIndex + 1) % len(screenResolutions)
		}
		screenWidth = screenResolutions[currentResolutionIndex].width
		screenHeight = screenResolutions[currentResolutionIndex].height
		ebiten.SetWindowSize(screenWidth, screenHeight)
	case 2: // Controls
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			if controlScheme == "WASD" {
				controlScheme = "Arrow Keys"
			} else {
				controlScheme = "WASD"
			}
		}
	}
}

func drawSettings(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x10, 0x10, 0x10, 0xff})

	settingsText := "Settings:\n\n"
	for i, option := range settingsOptions {
		prefix := "  "
		if i == selectedOption {
			prefix = "> "
		}

		switch i {
		case 0:
			settingsText += fmt.Sprintf("%s%s: %d\n", prefix, option, volume)
		case 1:
			res := screenResolutions[currentResolutionIndex]
			settingsText += fmt.Sprintf("%s%s: %dx%d\n", prefix, option, res.width, res.height)
		case 2:
			settingsText += fmt.Sprintf("%s%s: %s\n", prefix, option, controlScheme)
		}
	}

	settingsText += "\nPress Escape to Return to Menu"
	drawCenteredText(screen, settingsText, 100, fontFace)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
