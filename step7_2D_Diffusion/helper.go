package main

func average(arr [][]float64) float64 {
	total := 0.0
	for i := range arr {
		for j := range arr[i] {
			total += arr[i][j]
		}
	}
	return total / float64(len(arr))
}

func sum(arr [][]float64) float64 {
	total := 0.0
	for i := range arr {
		for j := range arr[i] {
			total += arr[i][j]
		}
	}
	return total
}

func max(arr [][]float64) float64 {
	curr_max := 0.0
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] > curr_max {
				curr_max = arr[i][j]
			}
		}
	}
	return curr_max
}
