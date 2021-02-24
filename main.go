package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	b, err := getUsage("https://myaccount.alliantenergy.com/Portal/Usages.aspx/LoadUsage")
	if err != nil {
		log.Printf("%s\n", err)
	}
	fmt.Fprint(w, string(b))
}

func getUsage(s string) (b []byte, err error) {
	w, err := http.Get(s)
	if err != nil {
		log.Printf("%s\n", err)
	}
	defer w.Body.Close()
	b, err = ioutil.ReadAll(w.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
