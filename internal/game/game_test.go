package game

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGamestartingPosInput(t *testing.T) {
	writer := bytes.NewBuffer([]byte("1"))
	g := NewTTTGame(nil, nil)
	g.inPipe = writer
	v := g.startingPosInput()
	require.Equal(t, 1, v)
	writer.WriteString("2")
	v = g.startingPosInput()
	require.Equal(t, 2, v)

	r, w, _ := os.Pipe()
	os.Stdout = w
	writer.WriteString("3")

	go func() {
		g.startingPosInput()
	}()
	time.Sleep(time.Millisecond * 100)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	if !bytes.Contains(out, []byte("invalid postion choice")) {
		t.Fatal("expected an error")
	}

	r, w, _ = os.Pipe()
	os.Stdout = w
	writer.WriteString("Z")

	go func() {
		g.startingPosInput()
	}()
	time.Sleep(time.Millisecond * 100)

	w.Close()
	out, _ = ioutil.ReadAll(r)
	if !bytes.Contains(out, []byte("invalid postion choice")) {
		t.Fatal("expected an error")
	}
}

func TestGamesymbolInput(t *testing.T) {
	writer := bytes.NewBuffer([]byte("X"))
	g := NewTTTGame(nil, nil)
	g.inPipe = writer
	player, ai := g.symbolInput()
	require.Equal(t, PlayerX, player)
	require.Equal(t, PlayerO, ai)

	writer.WriteString("O")
	player, ai = g.symbolInput()
	require.Equal(t, PlayerO, player)
	require.Equal(t, PlayerX, ai)

	r, w, _ := os.Pipe()
	os.Stdout = w
	writer.WriteString("Z")

	go func() {
		g.symbolInput()
	}()
	time.Sleep(time.Millisecond * 100)
	w.Close()
	out, _ := ioutil.ReadAll(r)

	if !bytes.Contains(out, []byte("invalid symbol choice")) {
		t.Fatal("expected an error")
	}
}

func TestGamePlay(t *testing.T) {
	writer := bytes.NewBuffer([]byte("O"))
	pl := NewPlayersList()
	g := NewTTTGame(NewBoard(3), pl)
	g.inPipe = writer
	writer.WriteString("\n1")

	// r, w, _ := os.Pipe()
	// stdout := os.Stdout
	// os.Stdout = w

	writer.WriteString("\nzx2")
	writer.WriteString("\n1xz")
	writer.WriteString("\n1x4")
	writer.WriteString("\n1x1")
	writer.WriteString("\n1x2")

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w
	g.Play()
	// require.Equal(t, PlayerX, board.state[0][1])

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout
	log.Println(string(out))

	// w.Close()
	// out, _ := ioutil.ReadAll(r)
	// os.Stdout = stdout
	// log.Println(string(out))
}
