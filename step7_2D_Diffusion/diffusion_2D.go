package main

import (
	"fmt"
	"log"
	// "math"
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
		nx:     400, // Number of cells
		ny:     400, // Number of cells
		width:  2,   // meters of all cells
		height: 2,   // meters of all cells
		// steps: 1000, // Number of steps
		totaltime:    10, // seconds
		c:            1.0,
		nu:           0.05, // viscosity
		maxintensity: 2,
		// sigma:        0.1, // used to determine timestep via courant number
	}
	sc.dx = sc.width / float64(sc.nx-1)
	sc.dy = sc.height / float64(sc.ny-1)
	// sc.sigma = 1 / sc.maxintensity
	sc.sigma = 0.25
	sc.dt = sc.sigma * sc.dx * sc.dy / sc.nu
	// sc.steps = int(sc.totaltime / sc.dt)
	sc.steps = 200
	// if sc.steps > 2000 {
	// 	sc.steps = 2000
	// }

	fmt.Println("sigma = ", sc.sigma, ", dx = ", sc.dx, ", dt = ", sc.dt)
	fmt.Println("Total time = ", float64(sc.steps)*sc.dt, " seconds")

	fmt.Println(sc)

	// allocate arrays
	gridu := make([][]float64, sc.nx)
	gridu_tmp := make([][]float64, sc.nx)
	gridv := make([][]float64, sc.nx)
	gridv_tmp := make([][]float64, sc.nx)

	for i := range gridu {
		gridu[i] = make([]float64, sc.ny)
		gridu_tmp[i] = make([]float64, sc.ny)
		gridv[i] = make([]float64, sc.ny)
		gridv_tmp[i] = make([]float64, sc.ny)
	}

	// Seed gridu
	for i := range gridu {
		for j := range gridu[i] {
			gridu[i][j] = 1
			gridu_tmp[i][j] = 1
			gridv[i][j] = 1
			gridv_tmp[i][j] = 1
		}
	}
	fmt.Println(int(.5*sc.width/sc.dx), int(1*sc.width/sc.dx+1))
	for i := int(0.4 * float64(sc.nx)); i < int(0.6*float64(sc.nx)+1); i++ {
		for j := int(0.4 * float64(sc.ny)); j < int(0.6*float64(sc.ny)+1); j++ {
			gridu[i][j] = sc.maxintensity
			gridv[i][j] = sc.maxintensity
		}
	}

	// for i := int(0.7 * float64(sc.n)); i < int(0.72*float64(sc.n)+1); i++ {
	// 	gridu[i] = sc.maxintensity
	// }

	// Draw each 1D gridu as a row in a PNG image
	m := initPNG(sc.ny, sc.nx)
	updatePNG(m, gridu[:], sc.maxintensity)
	filename := fmt.Sprintf("%s/1d_sim%dx%d_%03dS.png", filedir, sc.nx, sc.ny, 0)
	writePNG(m, filename)

	for i := 0; i < sc.steps; i++ {
		step2DDiffusion(gridu[:], gridu_tmp[:], gridv[:], gridv_tmp[:], sc)
		gridu, gridu_tmp = gridu_tmp, gridu // Swap
		gridv, gridv_tmp = gridv_tmp, gridv

		// fmt.Printf("%.1f\n", gridu)
		updatePNG(m, gridu[:], sc.maxintensity)
		fmt.Println("Sum: ", sum(gridu[:]), ", Average: ", average(gridu[:]))
		filename := fmt.Sprintf("%s/1d_sim%dx%d_%03dS.png", filedir, sc.nx, sc.ny, i+1)
		writePNG(m, filename)
		// ShowMac(filename)
	}

	// fmt.Println(gridu_tmp)
	fmt.Println("u Sum: ", sum(gridu[:]), ", Average: ", average(gridu[:]))
	fmt.Println("v Sum: ", sum(gridv[:]), ", Average: ", average(gridv[:]))

	// ShowUbuntu(filename)
	// ShowMac(filename)
}

// Diffusion
func step2DDiffusion(arru_in [][]float64, arru_out [][]float64,
	arrv_in [][]float64, arrv_out [][]float64, sc Sim_constants) {

	for i := 1; i < sc.nx-1; i++ {
		for j := 1; j < sc.ny-1; j++ {
			arru_out[i][j] = arru_in[i][j] +
				sc.nu*sc.dt/(sc.dx*sc.dx)*
					(arru_in[i+1][j]-2*arru_in[i][j]+arru_in[i-1][j]) +
				sc.nu*sc.dt/(sc.dy*sc.dy)*
					(arru_in[i][j+1]-2*arru_in[i][j]+arru_in[i][j-1])
			// arrv_out[i][j] = arrv_in[i][j] -
			// 	arru_in[i][j]*sc.dt/sc.dx*(arrv_in[i][j]-arrv_in[i-1][j]) -
			// 	arrv_in[i][j]*sc.dt/sc.dy*(arrv_in[i][j]-arrv_in[i][j-1])
		}
	}
}

// Convection
func step2DConvection(arru_in [][]float64, arru_out [][]float64,
	arrv_in [][]float64, arrv_out [][]float64, sc Sim_constants) {
	for i := 1; i < sc.nx-1; i++ {
		for j := 1; j < sc.ny-1; j++ {
			arru_out[i][j] = arru_in[i][j] -
				arru_in[i][j]*sc.dt/sc.dx*(arru_in[i][j]-arru_in[i-1][j]) -
				arrv_in[i][j]*sc.dt/sc.dy*(arru_in[i][j]-arru_in[i][j-1])
			arrv_out[i][j] = arrv_in[i][j] -
				arru_in[i][j]*sc.dt/sc.dx*(arrv_in[i][j]-arrv_in[i-1][j]) -
				arrv_in[i][j]*sc.dt/sc.dy*(arrv_in[i][j]-arrv_in[i][j-1])
		}
	}
}

func stepConvection(arr_in [][]float64, arr_out [][]float64,
	arrv_in [][]float64, arrv_out [][]float64, sc Sim_constants) {
	for i := 1; i < sc.nx-1; i++ {
		for j := 1; j < sc.ny-1; j++ {
			arr_out[i][j] = arr_in[i][j] -
				sc.c*sc.dt/sc.dx*(arr_in[i][j]-arr_in[i-1][j]) - sc.c*sc.dt/sc.dy*(arr_in[i][j]-arr_in[i][j-1])
		}
	}
}

// Linear Convection
func stepLinearConvection(arr_in [][]float64, arr_out [][]float64, sc Sim_constants) {
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
