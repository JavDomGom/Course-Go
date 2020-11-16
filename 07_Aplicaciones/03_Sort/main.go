package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"Javier", "A", "B", "C", "Beatriz", "Z", "Y"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
