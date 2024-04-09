package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	fmt.Println("Server started with endpoint /hello on port 8080")
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Listen to port 8080 and wait
	// Front-Door-TLS
	log.Fatal(http.ListenAndServeTLS(":8443", "./certificates/cert.pem", "./certificates/key.pem", nil))
}
