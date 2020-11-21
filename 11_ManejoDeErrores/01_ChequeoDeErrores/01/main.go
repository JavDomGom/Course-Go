package main

import "fmt"

func main() {
	var res1, res2, res3 string

	fmt.Print("Nombre: ")
	_, err := fmt.Scan(&res1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Comida favorita: ")
	_, err = fmt.Scan(&res2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Número favorito: ")
	_, err = fmt.Scan(&res3)
	if err != nil {
		panic(err)
	}

	fmt.Println(res1, res2, res3)
}
