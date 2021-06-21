package game

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ai_PlayMove(t *testing.T) {
	board := NewBoard(3)
	board.state = [][]State{
		{PlayerO, PlayerX, Empty},
		{PlayerX, PlayerO, PlayerO},
		{PlayerX, PlayerO, Empty},
	}
	board.turn = 7
	ai := NewAiPlayer(PlayerX, board).(*ai)
	ai.PlayMove()
	assert.Equal(t, PlayerX, board.state[2][2]) // block

	board.Reset()

	ai.PlayMove()
	assert.Equal(t, PlayerX, board.state[0][0]) // first move always in the corner

	board.Reset()
	board.state = [][]State{
		{PlayerX, PlayerO, Empty},
		{PlayerX, Empty, Empty},
		{Empty, Empty, PlayerO},
	}
	board.turn = 4

	ai.PlayMove()
	assert.Equal(t, PlayerX, board.state[1][1]) // fork
}

func Test_player_PlayMove(t *testing.T) {
	writer := bytes.NewBuffer([]byte("1z2"))

	board := NewBoard(3)
	board.Move(0, 0, PlayerX)
	player := NewPlayer(PlayerX, board, writer).(*player)

	writer.WriteString("\nzx2")
	writer.WriteString("\n1xz")
	writer.WriteString("\n1x4")
	writer.WriteString("\n1x1")
	writer.WriteString("\n1x2")
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w
	player.PlayMove()
	require.Equal(t, PlayerX, board.state[0][1])

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout
	log.Println(string(out))
	assert.Contains(t, string(out), "move format")
	assert.Contains(t, string(out), "row")
	assert.Contains(t, string(out), "column")
	assert.Contains(t, string(out), "boundaries")
	assert.Contains(t, string(out), "empty")
}

// test ai vs ai, this should always produce a tie if the negamax algo is properly implemented
func TestAIvsAI(t *testing.T) {
	board := NewBoard(3)
	// playerList := NewPlayersList()
	ai1 := NewAiPlayer(PlayerX, board)
	ai2 := NewAiPlayer(PlayerO, board)

	for board.Score() == 0 {
		ai1.PlayMove()
		ai2.PlayMove()
	}
	require.Equal(t, Tie, board.Score())
	board.Reset()

	for board.Score() == 0 {
		ai2.PlayMove()
		ai1.PlayMove()
	}
	require.Equal(t, Tie, board.Score())
}
