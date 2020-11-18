package main

import "fmt"

func main() {
	// Buffered channel, bidirectional
	c1 := make(chan int, 2)
	fmt.Printf("%T\n", c1)

	// Buffered channel, send-only
	c2 := make(chan<- int, 2)
	fmt.Printf("%T\n", c2)

	// Buffered channel, receive-only
	c3 := make(<-chan int, 2)
	fmt.Printf("%T\n", c3)
}
