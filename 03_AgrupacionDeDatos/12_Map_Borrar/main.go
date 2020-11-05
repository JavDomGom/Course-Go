package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	delete(m, "Robin")
	fmt.Println(m)

	if v, ok := m["Batman"]; ok {
		fmt.Println("Se borró la clave con valor", v)
		delete(m, "Batman")
	}
	fmt.Println(m)
}
