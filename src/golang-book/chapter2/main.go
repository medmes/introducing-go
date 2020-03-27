package main

import "fmt"

/**
* Go’s integer types are unsigned: uint8, uint16, uint32, uint64. signed: int8, int16, int32, and int64.
* Go has two floating-point types: float32 and float64
* (also often referred to as single precision and double precision, respectively).
* it also has an additional data type for representing complex numbers: complex64 and complex128.
 */
func main() {

	fmt.Println("1 + 1= ", 1.0+1.0) // 2

	fmt.Println(len("Hello, World!")) // 13
	fmt.Println("Hello, World"[1])    // will show 101 as byte number instead of e

	fmt.Println(true && true)                                            // true
	fmt.Println(true && false)                                           // false
	fmt.Println(true || true)                                            // true
	fmt.Println(true || false)                                           // true
	fmt.Println(!true)                                                   // false
	fmt.Println((true && false) || (false && true) || !(false && false)) // true
}
