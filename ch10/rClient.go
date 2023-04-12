package main

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type cUser struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u1 = cUser{
	Username: "admin",
	Password: "pass",
}
var u2 = cUser{
	Username: "matt",
	Password: "pass",
}
var u3 = cUser{
	Username: "",
	Password: "pass",
}

const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

func deleteEndpoint(server string, user cUser) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)
	r, err := http.NewRequest(http.MethodDelete, server+deleteEndPoint, u)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	r.Header.Set("Content-Type", "application/json")
	c := &http.Client{Timeout: 15 * time.Second}

	resp, err := c.Do(r)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
	}

	if resp == nil {
		return http.StatusNotFound
	}

	data, err := io.ReadAll(resp.Body)
	log.Println("Delete: ", string(data))

	if err != nil {
		log.Println(err)
	}

	return resp.StatusCode
}

func getEndpoint(server string, user cUser) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)
	r, err := http.NewRequest(http.MethodGet, server+getEndPoint, u)
	r.Header.Set("Content-Type", "application/json")
	c := &http.Client{Timeout: 3 * time.Second}
	resp, err := c.Do(r)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if resp == nil {
		return http.StatusNotFound
	}

	data, err := io.ReadAll(resp.Body)
	log.Println("Get: ", string(data))
	if err != nil {
		log.Println(err)
	}

	return resp.StatusCode
}

func addEndpoint(server string, user cUser) int {
	uM, _ := json.Marshal(user)
	u := bytes.NewReader(uM)
	r, err := http.NewRequest(http.MethodPost, server+addEndPoint, u)
	r.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := c.Do(r)
	if err != nil {
		log.Println(err)
	}

	if resp == nil {
		return http.StatusNotFound
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	log.Println("Add: ", string(data))
	if err != nil {
		log.Println(err)
	}

	return resp.StatusCode
}

func timeEndpoint(server string) (int, string) {
	r, err := http.NewRequest(http.MethodGet, server+timeEndPoint, nil)
	c := &http.Client{Timeout: 3 * time.Second}

	resp, err := c.Do(r)

	if err != nil {
		log.Println(err)
	}

	if resp == nil {
		return http.StatusNotFound, ""
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	return resp.StatusCode, string(data)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server")
		return
	}
	server := os.Args[1]

	fmt.Println("/add")
	HTTPcode := addEndpoint(server, u1)
	if HTTPcode != http.StatusOK {
		fmt.Println("u1 Return code:", HTTPcode)
	} else {
		fmt.Println("u1 Data added:", u1, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u2)
	if HTTPcode != http.StatusOK {
		fmt.Println("u2 Return code:", HTTPcode)
	} else {
		fmt.Println("u2 Data added:", u2, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u3)
	if HTTPcode != http.StatusOK {
		fmt.Println("u3 Return code:", HTTPcode)
	} else {
		fmt.Println("u3 Data added:", u3, HTTPcode)
	}

	fmt.Println("/get")
	HTTPcode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u2)
	fmt.Println("/get u2 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u3)
	fmt.Println("/get u3 return code:", HTTPcode)

	fmt.Println("/delete")
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u2)
	fmt.Println("/delete u2 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u3)
	fmt.Println("/delete u3 return code:", HTTPcode)

	fmt.Println("/time")
	HTTPcode, myTime := timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)
	time.Sleep(time.Second)
	HTTPcode, myTime = timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)
}
