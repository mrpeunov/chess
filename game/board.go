package game

type Piece string
type Square string

// ChessBoard Нотация Форсайта — Эдвардса (https://ru.wikipedia.org/wiki/Нотация_Форсайта_—_Эдвардса)
type ChessBoard struct {
	squares         *[8][8]Square
	isWhiteMove     bool
	castling        [4]string
	enPassant       string
	halfMoveCounter int
	moveNumber      int
}

func GetSquares() [8][8]Square {
	return [8][8]Square{
		{"R", "P", "-", "-", "-", "-", "p", "r"},
		{"N", "P", "-", "-", "-", "-", "p", "n"},
		{"B", "P", "-", "-", "-", "-", "p", "b"},
		{"K", "P", "-", "-", "-", "-", "p", "k"},
		{"Q", "P", "-", "-", "-", "-", "p", "q"},
		{"B", "P", "-", "-", "-", "-", "p", "b"},
		{"N", "P", "-", "-", "-", "-", "p", "n"},
		{"R", "P", "-", "-", "-", "-", "p", "r"},
	}
}

func CreateChessBoard() ChessBoard {
	squares := GetSquares()
	return ChessBoard{
		squares:         &squares,
		isWhiteMove:     true,
		castling:        [4]string{"r", "n", "b", "q"},
		enPassant:       "r",
		halfMoveCounter: 0,
		moveNumber:      1,
	}
}

func (board *ChessBoard) Get(i, j int) Square {
	return board.squares[i][j]
}

func (board *ChessBoard) Set(i, j int, p Square) {
	board.squares[i][j] = p
}
