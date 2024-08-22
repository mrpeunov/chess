package main

import (
	"fmt"
	"strconv"
	"strings"
)

const row = "abcdefgh"

type Piece string

// ChessBoard Нотация Форсайта — Эдвардса (https://ru.wikipedia.org/wiki/Нотация_Форсайта_—_Эдвардса)
type ChessBoard struct {
	positions       *[8][8]Piece
	isWhiteMove     bool
	castling        [4]string
	enPassant       string
	halfMoveCounter int
	moveNumber      int
}

func createChessBoard() ChessBoard {
	positions := [8][8]Piece{
		{"R", "N", "B", "Q", "K", "B", "N", "R"},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"r", "n", "b", "q", "k", "b", "n", "r"},
	}
	return ChessBoard{
		positions:       &positions,
		isWhiteMove:     true,
		castling:        [4]string{"r", "n", "b", "q"},
		enPassant:       "r",
		halfMoveCounter: 0,
		moveNumber:      1,
	}
}

func (board *ChessBoard) show() {
	if board.isWhiteMove {
		for i := 7; i >= 0; i-- {
			fmt.Print(i + 1)
			for j := 7; j >= 0; j-- {
				fmt.Print(board.positions[i][j])
			}
			fmt.Println()
		}
		fmt.Println("-abcdefgh")
	} else {
		for i := 0; i < 8; i++ {
			fmt.Print(i + 1)
			for j := 7; j >= 0; j-- {
				fmt.Print(board.positions[i][j])
			}
			fmt.Println()
		}
		fmt.Println("-hgfedcba")
	}
}

func (board *ChessBoard) move(startPosition string, finishPosition string) {
	startI, startJ := parsePosition(startPosition)
	finishI, finishJ := parsePosition(finishPosition)
	board.positions[finishI][finishJ] = board.positions[startI][startJ]
	board.positions[startI][startJ] = "-"
	board.isWhiteMove = !board.isWhiteMove
	board.moveNumber++
}

func parsePosition(position string) (int, int) {
	i := strings.Index(row, position[0:1])
	j, err := strconv.Atoi(position[1:2])
	if err != nil {
		panic(err)
	}
	return j - 1, i
}

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func main() {
	var from, to string

	board := createChessBoard()

	for {
		board.show()

		from = getInput("Enter from: ")
		to = getInput("Enter to: ")

		if from == "end" || to == "end" {
			break
		}

		board.move(from, to)
	}
}
