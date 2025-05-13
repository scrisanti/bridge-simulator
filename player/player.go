package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Trick struct {
	Cards []Card // cards played in this trick
}

type Player interface {
	Bid() string
	PlayCard(trick Trick) Card
	ReceiveHand(hand []Card)
	EvaluateHand(hand []Card) int
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
	Logger.Debug(fmt.Sprintf("This hand has %v cards", len(hand)))
	// Evaluate
}

func (rb *RandomBot) EvaluateHand(hand []Card) int {
	hcp := 0

	highCardPoints := map[string]int{
		"J": 1,
		"Q": 2,
		"K": 3,
		"A": 4,
	}

	suitDistribution := map[string]int{
		"C": 0,
		"D": 0,
		"H": 0,
		"S": 0,
	}

	for _, card := range hand {
		// High Card Points
		val, ok := highCardPoints[string(card.Rank)]
		if ok {
			hcp += val
		}
		// Suit distribution
		suitDistribution[string(card.Suit)] += 1
	}

	var kv_pairs []string
	for k, v := range suitDistribution {
		kv_pairs = append(kv_pairs, fmt.Sprintf("%d%s", v, k))
	}
	Logger.Info(fmt.Sprintf("Player was dealt %v points", hcp))
	Logger.Info(fmt.Sprintf("Suit Distribution: %v", strings.Join(kv_pairs, ", ")))
	return hcp
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
	Logger.Debug(fmt.Sprintf("%v", rb.hand))
	return card
}
