package main

import "fmt"

func main() {
	p1 := struct {
		nombre   string
		apellido string
		edad     int
	}{
		nombre:   "Javier",
		apellido: "Domínguez",
		edad:     39,
	}

	fmt.Println(p1)

	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
}
