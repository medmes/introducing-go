package main

import "fmt"

func main() {

	//Functions
	numbers := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(numbers))
	x, y := returnTwoValues()
	fmt.Println(x, y)

	//Closure in Golang
	i := 0
	increment := func() int {
		i++
		return i
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(factorial(10))
}

/**
* Calculate Average of numbers, named Return
* Param: []float64
* return float64
 */
func average(xs []float64) (total float64) {
	t := 0.0
	for _, v := range xs {
		t += v
	}
	total = t / float64(len(xs))
	return
}

func returnTwoValues() (int, int) {
	return 10, 20
}

//Recursion Function
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
