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
		fmt.Println("USAGW | PORT")
		return
	}

	PORT := ":" + os.Args[1]

	l, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println("Listen: ", err)
		return
	}

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Stopping Bye!")
			return
		}

		fmt.Println("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))

	}
}
