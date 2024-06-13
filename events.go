package main

import (
	"math/rand"
)

type Event struct {
	Description string
	Effect      func()
}

var events []Event

func initializeEvents() {
	events = []Event{
		{"A rival hacker has stolen some of your money!", func() {
			player.Money -= 100
			if player.Money < 0 {
				player.Money = 0
			}
			currentMessage = "A rival hacker has stolen $100 from you!"
		}},
		{"You found a vulnerability in a major system!", func() {
			player.Money += 200
			currentMessage = "You found a vulnerability and earned $200!"
		}},
		{"Your malware was detected and removed.", func() {
			player.Reputation -= 50
			if player.Reputation < 0 {
				player.Reputation = 0
			}
			currentMessage = "Your malware was detected and your reputation decreased by 50."
		}},
	}
}

func triggerRandomEvent() {
	event := events[rand.Intn(len(events))]
	event.Effect()
}
