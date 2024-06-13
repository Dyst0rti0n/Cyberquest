package main

import (
	"fmt"
	"time"
)

type Mission struct {
	ID            int
	Name          string
	Details       string
	RequiredSkill int
	Reward        float64
	Penalty       float64
	IsCompleted   bool
	NextMissions  []int
}

var missions []Mission
var currentMission Mission

func initializeMissions() {
	missions = []Mission{
		{1, "Infiltrate Local Server", "Gain access to the local network of the corrupt corporation.", 0, 500.0, -100.0, false, []int{2}},
		{2, "Retrieve Sensitive Data", "Find and exfiltrate sensitive data from the corporation's database.", 1, 1000.0, -200.0, false, []int{3}},
		{3, "Disable Security Systems", "Hack into and disable the corporation's security systems.", 2, 750.0, -150.0, false, []int{4}},
		{4, "Expose Corruption", "Gather evidence and expose the corruption to the public.", 3, 2000.0, -300.0, false, []int{}},
		{5, "Hacker Recruitment", "Recruit new hackers to join your team.", 2, 500.0, -100.0, false, []int{}},
		{6, "Create Malware", "Develop and deploy malware to gather intel.", 3, 1200.0, -250.0, false, []int{}},
		{7, "Hack Rival Hackers", "Take down rival hacking groups to gain control.", 4, 1500.0, -300.0, false, []int{}},
		{8, "Defend Against Attacks", "Defend your systems against incoming attacks.", 4, 800.0, -200.0, false, []int{}},
	}
}

func startMission(index int) {
	currentMission = missions[index]
	transitionToState(StateMission)
}

func completeMission(success bool) {
	if success {
		currentMission.IsCompleted = true
		player.Score += 10
		player.Money += currentMission.Reward
		levelUp()

		currentMessage = "Mission Accomplished!"
		showMissionReport()

		if len(currentMission.NextMissions) > 0 {
			nextMissionIndex := currentMission.NextMissions[0]
			time.AfterFunc(time.Second*3, func() {
				startMission(nextMissionIndex)
			})
		} else {
			transitionToState(StateSuccess)
		}

		savePlayerData()
	} else {
		player.Money += currentMission.Penalty
		if player.Money < 0 {
			player.Money = 0
		}
		currentMessage = "Mission Failed!"
		transitionToState(StateFailure)
		savePlayerData()
	}
}

func showMissionReport() {
	// Generate a detailed report of the mission
	report := fmt.Sprintf("Mission Report: %s\n\nDetails: %s\n\nReward: $%.2f\n", currentMission.Name, currentMission.Details, currentMission.Reward)
	if currentMission.IsCompleted {
		report += "Status: Successful\n"
	} else {
		report += "Status: Failed\n"
	}

	currentMessage = report
}