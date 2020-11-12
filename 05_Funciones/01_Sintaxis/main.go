package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := saludar("Javier", "Domínguez")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("Hola desde foo.")
}

func bar(s string) {
	fmt.Println("Hola", s)
}

func woo(s string) string {
	return fmt.Sprint("Hola ", s)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, " dice hola.")
	q := true

	return p, q
}
