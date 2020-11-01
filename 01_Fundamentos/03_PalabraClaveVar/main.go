package main

import "fmt"

// scope: package
var z = 41

func main() {
	x := 43
	y := 2 + 4

	fmt.Println(x, y, z)
}

func numero() {
	fmt.Println(z)
}
