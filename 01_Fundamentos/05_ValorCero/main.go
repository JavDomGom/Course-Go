package main

import "fmt"

// scope: package
var a int
var b float32
var c string

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	a = 1
	b = 3.14
	c = "Hola"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
