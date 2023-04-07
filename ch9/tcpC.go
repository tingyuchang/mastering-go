package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("USAGE | HOST:PORT ")
		return
	}

	connect := os.Args[1]
	c, err := net.Dial("tcp", connect)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(">>")
		text, _ := reader.ReadString('\n')
		// write to connect
		fmt.Fprintf(c, text+"\n")

		// received msg from remote

		message, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Println(">: ", message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Bye!")
			return
		}
	}
}
