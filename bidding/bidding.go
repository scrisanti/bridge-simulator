// bidding/bidding.go
package bidding

import (
	"github.com/scrisanti/bridge-simulator/card"
	"github.com/scrisanti/bridge-simulator/log"
)

type HandFeatures struct {
	HCP           int
	SuitLengths   map[card.Suit]int
	IsBalanced    bool
	LongestSuit   card.Suit
	NumVoids      int
	NumSingletons int
	NumDoubletons int
}

func AnalyzeHand(hand []card.Card) HandFeatures {
	log.Logger.Debug("Analyzing Hand...")

	suitLengths := map[card.Suit]int{
		card.Spades: 0, card.Hearts: 0, card.Diamonds: 0, card.Clubs: 0,
	}

	hcp := 0
	for _, c := range hand {
		suitLengths[c.Suit]++
		switch c.Value {
		case "A":
			hcp += 4
		case "K":
			hcp += 3
		case "Q":
			hcp += 2
		case "J":
			hcp += 1
		}
	}

	voids, singletons, doubletons := 0, 0, 0
	maxLen := 0
	longest := card.Spades
	for suit, length := range suitLengths {
		switch length {
		case 0:
			voids++
		case 1:
			singletons++
		case 2:
			doubletons++
		}
		if length > maxLen {
			maxLen = length
			longest = suit
		}
	}

	balanced := (voids == 0 && singletons == 0 && doubletons <= 1)
	return HandFeatures{
		HCP:           hcp,
		SuitLengths:   suitLengths,
		IsBalanced:    balanced,
		LongestSuit:   longest,
		NumVoids:      voids,
		NumSingletons: singletons,
		NumDoubletons: doubletons,
	}
}
