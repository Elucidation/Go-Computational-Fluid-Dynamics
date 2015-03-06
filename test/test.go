package main

import "fmt"

func main() {
	fmt.Printf("Run Test\n")

	const n = 10
	// The main array
	var a [n]int

	for i, _ := range a {
		a[i] = i
	}

	fmt.Println(a)
	fmt.Println(a[:5], a[5:])
}
