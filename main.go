package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)
var tlsCert = /path/to/cert
var tlsKey = /path/to/key

func main() {
	srv := &http.Server{
		Addr:         ":443",
		Handler:      nil,
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("*"))))
	err := srv.ListenAndServeTLS("tlsCert", "tlsKey")
	if err != nil {
		fmt.Printf("could not serve tls")
	}
}