package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("archivo_que_no_existe.txt")
	if err != nil {
		log.Println("Ocurrió un error:", err)
	}
	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt.")
}
