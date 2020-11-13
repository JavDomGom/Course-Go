package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentar() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Javier",
			apellido: "Domínguez",
		},
		lpm: false,
	}

	fmt.Println(as1)
	as1.presentar()

	fmt.Println(as2)
	as2.presentar()
}
