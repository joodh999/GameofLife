package main

import (
	"flag"
	"math/rand"
	"time"
)

var gridstate [20][20]int
var gen = 0

func main() {

	state := flag.Int("state", 0, "set state (0-3)")
	gen := flag.Int("gen", 10, "set generation")

	print(*state)
	println(*gen)
	// setlife(3)
	// display()

}

func display() {
	time.Sleep(1 * time.Second)
	print("\033[H\033[2J")
	println(" generation", gen)
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			print("")
			if gridstate[i][j] == 0 {
				print("\033[37m ·")
			} else {
				print("\033[92m O")
			}

		}
		println("")

	}
}

func setlife(life int) {

	switch life {
	case 0:
		for i := 0; i < 20; i++ {
			for j := 0; j < 20; j++ {
				gridstate[i][j] = rand.Intn(2)
			}
		}
	case 1:
		gridstate[8][9] = 1
		gridstate[9][9] = 1
		gridstate[10][9] = 1
	case 2:
		// row 1
		gridstate[9][8] = 1
		gridstate[9][9] = 1
		gridstate[9][10] = 1
		// row 2
		gridstate[8][9] = 1
		gridstate[8][8] = 1
		gridstate[8][7] = 1
	case 3:
		gridstate[3][1] = 1
		gridstate[3][2] = 1
		gridstate[3][3] = 1
		gridstate[2][3] = 1
		gridstate[1][2] = 1
	}

}
