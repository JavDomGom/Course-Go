package main

import (
	"os"
)

func main() {
	_, err := os.Open("archivo_que_no_existe.txt")
	if err != nil {
		panic(err)
	}
}
