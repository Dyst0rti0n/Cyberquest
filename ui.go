package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Button struct {
	X, Y, W, H int
	Label      string
	OnClick    func()
	Color      color.NRGBA 
}


var (
	fontFace       font.Face
	fontFaceTitle  font.Face
	fontFaceSmall  font.Face
	buttons        []Button
	tooltipText    string
	tooltipVisible bool
)

func loadFont() error {
	fontBytes, err := os.ReadFile("DejaVuSans.ttf")
	if err != nil {
		return fmt.Errorf("could not read font file: %w", err)
	}
	ttf, err := opentype.Parse(fontBytes)
	if err != nil {
		return fmt.Errorf("could not parse font: %w", err)
	}
	const dpi = 72
	fontFace, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return fmt.Errorf("could not create font face: %w", err)
	}
	fontFaceTitle, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return fmt.Errorf("could not create title font face: %w", err)
	}
	fontFaceSmall, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return fmt.Errorf("could not create small font face: %w", err)
	}
	return nil
}

func drawCenteredText(screen *ebiten.Image, textContent string, y int, face font.Face) {
	lines := strings.Split(textContent, "\n")
	for i, line := range lines {
		text.Draw(screen, line, face, (800-text.BoundString(face, line).Dx())/2, y+(i*30), color.White)
	}
}

func drawTerminalBorder(screen *ebiten.Image) {
	for x := 0; x < 800; x++ {
		screen.Set(x, 40, color.White)
		screen.Set(x, 560, color.White)
	}
	for y := 0; y < 600; y++ {
		screen.Set(0, y, color.White)
		screen.Set(799, y, color.White)
	}
}

func drawButton(screen *ebiten.Image, btn Button) {
	rect := image.Rect(btn.X, btn.Y, btn.X+btn.W, btn.Y+btn.H)
	btnColor := btn.Color

	// Animate button color
	animationFactor := (math.Sin(float64(ebiten.CurrentTPS())*0.1) + 1) / 2 // Oscillates between 0 and 1
	animatedColor := lightenColor(btn.Color, animationFactor*0.2)

	if btn.isHovered() {
		btnColor = lightenColor(animatedColor, 0.4)
	} else {
		btnColor = animatedColor
	}
	draw.Draw(screen, rect, &image.Uniform{btnColor}, image.Point{}, draw.Src)

	labelBounds, _ := font.BoundString(fontFace, btn.Label)
	labelWidth := (labelBounds.Max.X - labelBounds.Min.X).Ceil()
	labelHeight := (labelBounds.Max.Y - labelBounds.Min.Y).Ceil()

	labelX := btn.X + (btn.W-labelWidth)/2
	labelY := btn.Y + (btn.H+labelHeight)/2

	text.Draw(screen, btn.Label, fontFace, labelX, labelY, color.White)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && btn.isHovered() {
		playClickSound()
		btn.OnClick()
	}
}

func lightenColor(c color.NRGBA, factor float64) color.NRGBA {
	r := float64(c.R) + (255-float64(c.R))*factor
	g := float64(c.G) + (255-float64(c.G))*factor
	b := float64(c.B) + (255-float64(c.B))*factor
	return color.NRGBA{uint8(r), uint8(g), uint8(b), c.A}
}

func (btn Button) isHovered() bool {
	cx, cy := ebiten.CursorPosition()
	return btn.X <= cx && cx <= btn.X+btn.W && btn.Y <= cy && cy <= btn.Y+btn.H
}

func drawButtons(screen *ebiten.Image) {
	for _, btn := range buttons {
		drawButton(screen, btn)
	}
	buttons = []Button{}
}

func updateButtons() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		for _, btn := range buttons {
			if btn.X <= cx && cx <= btn.X+btn.W && btn.Y <= cy && cy <= btn.Y+btn.H {
				btn.OnClick()
			}
		}
	}
}

func drawStatus(screen *ebiten.Image) {
	statusText := fmt.Sprintf("Cash: $%.2f | Reputation: %d | %s", player.Money, player.Reputation, getCurrentGameStateName())
	text.Draw(screen, statusText, fontFace, 20, 30, color.NRGBA{0x00, 0xff, 0x00, 0xff})
}

func getCurrentGameStateName() string {
	switch gameState {
	case StateMainMenu:
		return "Main Menu"
	case StateNewPlayer:
		return "New Player"
	case StateStory:
		return "Story"
	case StateMissionSelect:
		return "Mission Select"
	case StateMission:
		return "Mission"
	case StateHackMethod:
		return "Hacking Method"
	case StateSuccess:
		return "Mission Success"
	case StateFailure:
		return "Mission Failure"
	case StateShop:
		return "Shop"
	case StateSettings:
		return "Settings"
	case StateHelp:
		return "Help"
	case StateCredits:
		return "Credits"
	case StateAbilities:
		return "Abilities"
	case StateMessaging:
		return "Messaging"
	case StateBruteForceMiniGame:
		return "Brute Force Hack"
	case StateSocialEngineeringMiniGame:
		return "Social Engineering Hack"
	case StatePhishingMiniGame:
		return "Phishing Hack"
	case StateZeroDayMiniGame:
		return "Zero-Day Exploit Hack"
	case StateDDoSMiniGame:
		return "DDoS Hack"
	case StateSQLInjectionMiniGame:
		return "SQL Injection Hack"
	default:
		return "Unknown"
	}
}

func drawMainMenu(screen *ebiten.Image) {
	// Draw the title with hacker-style introduction
	title := "CyberQuest: The Hacker's Journey"
	drawCenteredText(screen, title, 80, fontFaceTitle) // Adjusted title position

	// Draw menu options
	menuOptions := []struct {
		Label       string
		Action      func()
		Description string
	}{
		{"Start Game", func() { transitionToState(StateStory) }, "Begin your hacking journey"},
		{"Settings", func() { transitionToState(StateSettings) }, "Adjust your preferences"},
		{"Help", func() { transitionToState(StateHelp) }, "Get assistance"},
		{"Credits", func() { transitionToState(StateCredits) }, "See the credits"},
		{"Abilities", func() { transitionToState(StateAbilities) }, "View your abilities"},
		{"Messaging", func() { transitionToState(StateMessaging) }, "Check your messages"},
		{"Shop", func() { transitionToState(StateShop) }, "Purchase items"},
	}

	for i, option := range menuOptions {
		y := 160 + i*50 // Adjusted y position for more space
		drawButton(screen, Button{
			X:      300,
			Y:      y,
			W:      200,
			H:      40,
			Label:  option.Label,
			Color:  color.NRGBA{0x00, 0xff, 0x00, 0xff}, // Hacker green color
			OnClick: option.Action,
		})
	}
}


func drawNewPlayer(screen *ebiten.Image) {
	drawCenteredText(screen, fmt.Sprintf("Enter your name: %s\n\nPress Enter to Continue", playerName), 100, fontFace)
}

func drawStory(screen *ebiten.Image) {
	storyText := `You are a novice hacker, just starting to learn the ropes of the cyber world.
Your journey begins with simple tasks and progresses to more complex challenges.
Use your skills wisely, and remember, every choice has consequences.

Press Enter to begin your mission.`
	drawCenteredText(screen, storyText, 100, fontFace)
}

func drawMissionSelect(screen *ebiten.Image) {
	missionText := "Select your mission:\n\n"
	for i, mission := range missions {
		missionText += fmt.Sprintf("%d. %s\n", i+1, mission.Name)
	}
	missionText += "\nPress 1-8 to select a mission"
	drawCenteredText(screen, missionText, 100, fontFace)
}

func drawMission(screen *ebiten.Image) {
	missionDetails := fmt.Sprintf("Mission: %s\n\n%s\n\nObjective: %s\nReward: $%.2f\n\nPress Enter to Start Hacking\nPress Escape to Return to Mission Select", currentMission.Name, currentMission.Details, currentMission.Details, currentMission.Reward)
	drawCenteredText(screen, missionDetails, 100, fontFace)
}

func drawHackMethod(screen *ebiten.Image) {
	methodText := "Choose your hacking method:\n"
	if player.Skills["Brute Force Attack"] {
		methodText += "1. Brute Force Attack (Low success rate, high detection risk)\n"
	}
	if player.Skills["Social Engineering"] {
		methodText += "2. Social Engineering (Medium success rate, low detection risk)\n"
	}
	if player.Skills["Phishing Attack"] {
		methodText += "3. Phishing Attack (High success rate, medium detection risk)\n"
	}
	if player.Skills["Zero-Day Exploit"] {
		methodText += "4. Zero-Day Exploit (Very high success rate, low detection risk)\n"
	}
	if player.Skills["DDoS Attack"] {
		methodText += "5. DDoS Attack (Medium success rate, high detection risk)\n"
	}
	if player.Skills["SQL Injection"] {
		methodText += "6. SQL Injection (High success rate, medium detection risk)\n"
	}
	methodText += "Press the corresponding number to choose a method or Escape to go back"
	drawCenteredText(screen, methodText, 100, fontFace)
}

func drawMissionResult(screen *ebiten.Image) {
	drawCenteredText(screen, currentMessage+"\n\nPress Enter to Return to Mission Select\nPress Escape to Return to Main Menu", 100, fontFace)
}

func drawHelp(screen *ebiten.Image) {
	helpText := `Welcome to CyberQuest!
Here are some tips to get you started:
1. Use the Start Game button to begin your adventure.
2. Navigate through the levels to learn about different cybersecurity topics.
3. Use the hints if you get stuck.

Press Escape to Return to Menu`
	drawCenteredText(screen, helpText, 100, fontFace)
}

func drawCredits(screen *ebiten.Image) {
	creditsText := `CyberQuest: The Hacker's Journey
Developed by [Your Name]
Special Thanks to the Open Source Community
and Cybersecurity Experts for their invaluable input.

Press Escape to Return to Menu`
	drawCenteredText(screen, creditsText, 100, fontFace)
}

func drawAbilities(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	drawCenteredText(screen, "Abilities and Skills", 60, fontFace)

	abilitiesText := "Your Skills:\n\n"
	skillIndex := 1
	for _, skill := range playerSkillTree.Skills {
		status := "Locked"
		if skill.Level > 0 {
			status = "Unlocked"
		}
		abilitiesText += fmt.Sprintf("%d. %s: %s\n", skillIndex, skill.Name, status)
		skillIndex++
	}

	drawCenteredText(screen, abilitiesText, 150, fontFace)

	buttonIndex := 0
	for _, skill := range playerSkillTree.Skills {
		if skill.Level == 0 {
			yPosition := 180 + buttonIndex*40
			drawButton(screen, Button{
				X: 340, Y: yPosition, W: 120, H: 40,
				Label: "Unlock " + skill.Name,
				OnClick: func(skillName string) func() {
					return func() {
						if player.Money >= 500 { // Example cost
							player.Money -= 500
							playerSkillTree.UnlockSkill(skillName)
						}
					}
				}(skill.Name),
			})
			buttonIndex++
		}
	}

	drawButton(screen, Button{X: 340, Y: 480, W: 120, H: 40, Label: "Back", OnClick: func() { transitionToState(StateMainMenu) }})
}


func drawMessaging(screen *ebiten.Image) {
	messagingText := "Direct Messages:\n\n"
	y := 100
	for i, message := range messages {
		messagingText += fmt.Sprintf("From: %s\n%s\n\n", message.From, message.Content)
		// Add button for NPC interaction if applicable
		for j, npc := range npcs {
			if message.From == npc.Name {
				drawButton(screen, Button{
					X: 340, Y: y + i*100, W: 120, H: 40,
					Label: "Interact",
					OnClick: func() {
						interactWithNPC(j)
					},
				})
			}
		}
		y += 100
	}
	messagingText += "\nPress Escape to Return to Menu"
	drawCenteredText(screen, messagingText, 50, fontFace)
}
