package main

import (
	"math/rand"
)

type Trick struct {
	Cards []Card // cards played in this trick
}

type Player interface {
	Bid() string
	PlayCard(trick Trick) Card
	ReceiveHand(hand []Card)
	Name() string
}

// Bot Logic Implementation

type RandomBot struct {
	hand []Card
	name string
}

func (rb *RandomBot) Name() string {
	return rb.name
}

func (rb *RandomBot) ReceiveHand(hand []Card) {
	rb.hand = hand
}

func (rb *RandomBot) Bid() string {
	return "Pass"
}

func (rb *RandomBot) PlayCard(trick Trick) Card {
	if len(rb.hand) == 0 {
		panic("no cards left to play")
	}
	i := rand.Intn(len(rb.hand))
	card := rb.hand[i]
	rb.hand = append(rb.hand[:i], rb.hand[i+1:]...) // remove played card
	return card
}
