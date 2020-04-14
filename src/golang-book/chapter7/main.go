package main

import "fmt"

// Pointers

func main() {

	x := 10
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)

	a := 200
	b := &a
	*b++
	fmt.Println(*b)
	fmt.Println(a)

}

func zero(ptr *int) {
	*ptr = 0
}
