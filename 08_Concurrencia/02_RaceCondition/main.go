package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Gorutines:", runtime.NumGoroutine())

	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			v++
			runtime.Gosched() // Yield
			contador = v
			wg.Done()
		}()
		fmt.Println("Gorutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
