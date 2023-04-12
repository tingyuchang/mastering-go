package main

import "fmt"

func main() {
	// list user

	users := ListUsers()

	if len(users) == 0 {
		fmt.Println("Can't query users from db")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}

	user := User{
		Username: "Mika",
		Password: "Chen",
	}
	err := InsertUser(user)

	if err != nil {
		fmt.Println(err)
	}

}
