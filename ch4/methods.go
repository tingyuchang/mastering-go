package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] += b[i][j]
		}
	}
}

func (a *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] -= b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		fmt.Println("Need 8 numbers!")
		return
	}

	k := [8]int{}

	for i, v := range os.Args[1:] {
		v, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[i] = v
	}

	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}
	a.Add(b)
	fmt.Println("a + b ", a)
	a.Subtract(b)
	fmt.Println("a-b ", b)
	a.Multiply(b)
	fmt.Println("a * b ", a)

}
