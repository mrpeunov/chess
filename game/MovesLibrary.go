package game

type Offset struct {
	i, j int
}

type PieceSetting struct {
	offsets []Offset
	isLong  bool
}

var RookSettings = PieceSetting{
	offsets: []Offset{{0, 1}, {1, 0}, {-1, 0}, {0, -1}},
	isLong:  true,
}

var BishopSettings = PieceSetting{
	offsets: []Offset{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}},
	isLong:  true,
}

var QueenSettings = PieceSetting{
	offsets: []Offset{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}},
	isLong:  true,
}

var KnightSettings = PieceSetting{
	offsets: []Offset{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}},
	isLong:  false,
}

var KingSettings = PieceSetting{
	offsets: []Offset{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}},
	isLong:  false,
}

var PawnSettingsWhite = PieceSetting{
	offsets: []Offset{{0, 1}, {1, 1}, {-1, 1}},
	isLong:  false,
}

var PawnSettingsBlack = PieceSetting{
	offsets: []Offset{{0, 1}, {1, 1}, {-1, 1}},
	isLong:  false,
}

var MovesMap = map[Square]PieceSetting{
	"R": RookSettings,
	"N": KnightSettings,
	"B": BishopSettings,
	"K": KingSettings,
	"Q": QueenSettings,
	"P": PawnSettingsWhite,
	"r": RookSettings,
	"n": KnightSettings,
	"b": BishopSettings,
	"k": KingSettings,
	"q": QueenSettings,
	"p": PawnSettingsBlack,
}

// King — король.
// Queen — королева.
// Bishop — слон.
// Knight — конь.
// Rook — ладья.
// Pawn — пешка.
