package game

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	board       *Board
	playersList *PlayersList
	inPipe      io.Reader
}

func NewTTTGame(board *Board, playersList *PlayersList) *Game {
	return &Game{board: board, playersList: playersList, inPipe: os.Stdin}
}

func (g *Game) startingPosInput() (playerPos int) {
	fmt.Println("Chose your starting position, 1 or 2:")
	scanner := bufio.NewScanner(g.inPipe)
	for scanner.Scan() {
		textPos := scanner.Text()
		position, err := strconv.Atoi(strings.TrimSpace(textPos))
		if err != nil {
			fmt.Printf("invalid postion choice: %s \n", textPos)
			continue
		}
		switch position {
		case 1:
			return 1
		case 2:
			return 2
		default:
			fmt.Printf("invalid postion choice: %v \n", position)
			continue
		}
	}
	return
}

func (g *Game) symbolInput() (player, ai State) {
	fmt.Println("Chose your symbol, X or O:")
	scanner := bufio.NewScanner(g.inPipe)
	for scanner.Scan() {

		symbol := strings.TrimSpace(strings.ToUpper(scanner.Text()))
		switch symbol {
		case PlayerX.String():
			return PlayerX, PlayerO
		case PlayerO.String():
			return PlayerO, PlayerX
		default:
			fmt.Printf("invalid symbol choice: %s \n", symbol)
			continue
		}
	}
	return
}

// Play - main game loop
func (g *Game) Play() error {
	fmt.Println("Starting a new game")
	player, ai := g.symbolInput()
	playerPos := g.startingPosInput()

	if playerPos == 1 {
		g.playersList.Add(NewPlayer(player, g.board, g.inPipe))
		g.playersList.Add(NewAiPlayer(ai, g.board))
	} else {
		g.playersList.Add(NewAiPlayer(ai, g.board))
		g.playersList.Add(NewPlayer(player, g.board, g.inPipe))
	}

	for g.board.Score() == Empty {
		player, err := g.playersList.Next()
		if err != nil {
			return err
		}
		err = player.PlayMove()
		if err != nil {
			return err
		}
	}

	winner := g.board.Score()
	switch winner {
	case Tie:
		fmt.Println("Game ended in a tie")
		g.board.Print()
	default:
		fmt.Printf("%s won the game \n", winner)
		g.board.Print()
	}
	return nil
}

// Start - play the game loop forever
func (g *Game) Start() {
	for {
		err := g.Play()
		if err != nil {
			log.Printf("game error: %s \n", err)
		}
		g.Reset()
	}
}

func (g *Game) Reset() {
	g.board.Reset()
	g.playersList.Reset()
}
