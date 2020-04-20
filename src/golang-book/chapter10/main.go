package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// We will run this as a goroutine
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Channels provide a way for two goroutines to communicate with each other and synchronize
// their execution.

/*
	Send "ping" to channel `c`

	We can specify a direction on a channel type, thus restricting it to either sending or
	receiving. For example, pinger’s function signature can be changed to this:

		func pinger(c chan<- string)

	Now pinger is only allowed to send to c. Attempting to receive from c will result in a
	compile-time error.
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // Send "ping" through channel `c`
	}
}

// Send "pong" to channel `c`
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // Send "pong" through channel `c`
	}
}

// Receive from channel `c`
func printer(c chan string) {
	for {
		msg := <-c // Receive ping through channel `c`
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	// Using Goroutine
	for i := 0; i < 10; i++ {
		go f(i) // Goroutine
	}

	// Using channels
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	go ponger(c)

	// Using `select` which is basically switch but for channels
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second): // time.After creates a channel, and after the given duration, will send the current time on it
				fmt.Println("timeout")
			default:
				// do nothing lol
			}

		}
	}()

	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}
	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest home page:", biggest.URL)
	// This prevents the program from exiting immediately because goroutines let the next line run immediately
	var input string
	fmt.Scanln(&input)
}
