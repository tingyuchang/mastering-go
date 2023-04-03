package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	ID  int
	Val int
}

type Result struct {
	Job Client
	Val int
}

var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {

	for c := range clients {
		output := Result{c, c.Val * c.Val}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{
			ID:  i,
			Val: i,
		}
		clients <- c
	}

	close(clients)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)

	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.Job.ID)
			fmt.Printf("%d\tsquare: %d\n", d.Job.Val, d.Val)
		}
		finished <- true
	}()

	var wg sync.WaitGroup

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(data)

	fmt.Printf("Finished: %v\n", <-finished)

}
