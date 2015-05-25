package main

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

func max(arr []float64) float64 {
	curr_max := 0.0
	for _, v := range arr {
		if v > curr_max {
			curr_max = v
		}
	}
	return curr_max
}
