package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	a := T{
		F1: 1,
		F2: "F2",
		F3: 3.0,
	}

	fmt.Println("a: ", a)

	r := reflect.ValueOf(&a).Elem()

	fmt.Println("String value", r.String())
	typeOfA := r.Type()

	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Updated!")
		}
	}

	fmt.Println("a: ", a)
}
