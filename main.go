package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	content, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces, ok := parsePieces(content)
	if !ok {
		fmt.Println("ERROR")
		return
	}

	size := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))

	for {
		grid := makeGrid(size)
		if solve(grid, pieces, 0, size) {
			printGrid(grid)
			return
		}
		size++
	}
}