package player

import "github.com/scrisanti/bridge-simulator/card"

type Player interface {
	ReceiveCard(card.Card)
	PlayCard() card.Card
	GetHand() []card.Card
}

type BasicPlayer struct {
	Hand []card.Card
}

func (p *BasicPlayer) ReceiveCard(c card.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *BasicPlayer) PlayCard() card.Card {
	card := p.Hand[0]
	p.Hand = p.Hand[1:]
	return card
}

func (p *BasicPlayer) GetHand() []card.Card {
	return p.Hand
}
