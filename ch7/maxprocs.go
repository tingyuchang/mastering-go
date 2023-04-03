package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH, " machine")
	fmt.Println("Go version: ", runtime.Version())
	fmt.Println("processes: ", runtime.GOMAXPROCS(0))
}
