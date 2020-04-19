package main

import (
	m "./math"
	"fmt"
)

func main() {

	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)

	// Exercises: Min, Max
	min := m.Min(xs)
	fmt.Println(min)

	max := m.Max(xs)
	fmt.Println(max)

	// How would you document the functions you created in #4?
	// By commenting the function Min, Max and generate the documentation using:
	// godoc golang-book/chapter8/math Min
	// godoc golang-book/chapter8/math Max
	// After generation,  the documentation is also available in web form by running this command:
	// godoc -http=":6060"
}
