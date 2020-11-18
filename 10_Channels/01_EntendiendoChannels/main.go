package main

import "fmt"

func main() {
	// Unbuffered channel
	c1 := make(chan int)

	// Al ser unbuffered es necesario que una goroutine
	// emisora envíe a otra goroutine receptora.
	go func() {
		c1 <- 3
	}()

	fmt.Println(<-c1)

	// Buffered channel
	c2 := make(chan int, 2)
	c2 <- 9
	c2 <- 11
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
