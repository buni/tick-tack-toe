package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayersList(t *testing.T) {
	pList := NewPlayersList()
	nextPlayer, err := pList.Next()

	require.Nil(t, nextPlayer)
	require.Error(t, err, ErrPlayersListIsEmpty)

	pList.Add(NewPlayer(0, nil, nil))
	pList.Add(NewPlayer(1, nil, nil))

	nextPlayer, err = pList.Next()
	require.Nil(t, err)
	v, ok := nextPlayer.(*player)
	require.True(t, ok)
	require.Equal(t, v.symbol, State(0))

	nextPlayer, err = pList.Next()
	require.Nil(t, err)
	v, ok = nextPlayer.(*player)
	require.True(t, ok)
	require.Equal(t, v.symbol, State(1))

	nextPlayer, err = pList.Next()
	require.Nil(t, err)
	v, ok = nextPlayer.(*player)
	require.True(t, ok)
	require.Equal(t, v.symbol, State(0))

	pList.Reset()
	require.Empty(t, pList.data)
}
