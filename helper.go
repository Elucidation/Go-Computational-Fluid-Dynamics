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

func step(arr_in []float64, arr_out []float64) {
	for i := range arr_in {
		switch i {
		case 0: // Near edge
			arr_out[i] = (arr_in[i]*2 + arr_in[i+1]) / 3
		case len(arr_in) - 1: // Far edge
			arr_out[i] = (arr_in[i]*2 + arr_in[i-1]) / 3
		default:
			arr_out[i] = (arr_in[i-1] + arr_in[i] + arr_in[i+1]) / 3
		}
	}
}
