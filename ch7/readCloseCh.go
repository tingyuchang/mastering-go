package main

import "fmt"

func main() {
	willClose := make(chan complex64, 10)

	willClose <- -1
	willClose <- 1i
	fmt.Println(<-willClose)
	fmt.Println(<-willClose)
	close(willClose)

	read := <-willClose
	fmt.Println(read)
}
