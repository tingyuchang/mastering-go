package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return c.R * 2 * math.Pi
}

func main() {
	a := circle{
		1.5,
	}

	fmt.Printf("R %.2f -> Perimeter %.3f\n", a.R, a.Perimeter())

	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is Shape2D")
	}
}
