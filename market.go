package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	shopScrollY    = 0
	marketScrollY  = 0
	scrollSpeed    = 20
	currentTab     = "Tools"
	tabs           = []string{"Tools", "Services", "Hardware", "Drugs"}
	tabWidth       = 200
	selectedTabIdx  = 0
	lastMarketUpdate time.Time
)

type MarketItem struct {
	Name        string
	Description string
	Category    string
	Price       float64
	Availability int
}

var marketItems []MarketItem
var categories = []string{"Tools", "Services", "Hardware", "Drugs"}
var tools = []MarketItem{
	{"Brute Force Tool", "Tool for brute force attacks.", "Tools", 100.0, 10},
	{"Phishing Kit", "Kit for crafting phishing emails.", "Tools", 200.0, 10},
	{"Zero-Day Exploit", "Exploit for zero-day vulnerabilities.", "Tools", 500.0, 5},
	{"Firewall Bypass Tool", "Tool for bypassing firewalls.", "Tools", 150.0, 10},
	{"Advanced Hacking Tool", "Advanced tools for sophisticated hacks.", "Tools", 300.0, 8},
}
var services = []MarketItem{
	{"Hire a Hacker", "Hire a professional hacker for your needs.", "Services", 1000.0, 3},
	{"DDoS Attack", "Distributed Denial of Service attack service.", "Services", 800.0, 5},
}
var hardware = []MarketItem{
	{"New PC", "High-performance personal computer.", "Hardware", 1500.0, 4},
	{"GPU", "High-end graphics processing unit.", "Hardware", 700.0, 6},
}
var drugs = []MarketItem{
	{"Adderall", "Increases focus and productivity.", "Drugs", 50.0, 20},
	{"Modafinil", "Enhances wakefulness and cognitive function.", "Drugs", 60.0, 15},
}

func initializeMarket() {
	rand.Seed(time.Now().UnixNano())
	generateMarketListings()
	updateMarket()
}

func generateMarketListings() {
	marketItems = append(marketItems, tools...)
	marketItems = append(marketItems, services...)
	marketItems = append(marketItems, hardware...)
	marketItems = append(marketItems, drugs...)

	// Generate random items for realism
	for i := 0; i < 50; i++ {
		category := categories[rand.Intn(len(categories))]
		name := fmt.Sprintf("Random %s Item %d", category, i+1)
		description := fmt.Sprintf("A randomly generated %s item.", category)
		price := rand.Float64() * 1000
		availability := rand.Intn(20) + 1

		marketItems = append(marketItems, MarketItem{
			Name:        name,
			Description: description,
			Category:    category,
			Price:       price,
			Availability: availability,
		})
	}
}

func updateMarket() {
	if time.Since(lastMarketUpdate) >= 3*time.Minute {
		for i := range marketItems {
			change := (rand.Float64()*0.2 - 0.1) * marketItems[i].Price
			marketItems[i].Price += change
			if marketItems[i].Price < 1 {
				marketItems[i].Price = 1
			}
			// Randomly increase or decrease availability
			marketItems[i].Availability += rand.Intn(3) - 1
			if marketItems[i].Availability < 1 {
				marketItems[i].Availability = 1
			}
		}
		lastMarketUpdate = time.Now()
	}
}

func displayMarket() {
	fmt.Println("Dark Web Marketplace:")
	for _, item := range marketItems {
		fmt.Printf("Name: %s\nDescription: %s\nCategory: %s\nPrice: $%.2f\nAvailability: %d\n\n",
			item.Name, item.Description, item.Category, item.Price, item.Availability)
	}
}

func drawTabs(screen *ebiten.Image) {
	for i, tab := range tabs {
		tabColor := color.NRGBA{0x20, 0x20, 0x20, 0xff}
		if tab == currentTab {
			tabColor = color.NRGBA{0x00, 0xff, 0x00, 0xff}
		}
		tabRect := image.Rect(i*tabWidth, 60, (i+1)*tabWidth, 100)
		draw.Draw(screen, tabRect, &image.Uniform{tabColor}, image.Point{}, draw.Src)
		text.Draw(screen, tab, fontFace, i*tabWidth+10, 90, color.White)
	}
}

func getItemsByCategory(category string) []MarketItem {
	var items []MarketItem
	for _, item := range marketItems {
		if item.Category == category {
			items = append(items, item)
		}
	}
	return items
}

func handleMarketInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		transitionToState(StateMainMenu)
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		marketScrollY -= scrollSpeed
		if marketScrollY < 0 {
			marketScrollY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		marketScrollY += scrollSpeed
	}
	updateButtons()
}

func drawMarket(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x10, 0x10, 0x10, 0xff})

	drawCenteredText(screen, "Dark Web Marketplace", 30, fontFaceTitle)

	y := 100 - marketScrollY
	for i, item := range marketItems {
		if y > 600 {
			break
		}
		if y >= 0 {
			drawMarketItem(screen, item, i, y)
		}
		y += 80
	}

	drawCenteredText(screen, "Press Escape to return to the main menu.", 580, fontFace)
	drawButtons(screen)
}

func drawMarketItem(screen *ebiten.Image, item MarketItem, index, y int) {
	itemRect := image.Rect(80, y, 720, y+60)
	draw.Draw(screen, itemRect, &image.Uniform{color.NRGBA{0x20, 0x20, 0x20, 0xff}}, image.Point{}, draw.Src)

	text.Draw(screen, fmt.Sprintf("%d. %s", index+1, item.Name), fontFace, 100, y+20, color.White)
	text.Draw(screen, fmt.Sprintf("$%.2f", item.Price), fontFace, 500, y+20, color.White)
	text.Draw(screen, item.Description, fontFaceSmall, 100, y+50, color.Gray{0xaa})

	// Add a button to purchase the item
	buttons = append(buttons, Button{
		X: 650, Y: y + 10, W: 60, H: 40,
		Label: "Buy",
		OnClick: func() {
			if playerMoney >= item.Price {
				playerMoney -= item.Price
				currentMessage = fmt.Sprintf("Purchased %s for $%.2f", item.Name, item.Price)
				savePlayerData()
				showPurchasePopup(item.Name, item.Price)
			} else {
				currentMessage = "Not enough money to purchase this item"
			}
		},
	})
}