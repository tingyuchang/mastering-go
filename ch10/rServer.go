package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var user User

var PORT = ":1234"

var DATA = make(map[string]string)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, " from: ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	body := "Thanks for visiting!"
	fmt.Fprintf(w, "%s\n", body)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, " from: ", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := "The current time is " + t
	fmt.Fprintf(w, "%s\n", body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, " from: ", r.Host)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed!\n")
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	if user.Username != "" {
		DATA[user.Username] = user.Password
		log.Println(DATA)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, " form: ", r.Host)
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	_, ok := DATA[user.Username]

	if ok && user.Username != "" {
		log.Println("Found user: ", user.Username)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not found: ", user.Username)
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map resource not fond", http.StatusNotFound)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, " from: ", r.Host)
	d, err := io.ReadAll(r.Body)

	if r.Method != http.MethodDelete {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	log.Println(user)

	_, ok := DATA[user.Username]

	if ok && user.Username != "" {
		delete(DATA, user.Username)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not found: ", user.Username)
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map resource not fond", http.StatusNotFound)
	}
}

func main() {
	arguments := os.Args

	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/delete", deleteHandler)

	err := s.ListenAndServe()

	if err != nil {
		log.Println(err)
		return
	}
}
