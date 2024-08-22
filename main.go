package main

import (
	"chess/core"
	"fmt"
)

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

// вынести
func validatePosition(input string) error {
	if input == "end" {
		panic("Game was ended by user")
	}
	isValid, err := core.IsPositionValid(input)
	if isValid != true || err != nil {
		return err
	}
	return nil
}

func getPosition(message string) string {
	for {
		input := getInput(message)
		err := validatePosition(input)
		if err != nil {
			fmt.Println("Incorrect!", err)
			continue
		}
		return input
	}
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
