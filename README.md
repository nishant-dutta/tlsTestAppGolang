# Overview
This project demonstrates the use of "net/http" package to create an HTTP server and HTTP client.

The project has three branches:
* non-tls: It demonstrates a simple HTTP connection
* front-door-tls: Also known as One-way-TLS, it demonstrates only server verification
* mutual-tls: Also known as Two-way-TLS or mTLS, it demonstrates both the client and server certificate verification. Although, the demonstration uses only one set of certificates for this but a real production application will use two sets(Just replace the certificate path)

### Sample Project using gentleman as the client
https://github.com/nishant-dutta/tlsTestAppGentleman

## Generating Self-Signed Certificates
A self-signed certificate can be generated using the command mentioned in certificates/generate_certificates.sh. It has two set of commands which can be used depending on the Operating System:
* Ubuntu
* wsl/wsl-2(Windows)
