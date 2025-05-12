# Bridge Simulator

This is a bridge simulator written in Go. The user can play as one of the four players, while the other 3 chairs are robots.

## General Information on Bridge

Bridge is a trick-taking card game that is played around the world. It is widely considered the "most fair" card game (espeacially when played in the style of duplicate bridge) because it can be scored in such a way that only the outcomes of bidding and card playing impact the final score that a team gets.

There are 4 playes in a bridge game, split into two pairs of teams. The players and teams are refered to by cardinal directions: North and South Play together against East and West.

# Core Architecture

The game is loosly split into three sections: Dealing, Bidding, and Playing Cards. The underlying logic the robots use to interact with the player is complicated but in general, the robots are aware of counting high card points, identifying long suits, singletons, and overarchingly want to get "Game"

## Dealing

## Bidding

## Playing
