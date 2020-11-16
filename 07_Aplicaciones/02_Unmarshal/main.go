package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"Javier","Apellido":"Domínguez","Edad":39},{"Nombre":"Beatriz","Apellido":"Giménez","Edad":47}]`
	sb := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", sb)

	var personas []persona

	err := json.Unmarshal(sb, &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(personas)

	for _, v := range personas {
		fmt.Println(v)
	}
}
