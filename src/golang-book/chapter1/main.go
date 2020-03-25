package main

import (
	"fmt"
	"log"
	"os"
)

// This is a comment,  Hello World!

func main() {
	fmt.Println("Hello, World")
	// Modify name, with argument:
	if len(os.Args) > 1 {
		fmt.Println("Hello " + os.Args[1])
	} else {
		log.Println("No argument has been given!")
	}
}
