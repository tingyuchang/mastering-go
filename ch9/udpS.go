package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("USAGE | PORT")
		return
	}

	address := "localhost:" + os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	conn, err := net.ListenUDP("udp4", udpAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())
	for {
		n, err := conn.Read(buffer)
		fmt.Println(">>: ", string(buffer[:n]))

		if strings.TrimSpace(string(buffer[:n])) == "STOP" {
			fmt.Println("bye!")
			return
		}

		data := []byte(strconv.Itoa(rand.Intn(1000)))
		_, err = conn.WriteToUDP(data, udpAddr)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}