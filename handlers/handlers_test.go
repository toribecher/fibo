package handlers

import (
	"net/http"
	"testing"
)

func TestFibHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	_, err := http.NewRequest("GET", "/fibonacci", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMemoHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	_, err := http.NewRequest("GET", "/memoizedresults", nil)
	if err != nil {
		t.Fatal(err)
	}
}
