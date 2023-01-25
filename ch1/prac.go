package ch1

import (
	"fmt"
	"os"
	"strconv"
)

func checkArgs() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("unexpected inputs")
	}
	return nil
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}

func argsHandler(arg string) {
	value, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println(err)
	}

	switch {
	case value == 1:
		fmt.Println("option 1")
		loopIdiomatic()
	case value == 2:
		fmt.Println("option 2")
		printUserInput()
	case value == 3:
		fmt.Println("Find max & min value")
		max, min := findMaxAndMinValue(os.Args[2:])
		fmt.Printf("max: %.2f min: %.2f\n", max, min)
	case value < 0:
		fmt.Println("Error: ")
	default:
		fmt.Println("unexpected value: ", value)
	}
}

func loopIdiomatic() {
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range aSlice {
		fmt.Printf("%d: %d\n", i, v)
	}

}

func printUserInput() {
	var value string
	fmt.Println("Pleaes enter your name!")
	fmt.Scanln(&value)
	fmt.Println("Your name is ", value)

}

func findMaxAndMinValue(numbers []string) (max, min float64) {
	for i, v := range numbers {
		vf, err := strconv.ParseFloat(v, 64)

		if err != nil {
			panic(err)
		}

		if i == 0 {
			min = vf
			max = vf
		} else {
			if vf > max {
				max = vf
			}
			if vf < min {
				min = vf
			}
		}

	}

	return max, min
}
