package main

import (
	"fmt"
)

type aStructure struct {
	field1 complex128
	field2 int
}

func main() {
	//aSlice := []int{1, 2, 3, 4, 5, 6}
	//add(aSlice, 1)
	//extend2(&aSlice)
	//fmt.Println("main:", aSlice)

	//f := 3.14
	//processPointer(&f)
	//fmt.Println(&f, f)

	var k *aStructure
	var v aStructure

	fmt.Println(k) // nil
	fmt.Println(v) // {(0+0i) 0}
	fmt.Printf("%+v\n", v)
}

func add(nums []int, n int) {
	for i, v := range nums {
		nums[i] = v + n
	}
	fmt.Println("Add ", nums)
}

func extend(nums []int) {
	nums = append(nums, 9, 8, 7, 8, 9, 0, 0)
	fmt.Println("extend: ", nums)

}
func extend2(nums *[]int) {
	*nums = append(*nums, 9, 8, 7, 8, 9, 0, 0)
	fmt.Println("extend2: ", *nums)

}

func processPointer(x *float64) {
	*x = *x * *x
}
