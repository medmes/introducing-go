/**
* custom package math
 */
package math

// Calculate average of numbers list:
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Calculate min of numbers list:
func Min(numbers []float64) (min float64) {
	min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return
}

// Calculate max of numbers list:
func Max(numbers []float64) (max float64) {
	max = numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return
}
