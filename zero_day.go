package main

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	vulnerabilities []string
	foundVulnerability bool
	selectedVulnerability int
)

func initZeroDayMiniGame() {
	vulnerabilities = []string{
		"Buffer Overflow",
		"SQL Injection",
		"Cross-Site Scripting",
		"Path Traversal",
	}
	foundVulnerability = false
	selectedVulnerability = rand.Intn(len(vulnerabilities))
}

func drawZeroDayMiniGame(screen *ebiten.Image) {
	if foundVulnerability {
		drawCenteredText(screen, fmt.Sprintf("Exploit found: %s\nPress Enter to exploit\nPress Escape to fail", vulnerabilities[selectedVulnerability]), 100, fontFace)
	} else {
		drawCenteredText(screen, "Searching for vulnerabilities...\nPress Enter to search\nPress Escape to fail", 100, fontFace)
	}
}

func handleZeroDayMiniGameInput() {
	if foundVulnerability {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			completeMission(true)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			completeMission(false)
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if rand.Float64() < 0.5 { // Example probability
				foundVulnerability = true
			} else {
				completeMission(false)
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			completeMission(false)
		}
	}
}
