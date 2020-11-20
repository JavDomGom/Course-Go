package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:\t\t", ctx)
	fmt.Println("error de context:\t", ctx.Err())
	fmt.Printf("tipo de context:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, _ = context.WithCancel(ctx)
	fmt.Println("context:\t\t", ctx)
	fmt.Println("error de context:\t", ctx.Err())
	fmt.Printf("tipo de context:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("context:\t\t", ctx)
	fmt.Println("error de context:\t", ctx.Err())
	fmt.Printf("tipo de context:\t%T\n", ctx)
	fmt.Println("cancel:\t\t\t", cancel)
	fmt.Printf("tipo de cancel:\t\t%T\n", cancel)
	fmt.Println("----------")

	cancel()
}
