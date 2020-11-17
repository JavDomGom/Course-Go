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
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			runtime.Gosched() // Yield
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Gorutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
