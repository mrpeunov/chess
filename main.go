package main

import (
	"chess/core"
	"errors"
	"fmt"
)

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func getPosition(message string) (core.Position, error) {
	for {
		input := getInput(message)
		if input == "end" {
			return core.Position{}, errors.New("game was ended by user")
		}

		position, err := core.NewPosition(input)
		if err != nil {
			fmt.Println("Incorrect:", err)
			fmt.Println("Try again.")
			continue
		}
		return position, nil
	}
}

func getMove() (core.Move, error) {
	move := core.Move{}
	from, err := getPosition("Enter from: ")
	if err != nil {
		return move, err
	}

	to, err := getPosition("Enter to: ")
	if err != nil {
		return move, err
	}

	move.From = from
	move.To = to
	return move, nil
}

func main() {
	board := core.CreateChessBoard()

	for {
		board.Show()
		move, err := getMove()
		if err != nil {
			fmt.Println(err)
			break
		}
		err = board.ProcessMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
