package c4

// Win Conditions

// isWin returns true if board has four markers of the same color in a row vertically, horizontally or diagonally
func IsWin(board [][]Colour) bool {
	return fourHorizontal(board) || fourVertical(board) || fourDiagonal(board)
}

func fourHorizontal(board [][]Colour) bool {

}

func fourVertical(board [][]Colour) bool {
	transposed := Transpose(board)
	return fourHorizontal(transposed)
}
func fourDiagonal(board [][]Colour) bool {

}

func fourInARow([]Colour) bool {

}
