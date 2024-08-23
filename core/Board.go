package core

import (
	"errors"
	"fmt"
	"slices"
)

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

var piecesViewMap = map[Piece]string{
	"R": "♜",
	"N": "♞",
	"B": "♝",
	"K": "♚",
	"Q": "♛",
	"P": "♟",
	"r": "♖",
	"n": "♘",
	"b": "♗",
	"k": "♔",
	"q": "♕",
	"p": "♙",
	"-": "-",
}

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

func printField(i, j int, piece Piece) {
	symbol := piecesViewMap[piece]
	if symbol == "-" {
		if (i+j)%2 == 0 {
			fmt.Print("□ ")
		} else {
			fmt.Print("■ ")
		}
	} else {
		fmt.Print(symbol + " ")
	}
}

func (board *ChessBoard) Show() {
	if board.isWhiteMove {
		for i := 7; i >= 0; i-- {
			fmt.Print(i + 1)
			fmt.Print(" ")
			for j := 7; j >= 0; j-- {
				printField(i, j, board.positions[i][j])
			}
			fmt.Println()
		}
		fmt.Println("- a b c d e f g h ")
	} else {
		for i := 0; i < 8; i++ {
			fmt.Print(i + 1)
			fmt.Print(" ")
			for j := 7; j >= 0; j-- {
				printField(i, j, board.positions[i][j])
			}
			fmt.Println()
		}
		fmt.Println("- h g f e d c b a ")
	}
}

func (board *ChessBoard) ProcessMove(move Move) error {
	availableMoves := board.getAvailablePositionForMove(move.From)
	if !slices.Contains(availableMoves, move.To) {
		return errors.New("move not available")
	}
	board.processMove(move)
	return nil
}

func (board *ChessBoard) getAvailablePositionForMove(startPosition Position) []Position {
	result := make([]Position, 3)
	position, _ := NewPosition("e4")
	result = append(result, position)
	return result
}

func (board *ChessBoard) processMove(move Move) {
	board.positions[move.To.i][move.To.j] = board.positions[move.From.i][move.From.j]
	board.positions[move.From.i][move.From.j] = "-"
	board.isWhiteMove = !board.isWhiteMove
	board.moveNumber++
}
