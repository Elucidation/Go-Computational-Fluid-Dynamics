package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Sim_constants struct {
	dt, dx, c float64
}

func main() {
	fmt.Printf("Run Test\n")

	const n = 1001
	const Width = 2    // meters
	const T_final = 10 // seconds
	const steps = 1000

	sc := Sim_constants{
		dt: float64(0.0025),
		dx: float64(Width) / (n - 1),
		c:  float64(0.1),
	}
	fmt.Println(sc)

	var (
		// The main array
		grid [n]float64
		// Temp array used by step for update
		grid_tmp [n]float64
	)

	// Seed grid

	for i := range grid {
		grid[i] = 1
	}
	fmt.Println(int(.5/sc.dx), int(1/sc.dx+1))
	for i := int(.5 / sc.dx); i < int(1/sc.dx+1); i++ {
		grid[i] = 2
	}
	// grid[int(1/sc.dx)] = 10

	maxInitVal := max(grid[:])

	// Draw each 1D grid as a row in a PNG image
	m := initPNG(steps, n)
	updatePNG(m, 0, grid[:], maxInitVal)

	for i := 0; i < steps; i++ {
		step(grid[:], grid_tmp[:], sc)
		grid, grid_tmp = grid_tmp, grid // Swap

		// fmt.Printf("%.1f\n", grid)
		updatePNG(m, i+1, grid[:], maxInitVal)
	}

	// fmt.Println(grid_tmp)
	fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))

	filename := fmt.Sprintf("1d_simN%dx%dS.png", n, steps)
	writePNG(m, filename)

	ShowUbuntu(filename)
}

// Iterates one step of diffusion for the grid
func step(arr_in []float64, arr_out []float64, sc Sim_constants) {
	for i := range arr_in {
		switch i {
		case 0: // Near edge
			arr_out[i] = arr_in[i]
		default:
			arr_out[i] = arr_in[i] - sc.c*sc.dt/sc.dx*(arr_in[i]-arr_in[i-1])
		}
	}
}

// show  a specified file by Preview.app for OS X(darwin)
func ShowUbuntu(name string) {
	command := "eog"
	cmd := exec.Command(command, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// show  a specified file by Preview.app for OS X(darwin)
func ShowMac(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
