package main

import (
	"fmt"
)

func main() {
	n := 3

	// This is where we "make" the channel, which can be used
	// to move the 'int' datatype
	out := make(chan int)

	// We still run this func as a goroutine, but this time,
	// the channel that we made is also provided
	go multiplyByTwo(n, out)

	// Once any output is received on the channel, print it to the console
	// and proceed
	fmt.Println(<-out)
}

// This func now accepts a channel as its second argument...
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2

	out <- result
}
