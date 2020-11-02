package main

import "fmt"

func main() {
	switch "Manzana" {
	case "Pera", "Limón":
		fmt.Println("No debe imprimir nada 1")
	case "Manzana", "Ciruela", "Fresas":
		fmt.Println("Debe imprimir algo")
		fallthrough // fuerza a imprimir también el caso siguiente.
	default:
		fmt.Println("Debe imprimir algo por default")
	}
}
