package main

import (
	"github.com/scrisanti/bridge-simulator/game"
	"github.com/scrisanti/bridge-simulator/log"
)

//  This is the main game entry point

func main() {
	log.InitLogger("bridge.log")
	log.Logger.Info("------- Starting Game -------- ")
	game.Start()
}
