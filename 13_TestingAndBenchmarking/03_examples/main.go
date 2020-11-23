package main

import (
	"fmt"

	"github.com/JavierDominguezGomez/Course-Go/13_TestingAndBenchmarking/03_examples/mate"
)

func main() {
	fmt.Println(mate.Sum(2, 4, 5))
	fmt.Println(mate.Sum(4, 7, 8, 0))
}
