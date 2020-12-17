package main

import (
	"flag"
	"math/rand"
	"time"
)

var grid [20][20]int
var gen = 0

func main() {

	state := flag.Int("S", 0, "set state (0-3)")
	gens := flag.Int("G", 10, "set generation")
	flag.Parse()

	println(*state)
	println(*gens)

	grid[9][9] = 1

	grid[8][9] = 1
	grid[10][9] = 1

	grid[10][10] = 1
	grid[10][8] = 1

	// display()
	nextgen()
	// display()

}

func nextgen() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if grid[i][j] == 1 {

				// Each cell with one or no neighbors dies, as if by solitude.
				if getneighbors(i, j) == 0 || getneighbors(i, j) == 1 {
					grid[i][j] = 0
				}

				// Each cell with four or more neighbors dies, as if by overpopulation.
				if getneighbors(i, j) <= 4 {
					grid[i][j] = 0
				}

				// Each cell with two or three neighbors survives
				if getneighbors(i, j) == 2 || getneighbors(i, j) == 3 {
					grid[i][j] = 1
				}

			} else if grid[i][j] == 0 {
				// Each cell with three neighbors becomes populated.
				if getneighbors(i, j) == 3 {
					grid[i][j] = 1
				}
			}

		}
	}
}

func getneighbors(x, y int) int {
	return grid[x][y+1] + grid[x][y-1] + grid[x-1][y] + grid[x+1][y] + grid[x-1][y-1] + grid[x-1][y+1] + grid[x+1][y-1] + grid[x+1][y-1]

}

// func checknegative(num int) int {
// 	if num < 0 {
// 		return 0
// 	}
// 	return num

// }

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

/*
	grid[9][10] = 1 // up
	grid[9][8] = 1  //down

	// left & right
	grid[8][9] = 1
	grid[10][9] = 1

	// sad
	grid[8][10] = 1
	grid[8][8] = 1

	grid[10][10] = 1
	grid[10][8] = 1
	grid[10]
*/
