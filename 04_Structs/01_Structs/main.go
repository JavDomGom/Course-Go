package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := persona{
		nombre:   "Javier",
		apellido: "Domínguez",
		edad:     39,
	}

	p2 := persona{
		nombre:   "Beatriz",
		apellido: "Giménez",
		edad:     47,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	fmt.Println(p2.edad)
}
