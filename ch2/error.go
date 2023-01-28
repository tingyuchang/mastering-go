package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println(check(0, 0))
	fmt.Println(formattedError(0, 0))
}

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d . UserId: %d", a, b, os.Getuid())
	}
	return nil
}
