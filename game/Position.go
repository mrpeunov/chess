package game

import (
	"errors"
	"strconv"
	"strings"
)

type Position struct {
	raw string
	i   int
	j   int
}

func NewPosition(rawPosition string) (Position, error) {
	position := Position{}
	if len(rawPosition) != 2 {
		return position, errors.New("invalid position: len != 2")
	}

	i, err := parseDigit(rawPosition)
	if err != nil {
		return position, err
	}

	j, err := parseLetter(rawPosition)
	if err != nil {
		return position, err
	}

	position.raw = rawPosition
	position.i = i
	position.j = j

	return position, nil
}

func NewPositionByIndexes(i, j int) Position {
	digit := strconv.Itoa(i + 1)
	letter := ROW[j : j+1]
	return Position{letter + digit, i, j}
}

func parseDigit(rawPosition string) (int, error) {
	index, err := strconv.Atoi(rawPosition[1:])
	if err != nil {
		return 0, errors.New("invalid position: can't parse digit: " + rawPosition)
	}
	if index < 1 || index > 8 {
		return 0, errors.New("invalid position: number < 1 or number > 8: " + rawPosition)
	}
	return index - 1, nil
}

func parseLetter(rawPosition string) (int, error) {
	letter := rawPosition[:1]
	index := strings.Index(ROW, letter)
	if index == -1 {
		return 0, errors.New("invalid position: incorrect letter: " + rawPosition)
	}
	return index, nil
}
