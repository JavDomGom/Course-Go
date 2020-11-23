package main

import (
	"fmt"

	"github.com/JavierDominguezGomez/Course-Go/09_Paquetes/gato"
)

func main() {
	fmt.Println("Hola desde perro.")

	gato.Hola()
	gato.Comen()
}
