package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []byte("A String")
	fmt.Printf("%v Type: %T\n", a, a)
	fmt.Printf("%v Type: %T\n", a[0], a[0])
	fmt.Printf("%v Type: %T\n", string(a), string(a))
	fmt.Printf("A String Type: %T\n", "A String")

	r := '。'
	fmt.Printf("%v %c Type: %T\n", r, r, r)
	r2 := "A string。"
	for _, v := range r2 {
		fmt.Printf("%x %s, %v, %c  %T\n", v, v, v, v, v)
	}

	fmt.Println("-- string convert --")
	n := 100
	unicodeString := string(n)
	iotaString := strconv.Itoa(n)
	formatString := strconv.FormatInt(int64(n), 10)

	fmt.Printf("%s %T\n", unicodeString, unicodeString)
	fmt.Printf("%s %T\n", iotaString, iotaString)
	fmt.Printf("%s %T\n", formatString, formatString)

	fmt.Printf("Equal Fold: %v\n", strings.EqualFold("Matt", "matt"))
	fmt.Println(strings.Fields("This is a string!"))
	fmt.Println(strings.Fields("ThisIs a\tstring!"))
}
