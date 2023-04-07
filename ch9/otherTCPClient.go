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
		fmt.Println("USAGE || HOST:PORT")
		return
	}

	connect := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", connect)
	if err != nil {
		fmt.Println("ResolveAddr: ", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP: ", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("->: ", message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Bye!")
			conn.Close()
			return
		}
	}

}
