package core

import (
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

func parseIndexes(position string) (int, int) {
	i := strings.Index(ROW, position[0:1])
	j, err := strconv.Atoi(position[1:2])
	if err != nil {
		panic(err)
	}
	return j - 1, i
}
