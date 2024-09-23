package game

import (
	"testing"
)

func TestGetSquareView(t *testing.T) {
	var tests = []struct {
		i, j     int
		square   Square
		expected string
	}{
		{0, 0, "-", "□"},
		{0, 0, "R", "♜"},
		{0, 0, "r", "♖"},
		{1, 0, "N", "♞"},
		{1, 0, "-", "■"},
		{3, 0, "Q", "♛"},
		{4, 0, "K", "♚"},
		{4, 0, "-", "□"},
		{3, 7, "q", "♕"},
		{4, 7, "k", "♔"},
		{0, 7, "-", "■"},
	}
	for _, test := range tests {
		got := getSquareView(test.i, test.j, test.square)
		if got != test.expected {
			t.Errorf("getSquareView(%d, %d, %q), result - %q; expected - %q",
				test.i, test.j, test.square, got, test.expected)
		}
	}

}
