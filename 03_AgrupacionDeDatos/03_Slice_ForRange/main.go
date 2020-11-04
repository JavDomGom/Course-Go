package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 42}
	fmt.Println(cap(x))

	for key, value := range x {
		fmt.Println(key, value)
	}
}
