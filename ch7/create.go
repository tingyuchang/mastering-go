package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 10
	var wg sync.WaitGroup

	fmt.Printf("Going to create %d goroutines\n", count)
	for i := 0; i < count; i++ {
		go func(n int) {
			fmt.Println(n)
			defer wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Exiting...")
}
