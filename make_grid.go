package main

func makeGrid(size int) [][]byte {
	grid := make([][]byte, size)
	for i := range grid {
		grid[i] = make([]byte, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}