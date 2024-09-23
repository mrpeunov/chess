package game

import (
	"errors"
	"fmt"
	"slices"
)

type Move struct {
	from Position
	to   Position
}

func NewMove(from Position, to Position) Move {
	return Move{from, to}
}

func (board *ChessBoard) ProcessMove(move Move) error {
	availablePositions, err := board.GetAvailablePosition(move.from)
	if err != nil {
		return err
	}

	if !slices.Contains(availablePositions, move.to) {
		return errors.New("move not available")
	}

	piece := board.Get(move.from.i, move.from.j)
	board.Set(move.to.i, move.to.j, piece)
	board.Set(move.from.i, move.from.j, "-")

	board.isWhiteMove = !board.isWhiteMove
	board.moveNumber++
	return nil
}

var BlackPawn Piece = "p"
var WhitePawn Piece = "P"
var EmptySquare Square = "-"

func (board *ChessBoard) GetAvailablePosition(startPosition Position) ([]Position, error) {
	var availablePositions []Position
	var err error
	fmt.Println(startPosition)
	square := board.Get(startPosition.i, startPosition.j)

	if square == EmptySquare {
		return make([]Position, 0), errors.New("empty square")
	}

	// todo: проверка цвета фигуры и игрока

	piece := Piece(square)

	switch piece {
	case BlackPawn:
	case WhitePawn:
		availablePositions, err = board.getAvailablePositionForPawn(startPosition)
	}

	if err != nil {
		return make([]Position, 0), err
	}

	return availablePositions, nil
}
