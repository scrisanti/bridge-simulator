package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func NewDeck() []Card {
	var deck []Card
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func Shuffle(deck []Card) {
	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	Logger.Info("The Deck has been shuffled")
}

func Deal(players []Player) {
	deck := NewDeck()
	Logger.Info("Retrieved a new deck...")
	Shuffle(deck)

	hands := make([][]Card, len(players))
	for i, card := range deck {
		hands[i%len(players)] = append(hands[i%len(players)], card)
	}

	for i, player := range players {
		Logger.Info(fmt.Sprintf("Initializing %s", player.Name()))
		player.ReceiveHand(hands[i])
		player.EvaluateHand(hands[i])
	}

	Logger.Info(fmt.Sprintf("Dealt %d cards", len(deck)))

}
