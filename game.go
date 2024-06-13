package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct{}

func (g *Game) Update() error {
	handleTooltipUpdate()
	handleSkillUnlock()

	mouseX, mouseY = ebiten.CursorPosition()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
	}

	switch gameState {
	case StateMainMenu:
		handleMainMenuInput()
	case StateNewPlayer:
		handleNewPlayerInput()
	case StateStory:
		handleStoryInput()
	case StateMissionSelect:
		handleMissionSelectInput()
	case StateMission:
		handleMissionInput()
	case StateHackMethod:
		handleHackMethodInput()
	case StateBruteForceMiniGame:
		handleBruteForceMiniGameInput()
	case StateSocialEngineeringMiniGame:
		handleSocialEngineeringMiniGameInput()
	case StatePhishingMiniGame:
		handlePhishingMiniGameInput()
	case StateZeroDayMiniGame:
		handleZeroDayMiniGameInput()
	case StateDDoSMiniGame:
		handleDDoSMiniGameInput()
	case StateSQLInjectionMiniGame:
		handleSQLInjectionMiniGameInput()
	case StateSuccess, StateFailure:
		handleMissionResultInput()
	case StateShop:
		handleShopInput()
	case StateSettings:
		handleSettingsInput()
	case StateHelp, StateCredits, StateAbilities, StateMessaging:
		handleGeneralInput()
	case StateNPCInteraction:
		handleNPCInteractionInput()
	case StateNetwork:
		handleNetworkInput()
	case StateWorldMap:
		handleWorldMapInput()
	case StateMarket:
		handleMarketInput()
	}

	updateButtons()
	updateTransition()
	updateMarket()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	drawTerminalBorder(screen)
	drawStatus(screen)

	switch gameState {
	case StateMainMenu:
		drawMainMenu(screen)
	case StateNewPlayer:
		drawNewPlayer(screen)
	case StateStory:
		drawStory(screen)
	case StateMissionSelect:
		drawMissionSelect(screen)
	case StateMission:
		drawMission(screen)
	case StateHackMethod:
		drawHackMethod(screen)
	case StateBruteForceMiniGame:
		drawBruteForceMiniGame(screen)
	case StateSocialEngineeringMiniGame:
		drawSocialEngineeringMiniGame(screen)
	case StatePhishingMiniGame:
		drawPhishingMiniGame(screen)
	case StateZeroDayMiniGame:
		drawZeroDayMiniGame(screen)
	case StateDDoSMiniGame:
		drawDDoSMiniGame(screen)
	case StateSQLInjectionMiniGame:
		drawSQLInjectionMiniGame(screen)
	case StateSuccess, StateFailure:
		drawMissionResult(screen)
	case StateShop:
		drawShop(screen)
	case StateSettings:
		drawSettings(screen)
	case StateHelp:
		drawHelp(screen)
	case StateCredits:
		drawCredits(screen)
	case StateAbilities:
		drawAbilities(screen)
	case StateMessaging:
		drawMessaging(screen)
	case StateNPCInteraction:
		drawNPCInteraction(screen)
	case StateNetwork:
		drawNetwork(screen)
	case StateWorldMap:
		drawWorldMap(screen)
	case StateMarket:
		drawMarket(screen)
	}

	drawButtons(screen)
	drawTooltip(screen)
	drawTransitionEffect(screen) // Draw transition effect

	drawMouseCursor(screen)
}

func drawMouseCursor(screen *ebiten.Image) {
	cursor := ebiten.NewImage(10, 10)
	cursor.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(mouseX), float64(mouseY))
	screen.DrawImage(cursor, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func handleMainMenuInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		transitionToState(StateStory)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		transitionToState(StateSettings)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		transitionToState(StateHelp)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		transitionToState(StateCredits)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		transitionToState(StateAbilities)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		transitionToState(StateMessaging)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) { // Use 'P' key for shop
		transitionToState(StateShop)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		interactWithNPC(0) // Interact with first NPC for demo
	} else if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		transitionToState(StateNetwork)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		transitionToState(StateWorldMap)
	}
}

func handleNewPlayerInput() {
	for _, r := range ebiten.InputChars() {
		updatePlayerName(r)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && playerName != "" {
		player.Name = playerName
		player.CurrentLevel = 1
		player.Score = 0
		player.Money = 100.0
		player.Skills = make(map[string]bool)
		player.Skills["Brute Force Attack"] = true
		player.Skills["Social Engineering"] = false
		player.Skills["Phishing Attack"] = false
		player.Skills["Zero-Day Exploit"] = false
		player.Reputation = 100
		player.Level = 1
		savePlayerData()
		transitionToState(StateMainMenu)
	}
}

func handleStoryInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		transitionToState(StateMissionSelect)
	}
}

func handleMissionSelectInput() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		startMission(0)
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		startMission(1)
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		startMission(2)
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		startMission(3)
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		startMission(4)
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
		startMission(5)
	} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
		startMission(6)
	} else if inpututil.IsKeyJustPressed(ebiten.Key8) {
		startMission(7)
	}
}

func handleMissionInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		transitionToState(StateHackMethod)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMissionSelect)
	}
}


func handleHackMethodInput() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) && player.Skills["Brute Force Attack"] {
		startBruteForceHack()
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) && player.Skills["Social Engineering"] {
		startSocialEngineeringHack()
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) && player.Skills["Phishing Attack"] {
		startPhishingHack()
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) && player.Skills["Zero-Day Exploit"] {
		startZeroDayExploitHack()
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) && player.Skills["DDoS Attack"] {
		startDDoSHack()
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) && player.Skills["SQL Injection"] {
		startSQLInjectionHack()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMission)
	}
}

func startBruteForceHack() {
	initBruteForceMiniGame()
	transitionToState(StateBruteForceMiniGame)
}

func startSocialEngineeringHack() {
	initSocialEngineeringMiniGame()
	transitionToState(StateSocialEngineeringMiniGame)
}

func startPhishingHack() {
	initPhishingMiniGame()
	transitionToState(StatePhishingMiniGame)
}

func startZeroDayExploitHack() {
	initZeroDayMiniGame()
	transitionToState(StateZeroDayMiniGame)
}

func startDDoSHack() {
	initDDoSMiniGame()
	transitionToState(StateDDoSMiniGame)
}

func startSQLInjectionHack() {
	initSQLInjectionMiniGame()
	transitionToState(StateSQLInjectionMiniGame)
}

var randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))

func executeHack(successRate float64, penalty int) {
	if randomgen.Float64() < successRate {
		completeMission(true)
	} else {
		completeMission(false)
		player.Reputation += penalty
		if player.Reputation < 0 {
			player.Reputation = 0
		}
	}
}

func handleMissionResultInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		transitionToState(StateMissionSelect)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
	}
}


func handleGeneralInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
	}
}

func handleTooltipUpdate() {
	cx, cy := ebiten.CursorPosition()
	if cx > 340 && cx < 460 && cy > 480 && cy < 520 {
		tooltipText = "Open the shop to buy tools and upgrades."
		tooltipVisible = true
	} else {
		tooltipVisible = false
	}
}

func drawTooltip(screen *ebiten.Image) {
	if tooltipVisible {
		text.Draw(screen, tooltipText, fontFace, 10, 570, color.White)
	}
}
