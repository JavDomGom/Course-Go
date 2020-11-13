package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Estoy en la función anónima.")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(42)
}

func foo() {
	fmt.Println("Estoy en foo.")
}
