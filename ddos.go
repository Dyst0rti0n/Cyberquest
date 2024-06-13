package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ddosTargets []DDoSTarget
var ddosSuccessRate float64

type DDoSTarget struct {
	Name      string
	IsHit     bool
	X, Y      float64
	Width, Height int
}

func initDDoSMiniGame() {
	ddosTargets = []DDoSTarget{
		{"Target 1", false, 100, 100, 50, 50},
		{"Target 2", false, 200, 200, 50, 50},
		{"Target 3", false, 300, 150, 50, 50},
		{"Target 4", false, 400, 100, 50, 50},
	}
	ddosSuccessRate = 0.0
}

func handleDDoSMiniGameInput() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		for i := range ddosTargets {
			target := &ddosTargets[i]
			if cx >= int(target.X) && cx <= int(target.X)+target.Width &&
				cy >= int(target.Y) && cy <= int(target.Y)+target.Height {
				target.IsHit = true
				ddosSuccessRate += 0.25 // Increase success rate for each hit
			}
		}
	}

	if ddosSuccessRate >= 1.0 {
		completeMission(true)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMissionSelect)
	}
}

func drawDDoSMiniGame(screen *ebiten.Image) {
	for _, target := range ddosTargets {
		col := color.RGBA{255, 0, 0, 255}
		if target.IsHit {
			col = color.RGBA{0, 255, 0, 255}
		}
		ebitenutil.DrawRect(screen, target.X, target.Y, float64(target.Width), float64(target.Height), col)
	}
}
