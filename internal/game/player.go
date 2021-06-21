package game

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Player interface {
	PlayMove() error
}

type ai struct {
	board  *Board
	symbol State
}

func NewAiPlayer(symbol State, board *Board) Player {
	return &ai{symbol: symbol, board: board}
}

func (p *ai) negamax(alpha int, beta int, color int) (score int, x, y int) {
	switch p.board.Score() {
	case p.symbol:
		score = 1
		return color * score, -1, -1
	case -p.symbol:
		score = -1
		return color * score, -1, -1
	case Tie:
		return 0, -1, -1
	}

	for _, spot := range p.board.EmptySpots() {
		symbol := p.symbol
		if color == -1 {
			symbol = -p.symbol
		}
		if err := p.board.Move(spot[0], spot[1], symbol); err != nil {
			panic(err)
		}
		score, _, _ := p.negamax(-beta, -alpha, -color)
		score = -score
		if err := p.board.Undo(spot[0], spot[1]); err != nil {
			panic(err)
		}

		if score >= beta {
			return score, spot[0], spot[1]
		}
		if score > alpha {
			alpha = score
			x, y = spot[0], spot[1]
		}
	}
	return alpha, x, y
}

func (p *ai) PlayMove() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic during AI move execution %w", r)
		}
	}()

	_, x, y := p.negamax(-100000, 100000, 1)

	err = p.board.Move(x, y, p.symbol)
	if err != nil {
		return err
	}

	p.board.Print()
	return
}

type player struct {
	inPipe io.Reader
	board  *Board
	symbol State
}

func NewPlayer(symbol State, board *Board, inPipe io.Reader) Player {
	return &player{symbol: symbol, board: board, inPipe: inPipe}
}

func (p *player) PlayMove() error {
	scanner := bufio.NewScanner(p.inPipe)
	fmt.Println("Chose your next move, in format 1x2 (rowXcolumn)")
	for scanner.Scan() {

		move := strings.TrimSpace(strings.ToLower(scanner.Text()))

		cords := strings.Split(move, "x")
		if len(cords) != 2 {
			fmt.Printf("invalid move format: %s \n", move)
			continue
		}

		row, err := strconv.Atoi(cords[0])
		if err != nil {
			fmt.Printf("invalid format for row: %s \n", cords[0])
			continue
		}

		column, err := strconv.Atoi(cords[1])
		if err != nil {
			fmt.Printf("invalid format for column: %s \n", cords[1])
			continue
		}

		if p.board.IsOutOfBounds(row-1, column-1) {
			fmt.Printf("position is out of the board boundaries %v \n", move)
			continue
		}

		if !p.board.IsEmpty(row-1, column-1) {
			fmt.Printf("position is not empty %v \n", move)
			continue
		}

		if err = p.board.Move(row-1, column-1, p.symbol); err != nil {
			fmt.Printf("player move err: %s \n", err)
			continue
		}

		break
	}
	return nil
}
