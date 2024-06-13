package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type NPC struct {
	Name     string
	Role     string
	Offers   []Offer
	Messages []Message
}

type Offer struct {
	Description string
	Price       float64
}

type Message struct {
	From    string
	Content string
}

var npcs []NPC
var messages []Message
var currentNPC *NPC
var messageScrollY = 0

func loadMessages() {
	messages = []Message{
		{"Unknown", "Hey there, newbie. Welcome to the dark web. We've got jobs for you if you're up for the challenge. Complete them, and you'll earn some serious cash. Fail, and there'll be consequences. Good luck!"},
	}
}

func addNPCMessages() {
	for _, npc := range npcs {
		messages = append(messages, npc.Messages...)
	}
}

func initializeNPCs() {
	npcs = []NPC{
		{
			Name: "Alice",
			Role: "Exploit Dealer",
			Offers: []Offer{
				{"Zero-Day Exploit", 1000.0},
				{"Phishing Kit", 500.0},
			},
			Messages: []Message{
				{"Alice", "Hey, I've got some new exploits for sale. Interested?"},
			},
		},
		{
			Name: "Bob",
			Role: "Job Provider",
			Offers: []Offer{
				{"Corporate Espionage Job", 2000.0},
				{"Network Breach Job", 1500.0},
			},
			Messages: []Message{
				{"Bob", "Need some help with a job? I've got a few offers."},
			},
		},
	}
}

func interactWithNPC(npcIndex int) {
	if npcIndex < 0 || npcIndex >= len(npcs) {
		fmt.Println("Invalid NPC index.")
		return
	}
	npc := npcs[npcIndex]
	currentNPC = &npc
	transitionToState(StateNPCInteraction)
}

func handleNPCInteractionInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		messageScrollY -= scrollSpeed
		if messageScrollY < 0 {
			messageScrollY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		messageScrollY += scrollSpeed
	}

	for i := range currentNPC.Offers {
		if inpututil.IsKeyJustPressed(ebiten.Key(i + 1)) {
			selectedOffer := currentNPC.Offers[i]
			if player.Money >= selectedOffer.Price {
				player.Money -= selectedOffer.Price
				currentMessage = fmt.Sprintf("Purchased %s for $%.2f", selectedOffer.Description, selectedOffer.Price)
				currentNPC.Messages = append(currentNPC.Messages, Message{currentNPC.Name, "Thanks for your business!"})
				messages = append(messages, Message{currentNPC.Name, fmt.Sprintf("You purchased %s for $%.2f", selectedOffer.Description, selectedOffer.Price)})
				savePlayerData()
			} else {
				currentMessage = "Not enough money to purchase this item"
			}
			break
		}
	}
}

func drawNPCInteraction(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x10, 0x10, 0x10, 0xff})

	// Draw NPC Info
	npcInfo := fmt.Sprintf("Interacting with %s (%s)", currentNPC.Name, currentNPC.Role)
	drawCenteredText(screen, npcInfo, 50, fontFaceTitle)

	// Draw Messages in a Forum-like Style
	messageY := 100 - messageScrollY
	for _, msg := range currentNPC.Messages {
		if messageY > 600 {
			break
		}
		if messageY >= 100 {
			drawForumMessage(screen, msg, messageY)
		}
		messageY += 100
	}

	// Draw Offers
	offerY := 300
	for i, offer := range currentNPC.Offers {
		drawOffer(screen, offer, i, offerY)
		offerY += 80
	}

	drawCenteredText(screen, "Press the corresponding number to buy an item.\nPress Escape to return to the main menu.", offerY+40, fontFace)
}

func drawForumMessage(screen *ebiten.Image, message Message, y int) {
	messageRect := image.Rect(80, y, 720, y+80)
	draw.Draw(screen, messageRect, &image.Uniform{color.NRGBA{0x20, 0x20, 0x20, 0xff}}, image.Point{}, draw.Src)
	text.Draw(screen, fmt.Sprintf("%s: %s", message.From, message.Content), fontFace, 100, y+20, color.White)
}

func drawOffer(screen *ebiten.Image, offer Offer, index, y int) {
	offerRect := image.Rect(80, y, 720, y+60)
	draw.Draw(screen, offerRect, &image.Uniform{color.NRGBA{0x20, 0x20, 0x20, 0xff}}, image.Point{}, draw.Src)

	text.Draw(screen, fmt.Sprintf("%d. %s - $%.2f", index+1, offer.Description, offer.Price), fontFace, 100, y+20, color.White)
}


