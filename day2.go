package main

import "fmt"

// The main array

func average(arr []float64) float64 {
	panic("Not done")
}

func main() {
	fmt.Printf("Run Test\n")

	const n = 10
	var a [n]float64

	for i, _ := range a {
		a[i] = float64(i)
	}

	// total := 0.0
	// for _, v := range a {
	// 	total += v
	// }
	// avg := total / float64(len(a))
	avg := average(a[:])

	fmt.Println(a)
	fmt.Println("Average: ", avg)
}
