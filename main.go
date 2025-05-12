package main

import (
	"fmt"
)

//  This is the main game entry point

func main() {

	players := []Player{
		&RandomBot{name: "North"},
		&RandomBot{name: "East"},
		&RandomBot{name: "South"},
		&RandomBot{name: "West"},
	}

	Deal(players)

	for _, p := range players {
		Logger.Info(fmt.Sprintf("%s bids: %s", p.Name(), p.Bid()))
	}

	// Play 13 tricks (simplified)
	for i := 0; i < 13; i++ {
		var trick Trick
		for _, p := range players {
			card := p.PlayCard(trick)
			Logger.Info(fmt.Sprintf("%s plays %s", p.Name(), card))
			trick.Cards = append(trick.Cards, card)
		}
		fmt.Println("--- End of Trick ---")
	}

	fmt.Println("Game Over")
}
