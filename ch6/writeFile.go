package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	argements := os.Args
	if len(argements) != 3 {
		fmt.Println("Usage | filename | data")
		return
	}

	filePath := argements[1]
	data := argements[2]
	F4(filePath, data)
}

func F1(filePath, data string) {
	buffer := []byte(data)
	f1, err := os.Create(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	n, err := fmt.Fprintf(f1, string(buffer))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d data to %s", n, filePath)
}

func F2(filePath, data string) {
	buffer := []byte(data)
	f2, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	n, err := f2.WriteString(string(buffer))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d data to %s", n, filePath)

}

func F3(filePath, data string) {
	buffer := []byte(data)
	f3, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err := w.WriteString(string(buffer))
	w.Flush()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Write %d data to %s", n, filePath)
}

func F4(filePath, data string) {
	buffer := []byte(data)
	f4, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	for i := 0; i < 5; i++ {
		n, err := io.WriteString(f4, string(buffer))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d Write %d data to %s", i, n, filePath)
	}

	f4, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()
	n, err := f4.Write([]byte("Append data to exist file"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Write %d data to %s", n, filePath)

}
