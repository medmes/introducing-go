package main

import (
	"fmt"
)

var sharedVariable string = "SHARED VARIABLE!" // Scope

func main() {

	var x string
	x = "Hello World!"
	x += " and GoodBye!"
	fmt.Println(x)

	y := "without type declaration!"
	fmt.Println(y)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	fmt.Println(sharedVariable)

	f()

	// Go support Constants types , for Immutability

	const constantVariable string = "constannts!"
	const constantVariable2 = "constannts!"
	const constantVariable3 = 30.8789
	fmt.Println(constantVariable)
	fmt.Println(constantVariable2)
	fmt.Println(constantVariable3)

	//Go also offer to declare multiple variable in one shot, with another shorthand:
	var (
		a = 14
		b = " - "
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Scanf as C language for reading input data
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}

func f() {
	fmt.Println(sharedVariable)
}
