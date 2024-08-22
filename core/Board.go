package core

import (
	"fmt"
)

const ROW = "abcdefgh"

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

// CreateChessBoard Function for creating starting position of game
func CreateChessBoard() ChessBoard {
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

func (board *ChessBoard) Show() {
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

func (board *ChessBoard) ProcessMove(move Move) {
	startI, startJ := move.getFromIndexes()
	finishI, finishJ := move.getToIndexes()
	board.positions[finishI][finishJ] = board.positions[startI][startJ]
	board.positions[startI][startJ] = "-"
	board.isWhiteMove = !board.isWhiteMove
	board.moveNumber++
}
