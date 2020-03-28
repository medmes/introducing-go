package main

import "fmt"

func main() {

	// Go has only one loop compared to some programming language which offer a variety of options: while, do-while, until, foreach...
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// if-statement,
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// Switch cases
	var i int
	println("Please enter a number?3")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to scan, due to: ", err)
		return
	}
	// other way, in a local way we can do :
	/*	var j int
		if _, err := fmt.Scanf("%d", &j);  err != nil {
			fmt.Println("failed to scan, due to: ", err)
			return
		}*/

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// Exercises//
	//1. Response: Small
	//2. Write a program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
	}
	//3. Write a program that prints the numbers from 1 to 100, but for multiples of three, print “Fizz” instead of the number, and for the multiples of five, print “Buzz.” For numbers that are multiples of both three and five, print “FizzBuzz.”
	// FizzBuzz
	fmt.Println()

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz ")
		}
		if i%5 == 0 {
			fmt.Print("Buzz ")
		}
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Print("FizzBuzz ")
		}
	}
}
