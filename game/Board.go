package game

type Piece string

// ChessBoard Нотация Форсайта — Эдвардса (https://ru.wikipedia.org/wiki/Нотация_Форсайта_—_Эдвардса)
type ChessBoard struct {
	fields          *[8][8]Piece
	isWhiteMove     bool
	castling        [4]string
	enPassant       string
	halfMoveCounter int
	moveNumber      int
}

func CreateChessBoard() ChessBoard {
	fields := [8][8]Piece{
		{"R", "N", "B", "Q", "K", "B", "N", "R"},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"r", "n", "b", "q", "k", "b", "n", "r"},
	}
	return ChessBoard{
		fields:          &fields,
		isWhiteMove:     true,
		castling:        [4]string{"r", "n", "b", "q"},
		enPassant:       "r",
		halfMoveCounter: 0,
		moveNumber:      1,
	}
}

// Get получает фигуру по правильным индексам, а не по хранению
func (board *ChessBoard) get(i, j int) Piece {
	return board.fields[i][j]
}

// Set устанавливает фигуру по правильным индексам, а не по хранению
func (board *ChessBoard) set(i, j int, p Piece) {
	board.fields[i][j] = p
	// todo: подумать над выходом за пределы массива
}

// todo: нужно решить проблему невнятного индексирования
// todo: возможно стоит использовать x, y нотацию и передавать позицию, а не индексы
// todo: переименовать Core в Game
// todo: вынести консольное отображение в отдельный слой
