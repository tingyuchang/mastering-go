package main

import "fmt"

func returnAValue() interface{} {
	return 1
}

func main() {
	a := returnAValue()

	num, _ := a.(int)
	fmt.Println(num)
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("not string")
	}
}
