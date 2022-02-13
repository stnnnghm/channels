package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan int)
	out2 := make(chan int)

	// Start both fast and slow
	// goroutines with different channels
	go fast(2, out1)
	go slow(3, out2)

	// Perform some action depending on which channel
	// receives information first
	select {
	case res := <-out1:
		fmt.Println("fast finished first, result:", res)
	case res := <-out2:
		fmt.Println("slow finished first, result:", res)
	}
}

func fast(num int, out chan<- int) {
	result := num * 2
	time.Sleep(5 * time.Millisecond)
	out <- result
}

func slow(num int, out chan<- int) {
	result := num * 2
	time.Sleep(15 * time.Millisecond)
	out <- result
}
