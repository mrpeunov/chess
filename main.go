package main

import (
	"chess/game"
	"errors"
	"fmt"
)

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func getPosition(message string) (game.Position, error) {
	for {
		input := getInput(message)
		if input == "end" {
			return game.Position{}, errors.New("game was ended by user")
		}

		position, err := game.NewPosition(input)
		if err != nil {
			fmt.Println("Incorrect:", err)
			fmt.Println("Try again.")
			continue
		}
		return position, nil
	}
}

func getMove() (game.Move, error) {
	move := game.Move{}
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
	board := game.CreateChessBoard()

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
