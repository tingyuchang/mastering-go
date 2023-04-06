package main

import (
	"fmt"
	"log"
	"net/http"
)

var FILEPORT = ":8765"

func defaultFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from: ", r.Host)
	w.WriteHeader(http.StatusOK)
	body := "Thanks for visiting!"
	fmt.Fprintf(w, "%s", body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultFileHandler)
	fileServer := http.FileServer(http.Dir("/Users/matt/go/src/mastering-go/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/getContents/", getFileHandler)

	fmt.Println(http.ListenAndServe(FILEPORT, mux))
}
