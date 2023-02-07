package main

import "fmt"

type Digit int

const PI = 3.1415926

const (
	C1 = "C1"
)

const (
	ZERO = iota
	ONE
	TWO
	THREE
)

const (
	p2_0 Digit = 1 << iota
	_
	P2_2
	_
	p2_4
	_
	p2_6
)

func main() {

	var n Digit
	n = 5
	fmt.Println(n)
	fmt.Println(PI)
	fmt.Println(C1)
	fmt.Println(ZERO, ONE, TWO)

	fmt.Println(p2_0)
}
