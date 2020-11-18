package main

import "fmt"

func main() {
	// Unbuffered channel, bidirectional
	c := make(chan int)

	// Send
	go enviar(c)

	// Receive
	recibir(c)

	fmt.Println("Finalizando.")
}

func enviar(c chan<- int) {
	c <- 3
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
