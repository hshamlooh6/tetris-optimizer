package main

import "fmt"

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}