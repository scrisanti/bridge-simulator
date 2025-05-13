package player

import "github.com/scrisanti/bridge-simulator/card"

type Player interface {
	ReceiveCard(card.Card)
	PlayCard() card.Card
	GetHand() []card.Card
	GetName() string
}

type BasicPlayer struct {
	Hand []card.Card
	Name string
}

func (p *BasicPlayer) ReceiveCard(c card.Card) {
	p.Hand = append(p.Hand, c)
}

func NewBasicPlayer(name string) *BasicPlayer {
	return &BasicPlayer{Name: name}
}
func (p *BasicPlayer) PlayCard() card.Card {
	card := p.Hand[0]
	p.Hand = p.Hand[1:]
	return card
}

func (p *BasicPlayer) GetHand() []card.Card {
	return p.Hand
}

func (p *BasicPlayer) GetName() string {
	return p.Name
}
