package main

import (
	"chess/core"
	"fmt"
	"strconv"
	"strings"
)

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func validatePosition(input string) {
	if input == "end" {
		panic("Game was ended by user")
	}
	if len(input) != 2 {
		panic("Invalid position: len != 2")
	}
	if !strings.Contains(core.ROW, input[:1]) {
		panic("Invalid position: incorrect letter")
	}
	val, err := strconv.Atoi(input[1:])
	if err != nil {
		panic("Invalid position")
	}
	if val < 1 || val > 8 {
		panic("Invalid position")
	}
}

func getPosition(message string) string {
	input := getInput(message)
	validatePosition(input)
	return input
}

func getMove() core.Move {
	var from, to string
	from = getPosition("Enter from: ")
	to = getPosition("Enter to: ")
	return core.Move{From: from, To: to}
}

func main() {
	board := core.CreateChessBoard()

	for {
		board.Show()
		move := getMove()
		board.ProcessMove(move)
	}
}
