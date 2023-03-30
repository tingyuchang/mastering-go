package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)

	if err == io.EOF {
		return nil
	} else if err != nil {
		return nil
	}

	return buffer[:n]
}

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		return
	}

	f, err := os.Open(arguments[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	size, _ := strconv.Atoi(arguments[1])

	fmt.Println(string(readSize(f, size)))
}
