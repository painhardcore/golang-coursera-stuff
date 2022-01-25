package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 1
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		x++
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		x = x * 10
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}

// The race condition happens when we access x variable from the 2 goroutines without synchronizing the access to it.

