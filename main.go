package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit string
type Rank string

const (
	Spades   Suit = "S"
	Hearts   Suit = "H"
	Diamonds Suit = "D"
	Clubs    Suit = "C"
)

var Suits = []Suit{Spades, Hearts, Clubs, Diamonds}
var Ranks = []Rank{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	fmt.Sprintf("%s%s", c.Rank, c.Suit)
}
