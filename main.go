package main

import (
	"runtime"
	"sync"
)

func goroutine() {
	x := 0

	for {
		x++
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(runtime.NumCPU())

	for i := 0; i <= runtime.NumCPU(); i++ {
		go goroutine()
	}

	wg.Wait()
}
