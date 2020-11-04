package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y [6]int
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	x[3] = 42
	fmt.Println(x)
}
