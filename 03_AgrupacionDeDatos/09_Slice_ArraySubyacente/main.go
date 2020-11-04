package main

import "fmt"

func main() {
	jd := []string{"Javier", "Domínguez", "Basket", "Patinaje", "Atletismo"}
	fmt.Println(jd)

	bg := []string{"Beatriz", "Giménez", "Rugby", "Polo", "Cricket"}
	fmt.Println(bg)

	vp := [][]string{jd, bg}
	fmt.Println(vp)
}
