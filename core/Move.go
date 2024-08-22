package core

import (
	"errors"
	"strconv"
	"strings"
)

type Move struct {
	From string
	To   string
}

func (move *Move) getFromIndexes() (int, int) {
	return parseIndexes(move.From)
}

func (move *Move) getToIndexes() (int, int) {
	return parseIndexes(move.To)
}

func IsPositionValid(position string) (bool, error) {
	if len(position) != 2 {
		return false, errors.New("invalid position: len != 2")
	}
	if !strings.Contains(ROW, position[:1]) {
		return false, errors.New("invalid position: incorrect letter")
	}
	val, err := strconv.Atoi(position[1:])
	if err != nil || val < 1 || val > 8 {
		return false, errors.New("invalid position: invalid number")
	}
	return true, nil
}

func parseIndexes(position string) (int, int) {
	isValid, err := IsPositionValid(position)
	if err != nil || !isValid {
		panic(err)
	}
	i := strings.Index(ROW, position[0:1])
	j, _ := strconv.Atoi(position[1:2])
	return j - 1, i
}
