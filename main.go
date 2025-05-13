package main

import (
	"github.com/scrisanti/bridge-simulator/game"
	"github.com/scrisanti/bridge-simulator/log"
)

//  This is the main game entry point

func main() {
	log.Init("bridge.log")
	log.Logger.Info("------- Let's Play Bridge! -------- ")
	game.Start()
}
