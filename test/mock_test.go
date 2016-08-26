package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var json = `{
  "first_name" : "Eko Kurniawan",
  "last_name" : "Khannedy"
}`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, json)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload testing mock server
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to download json")
	{
		t.Log("When checking", server.URL)
		{
			t.Log("Should return status code", statusCode)
			response, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("Error get", err)
			}

			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				t.Fatal("Status code not equal with", http.StatusOK)
			}
		}
	}
}
