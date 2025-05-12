package main

import (
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func Deal(players []Player) {
	deck := NewDeck()
	Shuffle(deck)
	for i, card := range deck {
		player := players[i%len(players)]
		player.ReceiveHand(append([]Card{}, card)) // send a copy of each card
	}
}
