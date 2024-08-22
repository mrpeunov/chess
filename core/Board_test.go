package core

import "testing"

func TestCreateChessBoard(t *testing.T) {
	board := CreateChessBoard()
	if board.isWhiteMove != true {
		t.Errorf("White move should be true")
	}
	if board.moveNumber != 1 {
		t.Errorf("Move number should be 1")
	}
	if board.halfMoveCounter != 0 {
		t.Errorf("Half move counter should be 0")
	}
}
