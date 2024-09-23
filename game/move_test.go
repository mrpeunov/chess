package game

import (
	"sort"
	"testing"
)

func positionsToRaw(positions []Position) []string {
	rawPositions := make([]string, 0)
	for _, position := range positions {
		rawPositions = append(rawPositions, position.raw)
	}
	return rawPositions
}

func compareRawPositions(first, second []string) bool {
	sort.Strings(first)
	sort.Strings(second)
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func TestGetAvailablePositionForMove_StartPawn(t *testing.T) {
	board := CreateChessBoard()

	var tests = []struct {
		rawPosition         string
		expectedRawPosition []string
	}{
		{"e2", []string{"e3", "e4"}},
		{"e7", []string{"e6", "e5"}},
	}

	for _, test := range tests {
		position, _ := NewPosition(test.rawPosition)
		gotPositions, _ := board.GetAvailablePosition(position)
		gotRawPositions := positionsToRaw(gotPositions)

		if !compareRawPositions(gotRawPositions, test.expectedRawPosition) {
			t.Errorf("%q != %q", gotRawPositions, test.expectedRawPosition)
		}
	}
}
