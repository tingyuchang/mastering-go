package main

import (
	"fmt"
	"mastering-go/dbpost"
	"math/rand"
	"time"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	starChar := "A"
	temp := ""

	var i int64 = 1

	for {
		myRand := random(MIN, MAX)
		newChar := string(starChar[0] + byte(myRand))
		temp += newChar
		if i == length {
			break
		}

		i++
	}
	return temp
}

func main() {
	dbpost.Hostname = "localhost"
	dbpost.Port = 5432
	dbpost.Username = "matt"
	dbpost.Password = "pass"
	dbpost.Database = "go"

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	random_username := getString(5)

	t := dbpost.UserData{
		Username:    random_username,
		Name:        "Matt",
		Surname:     "Chang",
		Description: "your boss!",
	}

	id := dbpost.AddUser(t)

	if id == -1 {
		fmt.Println("There is an error adding user")
	}

	data, err := dbpost.ListUser()

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}
}
