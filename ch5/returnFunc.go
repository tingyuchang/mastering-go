package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	} else {
		return func(k int) int {
			return k * k
		}
	}
}

func main() {
	n := 10

	i := funRet(n)
	j := funRet(-1)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(i(10))
	fmt.Println(j(10))

}
