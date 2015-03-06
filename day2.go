package main

import "fmt"

// The main array

func average(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}

func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

func step(arr_in []float64, arr_out []float64) {
	for i := range arr_in {
		switch i {
		case 0: // Near edge
			arr_out[i] = (arr_in[i]*2 + arr_in[i+1]) / 3
		case len(arr_in) - 1: // Far edge
			arr_out[i] = (arr_in[i]*2 + arr_in[i-1]) / 3
		default:
			arr_out[i] = (arr_in[i-1] + arr_in[i] + arr_in[i+1]) / 3
		}
	}
}

func main() {
	fmt.Printf("Run Test\n")

	const n = 11
	var (
		grid     [n]float64
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
