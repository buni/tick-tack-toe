package game

import (
	"errors"
)

var ErrPlayersListIsEmpty = errors.New("players list is empty")

type PlayersList struct {
	data  []Player
	index int
}

func NewPlayersList() *PlayersList {
	return &PlayersList{data: make([]Player, 0, 2)}
}

// Add - add player to the list
func (p *PlayersList) Add(player Player) {
	p.data = append(p.data, player)
}

//Next - get the next player in line
//functions like an iterator that starts from the beging when getting to the end 
func (p *PlayersList) Next() (Player, error) {
	if len(p.data) == 0 {
		return nil, ErrPlayersListIsEmpty
	}

	if p.index < len(p.data) {
		player := p.data[p.index]
		p.index++
		return player, nil
	}
	p.index = 1
	return p.data[0], nil
}

func (p *PlayersList) Reset() {
	p.data = make([]Player, 0, 2)
	p.index = 0
}
