package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name     string `json:"username"`
	Surname  string `json:"surname"`
	Year     int    `json:"year,omitempty"`
	Pass     bool   `json:"-"`
	CreateAt string `json:"createAt,omitempty"`
}

func main() {
	useall := UseAll{
		Name:    "Mika",
		Surname: "Chen",
		Year:    2022,
	}

	t, err := json.Marshal(&useall)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value: %s\n", t)
	}

	str := `{"username":"Roll","surname":"Welly","year":1999}`
	temp := UseAll{}

	err = json.Unmarshal([]byte(str), &temp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value: %+v\n", temp)
	}
}
