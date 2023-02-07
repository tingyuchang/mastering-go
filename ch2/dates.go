package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start.Unix())

	dateString := os.Args[1]

	// WTF
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Day(), d.Hour(), d.Month(), d.YearDay())
	}

	d = time.Unix(start.Unix(), 0)
	fmt.Println("Epoch: ", d)

	duration := time.Since(start)
	fmt.Println("End: ", duration)

}
