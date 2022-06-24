package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSumAandB(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, _ := http.NewRequest("POST", "/headers", nil)

	req.Header["A"] = []string{"2"}
	req.Header["B"] = []string{"3"}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SumAandB)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response header is what we expect.
	expected := `5`
	if rr.Header()["a+b"][0] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Header(), expected)
	}
}
