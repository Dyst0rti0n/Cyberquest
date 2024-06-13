package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DialogueOption struct {
	Text    string
	IsCorrect bool
}

var (
	socialEngineeringStep int
	dialogueOptions       [][]DialogueOption
)

func initSocialEngineeringMiniGame() {
	socialEngineeringStep = 0
	dialogueOptions = [][]DialogueOption{
		{
			{"Ask about their day", false},
			{"Ask for their password", false},
			{"Mention a mutual contact", true},
		},
		{
			{"Ask about their work", true},
			{"Ask for sensitive info", false},
			{"Talk about the weather", false},
		},
		// Add more steps as needed
	}
}

func drawSocialEngineeringMiniGame(screen *ebiten.Image) {
	if socialEngineeringStep >= len(dialogueOptions) {
		completeMission(true)
		return
	}

	options := dialogueOptions[socialEngineeringStep]
	conversationText := fmt.Sprintf("Step %d: Choose your dialogue option", socialEngineeringStep+1)
	drawCenteredText(screen, conversationText, 100, fontFace)
	for i, option := range options {
		drawCenteredText(screen, fmt.Sprintf("%d. %s", i+1, option.Text), 150+(i*30), fontFace)
	}
	drawCenteredText(screen, "Press 1, 2, or 3 to choose\nPress Escape to fail", 250, fontFace)
}

func handleSocialEngineeringMiniGameInput() {
	if socialEngineeringStep >= len(dialogueOptions) {
		return
	}

	options := dialogueOptions[socialEngineeringStep]
	if inpututil.IsKeyJustPressed(ebiten.Key1) && options[0].IsCorrect {
		socialEngineeringStep++
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) && options[1].IsCorrect {
		socialEngineeringStep++
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) && options[2].IsCorrect {
		socialEngineeringStep++
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		completeMission(false)
	} else {
		completeMission(false)
	}
}
