package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

func processA(a IntA) {
	fmt.Printf("%T\n", a)
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

type compose struct {
	field1 int
	a
}

func (varC c) foo() {
	fmt.Println("foo")
}

func (varC c) bar() {
	fmt.Println("bar")
}

func (A a) A() {
	fmt.Println("function A() for a")
}

func (B b) A() {
	fmt.Println("function A() for b")
}

func main() {
	var iC c = c{a{120, 12}, b{"-12", -12}}
	iC.A.A()
	iC.B.A()

	iCom := compose{
		field1: 123,
		a: a{
			456, 789,
		},
	}
	fmt.Println(iCom.XX, iCom.YY)

	iC.bar()
	processA(iC)
}