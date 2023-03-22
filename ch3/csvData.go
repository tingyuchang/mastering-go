package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

//var index map[string]int
//var CSVFIlE = "/Users/matt/go/src/mastering-go/ch3/csv.data"

func read(filePath string) error {
	_, err := os.Stat(filePath)

	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Record{
			line[0],
			line[1],
			line[2],
			line[3],
		}
		myData = append(myData, temp)
	}

	return nil
}

func save(filePath string) error {
	csvfile, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = ','
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("CSV data input output!")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	err := read(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	err = save(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}
