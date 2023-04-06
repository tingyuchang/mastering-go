package main

import (
	"fmt"
	"net/http"
	"os"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	var tmpFileName string

	f, err := os.CreateTemp("", "data*.txt")
	tmpFileName = f.Name()

	defer os.Remove(tmpFileName)

	err = saveJSONFile(tmpFileName)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Can't create %s\n", tmpFileName)
		return
	}

	fmt.Println("Serving: ", tmpFileName)
	http.ServeFile(w, r, tmpFileName)
}
