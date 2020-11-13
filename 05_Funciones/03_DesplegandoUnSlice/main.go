package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("Hola", xi...)
	fmt.Println(x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Println(s)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}

	return suma
}
