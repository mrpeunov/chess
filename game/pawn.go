package game

import (
	"errors"
)

func (board *ChessBoard) getAvailablePositionForWhitePawn(startPosition Position) ([]Position, error) {
	// проверить, что пешка
	// обработка прямых ходов
	// обработка съедания
	//return make([]Position, 0), nil

	positions := make([]Position, 0)

	square := board.Get(startPosition.i, startPosition.j)

	if square != "P" {
		return positions, errors.New("piece is not pawn")
	}

	position := NewPositionByIndexes(startPosition.i, startPosition.j+1)
	positions = append(positions, position)

	if startPosition.j == 1 {
		positions = append(positions, NewPositionByIndexes(position.i, position.j+1))
	}

	return positions, nil
}

func (board *ChessBoard) getAvailablePositionForBlackPawn(startPosition Position) ([]Position, error) {
	positions := make([]Position, 0)
	square := board.Get(startPosition.i, startPosition.j)
	if square != "p" {
		return positions, errors.New("piece is not pawn")
	}

	position := NewPositionByIndexes(startPosition.i, startPosition.j-1)
	positions = append(positions, position)

	if startPosition.j == 6 {
		positions = append(positions, NewPositionByIndexes(position.i, position.j-1))
	}

	// todo: подумать как объединить с верхней функцией

	return positions, nil
}
