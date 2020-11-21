package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrMat is a var.
var ErrMat = errors.New("matemática elemental: No hay raíz cuadarada real de un número negativo")

func main() {
	fmt.Printf("%T\n", ErrMat)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMat
	}
	return 42, nil
}
