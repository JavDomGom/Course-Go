package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Gorutines:", runtime.NumGoroutine())
	fmt.Println("GOARCH:", runtime.GOARCH)

	var contador int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched() // Yield
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Gorutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
