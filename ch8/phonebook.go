package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"tel"`
	LastAccess string `json:"lastaccess"`
}

// JSONFILE resides in the current directory
var JSONFILE = "./data.json"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int

// Implement sort.Interface
func (a PhoneBook) Len() int {
	return len(a)
}

// First based on surname. If they have the same
// surname take into account the name.
func (a PhoneBook) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = Serialize(&data, f)
	if err != nil {
		return err
	}

	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
}

func setJSONFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		JSONFILE = filepath
	}

	_, err := os.Stat(JSONFILE)
	if err != nil {
		fmt.Println("Creating", JSONFILE)
		f, err := os.Create(JSONFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	fileInfo, err := os.Stat(JSONFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", JSONFILE)
	}
	return nil
}

func initData() {
	err := setJSONFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readJSONFile(JSONFILE)
	// io.EOF is fine because it means the file is empty
	if err != nil && err != io.EOF {
		return
	}
	createIndex()
}

/*

/list: This lists all available entries.
/insert/name/surname/telephone/: This inserts a new entry. Later on, we are going to see how to extract the desired information from a URL that contains user data.
/delete/telephone/: This deletes an entry based on the value of telephone.
/search/telephone/: This searches for an entry based on the value of telephone.
/status: This is an extra URL that returns the number of entries in the phone book.

*/

func main() {
	initData()
	PORT := ":1234"

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:              PORT,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert/", http.HandlerFunc(insertHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := "test\n"
		_, _ = fmt.Fprintf(w, "%s", body)
	})

	err := s.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body := "Thanks for visiting!\n"
	_, _ = fmt.Fprintf(w, "%s", body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path: ", paramStr)
	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	telephone := paramStr[2]
	err := deleteEntry(telephone)

	if err != nil {
		fmt.Println(err)
		body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "%s", body)
		return
	}

	body := telephone + " deleted\n"
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s", body)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body, err := list()

	if err != nil {
		fmt.Println(err)
		body := err.Error() + "\n"
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%s", body)
		return
	}
	_, _ = fmt.Fprintf(w, "%s", body)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := fmt.Sprintf("Total entries: %d\n", len(data))
	_, _ = fmt.Fprintf(w, "%s", Body)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path: ", paramStr)
	if len(paramStr) < 5 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}
	name := paramStr[2]
	surname := paramStr[3]
	tel := paramStr[4]
	t := strings.ReplaceAll(tel, "-", "")
	if !matchTel(t) {
		fmt.Println("Not a valid telephone number:", tel)
		return
	}

	temp := &Entry{
		Name:    name,
		Surname: surname,
		Tel:     tel,
	}

	err := insert(temp)
	var body string
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		body = "Insert Failure" + err.Error() + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		body = "Insert success\n"
	}

	_, _ = fmt.Fprintf(w, "%s", body)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path: ", paramStr)
	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, "Not enough arguments: "+r.URL.Path)
		return
	}
	var body string
	telephone := paramStr[2]
	t := search(telephone)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		body = "Could not bre found: " + telephone + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		body, _ = PrettyPrintJSONstream(t)
	}
	fmt.Println("Serving", r.URL.Path, "from", r.Host)
	_, _ = fmt.Fprintf(w, "%s", body)
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}
	data = append(data[:i], data[i+1:]...)

	// Update the index - key does not exist any more
	delete(index, key)

	err := saveJSONFile(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func list() (string, error) {
	sort.Sort(PhoneBook(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		return "", err
	}

	return text, nil
}

// PrettyPrintJSONstream pretty prints the contents of the phone book
func PrettyPrintJSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	createIndex()
	// Save the data
	err := saveJSONFile(JSONFILE)
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

	return &data[i]
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}
