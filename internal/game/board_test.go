package game

import (
	"math/rand"
	"reflect"
	"testing"
)

func fullBoardHelper(t *testing.T, rows int) *Board {
	t.Helper()
	b := NewBoard(3)
	for i := 0; i < rows; i++ {
		for j := 0; j < b.size; j++ {
			if j+i%2 == 1 {
				b.state[i][j] = PlayerX
			} else {
				b.state[i][j] = PlayerO
			}
			b.turn++
		}
	}
	b.Print()
	return b
}

func tieBoard(t *testing.T) *Board {
	t.Helper()
	b := NewBoard(3)
	b.state = [][]State{
		{PlayerO, PlayerX, PlayerO},
		{PlayerX, PlayerO, PlayerX},
		{PlayerX, PlayerO, PlayerX},
	}
	b.turn = 9
	b.Print()
	return b
}

func horizontalWinBoard(t *testing.T, winner, looser State) *Board {
	t.Helper()
	b := NewBoard(3)
	b.state = [][]State{
		{looser, winner, looser},
		{winner, winner, winner},
		{looser, winner, looser},
	}
	b.turn = 9
	b.Print()
	return b
}

func verticalWinBoard(t *testing.T, winner, looser State) *Board {
	t.Helper()
	b := NewBoard(3)
	b.state = [][]State{
		{winner, winner, looser},
		{looser, winner, winner},
		{winner, winner, looser},
	}
	b.turn = 9
	b.Print()
	return b
}

func diagonalWinBoard(t *testing.T, winner, looser State, upper bool) *Board {
	t.Helper()
	b := NewBoard(3)
	if upper {
		b.state = [][]State{
			{winner, winner, looser},
			{looser, winner, winner},
			{looser, looser, winner},
		}
	} else {
		b.state = [][]State{
			{looser, winner, winner},
			{looser, winner, winner},
			{winner, looser, looser},
		}
	}

	b.turn = 9

	b.Print()
	return b
}

func TestBoardIsOutOfBounds(t *testing.T) {
	tests := []struct {
		name   string
		b      *Board
		row    int
		column int
		want   bool
	}{
		{name: "positive out of bounds", b: NewBoard(3), row: 4, column: 4, want: true},
		{name: "negative out of bounds", b: NewBoard(3), row: -1, column: -1, want: true},
		{name: "not out of bounds", b: NewBoard(3), row: 0, column: 0, want: false},
		{name: "not out of bounds", b: NewBoard(3), row: 2, column: 2, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsOutOfBounds(tt.row, tt.column); got != tt.want {
				t.Errorf("Board.IsOutOfBounds() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    *Board
		row  int
		col  int
		want bool
	}{
		{name: "empty board", b: NewBoard(3), row: 0, col: 0, want: true},
		{name: "empty board random", b: NewBoard(3), row: rand.Intn(3), col: rand.Intn(3), want: true},
		{name: "full board", b: fullBoardHelper(t, 3), row: 1, col: 1, want: false},
		{name: "out of board bounds ", b: fullBoardHelper(t, 3), row: 5, col: 5, want: false},
		{name: "not full board ", b: fullBoardHelper(t, 2), row: 2, col: 2, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsEmpty(tt.row, tt.col); got != tt.want {
				t.Errorf("Board.IsEmpty() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardIsFull(t *testing.T) {
	tests := []struct {
		name string
		b    *Board
		want bool
	}{
		{name: "empty board", b: NewBoard(3), want: false},
		{name: "full board", b: fullBoardHelper(t, 3), want: true},
		{name: "semi full board", b: fullBoardHelper(t, 2), want: false},
		{name: "tied board", b: tieBoard(t), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsFull(); got != tt.want {
				t.Errorf("Board.IsFull() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardEmptySpots(t *testing.T) {
	tests := []struct {
		name          string
		b             *Board
		wantPositions [][]int
	}{
		{name: "empty board", b: NewBoard(3), wantPositions: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{name: "full board", b: fullBoardHelper(t, 3), wantPositions: nil},
		{name: "semi full board", b: fullBoardHelper(t, 1), wantPositions: [][]int{{1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPositions := tt.b.EmptySpots(); !reflect.DeepEqual(gotPositions, tt.wantPositions) {
				t.Errorf("Board.EmptySpots() = got %v, want %v", gotPositions, tt.wantPositions)
			}
		})
	}
}

func TestBoardMove(t *testing.T) {
	tests := []struct {
		name    string
		b       *Board
		row     int
		col     int
		symbol  State
		wantErr bool
	}{
		{name: "move on empty coords", b: NewBoard(3), row: 2, col: 2, symbol: PlayerX},
		{name: "move on occupied coords X", b: fullBoardHelper(t, 3), row: 2, col: 2, symbol: PlayerX},
		{name: "move on occupied coords O", b: fullBoardHelper(t, 3), row: 2, col: 2, symbol: PlayerO},
		{name: "move out of board bounds", b: fullBoardHelper(t, 3), row: 4, col: 4, symbol: PlayerX, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Move(tt.row, tt.col, tt.symbol); (err != nil) != tt.wantErr {
				t.Fatalf("Board.Move() = got %v, want %v", err, nil)
			}
			if !tt.wantErr {
				if (tt.b.state[tt.row][tt.col] != tt.symbol) != tt.wantErr {
					t.Errorf("Board.Move() = got %v, want symbol on cords %v", tt.b.state[tt.row][tt.col], tt.symbol)
				}
			}
		})
	}
}

func TestBoardUndo(t *testing.T) {
	tests := []struct {
		name    string
		b       *Board
		row     int
		col     int
		wantErr bool
	}{
		{name: "undo on empty coords", b: NewBoard(3), row: 2, col: 2},
		{name: "undo on occupied coords", b: fullBoardHelper(t, 3), row: 2, col: 2},
		{name: "undo out of board bounds", b: fullBoardHelper(t, 3), row: 4, col: 4, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Undo(tt.row, tt.col); (err != nil) != tt.wantErr {
				t.Fatalf("Board.Move() = got %v, want %v", err, nil)
			}
			if !tt.wantErr {
				if (tt.b.state[tt.row][tt.col] != Empty) != tt.wantErr {
					t.Errorf("Board.Move() = got %v, want symbol on cords %v", tt.b.state[tt.row][tt.col], Empty)
				}
			}
		})
	}
}

func TestBoardScore(t *testing.T) {
	tests := []struct {
		name      string
		b         *Board
		wantState State
	}{
		{name: "tie score", b: tieBoard(t), wantState: Tie},
		{name: "horizontal win score X", b: horizontalWinBoard(t, PlayerX, PlayerO), wantState: PlayerX},
		{name: "horizontal win score O", b: horizontalWinBoard(t, PlayerO, PlayerX), wantState: PlayerO},
		{name: "vertical win score X", b: verticalWinBoard(t, PlayerX, PlayerO), wantState: PlayerX},
		{name: "vertical win score O", b: verticalWinBoard(t, PlayerO, PlayerX), wantState: PlayerO},
		{name: "diagonal win score X lower", b: diagonalWinBoard(t, PlayerX, PlayerO, false), wantState: PlayerX},
		{name: "diagonal win score O lower", b: diagonalWinBoard(t, PlayerO, PlayerX, false), wantState: PlayerO},
		{name: "diagonal win score X upper", b: diagonalWinBoard(t, PlayerX, PlayerO, true), wantState: PlayerX},
		{name: "diagonal win score O upper", b: diagonalWinBoard(t, PlayerO, PlayerX, true), wantState: PlayerO},
		{name: "empty board score", b: NewBoard(3), wantState: Empty},
		{name: "semi filled board score", b: fullBoardHelper(t, 2), wantState: Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotState := tt.b.Score(); !reflect.DeepEqual(gotState, tt.wantState) {
				t.Error(int(gotState))
				t.Errorf("Board.Score() = got %v, want %v", gotState, tt.wantState)
			}
		})
	}
}

func TestBoardReset(t *testing.T) {
	tests := []struct {
		name      string
		b         *Board
		wantState [][]State
	}{
		{name: "reset board state", b: fullBoardHelper(t, 3), wantState: NewBoard(3).state},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Reset()
			if !reflect.DeepEqual(tt.b.state, tt.wantState) {
				t.Errorf("Board.Reset() = got %v, want %v", tt.b.state, tt.wantState)
			}
			if tt.b.turn != 0 {
				t.Errorf("Board.Reset() = got %v, want %v", tt.b.turn, 0)
			}
		})
	}
}
