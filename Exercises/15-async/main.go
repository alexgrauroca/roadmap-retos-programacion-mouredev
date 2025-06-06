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
}

func asyncFunc(delaySec float32, name string) {
	fmt.Println("Starting function", name, "with a duration of ", delaySec, "seconds")
	<-time.After(time.Duration(delaySec) * time.Second)
	fmt.Println("Finished function", name)
}
