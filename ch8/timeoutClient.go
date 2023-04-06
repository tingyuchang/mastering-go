package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var myUrl string
var delay int = 5
var wg sync.WaitGroup

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer wg.Done()

	data := make(chan myData, 1)

	tr := &http.Transport{}

	httpClient := &http.Client{Transport: tr}

	request, _ := http.NewRequest(http.MethodGet, myUrl, nil)
	request = request.WithContext(c)
	go func() {
		response, err := httpClient.Do(request)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
		} else {
			data <- myData{response, nil}
		}
	}()

	select {
	case <-c.Done():
		// using Request.withContext instead
		//tr.CancelRequest()
		<-data
		fmt.Println("The request was canceled.")
		return c.Err()
	case ok := <-data:
		err := ok.err

		if err != nil {
			fmt.Println(err)
			return err
		}

		response := ok.r

		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("Server Response: %s\n", data)
	}

	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("USAGE | URL")
		return
	}

	myUrl = os.Args[1]

	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay: ", delay)

	c := context.Background()

	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Millisecond)
	defer cancel()

	fmt.Println("Connect to ", myUrl)
	wg.Add(1)
	go connect(c)
	wg.Wait()

	fmt.Println("Exiting..")

}
