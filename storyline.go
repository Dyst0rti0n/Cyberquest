package main

import "fmt"

type Storyline struct {
	Missions       []Mission
	CurrentMission int
}

func (s *Storyline) Initialize() {
	s.Missions = []Mission{
		{1, "Infiltrate Local Server", "Gain access to the local network of the corrupt corporation.", 0, 500.0, -100.0, false, []int{2}},
		{2, "Retrieve Sensitive Data", "Find and exfiltrate sensitive data from the corporation's database.", 1, 1000.0, -200.0, false, []int{3}},
		{3, "Disable Security Systems", "Hack into and disable the corporation's security systems.", 2, 750.0, -150.0, false, []int{4}},
		{4, "Expose Corruption", "Gather evidence and expose the corruption to the public.", 3, 2000.0, -300.0, false, []int{}},
		// Additional missions for an open-world feel
		{5, "Hacker Recruitment", "Recruit new hackers to join your team.", 2, 500.0, -100.0, false, []int{}},
		{6, "Create Malware", "Develop and deploy malware to gather intel.", 3, 1200.0, -250.0, false, []int{}},
		{7, "Hack Rival Hackers", "Take down rival hacking groups to gain control.", 4, 1500.0, -300.0, false, []int{}},
		{8, "Defend Against Attacks", "Defend your systems against incoming attacks.", 4, 800.0, -200.0, false, []int{}},
	}
	s.CurrentMission = 0
}

func (s *Storyline) GetCurrentMission() Mission {
	return s.Missions[s.CurrentMission]
}

func (s *Storyline) AdvanceMission() {
	if s.CurrentMission < len(s.Missions)-1 {
		s.Missions[s.CurrentMission].IsCompleted = true
		s.CurrentMission++
	}
}

func (s *Storyline) GetMissionDetails() string {
	mission := s.GetCurrentMission()
	return fmt.Sprintf("Mission: %s\nDetails: %s\nRequired Skill: %d\nReward: $%.2f\nPenalty: $%.2f\n", mission.Name, mission.Details, mission.RequiredSkill, mission.Reward, mission.Penalty)
}
