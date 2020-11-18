package main

import "fmt"

func main() {
	// Unbuffered channel, bidirectional
	c := make(chan int)

	// Send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}
