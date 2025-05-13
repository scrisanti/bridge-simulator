package card

type Suit string

const (
	Spades   Suit = "S"
	Hearts   Suit = "H"
	Diamonds Suit = "D"
	Clubs    Suit = "C"
	NoTrump  Suit = "NT"
)

type Card struct {
	Suit  Suit
	Value string
}

func NewDeck() []Card {
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	var deck []Card
	for _, s := range suits {
		for _, v := range values {
			deck = append(deck, Card{Suit: s, Value: v})
		}
	}
	return deck
}
