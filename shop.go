package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var shopItems = []struct {
	Name        string
	Description string
	Price       float64
}{
	{"Brute Force Tool", "Tool for brute force attacks.", 100.0},
	{"Phishing Kit", "Kit for crafting phishing emails.", 200.0},
	{"Zero-Day Exploit", "Exploit for zero-day vulnerabilities.", 500.0},
	{"Firewall Bypass Tool", "Tool for bypassing firewalls.", 150.0},
	{"Advanced Hacking Tool", "Advanced tools for sophisticated hacks.", 300.0},
	{"DDoS Attack Kit", "Kit for launching DDoS attacks.", 400.0},
	{"SQL Injection Tool", "Tool for SQL injection attacks.", 350.0},
}

func handleShopInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		shopScrollY -= scrollSpeed
		if shopScrollY < 0 {
			shopScrollY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		shopScrollY += scrollSpeed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		selectedTabIdx = (selectedTabIdx - 1 + len(tabs)) % len(tabs)
		currentTab = tabs[selectedTabIdx]
		shopScrollY = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		selectedTabIdx = (selectedTabIdx + 1) % len(tabs)
		currentTab = tabs[selectedTabIdx]
		shopScrollY = 0
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if y >= 60 && y <= 100 {
			for i := range tabs {
				if x >= i*tabWidth && x <= (i+1)*tabWidth {
					selectedTabIdx = i
					currentTab = tabs[selectedTabIdx]
					shopScrollY = 0
					break
				}
			}
		}
	}

	updateButtons()
}

func drawShop(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x10, 0x10, 0x10, 0xff})

	drawCenteredText(screen, "Shop", 30, fontFaceTitle)

	drawTabs(screen)

	y := 100 - shopScrollY
	for i, item := range getItemsByCategory(currentTab) {
		if y > 600 {
			break
		}
		if y >= 0 {
			drawShopItem(screen, item, i, y)
		}
		y += 80
	}

	drawCenteredText(screen, "Press Escape to return to the main menu.", 580, fontFace)
	drawButtons(screen)
}

func drawShopItem(screen *ebiten.Image, item MarketItem, index, y int) {
	itemRect := image.Rect(80, y, 720, y+60)
	draw.Draw(screen, itemRect, &image.Uniform{color.NRGBA{0x20, 0x20, 0x20, 0xff}}, image.Point{}, draw.Src)

	text.Draw(screen, fmt.Sprintf("%d. %s", index+1, item.Name), fontFace, 100, y+20, color.White)
	text.Draw(screen, fmt.Sprintf("$%.2f", item.Price), fontFace, 500, y+20, color.White)
	text.Draw(screen, item.Description, fontFaceSmall, 100, y+50, color.Gray{0xaa})

	// Add a button to purchase the item
	buttons = append(buttons, Button{
		X: 650, Y: y + 10, W: 60, H: 40,
		Label:  "Buy",
		Color:  color.NRGBA{0x00, 0x80, 0x00, 0xff},
		OnClick: func(item MarketItem) func() {
			return func() {
				if player.Money >= item.Price {
					player.Money -= item.Price
					currentMessage = fmt.Sprintf("Purchased %s for $%.2f", item.Name, item.Price)
					savePlayerData()
					showPurchasePopup(item.Name, item.Price)
				} else {
					currentMessage = "Not enough money to purchase this item"
					buttons[len(buttons)-1].Color = color.NRGBA{0xff, 0x00, 0x00, 0xff} // Highlight in red
				}
			}
			}(item),
		})
}


func showPurchasePopup(itemName string, price float64) {
	popupMessage := fmt.Sprintf("Congratulations!\nYou have purchased %s for $%.2f", itemName, price)
	popupDuration := time.Second * 3 // Show popup for 3 seconds

	go func() {
		time.Sleep(popupDuration)
		currentMessage = ""
	}()
	currentMessage = popupMessage
}

