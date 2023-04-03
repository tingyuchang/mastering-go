package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type atomicCounter struct {
	val int64
}

func (a *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&a.val)
}

func (a *atomicCounter) Add() {
	time.Sleep(1 * time.Millisecond)
	a.val++
}

func main() {
	x := 100
	y := 4

	var wg sync.WaitGroup

	counter := atomicCounter{}

	for i := 0; i < x; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for i := 0; i < y; i++ {
				time.Sleep(1 * time.Millisecond)
				atomic.AddInt64(&counter.val, 1)
				//counter.Add()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Counter:", counter.Value())
}
