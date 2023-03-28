package main

import "fmt"

func d1() {
	a := 1
	for i := 3; i > 0; i-- {
		defer fmt.Print(a+i, "")
	}
	a = 100
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, "")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, "")
		}(i)
	}

}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
	/*
		// D3
		123
		// D2
		???
		// D1
		123
	*/

}
