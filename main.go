package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":443",
		Handler: nil,
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("*"))))
	err := httpServer.ListenAndServeTLS("/Cert", "/Key")
	if err != nil {
		fmt.Printf("%s", err)
	}
}
