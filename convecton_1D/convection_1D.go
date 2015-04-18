package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Sim_constants struct {
	n, width, steps int
	dt, dx, c       float64
}

func main() {
	fmt.Printf("Run Test\n")

	sc := Sim_constants{
		n:     1000, // Number of cells
		width: 2,    // meters of all cells
		// const T_final = 10 // seconds
		steps: 2000,             // Number of steps
		dt:    float64(0.00025), // Time step
		// dx:    float64(width) / (n - 1),
		c: float64(0.1),
	}
	sc.dx = float64(sc.width) / float64(sc.n-1)

	fmt.Println(sc)

	// The main array
	grid := make([]float64, sc.n)
	// Temp array used by step for update
	grid_tmp := make([]float64, sc.n)

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
	m := initPNG(sc.steps, sc.n)
	updatePNG(m, 0, grid[:], maxInitVal)

	for i := 0; i < sc.steps; i++ {
		step(grid[:], grid_tmp[:], sc)
		grid, grid_tmp = grid_tmp, grid // Swap

		// fmt.Printf("%.1f\n", grid)
		updatePNG(m, i+1, grid[:], maxInitVal)
		fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))
	}

	// fmt.Println(grid_tmp)
	fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))

	filename := fmt.Sprintf("1d_simN%dx%dS.png", sc.n, sc.steps)
	writePNG(m, filename)

	// ShowUbuntu(filename)
	ShowMac(filename)
}

// Iterates one step of diffusion for the grid
func step(arr_in []float64, arr_out []float64, sc Sim_constants) {
	for i := range arr_in {
		switch i {
		case 0: // Near edge
			arr_out[i] = arr_in[i] // - arr_in[i]*sc.dt/sc.dx*(arr_in[i])
		default:
			arr_out[i] = arr_in[i] - arr_in[i]*sc.dt/sc.dx*(arr_in[i]-arr_in[i-1])
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
