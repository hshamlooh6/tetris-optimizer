package main

func solve(grid [][]byte, pieces []Piece, index int, size int) bool {
	if index == len(pieces) {
		return true
	}

	letter := byte('A' + index)
	piece := pieces[index]

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if canPlace(grid, piece, row, col, size) {
				place(grid, piece, row, col, letter)

				if solve(grid, pieces, index+1, size) {
					return true
				}

				remove(grid, piece, row, col)
			}
		}
	}

	return false
}

func canPlace(grid [][]byte, piece Piece, row, col, size int) bool {
	for _, cell := range piece.cells {
		r := row + cell[0]
		c := col + cell[1]
		if r < 0 || r >= size || c < 0 || c >= size {
			return false
		}
		if grid[r][c] != '.' {
			return false
		}
	}
	return true
}

func place(grid [][]byte, piece Piece, row, col int, letter byte) {
	for _, cell := range piece.cells {
		grid[row+cell[0]][col+cell[1]] = letter
	}
}

func remove(grid [][]byte, piece Piece, row, col int) {
	for _, cell := range piece.cells {
		grid[row+cell[0]][col+cell[1]] = '.'
	}
}