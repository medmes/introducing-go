package main

import "fmt"

func main() {

	//Functions
	fmt.Println("______Functions______")

	numbers := []float64{98, 93, 77, 82, 83}
	defer deferred()
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
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	// Pointers
	fmt.Println("______Pointers______")
	xPtr := 10
	zero(&xPtr)
	fmt.Println(xPtr)
	fmt.Println(&xPtr)

	a := 200
	b := &a
	*b++
	fmt.Println(*b)
	fmt.Println(a)

	panic("PANIC")
	// Unreachable code, after this line!
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

func deferred() {
	fmt.Println("This function will be invoked after the main function exit!")
}

func zero(ptr *int) {
	*ptr = 0
}
