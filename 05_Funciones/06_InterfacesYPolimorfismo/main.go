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
	fmt.Println("Hola, soy", a.nombre, a.apellido, "el agente secreto se presenta.")
}

func (p persona) presentar() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta.")
}

type humano interface {
	presentar()
}

func bar(h humano) {
	switch h.(type) {
	case persona:
		fmt.Println("Fui pasado a la función bar (persona)", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("Fui pasado a la función bar (agenteSecreto)", h.(agenteSecreto).nombre)
	}
	fmt.Println("Fui pasado a la función bar", h)
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

	p := persona{
		nombre:   "Beatriz",
		apellido: "Giménez",
	}

	bar(as1)
	bar(as2)
	bar(p)
}
