package main

import "github.com/buni/ttt-task/internal/game"

func main() {
	board := game.NewBoard(3)
	playerList := game.NewPlayersList()
	ttt := game.NewTTTGame(board, playerList)

	ttt.Start()
}
