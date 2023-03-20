package main

import "fmt"

/*
Create a function that concatenates two arrays into a new slice.
Create a function that concatenates two arrays into a new array.
Create a function that concatenates two slices into a new array.
*/

func exercise1(a [5]int, b [6]int) []int {
	res := make([]int, len(a)+len(b))
	i := 0

	for _, v := range a {
		res[i] = v
		i++
	}

	for _, v := range b {
		res[i] = v
		i++
	}

	return res
}

func exercise2(a [5]int, b [6]int) [11]int {
	res := [11]int{}
	i := 0
	for _, v := range a {
		res[i] = v
		i++
	}

	for _, v := range b {
		res[i] = v
		i++
	}

	return res

}

func exercise3(a, b []int) [11]int {
	res := [11]int{}
	i := 0
	for _, v := range a {
		res[i] = v
		i++
	}

	for _, v := range b {
		res[i] = v
		i++
	}

	return res
}

func main() {
	a, b := [5]int{1, 2, 3, 4, 5}, [6]int{6, 7, 8, 9, 10, 11}
	fmt.Println(exercise1(a, b))
	fmt.Println(exercise2(a, b))
	c, d := []int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11}
	fmt.Println(exercise3(c, d))
}
