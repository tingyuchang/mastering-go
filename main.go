package main

import (
	"fmt"
	"mastering-go/ch1"
	"os"
	"path"
)

func main() {
	args := os.Args
	if !checkArguments(args) {
		return
	}

	option, argument, err := parseCommand(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	execute(option, argument)

}

func checkArguments(args []string) bool {
	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <argument>\n", exe)
		return false
	}
	return true
}

// when enter this phase, args should have at least 2 elements
func parseCommand(args []string) (option, argument string, err error) {
	option = args[1]

	switch option {
	case "search":
		if len(args) != 3 {
			return "", "", fmt.Errorf("please provides surname to search\n")
		}
		argument = args[2]
		return
	case "list":
		return
	default:
		return "", "", fmt.Errorf("Usage: search|list <argument>\n")
	}

	return
}

func execute(option, argument string) {
	switch option {
	case "search":
		result := ch1.Search(argument)
		if result == nil {
			fmt.Printf("Entity not found: %s", argument)
			return
		}
		fmt.Println(*result)
	case "list":
		ch1.List()
	default:
		fmt.Println("something wrong.")
	}
}
