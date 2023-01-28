package main

import "fmt"

func main() {
	x, k := 12, 5
	var m, n float64
	m = 1.234
	fmt.Println("m, n: ", m, n)

	y := 4.3 / 2

	fmt.Printf("Y: %T, %v\n", y, y)

	divFloat := float64(x) / float64(k)
	fmt.Printf("divFloat: %T %v \n", divFloat, divFloat)

}
