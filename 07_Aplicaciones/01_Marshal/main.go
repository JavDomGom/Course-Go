package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		Nombre:   "Javier",
		Apellido: "Domínguez",
		Edad:     39,
	}
	p2 := persona{
		Nombre:   "Beatriz",
		Apellido: "Giménez",
		Edad:     47,
	}
	personas := []persona{p1, p2}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
