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
		fmt.Println("USAGE | HOST:PORT")
		return
	}

	connect := os.Args[1]

	s, err := net.ResolveUDPAddr("udp4", connect)
	c, err := net.DialUDP("udp4", nil, s)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">>")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("bye!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Reply: %s\n", string(buffer[:n]))

	}

}
