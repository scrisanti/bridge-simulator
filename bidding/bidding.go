// bidding/bidding.go
package bidding

import (
	"github.com/scrisanti/bridge-simulator/card"
	"github.com/scrisanti/bridge-simulator/log"
)

type Bid struct {
	Level int
	Trump card.Suit
}
type HandFeatures struct {
	HCP               int
	SuitLengths       map[card.Suit]int
	IsBalanced        bool
	LongestSuit       card.Suit
	LongestSuitLength int
	NumVoids          int
	NumSingletons     int
	NumDoubletons     int
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

	// Init everythign to zero
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
		HCP:               hcp,
		SuitLengths:       suitLengths,
		IsBalanced:        balanced,
		LongestSuit:       longest,
		LongestSuitLength: maxLen,
		NumVoids:          voids,
		NumSingletons:     singletons,
		NumDoubletons:     doubletons,
	}
}

type BidRule interface {
	Apply(handEval HandFeatures) (Bid, bool)
}

func PassBid() Bid {
	return Bid{Level: 0, Trump: ""}
}

type OneNoTrumpRule struct{}

// 15-17 points and has a balanced hand
func (r OneNoTrumpRule) Apply(handEval HandFeatures) (Bid, bool) {
	if handEval.IsBalanced && handEval.HCP >= 15 && handEval.HCP <= 17 {
		return Bid{Level: 1, Trump: "NoTrump"}, true
	}
	return Bid{}, false
}

type FiveCardMajorRule struct{}

// 13 or more points with a 5 card major
func (r FiveCardMajorRule) Apply(f HandFeatures) (Bid, bool) {
	majors := []card.Suit{card.Spades, card.Hearts}
	for _, s := range majors {
		if f.HCP >= 13 && f.SuitLengths[s] >= 5 {
			return Bid{Level: 1, Trump: s}, true
		}
	}
	return Bid{}, false
}

type MinorOpening struct{}

func (r MinorOpening) Apply(f HandFeatures) (Bid, bool) {
	minors := []card.Suit{card.Clubs, card.Diamonds}
	for _, s := range minors {
		if f.HCP >= 11 && f.SuitLengths[s] >= 3 && f.LongestSuit == s {
			return Bid{Level: 1, Trump: s}, true
		}
	}
	return Bid{}, false
}

// TODO: 2 Clubs rule

func ChooseOpeningBid(features HandFeatures, rules []BidRule) Bid {
	for _, rule := range rules {
		if bid, ok := rule.Apply(features); ok {
			return bid
		}
	}
	return PassBid()
}
