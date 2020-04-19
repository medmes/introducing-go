package main

import (
	m "./math"
	"fmt"
)

func main() {

	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
