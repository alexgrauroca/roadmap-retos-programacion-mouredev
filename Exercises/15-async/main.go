package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start main")
	// Using waitgroup to ensure the main execution doesn't finish before the goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		asyncFunc(1, "first")
		wg.Done()
	}()
	fmt.Println("goroutine triggered")
	wg.Wait()

	parallelFuncs := []struct {
		name     string
		delaySec float32
	}{
		{
			name:     "C",
			delaySec: 3,
		},
		{
			name:     "B",
			delaySec: 2,
		},
		{
			name:     "A",
			delaySec: 1,
		},
	}

	for _, parallelFunc := range parallelFuncs {
		wg.Add(1)
		go func() {
			asyncFunc(parallelFunc.delaySec, parallelFunc.name)
			wg.Done()
		}()
	}
	wg.Wait()
	asyncFunc(1, "D")
}

func asyncFunc(delaySec float32, name string) {
	fmt.Println("Starting function", name, "with a duration of ", delaySec, "seconds")
	<-time.After(time.Duration(delaySec) * time.Second)
	fmt.Println("Finished function", name)
}
