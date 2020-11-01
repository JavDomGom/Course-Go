package main

import "fmt"

func main() {
	if true {
		fmt.Println(1)
	}

	if false {
		fmt.Println(2)
	}

	if !true {
		fmt.Println(3)
	}

	if !false {
		fmt.Println(4)
	}

	if 43 == 43 {
		fmt.Println(5)
	}

	if x := 43; x <= 43 {
		fmt.Println(6)
	}
}
