package main

import (
	"encoding/json"
	"io"
	"os"
)

type Player struct {
	Name        string
	CurrentLevel int
	Score       int
	Money       float64
	Reputation  int
	Skills      map[string]bool
	Level       int 
}

var (
	player     Player
	playerName string
)


func loadPlayerData() bool {
	if _, err := os.Stat("player.json"); err == nil {
		file, err := os.Open("player.json")
		if err == nil {
			defer file.Close()
			data, err := io.ReadAll(file)
			if err == nil {
				json.Unmarshal(data, &player)
				if player.Skills == nil {
					player.Skills = make(map[string]bool)
				}
				return true
			}
		}
	}
	player.Skills = make(map[string]bool)
	player.Skills["Brute Force Attack"] = true
	player.Skills["Social Engineering"] = false
	player.Skills["Phishing Attack"] = false
	player.Skills["Zero-Day Exploit"] = false
	player.Reputation = 100
	return false
}

func savePlayerData() {
	file, err := os.Create("player.json")
	if err == nil {
		defer file.Close()
		data, _ := json.MarshalIndent(player, "", "  ")
		file.Write(data)
	}
}

func updatePlayerName(char rune) {
	if char == '\b' {
		if len(playerName) > 0 {
			playerName = playerName[:len(playerName)-1]
		}
	} else if len(playerName) < 20 { // Limit player name length
		playerName += string(char)
	}
}
