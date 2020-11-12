package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Javier",
			apellido: "Domínguez",
			edad:     39,
		},
		lpm: true,
	}

	p2 := persona{
		nombre:   "Beatriz",
		apellido: "Giménez",
		edad:     47,
	}

	fmt.Println(as1)
	fmt.Println(p2)

	fmt.Println(as1.nombre)
	fmt.Println(as1.apellido)
	fmt.Println(p2.edad)
}
