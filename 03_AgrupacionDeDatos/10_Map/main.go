package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	fmt.Println(m["Batman"])
	fmt.Println(m["Javier"])

	v, ok := m["Javier"]

	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Javier"]; ok {
		fmt.Println("Imprimiendo desde el if", v)
	}
}
