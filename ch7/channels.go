package main

import (
	"fmt"
	"sync"
)

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(c chan bool, times int) {
	for i := 0; i < times; i++ {
		c <- true
	}
	close(c)
}

func main() {
	c := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeToChannel(c, 10)
		fmt.Println("exit")
	}(c)

	fmt.Println("Read: ", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is closed")
	}
	wg.Wait()
	var ch chan bool = make(chan bool)

	go printer(ch, 5)

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println()
	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}

	fmt.Println()
}
