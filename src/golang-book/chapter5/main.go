package main

import (
	"fmt"
	"sort"
)

func main() {

	//Arrays
	var arr [5]int
	arr[4] = 100
	fmt.Println(arr) // In Java , You need to override toString() function in order to print the entire array in a single println line

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

	// Initialize array
	xx := [5]float64{
		1,
		3,
		4,
		4,
		99.3, // Notice extra tailing
	}
	fmt.Println(xx)

	//Slice -> it's like arrayList in Java World!
	fmt.Println("_______Slices_______")
	//var slice []float64
	//slice = make([]float64, 5) // type, length
	myArray := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("myArray: ", myArray)
	fmt.Println("cap(mArray): ", len(myArray))
	fmt.Println("cap(mArray): ", cap(myArray))

	m := make([]int, 3)
	m = myArray[0:2]
	fmt.Println("m: ", m)
	fmt.Println("cap(m): ", len(m))
	fmt.Println("cap(m): ", cap(m))

	m[0] = 100
	fmt.Println("myArray: ", myArray)
	fmt.Println("m: ", m)

	m = append(m, 300)
	fmt.Println("myArray: ", myArray)
	fmt.Println("m: ", m)
	fmt.Println("cap(myArray): ", len(myArray))
	fmt.Println("cap(myArray): ", cap(myArray))
	fmt.Println("cap(m): ", len(m))
	fmt.Println("cap(m): ", cap(m))

	myArrayCopy := myArray
	myArrayCopy[0] = 1000
	fmt.Println("myArray: ", myArray)
	fmt.Println("myArrayCopy: ", myArrayCopy)

	sort.Ints(myArray)
	fmt.Println("Sorted myArray: ")
	fmt.Println("myArray: ", myArray)
	fmt.Println("myArrayCopy: ", myArrayCopy)

	//A map is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables, or dictionaries).
	//Hash Table
	fmt.Println("_______Maps_______")

	//var myMap map[int]string
	myMapInitialized := map[string]float64{
		"e":  2.71828,
		"pi": 3.1416,
		"eu": 0.57721,
	}

	//Iterating over Map
	for k, v := range myMapInitialized {
		fmt.Printf("key[%s] value[%f]\n", k, v)
	}

	// Creating map using make function
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	if name, isOK := elements["B"]; isOK {
		fmt.Println(name, isOK)
	}

	// Exes:
	xm := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(xm[2:5])
	fmt.Println(len(xm[2:5]))
	fmt.Println(cap(xm[2:5]))

	elems := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := elems[0]
	for i := 0; i < len(elems); i++ {
		if smallest > elems[i] {
			smallest = elems[i]
		}
	}
	fmt.Println(smallest)
}
