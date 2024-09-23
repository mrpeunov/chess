package game

import (
	"errors"
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

var BlackPawn Square = "p"
var WhitePawn Square = "P"
var EmptySquare Square = "-"

func (board *ChessBoard) GetAvailablePosition(startPosition Position) ([]Position, error) {
	var availablePositions []Position
	var err error
	square := board.Get(startPosition.i, startPosition.j)

	if square == EmptySquare {
		return make([]Position, 0), errors.New("empty square")
	}

	// todo: проверка цвета фигуры и игрока

	switch square {

	case BlackPawn:
		availablePositions, err = board.getAvailablePositionForBlackPawn(startPosition)

	case WhitePawn:
		availablePositions, err = board.getAvailablePositionForWhitePawn(startPosition)
	}

	if err != nil {
		return make([]Position, 0), err
	}

	return availablePositions, nil
}
