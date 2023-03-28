package main

import "fmt"

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := 0.0

	for _, v := range s {
		sum += v
	}

	return sum
}

func printEverything(input ...interface{}) {
	fmt.Println(input)
}

func main() {
	sum := addFloats("Hi", 1.1, 2.2, 3.3)
	printEverything(sum)
}