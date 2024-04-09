package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Request /hello over port 8080 via the GET method
	// NON_TLS
	r, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	log.Printf("%s\n", body)
}
