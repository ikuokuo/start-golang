// Redirecting HTTP to HTTPS
//  https://github.com/unrolled/secure#redirecting-http-to-https
package main

import (
	"log"
	"net/http"

	"github.com/unrolled/secure" // or "gopkg.in/unrolled/secure.v1"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
})

func main() {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     "localhost:8443", // This is optional in production. The default behavior is to just redirect the request to the HTTPS protocol. Example: http://github.com/some_page would be redirected to https://github.com/some_page.
	})

	app := secureMiddleware.Handler(myHandler)

	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(":8080", app))
	}()

	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", app))
}
