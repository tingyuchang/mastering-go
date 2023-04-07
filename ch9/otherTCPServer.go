package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("USAGE | PORT")
		return
	}

	PORT := "localhost:" + os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", PORT)
	if err != nil {
		fmt.Println("ResolveTCPAddr: ", err)
		return
	}

	l, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("Listen: ", err)
		return
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(buffer[:n])) == "STOP" {
			fmt.Println("Stopping Bye!")
			return
		}

		fmt.Println("-> ", string(buffer[:n]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
