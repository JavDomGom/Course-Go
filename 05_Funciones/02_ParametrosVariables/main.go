package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	print(x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}

	return suma
}
