package game

import (
	"fmt"
)

type Board struct {
	state [][]State
	size  int
	turn  int
}

// NewBoard -
func NewBoard(size int) *Board {
	board := &Board{}
	board.size = size
	board.make()
	return board
}

func (b *Board) make() {
	b.state = make([][]State, b.size)
	for k := range b.state {
		b.state[k] = make([]State, b.size)
	}
}

// IsOutOfBounds - checks whether the provided coordinates are out of the boundaries of the current board
func (b *Board) IsOutOfBounds(row, column int) bool {
	if !inRange(int(row), 0, b.size-1) || !inRange(int(column), 0, b.size-1) {
		return true
	}
	return false
}

// IsEmpty - checks whether the provided coordinates are empty
// if the move is out of the board boundaries false is returned
func (b *Board) IsEmpty(row, col int) bool {
	if b.IsOutOfBounds(row, col) {
		return false
	}
	return b.state[row][col] == Empty
}

// IsFull - checks wether the board is full or not
func (b *Board) IsFull() bool {
	return b.turn == b.size*b.size
}

// EmptySpots - returns an array with coordinates of all empty spots on the board
// returns nil if there are no empty spots left
func (b *Board) EmptySpots() (positions [][]int) {
	for i := 0; i < len(b.state); i++ {
		for j := 0; j < len(b.state[i]); j++ {
			if b.state[i][j] == Empty {
				positions = append(positions, []int{i, j})
			}
		}
	}
	return
}

// Move - put the given symbol on the provided coordinates
// this function can an will override non empty values in the given position
func (b *Board) Move(row, col int, symbol State) (err error) {
	if b.IsOutOfBounds(row, col) {
		return ErrOutOfBound
	}

	b.state[row][col] = symbol
	b.turn++
	return
}

// Undo - sets the state of the given position to Empty
// and decrements the internal turn counter
func (b *Board) Undo(row, col int) (err error) {
	if b.IsOutOfBounds(row, col) {
		return ErrOutOfBound
	}

	b.state[row][col] = Empty
	b.turn--
	return
}

// Print - prints the board state into stdout
func (b *Board) Print() {
	fmt.Print("\n")
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Printf(" %s ", b.state[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Score - returns the terminal state for the board
// excluding a tied result/full board without a winner
// returns zero/Empty state if there is no winner
func (b *Board) Score() State {
	// horizontal
	for i := 0; i < b.size; i++ {
		var horizontalSum int
		for j := 0; j < b.size; j++ {
			horizontalSum += int(b.state[i][j])
		}
		if state := sumSize(horizontalSum, b.size); state != Empty {
			return state
		}
	}

	// vertical
	verticalSum := make([]int, b.size)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			verticalSum[j] += int(b.state[i][j])
		}
	}

	for _, v := range verticalSum {
		if state := sumSize(v, 3); state != Empty {
			return state
		}
	}

	var diagonalUpper int
	var diagonalLower int

	// diagonals
	jLower := b.size - 1
	for i := 0; i < b.size; i++ {
		diagonalUpper += int(b.state[i][i])
		diagonalLower += int(b.state[i][jLower])
		jLower--
	}

	if state := sumSize(diagonalLower, b.size); state != Empty {
		return state
	}

	if state := sumSize(diagonalUpper, b.size); state != Empty {
		return state
	}

	if b.IsFull() {
		return Tie
	}
	return 0
}

// Reset - resets the board
func (b *Board) Reset() {
	b.make()
	b.turn = 0
}
