package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("USAGE | socket path")
		return
	}

	socketPath := os.Args[1]

	//_, err := os.Stat(socketPath)

	c, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">>")
		text, _ := reader.ReadString('\n')
		_, err := c.Write([]byte(text))

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)

		n, err := c.Read(buffer)

		fmt.Println("Read: ", string(buffer[:n]))

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Bye!")
			return
		}

		time.Sleep(5 * time.Second)
	}

}
