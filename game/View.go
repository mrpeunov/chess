package game

import "fmt"

func getSquareView(i, j int, square Square) string {
	symbol := squaresViewMap[square]
	if symbol == "-" {
		if (i+j)%2 == 0 {
			return "□"
		} else {
			return "■"
		}
	} else {
		return symbol
	}
}

func (board *ChessBoard) Show() {
	// todo: нужно писать в буфер, тестировать это
	// todo: и отдельная функция для вывода в буфер
	if board.isWhiteMove {
		for j := 7; j >= 0; j-- {
			fmt.Print(j + 1)
			fmt.Print(" ")
			for i := 7; i >= 0; i-- {
				fmt.Print(getSquareView(i, j, board.get(i, j)) + " ")
			}
			fmt.Println()
		}
		fmt.Println("- a b c d e f g h ")
	} else {
		for j := 0; j < 8; j++ {
			fmt.Print(j + 1)
			fmt.Print(" ")
			for i := 7; i >= 0; i-- {
				fmt.Print(getSquareView(i, j, board.get(i, j)) + " ")
			}
			fmt.Println()
		}
		fmt.Println("- h g f e d c b a ")
	}
}

var squaresViewMap = map[Square]string{
	"R": "♜",
	"N": "♞",
	"B": "♝",
	"K": "♚",
	"Q": "♛",
	"P": "♟",
	"r": "♖",
	"n": "♘",
	"b": "♗",
	"k": "♔",
	"q": "♕",
	"p": "♙",
	"-": "-",
}
