package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var sqlTargets []SQLTarget
var sqlSuccessRate float64

type SQLTarget struct {
	Query string
	IsHit bool
	X, Y  float64
	Width, Height int
}

func initSQLInjectionMiniGame() {
	sqlTargets = []SQLTarget{
		{"SELECT * FROM users WHERE 'a'='a'", false, 100, 100, 300, 30},
		{"SELECT * FROM orders WHERE '1'='1'", false, 200, 200, 300, 30},
		{"SELECT * FROM products WHERE 'x'='x'", false, 300, 150, 300, 30},
		{"SELECT * FROM logins WHERE 'admin'='admin'", false, 400, 100, 300, 30},
	}
	sqlSuccessRate = 0.0
}

func handleSQLInjectionMiniGameInput() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		for i := range sqlTargets {
			target := &sqlTargets[i]
			if cx >= int(target.X) && cx <= int(target.X)+target.Width &&
				cy >= int(target.Y) && cy <= int(target.Y)+target.Height {
				target.IsHit = true
				sqlSuccessRate += 0.25 // Increase success rate for each hit
			}
		}
	}

	if sqlSuccessRate >= 1.0 {
		completeMission(true)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMissionSelect)
	}
}

func drawSQLInjectionMiniGame(screen *ebiten.Image) {
	for _, target := range sqlTargets {
		col := color.RGBA{255, 0, 0, 255}
		if target.IsHit {
			col = color.RGBA{0, 255, 0, 255}
		}
		ebitenutil.DrawRect(screen, target.X, target.Y, float64(target.Width), float64(target.Height), col)
		ebitenutil.DebugPrintAt(screen, target.Query, int(target.X), int(target.Y)-10)
	}
}
