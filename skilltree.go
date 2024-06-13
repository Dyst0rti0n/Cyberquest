package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Skill struct {
	Name        string
	Description string
	Level       int
	Cost        int
}

type SkillTree struct {
	Skills map[string]*Skill
	Points int
}

var playerSkillTree SkillTree

func (st *SkillTree) Initialize() {
	st.Skills = map[string]*Skill{
		"Brute Force":         {"Brute Force", "Increases success rate of brute force attacks.", 0, 2},
		"Social Engineering":  {"Social Engineering", "Increases effectiveness of social engineering.", 0, 2},
		"Phishing":            {"Phishing", "Increases success rate of phishing attacks.", 0, 2},
		"Zero-Day Exploit":    {"Zero-Day Exploit", "Increases success rate of zero-day exploits.", 0, 3},
		"Network Analysis":    {"Network Analysis", "Increases information gathering on networks.", 0, 1},
		"Malware Development": {"Malware Development", "Increases effectiveness of malware.", 0, 3},
		"DDoS Attack":         {"DDoS Attack", "Increases effectiveness of DDoS attacks.", 0, 3},
		"SQL Injection":       {"SQL Injection", "Increases success rate of SQL Injection attacks.", 0, 3},
	}
	st.Points = 0
}

func (st *SkillTree) UnlockSkill(skillName string) {
	skill, exists := st.Skills[skillName]
	if !exists {
		currentMessage = "Skill does not exist."
		return
	}
	if skill.Level > 0 {
		currentMessage = "Skill already unlocked."
		return
	}
	if st.Points >= skill.Cost {
		st.Points -= skill.Cost
		skill.Level++
		currentMessage = fmt.Sprintf("Unlocked skill: %s", skillName)
		savePlayerData()
	} else {
		currentMessage = "Not enough skill points."
	}
}

func levelUp() {
	player.Level++
	playerSkillTree.Points++
	currentMessage = fmt.Sprintf("Level Up! You are now level %d", player.Level)
	savePlayerData()
}

func handleSkillUnlock() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) && !player.Skills["Social Engineering"] && player.Money >= 500 {
		player.Money -= 500
		player.Skills["Social Engineering"] = true
		savePlayerData()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) && !player.Skills["Phishing Attack"] && player.Money >= 1000 {
		player.Money -= 1000
		player.Skills["Phishing Attack"] = true
		savePlayerData()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) && !player.Skills["Zero-Day Exploit"] && player.Money >= 2000 {
		player.Money -= 2000
		player.Skills["Zero-Day Exploit"] = true
		savePlayerData()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) && !player.Skills["DDoS Attack"] && player.Money >= 1500 {
		player.Money -= 1500
		player.Skills["DDoS Attack"] = true
		savePlayerData()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key5) && !player.Skills["SQL Injection"] && player.Money >= 2500 {
		player.Money -= 2500
		player.Skills["SQL Injection"] = true
		savePlayerData()
	}
}
