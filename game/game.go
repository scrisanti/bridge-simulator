package game

import (
	"github.com/scrisanti/bridge-simulator/card"
	"github.com/scrisanti/bridge-simulator/log"
	"github.com/scrisanti/bridge-simulator/player"
)

func Start() {
	log.Logger.Info("Starting new game...")
	players := []player.Player{
		&player.BasicPlayer{},
		&player.BasicPlayer{},
		&player.BasicPlayer{},
		&player.BasicPlayer{},
	}
	deck := card.NewDeck()
	deal(deck, players)
}

func deal(deck []card.Card, players []player.Player) {
	log.Logger.Info("Dealing cards...")
	for i, c := range deck {
		players[i%len(players)].ReceiveCard(c)
	}
}
