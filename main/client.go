package main

import (
	"io"
	"log"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"os"
)

func main() {
	// Request /hello over port 8080 via the GET method
	// TLS
	caCert, err := os.ReadFile("./certificates/cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create an HTTPS client and supply the created CA pool
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	r, err := client.Get("https://localhost:8443/hello")
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
