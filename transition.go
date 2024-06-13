package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"sync"
	"time"
)

var (
	transitionOpacity float64
	transitionMutex   sync.Mutex
	fadingIn          bool
	fadingOut         bool
	targetState       GameState
)

// Update the transition effect
func updateTransition() {
	if fadingOut {
		transitionMutex.Lock()
		transitionOpacity += 0.02
		if transitionOpacity >= 1.0 {
			transitionOpacity = 1.0
			fadingOut = false
			// Switch state after fade out
			gameState = targetState
			// Start fading in
			go endTransitionEffect()
		}
		transitionMutex.Unlock()
	} else if fadingIn {
		transitionMutex.Lock()
		transitionOpacity -= 0.02
		if transitionOpacity <= 0.0 {
			transitionOpacity = 0.0
			fadingIn = false
		}
		transitionMutex.Unlock()
	}
}

func startTransitionEffect() {
	transitionMutex.Lock()
	transitionOpacity = 0.0
	fadingOut = true
	fadingIn = false
	transitionMutex.Unlock()
	for fadingOut {
		time.Sleep(16 * time.Millisecond) // Sleep for ~60 FPS
		updateTransition()
	}
}

func endTransitionEffect() {
	transitionMutex.Lock()
	transitionOpacity = 1.0
	fadingIn = true
	fadingOut = false
	transitionMutex.Unlock()
	for fadingIn {
		time.Sleep(16 * time.Millisecond) // Sleep for ~60 FPS
		updateTransition()
	}
}

func drawTransitionEffect(screen *ebiten.Image) {
	if transitionOpacity > 0.0 {
		transitionMutex.Lock()
		alpha := uint8(transitionOpacity * 255)
		transitionMutex.Unlock()
		screen.Fill(color.NRGBA{0, 0, 0, alpha})
	}
}

func transitionToState(state GameState) {
	transitionMutex.Lock()
	targetState = state
	transitionMutex.Unlock()
	go startTransitionEffect()
}
