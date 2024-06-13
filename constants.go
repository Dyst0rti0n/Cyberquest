package main

type GameState int

const (
	StateMainMenu GameState = iota
	StateNewPlayer
	StateStory
	StateMissionSelect
	StateMission
	StateHackMethod
	StateBruteForceMiniGame
	StateSocialEngineeringMiniGame
	StatePhishingMiniGame
	StateZeroDayMiniGame
	StateDDoSMiniGame
	StateSQLInjectionMiniGame
	StateSuccess
	StateFailure
	StateShop
	StateSettings
	StateHelp
	StateCredits
	StateAbilities
	StateMessaging
	StateNPCInteraction
	StateNetwork
	StateWorldMap
	StateMarket
)

var (
	gameState GameState = StateMainMenu
	currentMessage string
)

