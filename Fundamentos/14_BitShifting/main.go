package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a)
	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b)
	c := b << 1
	fmt.Printf("%d\t\t%b\n", c, c)

	fmt.Println()

	// kb := 1024
	// gb := kb * 1024
	// tb := gb * 1024
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t%b\n", tb, tb)
}
