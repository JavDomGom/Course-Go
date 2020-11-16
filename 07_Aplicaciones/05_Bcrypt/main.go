package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println()
	}
	fmt.Println(s)
	fmt.Println(string(bs))

	claveLogin := `clave123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("Las cadenas no coindicen.", err)
		return
	}
	fmt.Println("Las cadenas coinciden.")
}
