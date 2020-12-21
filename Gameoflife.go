package main

import (
	"flag"
	"math/rand"
	"time"
)

var grid [20][20]int
var gen = 0

func main() {

	s := flag.Int("S", 0, "set state (0-3)")
	g := flag.Int("G", 50, "set generation")
	flag.Parse()

	populate(*s)
	display()

	for i := 0; i < *g; i++ {
		gen++
		gameOfLife()
		display()
	}

}

func gameOfLife() {

	cgrid := grid

	for x := 0; x < len(cgrid); x++ {
		for y := 0; y < len(cgrid[x]); y++ {

			// GET live neighbors
			neibors := inBound(cgrid, x, y+1) + inBound(cgrid, x, y-1) + inBound(cgrid, x-1, y) + inBound(cgrid, x+1, y) + inBound(cgrid, x-1, y-1) + inBound(cgrid, x-1, y+1) + inBound(cgrid, x+1, y-1) + inBound(cgrid, x+1, y+1)

			// Any live cell with fewer than two live neighbors dies
			// Any live cell with more than three live neighbors dies
			if cgrid[x][y] == 1 && neibors < 2 || neibors > 3 {
				grid[x][y] = 0
			}

			// Any live cell with two or three live neighbors lives on to the next generation

			// Any dead cell with exactly three live neighbors becomes a live cell
			if cgrid[x][y] == 0 && neibors == 3 {
				grid[x][y] = 1
			}

		}
	}

}

/// check if x,y is in Bound
func inBound(grid [20][20]int, x, y int) int {
	if x <= -1 || y <= -1 || x >= len(grid) || y >= len(grid[x]) {
		return 0
	}
	return grid[x][y]
}

func display() {
	time.Sleep(1 * time.Second)
	print("\033[H\033[2J")
	println(" generation", gen)
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			print("")
			if grid[i][j] == 0 {
				print("\033[37m ·")
			} else {
				print("\033[92m O")
			}

		}
		println("")

	}
}

func populate(life int) {

	switch life {
	case 0:
		for i := 0; i < 20; i++ {
			for j := 0; j < 20; j++ {
				grid[i][j] = rand.Intn(2)
			}
		}
	case 1:
		grid[8][9] = 1
		grid[9][9] = 1
		grid[10][9] = 1
	case 2:
		// row 1
		grid[9][8] = 1
		grid[9][9] = 1
		grid[9][10] = 1
		// row 2
		grid[8][9] = 1
		grid[8][8] = 1
		grid[8][7] = 1
	case 3:
		grid[3][1] = 1
		grid[3][2] = 1
		grid[3][3] = 1
		grid[2][3] = 1
		grid[1][2] = 1
	}

}
