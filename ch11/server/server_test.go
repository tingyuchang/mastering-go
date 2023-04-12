package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimeHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TimeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("Status code is not 200")
	}

}