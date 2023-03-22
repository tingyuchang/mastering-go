package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var data = []Entry{}
var index map[string]int
var CSVFIlE = "/Users/matt/go/src/mastering-go/ch3/csv.data"

func readCSVFile(filePath string) error {
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
		temp := Entry{
			line[0],
			line[1],
			line[2],
			line[3],
		}
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filePath string) error {
	csvfile, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = ','
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	argements := os.Args

	if len(argements) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	// read

	_, err := os.Stat(CSVFIlE)
	if err != nil {
		fmt.Println("Creating: ", CSVFIlE)
		f, err := os.Create(CSVFIlE)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}

	fileInfo, err := os.Stat(CSVFIlE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFIlE, " not a regular file!")
		return
	}

	err = readCSVFile(CSVFIlE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Can't create index: ", err)
		return
	}

	switch argements[1] {
	case "insert":
		if len(argements) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(argements[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number: ", t)
			return
		}
		temp := initEntry(argements[2], argements[3], t)
		if temp != nil {
			err := insert(temp)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(argements) != 3 {
			fmt.Println("Usage delete number")
			return
		}
		t := strings.ReplaceAll(argements[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid number: ", t)
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(argements) != 3 {
			fmt.Println("Usage search number")
			return
		}
		t := strings.ReplaceAll(argements[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid number: ", t)
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found: ", t)
			return
		}
		fmt.Printf("%+v\n", *temp)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option.")
	}
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Number] = i
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]

	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}

	data = append(data[:i], data[i+1:]...)
	delete(index, key)
	err := saveCSVFile(CSVFIlE)
	if err != nil {
		return err
	}
	return nil
}

func insert(entry *Entry) error {
	_, ok := index[(*entry).Number]
	if ok {
		return fmt.Errorf("%s already exist", entry.Number)
	}

	data = append(data, *entry)

	_ = createIndex()
	err := saveCSVFile(CSVFIlE)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func matchTel(num string) bool {
	t := []byte(num)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func initEntry(name, surname, number string) *Entry {
	if name == "" || surname == "" {
		return nil
	}
	lastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{name, surname, number, lastAccess}
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}
