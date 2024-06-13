package main

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	phishingText    string
	emailPhrases    []string
	selectedPhrases []string
)

func initPhishingMiniGame() {
	phishingText = ""
	emailPhrases = []string{
		"Dear user,",
		"Your account has been compromised.",
		"Please click the link below to reset your password.",
		"Thank you for your attention.",
		"Sincerely, The Security Team",
	}
	selectedPhrases = []string{}
}

func drawPhishingMiniGame(screen *ebiten.Image) {
	assembledEmail := strings.Join(selectedPhrases, " ")
	drawCenteredText(screen, "Craft your Phishing Email:\nPress 1-5 to add phrases\nPress Enter to send\nPress Escape to fail", 50, fontFace)
	drawCenteredText(screen, assembledEmail, 150, fontFace)

	for i, phrase := range emailPhrases {
		drawCenteredText(screen, fmt.Sprintf("%d. %s", i+1, phrase), 300+(i*30), fontFace)
	}
}

func handlePhishingMiniGameInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if len(selectedPhrases) >= 3 { // Example minimum number of phrases for a valid email
			completeMission(true)
		} else {
			completeMission(false)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		completeMission(false)
	} else {
		for i := 1; i <= len(emailPhrases); i++ {
			if inpututil.IsKeyJustPressed(ebiten.Key(int(ebiten.Key1) + i - 1)) { // Updated key check
				selectedPhrases = append(selectedPhrases, emailPhrases[i-1])
			}
		}
	}
}
