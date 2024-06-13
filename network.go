package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Node struct {
	X, Y        int
	Name        string
	IsTarget    bool
	IsInfected  bool
	Difficulty  int
	Connections []int
}

var networkNodes []Node

func initializeNetwork() {
	networkNodes = []Node{
		{100, 100, "Router", false, false, 1, []int{1, 2}},
		{200, 200, "Workstation A", true, false, 2, []int{0, 3}},
		{300, 100, "Workstation B", true, false, 2, []int{0, 3, 4}},
		{400, 200, "Database Server", false, false, 3, []int{1, 2}},
		{500, 100, "Email Server", true, false, 4, []int{2}},
	}
}

func drawNetwork(screen *ebiten.Image) {
	for _, node := range networkNodes {
		col := color.RGBA{0, 255, 0, 255} // Green for normal nodes
		if node.IsTarget {
			col = color.RGBA{255, 0, 0, 255} // Red for target nodes
		}
		if node.IsInfected {
			col = color.RGBA{255, 255, 0, 255} // Yellow for infected nodes
		}
		ebitenutil.DrawRect(screen, float64(node.X), float64(node.Y), 20, 20, col)
		ebitenutil.DebugPrintAt(screen, node.Name, node.X, node.Y+25)

		// Draw connections
		for _, conn := range node.Connections {
			ebitenutil.DrawLine(screen, float64(node.X+10), float64(node.Y+10), float64(networkNodes[conn].X+10), float64(networkNodes[conn].Y+10), col)
		}
	}
}

func handleNetworkInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
		return
	}

	// Check for interactions with nodes
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		for i, node := range networkNodes {
			if cx > node.X && cx < node.X+20 && cy > node.Y && cy < node.Y+20 {
				interactWithNode(i)
				break
			}
		}
	}
}

func interactWithNode(index int) {
	node := &networkNodes[index]
	if node.IsTarget {
		if node.Difficulty <= player.Level {
			node.IsInfected = true
			currentMessage = fmt.Sprintf("Infected %s!", node.Name)
		} else {
			currentMessage = fmt.Sprintf("%s is too difficult to hack. Increase your level.", node.Name)
		}
	} else {
		currentMessage = fmt.Sprintf("%s is not a target.", node.Name)
	}
}
