package main

import (
	"fmt"
)

func main() {
	out := make(chan int)
	in := make(chan int)

	// Create 3 'multiplyByTwo' goroutines
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Send data into goroutines
	go func() {
		in <- 1
		in <- 2
		in <- 3
		in <- 4
	}()

	// Wait for each result to arrive
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

// This func now accepts a channel as its second argument...
func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	for {
		num := <-in
		result := num * 2
		out <- result
	}
}
