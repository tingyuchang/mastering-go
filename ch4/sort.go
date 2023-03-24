package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 float64
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (a S2slice) Len() int {
	return len(a)
}

func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S2{
		{F1: 1, F2: "One", F3: S1{1, "S1_1", 10}},
		{2, "Two", S1{5, "S1_2", 10}},
		{3, "Three", S1{3, "S1_3", 10}},
		{4, "Four", S1{-3, "S1_4", 10}},
	}
	fmt.Println("Before: ", data)
	sort.Sort(S2slice(data))
	fmt.Println("After: ", data)
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse: ", data)
}
