package main

import (
	"fmt"
	"net"
	"os"
)

func echo(c net.Conn) {
	defer c.Close()
	for {

		buffer := make([]byte, 1024)
		n, err := c.Read(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}
		data := buffer[:n]
		fmt.Println("Server got: ", string(data))
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("USAGE | socket path")
		return
	}

	socketPath := os.Args[1]

	_, err := os.Stat(socketPath)

	if err == nil {
		fmt.Println("Deleting existing ", socketPath)
		err = os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go echo(fd)
	}
}
