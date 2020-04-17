package main

import (
	"fmt"
	"math"
)

// Structs and Interfaces
// `type` keyword introduces a new type
// followed by name of the type, `Circle` in this case
// concluded with `struct` keyword to indicate we are defining a struct
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

type Shape interface { // Declaring an interface
	/*
	   But instead of defining fields, we define a method set.
	   A method set is a list of methods that a type must have
	   in order to implement the interface.
	*/
	area() float64
}

// Interfaces can also be used as fields. Consider a MultiShape that is made up of several smaller shapes:
type MultiShape struct {
	shapes []Shape
}

// function calculate circle area:
func calculateCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//`(c *Circle)`is a receiver and means that the `Circle` struct contains this method
// Circle method to calculate circle area
func (c *Circle) area() float64 { // Receiver Type
	fmt.Println("circle Area:")
	return math.Pi * c.r * c.r
}

/*func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}*/

// Rectangle Area method
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	fmt.Println("Rectangle Area:")
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
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

/*
	Advice oon interfaces from book:
	When building new programs, you often won’t know what your types should look
	like when you start—and that’s OK. In Go, you generally focus more on the behavior
	of your program than on a taxonomy of types. Create a few small structs that do what
	you want, add in methods that you need, and as you build your program, useful interfaces
	will tend to emerge. There’s no need to have them all figured out ahead of time.
	Interfaces are particularly useful as software projects grow and become more complex.
	They allow us to hide the incidental details of implementation (e.g., the fields of
	our struct), which makes it easier to reason about software components in isolation.
	In our example, as long as the area methods we defined continue to produce the
	same results, we’re free to change how a Circle or Rectangle is structured without
	having to worry about whether or not the totalArea function will continue to work.
*/

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

	fmt.Println("____Interface_____")
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	for _, shape := range multiShape.shapes {
		fmt.Println(shape.area())
	}
}
