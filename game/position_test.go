package game

import "testing"

var tests = []struct {
	rawPosition string
	i, j        int
}{
	{"a1", 0, 0},
	{"e2", 4, 1},
	{"h8", 7, 7},
	{"d5", 3, 4},
}

func TestNewPositionSuccess(t *testing.T) {
	for _, test := range tests {
		gotPosition, _ := NewPosition(test.rawPosition)
		if gotPosition.i != test.i || gotPosition.j != test.j {
			t.Errorf("NewPosition(%s): got (%d, %d), want (%d, %d)", test.rawPosition, gotPosition.i, gotPosition.j, test.i, test.j)
		}
	}
}

func TestNewPositionByIndexes(t *testing.T) {
	for _, test := range tests {
		gotPosition := NewPositionByIndexes(test.i, test.j)
		if gotPosition.raw != test.rawPosition {
			t.Errorf("NewPositionByIndexes(%d, %d): got(%s), want(%s)", test.i, test.j, gotPosition.raw, test.rawPosition)
		}
	}
}

func TestNewPositionFailure(t *testing.T) {
	var tests = []string{
		"aaa", "bb", "33", "c9", "a0",
	}
	for _, test := range tests {
		gotPosition, err := NewPosition(test)
		if err == nil {
			t.Errorf("NewPosition(%q): got (%d, %d), want error", test, gotPosition.i, gotPosition.j)
		}
	}
}
