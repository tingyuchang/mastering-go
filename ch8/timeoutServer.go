package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	fmt.Println("Serving on ", PORT)

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              PORT,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	fmt.Println(server.ListenAndServe())

}
