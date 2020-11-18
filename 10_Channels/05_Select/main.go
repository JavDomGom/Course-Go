package main

import "fmt"

func main() {
	// Unbuffered channel, bidirectional
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// Send
	go enviar(par, impar, salir)

	// Receive
	recibir(par, impar, salir)

	fmt.Println("Finalizando.")
}

func enviar(par, impar, salir chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	salir <- 0
}

func recibir(par, impar, salir <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v)
		case v := <-salir:
			fmt.Println("Desde el canal salir:", v)
			return
		}
	}
}
