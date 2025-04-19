package main

import (
	"fmt"
	"net/http"
)

func main() {
	f := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.Handle("/static/", http.StripPrefix("/static/", f))

	http.ListenAndServe(":80", nil)
}
