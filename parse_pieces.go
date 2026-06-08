package main

import "strings"

type Piece struct {
	cells [][2]int
}

func parsePieces(content string) ([]Piece, bool) {
	blocks := strings.Split(strings.TrimRight(content, "\n"), "\n\n")

	var pieces []Piece

	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		if len(lines) != 4 {
			return nil, false
		}

		var cells [][2]int

		for r, line := range lines {
			if len(line) != 4 {
				return nil, false
			}
			for c, ch := range line {
				if ch == '#' {
					cells = append(cells, [2]int{r, c})
				} else if ch != '.' {
					return nil, false
				}
			}
		}

		if len(cells) != 4 {
			return nil, false
		}

		if !isConnected(cells) {
			return nil, false
		}

		pieces = append(pieces, Piece{cells: cells})
	}

	if len(pieces) == 0 {
		return nil, false
	}

	return pieces, true
}

func isConnected(cells [][2]int) bool {
	set := map[[2]int]bool{}
	for _, c := range cells {
		set[c] = true
	}

	visited := map[[2]int]bool{}
	queue := [][2]int{cells[0]}
	visited[cells[0]] = true

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			next := [2]int{cur[0] + d[0], cur[1] + d[1]}
			if set[next] && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return len(visited) == 4
}