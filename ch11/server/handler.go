package server

import (
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body := time.Now().Format(time.RFC1123)
	w.Write([]byte(body))
}
