package game

import "fmt"

func printField(i, j int, piece Piece) {
	symbol := piecesViewMap[piece]
	if symbol == "-" {
		if (i+j)%2 == 0 {
			fmt.Print("□ ")
		} else {
			fmt.Print("■ ")
		}
	} else {
		fmt.Print(symbol + " ")
	}
}

func (board *ChessBoard) Show() {
	if board.isWhiteMove {
		for i := 7; i >= 0; i-- {
			fmt.Print(i + 1)
			fmt.Print(" ")
			for j := 7; j >= 0; j-- {
				printField(i, j, board.get(i, j))
			}
			fmt.Println()
		}
		fmt.Println("- a b c d e f g h ")
	} else {
		for i := 0; i < 8; i++ {
			fmt.Print(i + 1)
			fmt.Print(" ")
			for j := 7; j >= 0; j-- {
				printField(i, j, board.get(i, j))
			}
			fmt.Println()
		}
		fmt.Println("- h g f e d c b a ")
	}
}

var piecesViewMap = map[Piece]string{
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
