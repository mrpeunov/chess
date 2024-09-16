package game

import (
	"errors"
	"fmt"
	"slices"
)

func (board *ChessBoard) ProcessMove(move Move) error {
	availableMoves := board.getAvailablePositionForMove(move.From)
	if !slices.Contains(availableMoves, move.To) {
		return errors.New("move not available")
	}
	board.processMove(move)
	return nil
}

func (board *ChessBoard) getAvailablePositionForMove(startPosition Position) []Position {
	piece := board.get(startPosition.i, startPosition.j)
	availablePositions := make([]Position, 3)
	settings := MovesMap[piece]
	for _, offset := range settings.offsets {
		newI := startPosition.i + offset.i
		newJ := startPosition.j + offset.j
		if 0 <= newI && newI < 8 && 0 <= newJ && newJ < 8 {
			position := NewPositionByIndexes(newI, newJ)
			availablePositions = append(availablePositions, position)
		}
	}
	fmt.Println(availablePositions)
	return availablePositions
}

func (board *ChessBoard) processMove(move Move) {
	piece := board.get(move.From.i, move.From.j)
	board.set(move.To.i, move.To.j, piece)
	board.set(move.From.i, move.From.j, "-")
	board.isWhiteMove = !board.isWhiteMove
	board.moveNumber++
}
