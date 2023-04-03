package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newVal int) {
	writeValue <- newVal
}

func read2() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("reveive new value %v \n", newValue)
		case readValue <- value:

		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Please give a number!\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers\n", n)

	rand.Seed(time.Now().Unix())

	go monitor()

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Println("Last value: ", read2())
}