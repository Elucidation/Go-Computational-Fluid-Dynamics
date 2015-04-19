package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
)

type Sim_constants struct {
	n, steps                    int
	width, dt, dx, c, sigma     float64
	totaltime, maxintensity, nu float64
}

func main() {
	fmt.Printf("Run Test\n")

	sc := Sim_constants{
		n:     1000, // Number of cells
		width: 1,    // meters of all cells
		// steps: 1000, // Number of steps
		totaltime: 0.1, // seconds
		// c: float64(0.1),
		nu:           0.2, // viscosity
		maxintensity: 2,
		// sigma:        0.1, // used to determine timestep via courant number
	}
	sc.dx = sc.width / float64(sc.n-1)
	sc.sigma = 1 / sc.maxintensity
	sc.dt = sc.sigma * math.Pow(sc.dx, 2) / sc.nu
	sc.steps = int(sc.totaltime / sc.dt)
	if sc.steps > 2000 {
		sc.steps = 2000
	}

	fmt.Println("sigma = ", sc.sigma, ", dx = ", sc.dx, ", dt = ", sc.dt)
	fmt.Println("Total time = ", float64(sc.steps)*sc.dt, " seconds")

	fmt.Println(sc)

	// The main array
	grid := make([]float64, sc.n)
	// Temp array used by step for update
	grid_tmp := make([]float64, sc.n)

	// grid := make([][]int, row)
	// for i := range grid {
	// 	grid[i] = make([]int, col)
	// }
	// grid_tmp := make([][]int, row)
	// for i := range grid_tmp {
	// 	grid_tmp[i] = make([]int, col)
	// }

	// Seed grid
	for i := range grid {
		grid[i] = 1
		grid_tmp[i] = 1
	}
	fmt.Println(int(.5*sc.width/sc.dx), int(1*sc.width/sc.dx+1))
	for i := int(0.4 * float64(sc.n)); i < int(0.6*float64(sc.n)+1); i++ {
		grid[i] = sc.maxintensity
	}
	// for i := int(0.7 * float64(sc.n)); i < int(0.72*float64(sc.n)+1); i++ {
	// 	grid[i] = sc.maxintensity
	// }

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
	arr_end := len(arr_in) - 1
	for i := 1; i < arr_end; i++ {
		arr_out[i] = arr_in[i] +
			sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
				(arr_in[i+1]-2*arr_in[i]+arr_in[i-1]) - arr_in[i]*sc.dt/sc.dx*(arr_in[i]-arr_in[i-1])
	}

	// Wrap around
	arr_out[0] = arr_in[0] +
		sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
			(arr_in[1]-2*arr_in[0]+arr_in[arr_end]) - arr_in[0]*sc.dt/sc.dx*(arr_in[0]-arr_in[arr_end])

	arr_out[arr_end] = arr_in[arr_end] +
		sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
			(arr_in[0]-2*arr_in[arr_end]+arr_in[arr_end-1]) - arr_in[arr_end]*sc.dt/sc.dx*(arr_in[arr_end]-arr_in[arr_end-1])
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
