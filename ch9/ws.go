package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var PORT = ":1234"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome\n")
	fmt.Fprintf(w, "Please use /ws for web socket")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: ", r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()

		if err != nil {
			log.Println("From: ", r.Host, " Read: ", err)
			break
		}

		log.Println("Received: ", string(message))

		// echo
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write message: ", err)
			break
		}
	}
}

func main() {

	arguments := os.Args

	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:              PORT,
		Handler:           mux,
		IdleTimeout:       10 * time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
	}

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/ws", wsHandler)
	fmt.Println("Listen: ", PORT)
	fmt.Println(s.ListenAndServe())
}