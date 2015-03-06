package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Run Test\n")

	const n = 400
	const steps = 400

	var (
		// The main array
		grid [n]float64
		// Temp array used by step for update
		grid_tmp [n]float64
	)

	// Seed grid
	grid[n/4] = 100
	grid[n/2] = 100
	grid[3*n/4] = 50

	maxInitVal := max(grid[:])

	// Draw each 1D grid as a row in a PNG image
	m := initPNG(steps, n)
	updatePNG(m, 0, grid[:], maxInitVal)

	for i := 0; i < steps; i++ {
		step(grid[:], grid_tmp[:])
		grid, grid_tmp = grid_tmp, grid // Swap

		// fmt.Printf("%.2f\n", grid)
		updatePNG(m, i+1, grid[:], maxInitVal)
	}

	// fmt.Println(grid_tmp)
	fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))

	writePNG(m, fmt.Sprintf("1d_simN%dx%dS.png", n, steps))
}
