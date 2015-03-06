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

func main() {
	fmt.Printf("Run Test\n")

	const n = 10
	var a [n]float64

	for i, _ := range a {
		a[i] = float64(i)
	}

	avg := average(a[:])

	fmt.Println(a)
	fmt.Println("Average: ", avg)
}
