package game

func (board *ChessBoard) getAvailablePositionForPawn(startPosition Position) ([]Position, error) {
	// проверить, что пешка
	// обработка прямых ходов
	// обработка съедания
	//return make([]Position, 0), nil
	e3, _ := NewPosition("e3")
	e4, _ := NewPosition("e4")
	lst := make([]Position, 0)
	lst = append(lst, e3)
	lst = append(lst, e4)
	return lst, nil
}
