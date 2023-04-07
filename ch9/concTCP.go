package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))

		if temp == "STOP" {
			fmt.Println("BYE!")
			break
		}

		fmt.Println(temp)

		counter := "Client numbers: " + strconv.Itoa(count) + "\n"
		c.Write([]byte(counter))
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("USAGE | PORT")
		return
	}

	PORT := ":" + os.Args[1]

	l, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
		count++

	}

}
