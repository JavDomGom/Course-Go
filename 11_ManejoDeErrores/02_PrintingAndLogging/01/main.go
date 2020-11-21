package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("archivo_que_no_existe.txt")
	if err != nil {
		log.Println("Ocurrió un error:", err)
	}
}
