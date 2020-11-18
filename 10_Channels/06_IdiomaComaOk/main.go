package main

import "fmt"

func main() {
	// Unbuffered channel, bidirectional
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	// Send
	go enviar(par, impar, salir)

	// Receive
	recibir(par, impar, salir)

	fmt.Println("Finalizando.")
}

func enviar(par, impar chan<- int, salir chan<- bool) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	close(salir)
}

func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v)
		case i, ok := <-salir:
			if !ok {
				fmt.Println("Desde comma ok:", i)
				return
			} else {
				fmt.Println("Desde comma ok:", i)
			}
		}
	}
}
