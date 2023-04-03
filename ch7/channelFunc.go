package main

import "fmt"

func printer1(ch chan<- bool) {
	ch <- true
}

func writeToChannel1(ch chan<- int, n int) {
	fmt.Println("1", n)
	ch <- n
	fmt.Println("2", n)
}

func f2C(out <-chan int, in chan<- int) {
	x := <-out
	in <- x
}
