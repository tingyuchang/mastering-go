package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE: %s needs URL", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	c := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       1 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Status code", data.StatusCode)
	header, _ := httputil.DumpResponse(data, false)
	fmt.Print(string(header))

	contentType := data.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")

	if len(characterSet) > 1 {
		fmt.Println("Character Set: ", characterSet[1])
	}

	if data.ContentLength == -1 {
		fmt.Println("Content length unknow!")
	} else {
		fmt.Println("Content length: ", data.ContentLength)
	}

	length := 0

	var buffer [1024]byte

	r := data.Body
	for {
		n, err := r.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length += n
	}

	fmt.Println("data length: ", length)
}
