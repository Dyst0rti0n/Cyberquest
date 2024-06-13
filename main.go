package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext    *audio.Context
	clickSound      *audio.Player
	storyline       Storyline
	skillTree       SkillTree
	clickSoundFile  *os.File
	playerMoney 	= 1000.0
	mouseX          int
	mouseY          int
)

func main() {
	loadSettings()
	gameState = StateMainMenu
	storyline.Initialize()
	skillTree.Initialize()
	initializeMissions()
	initializeMarket()
	initializeNPCs()
	initializeNetwork()
	initializeWorldLocations()
	initializeEvents()

	if !loadPlayerData() {
		gameState = StateNewPlayer
	}
	loadMessages()
	addNPCMessages()
	if err := loadFont(); err != nil {
		log.Fatal(err)
	}
	if err := loadSounds(); err != nil {
		log.Fatal(err)
	}

	go marketUpdateLoop() // Start market update in a separate goroutine

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("CyberQuest: The Hacker's Journey")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println("Error running game:", err)
	}

	if clickSoundFile != nil {
		clickSoundFile.Close()
	}
}

func marketUpdateLoop() {
	for {
		updateMarket()
		displayMarket()
		time.Sleep(10 * time.Second)
	}
}

func loadSounds() error {
	var err error
	audioContext = audio.NewContext(44100)
	clickSound, clickSoundFile, err = loadSound("click.wav")
	if err != nil {
		return err
	}
	return nil
}

func loadSound(filename string) (*audio.Player, *os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	d, err := wav.DecodeWithSampleRate(44100, f)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	player, err := audio.NewPlayer(audioContext, d)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	return player, f, nil
}

func playClickSound() {
	if clickSound != nil {
		clickSound.Rewind()
		clickSound.Play()
	}
}
