package game

import (
	"fmt"
	"github.com/scrisanti/bridge-simulator/bidding"
	"github.com/scrisanti/bridge-simulator/card"
	"github.com/scrisanti/bridge-simulator/log"
	"github.com/scrisanti/bridge-simulator/player"
)

func Start() {
	log.Logger.Info("Starting new game...")
	players := []player.Player{
		player.NewBasicPlayer("North"),
		player.NewBasicPlayer("East"),
		player.NewBasicPlayer("South"),
		player.NewBasicPlayer("West"),
	}

	deck := card.NewDeck()
	deal(card.Shuffle(deck), players)

	rules := []bidding.BidRule{
		bidding.OneNoTrumpRule{},
		bidding.FiveCardMajorRule{},
		// Add more rules as needed
		bidding.MinorOpening{},
	}

	// Opening bidding
	for _, player := range players {
		log.Logger.Debug(fmt.Sprintf(" ---- Player %s ----", player.GetName()))
		evalResults := bidding.AnalyzeHand(player.GetHand())
		log.Logger.Debug(fmt.Sprintf("%+v", evalResults))
		bid := bidding.ChooseOpeningBid(evalResults, rules)
		log.Logger.Info(fmt.Sprintf("%v Bid: %v", player.GetName(), bid))

		if bid.Level > 0 {
			log.Logger.Debug("Entering bidding response mode...")
			break
		}
	}
}

func deal(deck []card.Card, players []player.Player) {
	log.Logger.Info("Dealing cards...")
	for i, c := range deck {
		players[i%len(players)].ReceiveCard(c)
	}
}
