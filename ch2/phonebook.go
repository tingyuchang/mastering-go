package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Entity struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entity{}
var MIN = 0
var MAX = 94

func random(min, max int) int {
	// generate [min, max)
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1

	for i < len {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp += newChar
		i++
	}

	return temp
}

func populate(n int, s []Entity) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		tel := strconv.Itoa(random(100, 199))
		data = append(data, Entity{name, surname, tel})
	}
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func search(key string) *Entity {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}

	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: search|list <arguments>")
		return
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	n := 100
	populate(n, data)
	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Tel number")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Not found: ", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
