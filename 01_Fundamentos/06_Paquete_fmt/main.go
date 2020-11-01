package main

import "fmt"

// scope: package
var a int
var b string = "Programa"

func main() {
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("El valor de la variable a es: %v\n", a)
	fmt.Printf("El valor de la variable b es: %v\n", b)

	fmt.Printf("El valor de la variable a es: %d\n", a)
	fmt.Printf("El valor de la variable b es: %s\n", b)

	fmt.Printf("El tipo de la variable a es: %T\n", a)
	fmt.Printf("El tipo de la variable b es: %T\n", b)

	s1 := fmt.Sprint("El ", b, " dice hola mundo.")
	fmt.Println(s1)
	fmt.Printf("El tipo de la variable s1 es: %T\n", s1)
}
