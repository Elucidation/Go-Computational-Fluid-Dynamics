package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Run Test\n")

	const n = 11
	var (
		// The main array
		grid [n]float64
		// Temp array used by step for update
		grid_tmp [n]float64
	)

	grid[n/2] = 10

	for i := 0; i < 10; i++ {
		step(grid[:], grid_tmp[:])
		grid, grid_tmp = grid_tmp, grid // Swap
		fmt.Printf("%.2f\n", grid)
	}

	// fmt.Println(grid_tmp)
	fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))
}
