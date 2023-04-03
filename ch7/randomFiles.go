package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	start := 0
	end := 5
	for i := start; i < end; i++ {
		wg.Add(1)

		filePath := fmt.Sprintf("/Users/matt/go/src/mastering-go/ch7/temp/%d", i)

		go func(filePath string) {
			defer wg.Done()
			f, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

		}(filePath)
	}

	wg.Wait()

}
