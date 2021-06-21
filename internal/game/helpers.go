package game

import "errors"

type State int

const (
	PlayerO State = iota - 1
	Empty
	PlayerX
	Tie
)

var ErrOutOfBound = errors.New("row or column out of bound")

func (mark State) String() string {
	switch mark {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	case Empty:
		return "."
	}
	return "."
}

func sumSize(sum, size int) State {
	if sum == int(PlayerO)*size {
		return PlayerO
	}

	if sum == int(PlayerX)*size {
		return PlayerX
	}

	return Empty
}

func inRange(x, lower, upper int) bool {
	if x >= lower && x <= upper {
		return true
	}
	return false
}
