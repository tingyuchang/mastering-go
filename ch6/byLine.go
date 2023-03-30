package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func lineByLine(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Print(line)
	}

	return nil
}

func wordByword(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		r := regexp.MustCompile("[^\\s]]+")

		words := r.FindAllString(line, -1)
		for _, word := range words {
			fmt.Println(word)
		}
	}
	return nil
}

func charByChar(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

func main() {
	_ = charByChar(os.Args[1])
}
