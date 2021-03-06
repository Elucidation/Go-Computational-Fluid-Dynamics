package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
)

const filedir = "output"

type Sim_constants struct {
	nx, ny, steps                       int
	width, height, dt, dx, dy, c, sigma float64
	totaltime, maxintensity, nu         float64
}

func main() {
	fmt.Printf("Run Test\n")

	sc := Sim_constants{
		nx:     100, // Number of cells
		ny:     100, // Number of cells
		width:  1,   // meters of all cells
		height: 1,   // meters of all cells
		// steps: 1000, // Number of steps
		totaltime:    0.1, // seconds
		c:            1.0,
		nu:           0.2, // viscosity
		maxintensity: 2,
		// sigma:        0.1, // used to determine timestep via courant number
	}
	sc.dx = sc.width / float64(sc.nx-1)
	sc.dy = sc.height / float64(sc.ny-1)
	sc.sigma = 1 / sc.maxintensity
	sc.dt = sc.sigma * math.Pow(sc.dx, 2) / sc.nu
	sc.steps = int(sc.totaltime / sc.dt)
	if sc.steps > 200 {
		sc.steps = 200
	}

	fmt.Println("sigma = ", sc.sigma, ", dx = ", sc.dx, ", dt = ", sc.dt)
	fmt.Println("Total time = ", float64(sc.steps)*sc.dt, " seconds")

	fmt.Println(sc)

	// The main array
	// grid := make([]float64, sc.n)
	// Temp array used by step for update
	// grid_tmp := make([]float64, sc.n)

	grid := make([][]float64, sc.nx)
	for i := range grid {
		grid[i] = make([]float64, sc.ny)
	}
	grid_tmp := make([][]float64, sc.nx)
	for i := range grid_tmp {
		grid_tmp[i] = make([]float64, sc.ny)
	}

	// Seed grid
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 1
			grid_tmp[i][j] = 1
		}
	}
	fmt.Println(int(.5*sc.width/sc.dx), int(1*sc.width/sc.dx+1))
	for i := int(0.4 * float64(sc.nx)); i < int(0.6*float64(sc.nx)+1); i++ {
		for j := int(0.4 * float64(sc.ny)); j < int(0.6*float64(sc.ny)+1); j++ {
			grid[i][j] = sc.maxintensity
		}
	}

	// for i := int(0.7 * float64(sc.n)); i < int(0.72*float64(sc.n)+1); i++ {
	// 	grid[i] = sc.maxintensity
	// }

	// Draw each 1D grid as a row in a PNG image
	m := initPNG(sc.ny, sc.nx)
	updatePNG(m, grid[:], sc.maxintensity)
	filename := fmt.Sprintf("%s/1d_sim%dx%d_%03dS.png", filedir, sc.nx, sc.ny, 0)
	writePNG(m, filename)

	for i := 0; i < sc.steps; i++ {
		step(grid[:], grid_tmp[:], sc)
		grid, grid_tmp = grid_tmp, grid // Swap

		// fmt.Printf("%.1f\n", grid)
		updatePNG(m, grid[:], sc.maxintensity)
		fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))
		filename := fmt.Sprintf("%s/1d_sim%dx%d_%03dS.png", filedir, sc.nx, sc.ny, i+1)
		writePNG(m, filename)
	}

	// fmt.Println(grid_tmp)
	fmt.Println("Sum: ", sum(grid[:]), ", Average: ", average(grid[:]))

	// ShowUbuntu(filename)
	ShowMac(filename)
}

// Iterates one step of diffusion for the grid
func step(arr_in [][]float64, arr_out [][]float64, sc Sim_constants) {
	for i := 1; i < sc.nx-1; i++ {
		for j := 1; j < sc.ny-1; j++ {
			// arr_out[i][j] = arr_in[i][j] +
			// 	sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
			// 		(arr_in[i+1]-2*arr_in[i]+arr_in[i-1]) - arr_in[i]*sc.dt/sc.dx*(arr_in[i]-arr_in[i-1])
			arr_out[i][j] = arr_in[i][j] -
				sc.c*sc.dt/sc.dx*(arr_in[i][j]-arr_in[i-1][j]) - sc.c*sc.dt/sc.dy*(arr_in[i][j]-arr_in[i][j-1])
		}
	}

	// Wrap around
	// arr_out[0] = arr_in[0] +
	// 	sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
	// 		(arr_in[1]-2*arr_in[0]+arr_in[arr_end]) - arr_in[0]*sc.dt/sc.dx*(arr_in[0]-arr_in[arr_end])

	// arr_out[arr_end] = arr_in[arr_end] +
	// 	sc.nu*sc.dt/(math.Pow(sc.dx, 2))*
	// 		(arr_in[0]-2*arr_in[arr_end]+arr_in[arr_end-1]) - arr_in[arr_end]*sc.dt/sc.dx*(arr_in[arr_end]-arr_in[arr_end-1])
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
