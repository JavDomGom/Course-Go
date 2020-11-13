package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Mi primera expresión función.")
	}
	f()

	g := func(x int) {
		fmt.Println("Colón llegó a America en", x)
	}
	g(1492)
}
