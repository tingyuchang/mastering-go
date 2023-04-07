package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var SERVER = ""
var PATH = ""
var TIMESWAIT = 0
var TIMESWAITMAX = 5

func getInput(input chan string) {
	result, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input <- result
}
func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("USAGE | SERVER | PATH")
		return
	}

	SERVER = arguments[1]
	PATH = arguments[2]
	fmt.Println("Connect to: ", SERVER, " at ", PATH)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)

	URL := url.URL{
		Scheme: "ws",
		Host:   SERVER,
		Path:   PATH,
	}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received: %s", string(message))
		}
	}()

	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("Please enter your data!")
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println(err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)
		case <-interrupt:
			log.Println("Caught interrupt signal, quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println(err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}

}
