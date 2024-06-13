# CyberQuest: The Hacker's Journey

Welcome to **CyberQuest: The Hacker's Journey** – an immersive and captivating hacking simulation game where you step into the shoes of a skilled hacker navigating the treacherous terrain of the dark web. Take on missions, interact with NPCs, explore the underground market, and test your hacking skills in various mini-games. Are you ready to embark on this thrilling journey?

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Gameplay](#gameplay)
  - [Main Menu](#main-menu)
  - [Shop](#shop)
  - [Messaging](#messaging)
  - [Missions](#missions)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Engaging Missions:** Take on various hacking missions with different levels of difficulty.
- **Dynamic Market:** Interact with the dark web marketplace, where prices and availability of items change every three minutes.
- **NPC Interactions:** Communicate with NPCs in a forum-like messaging system and buy special items or take on jobs.
- **Mini-Games:** Test your hacking skills in various mini-games such as Brute Force, Phishing, and SQL Injection.
- **Immersive UI:** Experience a hacker-style UI that brings the dark web to life.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Ebiten](https://ebiten.org/)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/Dyst0rti0n/Cyberquest.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Cyberquest
   ```
3. Install dependencies:
   ```bash
   go get ./...
   ```
4. Run the game:
   ```bash
   go run main.go
   ```

## Gameplay

### Main Menu

The main menu offers various options to start your journey, adjust settings, get help, view credits, and more. Navigate using the arrow keys or mouse clicks to select your desired option.

### Shop

Visit the dark web marketplace to purchase tools, services, hardware, and drugs. Use the arrow keys or mouse clicks to navigate between tabs and scroll through items. Prices and availability are updated every three minutes, so plan your purchases wisely!

### Messaging

Interact with NPCs in a beautifully designed forum-like messaging system. Read messages, respond to offers, and take on special jobs. Use the arrow keys to scroll through messages and offers.

### Missions

Take on various hacking missions with unique challenges. Each mission tests your skills in different areas such as brute force attacks, phishing, and SQL injections. Complete missions to earn money and reputation, but beware of the consequences of failure!

## Development

### Project Structure

- `main.go`: The main entry point of the game.
- `shop.go`: Handles the shop functionality and UI.
- `messaging.go`: Manages NPC interactions and messaging UI.
- `missions.go`: Contains mission-related logic and mini-games.

### Adding New Features

To add new features or improve existing ones, follow these steps:

1. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
2. Implement your changes.
3. Commit and push your changes:
   ```bash
   git commit -m "Description of your changes"
   git push origin feature-name
   ```
4. Open a pull request on GitHub.

## Contributing

We welcome contributions from the community! To contribute, follow the standard [GitHub flow](https://guides.github.com/introduction/flow/):

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Open a pull request.

Please ensure your code follows our coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to reach out if you have any questions or need further assistance. Enjoy your journey through the dark web with **CyberQuest: The Hacker's Journey**!

---
---