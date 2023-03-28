package main

import (
	"fmt"
	"sort"
)

type Grade struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grade{
		{"J.", "Lewis", 10},
		{"M.", "Tsoukalos", 7},
		{"D.", "Lewis", 9},
		{"Y.", "Lewis", 8},
	}

	dataSortLess := func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	}

	isSorted := sort.SliceIsSorted(data, dataSortLess)

	if isSorted {
		fmt.Println("It's sorted")
	} else {
		fmt.Println("Not sorted yet\n sorting...")

		sort.Slice(data, dataSortLess)
		fmt.Println(data)
	}
}


