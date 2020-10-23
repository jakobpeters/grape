package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("static/*.html"))
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
	http.Handle("/static", http.StripPrefix("static/", http.FileServer(http.Dir("/static/*"))))
	http.Handle("/js", http.StripPrefix("js/", http.FileServer(http.Dir("/js/*"))))

}

func index(w http.ResponseWriter, r *http.Request) {
	b, err := getUsage("https://myaccount.alliantenergy.com/Portal/Usages.aspx/LoadUsage")
	if err != nil {
		log.Printf("%s\n", err)
	} else {

		err := t.ExecuteTemplate(w, "index.html", string(b))
		if err != nil {
			log.Printf("%s\n", err)
		}
	}

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
