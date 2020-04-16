package main

import (
	"fmt"
	"math"
)

// Structs and Interfaces

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person Person
	Model  string
}

func main() {
	// 1- struct instance
	c := new(Circle) // -> nil struct value
	// affecting
	c.x = 2.0
	c.y = 4.4
	c.r = 3.5
	fmt.Println(c)

	// 2 - Struct reference with constructor
	cc := Circle{x: 20, r: 23.4}
	fmt.Println(cc)

	// 3- Struct reference , if we know the order
	ccc := Circle{10, 3.9, 5}
	fmt.Println(ccc)

	// Calculate Circle Area
	fmt.Println(calculateCircleArea(*c))

	// Passing the struct by default in Go is by value, so the struct is copied
	fmt.Println("Struct default passing!")

	fmt.Println(ccc)
	zeroRWithoutPointerAsParameter(ccc)
	fmt.Println(ccc)

	// Since , the default is passing by value, we want to change the struct passed itself
	fmt.Println("Struct passing by reference!")
	fmt.Println(ccc)
	zeroRPointerAsParameter(&ccc)
	fmt.Println(ccc)

	// We can define methods as functions specefic to a struct, by calling <struct>.function:
	fmt.Println(c.area())

	// Embedded types
	person := Person{Name: "Mo"}
	person.Talk()
	// In Go , support relationship is rather has!
	// Android is a Person , rather than Android Has a Person.
	android := new(Android) // var android = new(Android)
	android.Model = "X-MAN"
	fmt.Println(android.Person.Name)
	fmt.Println(android.Model)
	android.Person.Talk()

}

// function calculate circle area:
func calculateCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// Circle method to calculate circle area
func (c *Circle) area() float64 { // Receiver Type
	return math.Pi * c.r * c.r
}

func zeroRWithoutPointerAsParameter(c Circle) Circle {
	c.r = 0
	return c
}

func zeroRPointerAsParameter(c *Circle) *Circle {
	c.r = 0
	return c
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
