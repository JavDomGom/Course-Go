package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// fmt.Println()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println()

	i := 0
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
}
