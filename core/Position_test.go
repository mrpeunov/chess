package core

import "testing"

func TestNewPositionSuccess(t *testing.T) {
	var tests = []struct {
		rawPosition string
		i, j        int
	}{
		{"a1", 0, 0},
		{"h8", 7, 7},
		{"d5", 4, 3},
	}

	for _, test := range tests {
		gotPosition, _ := NewPosition(test.rawPosition)
		if gotPosition.i != test.i || gotPosition.j != test.j {
			t.Errorf("NewPosition(%q): got (%d, %d), want (%d, %d)", test, gotPosition.i, gotPosition.j, test.i, test.j)
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
