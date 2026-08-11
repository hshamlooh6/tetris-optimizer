# tetris-optimizer

A Go program that reads a file containing tetrominoes and assembles them into the smallest possible square.

## Usage

```bash
go run . <path_to_file>
```

## Example

```bash
go run . test.txt
```

Output:
```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

## Input File Format

Each tetromino is a 4×4 grid using `#` for filled cells and `.` for empty cells. Pieces are separated by a blank line.

```
...#
...#
...#
...#

....
....
....
####

....
##..
##..
....
```

## Rules

- Each piece must have exactly 4 connected cells
- Only `#` and `.` characters are allowed
- At least one piece must be in the file
- Invalid input prints `ERROR`

## How It Works

1. **Read** the file content
2. **Parse** each 4×4 block into a piece and validate it using BFS
3. **Calculate** the minimum square size needed
4. **Solve** using backtracking — place each piece one by one, undo if stuck
5. **Print** the result with letters A, B, C... and `.` for empty cells

## Algorithms

- **BFS** — validates that all 4 cells of each piece are connected
- **Backtracking** — finds the optimal arrangement of all pieces in the grid

## Project Structure

```
tetris-optimizer/
├── main.go          → ties everything together
├── read_file.go     → reads the input file
├── parse_pieces.go  → parses and validates pieces
├── make_grid.go     → creates the empty grid
├── solve.go         → backtracking algorithm
├── print_grid.go    → prints the final grid
└── go.mod           → Go module file
```

## Allowed Packages

Only standard Go packages — `fmt`, `os`, `math`, `strings`

## Author

hshamloo
