package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("archivo_que_no_existe.txt")
	if err != nil {
		log.Fatalln("Ocurrió un error:", err)
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren.")
}
