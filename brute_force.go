package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	bruteForceProgress int
	bruteForceTarget   int
	bruteForceKeys     []ebiten.Key
	currentKeyIndex    int
	bruteForceTimer    time.Time
	bruteForceTimeLimit time.Duration
)

func initBruteForceMiniGame() {
	bruteForceProgress = 0
	bruteForceTarget = 10
	bruteForceKeys = []ebiten.Key{ebiten.KeyA, ebiten.KeyS, ebiten.KeyD, ebiten.KeyF}
	currentKeyIndex = rand.Intn(len(bruteForceKeys))
	bruteForceTimer = time.Now()
	bruteForceTimeLimit = 5 * time.Second
}

func drawBruteForceMiniGame(screen *ebiten.Image) {
	timeRemaining := bruteForceTimeLimit - time.Since(bruteForceTimer)
	if timeRemaining < 0 {
		timeRemaining = 0
	}

	progressText := fmt.Sprintf("Press '%s' rapidly to hack: %d/%d\nTime Remaining: %.1f seconds", bruteForceKeys[currentKeyIndex], bruteForceProgress, bruteForceTarget, timeRemaining.Seconds())
	drawCenteredText(screen, progressText, 100, fontFace)

	if bruteForceProgress >= bruteForceTarget {
		completeMission(true)
	} else if timeRemaining <= 0 {
		completeMission(false)
	}
}

func handleBruteForceMiniGameInput() {
	if inpututil.IsKeyJustPressed(bruteForceKeys[currentKeyIndex]) {
		bruteForceProgress++
		currentKeyIndex = rand.Intn(len(bruteForceKeys))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		completeMission(false)
	}
}
