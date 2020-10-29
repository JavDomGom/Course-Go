package main

import "fmt"

func main() {
	s1 := "Hola mundo"
	s2 := `Esta es una línea
	partida.`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s1 es: %T\n", s2)

	bs := []byte(s1)

	fmt.Println(bs)
	fmt.Printf("El tipo de s1 es: %T\n", bs)

	fmt.Println()

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U\n", s1[i])
	}

	fmt.Println()

	for i, v := range s1 {
		fmt.Printf("En el indice %d el valor es %#X\n", i, v)
	}

	fmt.Println()

	fmt.Printf("Con el verbo %q indico que se imprima el string %s", "%s", s1)

}
