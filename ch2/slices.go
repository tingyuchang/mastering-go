package main

import "fmt"

func main() {
	a := [...]int{4, 2, 3, 1}

	fmt.Printf("%T, %v\n", a, a)
	a2 := double(a)
	fmt.Println("After double: ", a)
	fmt.Println("After double: ", a2)

	b := []int{}
	fmt.Printf("%T, %v\n", b, b)

	c := []int{1, 2, 3, 4}

	d := c

	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	d[0] = 5
	fmt.Println("after 5")
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	c = c[0:2:3]
	fmt.Println("after rearrange")
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	d[0] = 9
	fmt.Println("after 9")
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	c = append(c, 5, 6, 7)
	fmt.Println("after append")
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	d[0] = 1
	c[1] = 44
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
}

func double(nums [4]int) [4]int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}

	return nums
}

func deleteSliceV1(nums []int, n int) {
	nums = append(nums[:n], nums[n+1:]...)
}

func deleteSliceV2(nums []int, n int) {
	nums[n] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]
}
